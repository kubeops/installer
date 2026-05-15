# AGENTS.md

This file provides guidance to coding agents (e.g. Claude Code, claude.ai/code) when working with code in this repository.

## Repository purpose

Go module `kubeops.dev/installer` — the umbrella Helm chart repo for the `kubeops.dev` family of operators. Bundles charts for every kubeops project (panopticon, supervisor, scanner, sidekick, fargocd, external-dns-operator, cluster-connector, falco-ui-server, petset, taskqueue, pgoutbox, etc.) plus shared CRD bundles (cert-manager-crds, gatekeeper-library, prometheus-operator-crds, application-crds, kmodules-crds) and AppsCode-curated dashboards (gatekeeper, opencost). Also exposes a Go API package (`apis/installer/v1alpha1/`) describing chart values.

Charts shipped (under `charts/`):
- **kubeops operators**: `cluster-connector`, `external-dns-operator`, `falco-ui-server`, `fargocd`, `kube-ui-server`, `operator-shard-manager`, `panopticon`, `petset`, `pgoutbox`, `scanner`, `sidekick`, `supervisor`, `taskqueue`.
- **CSI driver**: `cert-manager-csi-driver-cacerts`.
- **CRD bundles**: `application-crds`, `cert-manager-crds`, `kmodules-crds`, `prometheus-operator-crds`.
- **OPA Gatekeeper**: `gatekeeper-library`, `gatekeeper-grafana-dashboards`.
- **Misc**: `ace-user-roles`, `config-syncer`, `opencost-grafana-dashboards`.

## Architecture

- `charts/` — one subdirectory per Helm chart, standard layout (`Chart.yaml`, `values.yaml`, `templates/`, generated `doc.yaml`/`README.md`/`values.openapiv3_schema.yaml`).
- `apis/installer/v1alpha1/` — Go types backing chart values; single API group `installer:v1alpha1`.
- `catalog/` — image catalog (`imagelist.yaml` + generated scripts and CVE report).
- `hack/scripts/` — codegen/release helpers (`update-catalog.sh`, `update-chart-dependencies.sh`, `import-crds.sh`, `update-local-repo.sh`, `ct.sh`, etc.).
- `stable/` — Helm repository index for the stable channel.
- `tests/` — chart-testing configuration.
- `lintconf.yaml` — YAML lint config.
- `RELEASE.md` — release process notes.
- `vendor/` — Go vendored deps.

## Common commands

All Make targets run inside `ghcr.io/appscode/golang-dev` — Docker must be running.

- `make gen` — full regen (clientset + manifests + values schemas + chart docs). Run after any change to `apis/installer/v1alpha1/*_types.go`.
- `make manifests` — `gen-crds gen-values-schema gen-chart-doc`.
- `make gen-values-schema` — regenerate `values.openapiv3_schema.yaml`.
- `make gen-chart-doc` — regenerate per-chart `README.md`.
- `make update-charts` — refresh chart-level metadata.
- `make update-local-repo` — refresh local helm repo (`stable/`) metadata.
- `make fmt`, `make lint`, `make unit-tests` / `make test` — standard.
- `make ct` — chart-testing lint+test.
- `make verify` — `verify-modules`; `go mod tidy && go mod vendor` must leave the tree clean.
- `make add-license` / `make check-license` — manage license headers.

Auxiliary helpers (invoked outside Make):

- `./hack/scripts/update-catalog.sh` — regenerate `catalog/` from `imagelist.yaml` via `kmodules.xyz/image-packer`.
- `./hack/scripts/import-crds.sh` — pull CRDs from sibling kubeops repos into each chart's `crds/`.
- `./hack/scripts/update-chart-dependencies.sh` — refresh `Chart.lock` / subchart pins.

Run a single Go test (requires a local Go toolchain):

```
go test ./apis/installer/v1alpha1/... -run TestName -v
```

## Conventions

- Module path is `kubeops.dev/installer` (vanity URL). Imports must use that.
- Edit `apis/installer/v1alpha1/*_types.go` to change a chart's values surface, then run `make gen` so the generated clientset, `values.openapiv3_schema.yaml`, and per-chart `README.md` stay in sync. Do not hand-edit `zz_generated.*.go`, generated chart `README.md` files, `values.openapiv3_schema.yaml`, or anything under `catalog/` except `imagelist.yaml`.
- License: `LICENSE.md`; use `make add-license` to apply headers to new Go files.
- Sign off commits (`git commit -s`); contributions follow the DCO (`DCO`).
- Vendor directory is checked in — `go mod tidy && go mod vendor` must leave the tree clean (enforced by `verify-modules`).
- CRDs for operator charts are **imported** from the operator's repo via `import-crds.sh`; do not hand-author `charts/*/crds/*.yaml`. The CRD-bundle charts (`application-crds`, `cert-manager-crds`, `prometheus-operator-crds`, etc.) import upstream third-party CRDs the same way.
- Adding a new chart: create `charts/<name>/` following the existing layout, add a values-schema type under `apis/installer/v1alpha1/`, list it in `imagelist.yaml` if it ships images, then run `make gen`.
- The `stable/` directory holds the published Helm index; `make update-local-repo` regenerates it. Do not hand-edit the index.
