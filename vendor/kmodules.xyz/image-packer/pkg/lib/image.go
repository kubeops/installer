/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package lib

import (
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"kmodules.xyz/client-go/tools/parser"
	"kmodules.xyz/go-containerregistry/name"

	shell "gomodules.xyz/go-sh"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
)

func ListDockerImages(rootDir string, values map[string]string) ([]string, error) {
	images, err := MapImages(rootDir, values)
	if err != nil {
		return nil, err
	}
	return ListImages(images), nil
}

func MapImages(rootDir string, values map[string]string) (map[string]string, error) {
	entries, err := os.ReadDir(rootDir)
	if err != nil {
		return nil, err
	}

	sh := shell.NewSession()
	sh.SetDir(rootDir)
	sh.ShowCMD = true

	images := map[string]string{}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		mapChartImages(rootDir, values, sh, entry, images)
	}
	return images, nil
}

func mapChartImages(rootDir string, values map[string]string, sh *shell.Session, entry os.DirEntry, images map[string]string) {
	chartName := entry.Name()
	if !strings.HasSuffix(chartName, "-certified") &&
		!strings.HasSuffix(chartName, "-certified-crds") {
		err := sh.SetDir(filepath.Join(rootDir, chartName)).Command("helm", "dependency", "update").Run()
		if err != nil {
			panic(err)
		}
	}

	args := []any{"template", chartName}

	content, ok := values[chartName]
	if ok {
		filename, err := writeTempValues(chartName, []byte(content))
		if err != nil {
			klog.Fatal(err)
		}
		defer os.Remove(filename) // nolint:errcheck

		args = append(args, "--values="+filename)
	}

	if _, err := os.Stat(filepath.Join(rootDir, chartName, "ci", "ci-values.yaml")); err == nil {
		args = append(args, "--values="+chartName+"/ci/ci-values.yaml")
	}

	if chartName == "cluster-manager-spoke" {
		args = append(args, "--dry-run=server")
	} else {
		if files, err := filepath.Glob(filepath.Join(rootDir, chartName, "*.sample.yaml")); err == nil && len(files) > 0 {
			for _, file := range files {
				args = append(args, "--values="+chartName+"/"+filepath.Base(file))
			}
		}
	}
	if out, err := sh.SetDir(rootDir).Command("helm", args...).Output(); err == nil {
		if err := CollectRenderedImages(out, images); err != nil {
			panic(err)
		}
	} else {
		klog.Infof("Skipping %s due to error: %v", chartName, err)
	}
}

func writeTempValues(chartName string, content []byte) (string, error) {
	tmpfile, err := os.CreateTemp("", chartName+"-val-*.yaml")
	if err != nil {
		return "", err
	}

	if _, err := tmpfile.Write(content); err != nil {
		tmpfile.Close()           // nolint:errcheck
		os.Remove(tmpfile.Name()) // nolint:errcheck
		return "", err
	}

	// The handle must be closed before helm reads the file.
	if err := tmpfile.Close(); err != nil {
		os.Remove(tmpfile.Name()) // nolint:errcheck
		return "", err
	}
	return tmpfile.Name(), nil
}

// CollectRenderedImages records every image referenced by the resources in a
// `helm template` output into images, keyed by image and valued by the GroupKind
// of the resource that referenced it.
func CollectRenderedImages(out []byte, images map[string]string) error {
	resources, err := parser.ListResources(out)
	if err != nil {
		return err
	}
	for _, ri := range resources {
		collectImages(ri.Object.UnstructuredContent(), images, ri.Object.GetObjectKind().GroupVersionKind().GroupKind().String())
	}
	return nil
}

// placeholderRE matches a shell-style ${...} template placeholder.
var placeholderRE = regexp.MustCompile(`\$\{[^}]*\}`)

// containerArgRE matches a --flag=value container argument.
var containerArgRE = regexp.MustCompile(`^--[A-Za-z0-9._-]+=(\S+)$`)

// imageFromContainerArg reports an image reference passed to a container as a
// flag, e.g. --acme-http01-solver-image=<ref> or --prometheus-config-reloader=<ref>.
// An operator that launches other workloads takes their image this way, and the
// reference appears nowhere else in the manifest.
func imageFromContainerArg(arg string) (string, bool) {
	m := containerArgRE.FindStringSubmatch(arg)
	if m == nil {
		return "", false
	}
	ref := m[1]

	// name.ParseReference is far too permissive on its own: it defaults the
	// registry to docker.io and the tag to latest, so it accepts most flag
	// values (--log-level=info parses). Demand an explicit registry host and an
	// explicit tag or digest, which every image passed this way carries.
	if strings.Contains(ref, "://") {
		return "", false
	}
	host, remainder, ok := strings.Cut(ref, "/")
	if !ok || (!strings.ContainsAny(host, ".:") && host != "localhost") {
		return "", false
	}
	if last := remainder[strings.LastIndex(remainder, "/")+1:]; !strings.ContainsAny(last, ":@") {
		return "", false
	}
	if _, err := name.ParseReference(ref); err != nil {
		return "", false
	}
	return ref, true
}

func collectImages(obj map[string]any, images map[string]string, srcGK string) {
	for k, v := range obj {
		if k == "image" {
			if s, ok := v.(string); ok && strings.ContainsRune(s, ':') {
				for _, img := range expandVersionedImage(s, obj) {
					images[img] = srcGK
				}
			}
		} else if k == "args" || k == "command" {
			if items, ok := v.([]any); ok {
				for _, item := range items {
					if s, ok := item.(string); ok {
						if img, ok := imageFromContainerArg(s); ok {
							images[img] = srcGK
						}
					}
				}
			}
		} else if m, ok := v.(map[string]any); ok {
			collectImages(m, images, srcGK)
		} else if items, ok := v.([]any); ok {
			for _, item := range items {
				if m, ok := item.(map[string]any); ok {
					collectImages(m, images, srcGK)
				}
			}
		}
	}
}

// expandVersionedImage expands a ${...} placeholder in an image reference using
// the sibling "availableVersions" list found on the same object (kubestash addon
// Function resources carry a "v0.29.0_${DB_VERSION}"-style tag alongside the list
// of versions the image is published for). If the reference has no placeholder,
// or no usable availableVersions is present, it is returned unchanged; unexpanded
// placeholders are dropped later by ListImages/GroupImages.
func expandVersionedImage(image string, obj map[string]any) []string {
	if !placeholderRE.MatchString(image) {
		return []string{image}
	}

	versions, ok := obj["availableVersions"].([]any)
	if !ok || len(versions) == 0 {
		return []string{image}
	}

	out := make([]string, 0, len(versions))
	for _, v := range versions {
		if ver, ok := v.(string); ok {
			out = append(out, placeholderRE.ReplaceAllString(image, ver))
		}
	}
	if len(out) == 0 {
		return []string{image}
	}
	return out
}

func GroupImages(images map[string]string) map[string][]string {
	result := map[string][]string{}
	for img, srcGK := range images {
		if strings.Contains(img, "${") {
			continue
		}
		result[srcGK] = append(result[srcGK], img)
	}
	return result
}

func ListImages(images map[string]string) []string {
	result := make([]string, 0, len(images))
	for img := range images {
		if strings.Contains(img, "${") {
			continue
		}
		result = append(result, img)
	}
	sort.Strings(result)

	return result
}

func HasGroupKind(images map[string]string, in schema.GroupKind) bool {
	for _, srcGK := range images {
		gk := schema.ParseGroupKind(srcGK)
		if gk.Group == in.Group && (in.Kind == "" || gk.Kind == in.Kind) {
			return true
		}
	}
	return false
}
