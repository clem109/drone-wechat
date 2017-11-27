FROM plugins/base:multiarch
MAINTAINER Clement Venard <cvenard@gmail.com>

LABEL org.label-schema.version=latest
LABEL org.label-schema.vcs-url="https://github.com/clem109/drone-wechat.git"
LABEL org.label-schema.name="Drone Wechat"
LABEL org.label-schema.schema-version="1.0"

ADD release/linux/amd64/drone-wechat /bin/
ENTRYPOINT ["/bin/drone-wechat"]
