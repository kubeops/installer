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

package v1alpha1_test

import (
	"os"
	"testing"

	"kubeops.dev/installer/apis/installer/v1alpha1"

	schemachecker "kmodules.xyz/schema-checker"
)

func TestDefaultValues(t *testing.T) {
	checker := schemachecker.New(os.DirFS("../../.."),
		schemachecker.TestCase{Obj: v1alpha1.AceUserRolesSpec{}},
		schemachecker.TestCase{Obj: v1alpha1.ClusterConnectorSpec{}},
		schemachecker.TestCase{Obj: v1alpha1.ConfigSyncerSpec{}},
		schemachecker.TestCase{Obj: v1alpha1.ExternalDnsOperatorSpec{}},
		schemachecker.TestCase{Obj: v1alpha1.FalcoUiServerSpec{}},
		schemachecker.TestCase{Obj: v1alpha1.GatekeeperGrafanaDashboardsSpec{}},
		schemachecker.TestCase{Obj: v1alpha1.GatekeeperLibrarySpec{}},
		schemachecker.TestCase{Obj: v1alpha1.KubeUiServerSpec{}},
		schemachecker.TestCase{Obj: v1alpha1.OperatorShardManagerSpec{}},
		schemachecker.TestCase{Obj: v1alpha1.PanopticonSpec{}},
		schemachecker.TestCase{Obj: v1alpha1.PetsetSpec{}},
		schemachecker.TestCase{Obj: v1alpha1.ScannerSpec{}},
		schemachecker.TestCase{Obj: v1alpha1.SidekickSpec{}},
		schemachecker.TestCase{Obj: v1alpha1.SupervisorSpec{}},
	)
	checker.TestAll(t)
}
