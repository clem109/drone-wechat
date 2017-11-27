# drone-wechat

Drone plugin to send build status notifications via WeChat for Work. For usage
information please look at [the docs](DOCS.md).

## Build

Build the binary with the following commands:

```
go build
```

## CHANGE BELOW

## Docker

Build the Docker image with the following commands:

```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -o release/linux/amd64/drone-webhook
docker build --rm -t plugins/webhook .
```

### Usage

```
docker run --rm \
  -e PLUGIN_URLS=https://bbc.com/... \
  -e PLUGIN_HEADERS="HEADER1=value1" \
  -e PLUGIN_USERNAME=drone \
  -e PLUGIN_PASSWORD=password \
  -e DRONE_REPO_OWNER=octocat \
  -e DRONE_REPO_NAME=hello-world \
  -e DRONE_COMMIT_SHA=7fd1a60b01f91b314f59955a4e4d4e80d8edf11d \
  -e DRONE_COMMIT_BRANCH=master \
  -e DRONE_COMMIT_AUTHOR=octocat \
  -e DRONE_BUILD_NUMBER=1 \
  -e DRONE_BUILD_STATUS=success \
  -e DRONE_BUILD_LINK=http://github.com/octocat/hello-world \
  -e DRONE_TAG=1.0.0 \
  plugins/webhook
```
