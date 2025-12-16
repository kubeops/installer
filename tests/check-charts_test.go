/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"path/filepath"
	"runtime"
	"testing"

	"kmodules.xyz/image-packer/pkg/lib"
)

var ignoreMissingList = []string{
	"ghcr.io/appscode/config-syncer:v0.15.4",
}

func Test_CheckImageArchitectures(t *testing.T) {
	dir, err := rootDir()
	if err != nil {
		t.Error(err)
	}

	if err := lib.CheckImageArchitectures([]string{
		filepath.Join(dir, "catalog", "imagelist.yaml"),
	}, nil, ignoreMissingList); err != nil {
		t.Errorf("CheckImageArchitectures() error = %v", err)
	}
}

func Test_CheckUBIImageArchitectures(t *testing.T) {
	dir, err := rootDir()
	if err != nil {
		t.Error(err)
	}

	const (
		//	ubiAll = `global:
		// distro:
		//  ubi: all`
		ubiOperator = `distro:
  ubi: operator`
		//	ubiCatalog = `distro:
		// ubi: catalog`
	)
	values := map[string]string{
		"cert-manager-csi-driver-cacerts": ubiOperator,
		"external-dns-operator":           ubiOperator,
		"falco-ui-server":                 ubiOperator,
		"fargocd":                         ubiOperator,
		"kube-ui-server":                  ubiOperator,
		"operator-shard-manager":          ubiOperator,
		"panopticon":                      ubiOperator,
		"petset":                          ubiOperator,
		"pgoutbox":                        ubiOperator,
		"scanner":                         ubiOperator,
		"sidekick":                        ubiOperator,
		"supervisor":                      ubiOperator,
		"taskqueue":                       ubiOperator,
	}
	if err := lib.CheckHelmChartImageArchitectures(filepath.Join(dir, "charts"), values, nil, ignoreMissingList); err != nil {
		t.Errorf("CheckUBIImageArchitectures() error = %v", err)
	}
}

func rootDir() (string, error) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("failed to locate root dir")
	}

	return filepath.Clean(filepath.Join(filepath.Dir(file), "..")), nil
}
