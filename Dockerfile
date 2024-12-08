FROM registry.access.redhat.com/ubi9/ubi-micro:9.5-1733126338

WORKDIR /go/src/app
COPY ./imageset-generator .

USER nobody
EXPOSE 8000
CMD ["./imageset-generator"]
