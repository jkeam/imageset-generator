# Imageset Generator

## Process

### Mandatory

1. Get specific version for overall version

    ```shell
    oc-mirror list releases --version=4.17
    ```

2. Download release and required operators from `https://mirror.openshift.com/pub/openshift-v4/clients/ocp/{ocp-version}/release.txt`

    ```shell
    wget https://mirror.openshift.com/pub/openshift-v4/clients/ocp/4.17.1/release.txt
    ```

### Operators

1. Get all catalogs, notice no z, eg) 4.17 instead of 4.17.1

    ```shell
    oc-mirror list operators --catalogs --version=4.17
    ```

2. Get operators from a catalog

    ```shell
    oc-mirror list operators --catalog=registry.redhat.io/redhat/redhat-operator-index:v4.17
    ```

### Imageset Config

```yaml
apiVersion: mirror.openshift.io/v1alpha2
kind: ImageSetConfiguration
archiveSize: 1
storageConfig:
  local:
    path: /home/user/workspace
```

```yaml
apiVersion: mirror.openshift.io/v1alpha2
kind: ImageSetConfiguration
storageConfig:
  registry:
    imageURL: localhost:5000/metadata:latest
    skipTLS: true
```

## Out of Scope

1. Anything other than linux/amd64

## Docs
1. [Release](https://mirror.openshift.com/pub/openshift-v4/clients/ocp/4.16.18/release.txt)
2. [Unfortunately, that initial release of this feature in 4.15.0 and the above-mentioned versions had a critical flaw (ref OCPBUGS-33305) which wasn't resolved until the 4.15.13 release of oc-mirror.](https://access.redhat.com/solutions/7013461)
3. [Docs](https://github.com/openshift/oc-mirror?tab=readme-ov-file#building-the-imageset-config)
