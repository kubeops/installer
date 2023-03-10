[![Slack](https://shields.io/badge/Join_Slack-salck?color=4A154B&logo=slack)](https://slack.appscode.com)
[![Twitter](https://img.shields.io/twitter/follow/kubeops.svg?style=social&logo=twitter&label=Follow)](https://twitter.com/intent/follow?screen_name=Kubeops)

# Kubeops Installer Charts

Helm charts for Kubeops components.

## Local Development

To use charts during local development, use the charts from `stable` directory. If you want to use your own branch, replace `master` with your branch name.
```bash
helm repo add kubeops-dev https://raw.githubusercontent.com/kubeops/installer/master/stable
helm repo update
```

To update the charts of your branch, run:

```bash
make update-chart-repo
```

## Support

To speak with us, please leave a message on [our website](https://appscode.com/contact/).

To receive product announcements, follow us on [Twitter](https://twitter.com/Kubeops).
