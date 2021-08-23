# Release process

## panopticon

1. Make sure `panopticon` docker image has been tagged following semver.

2. Tag `panopticon` installer chart

```bash
make chart-panopticon CHART_VERSION=vYYYY.MM.DD APP_VERSION=${panopticon_tag}
```

3. Tag this repo with vYYYY.MM.DD and push to Github.
