# Kubernetes ConfigMap Reload

[![license](https://img.shields.io/github/license/jimmidyson/configmap-reload.svg?maxAge=2592000)](https://github.com/jimmidyson/configmap-reload)

**configmap-reload** is a simple binary to trigger a reload when Kubernetes ConfigMaps are updated.
It watches mounted volume dirs and notifies the target process that the config map has been changed.
I changed the original code so it works with process names instead of making HTTP requests.

It is available as a Docker image at https://hub.docker.com/r/viniciusramosdefaria/configmap-reload

### Usage

```
Usage of ./configmap-reload:
  -volume-dir value
        the config map volume directory to watch for updates; may be used multiple times
  -process-path string
        the process path to send SIGHUP to when the specified config map volume directory has been updated
```

### License

This project is [Apache Licensed](LICENSE.txt)

