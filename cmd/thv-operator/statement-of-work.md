# Relevant references

The following statement of work depends on an understanding of multiple specifications, standards and tools.

## Specifications

* YAML specification v1.2.2 at https://yaml.org/spec/1.2.2
* OpenAPI Specification v3.1.0 at https://spec.openapis.org/oas/v3.1.0.html
* 'Uniform Resource Locators (URL)' Request For Comments at https://www.rfc-editor.org/rfc/rfc1738.txt 

# Persona

* You are a Senior Software Engineer.
* You have expertise in the areas of writing golang code and operators for kubernetes following the Operator pattern at https://kubernetes.io/docs/concepts/extend-kubernetes/operator.

# Preferred tools

* You work in a linux environment using the bash shell in the terminal when necessary.
* You prefer to utilize the buildah tool at https://github.com/containers/buildah to build OCI images.
* You prefer to utilize the podman tool at https://github.com/containers/podman to run OCI images.


# Other tools

* *The build system yaml utilizes go-task also known as task from https://taskfile.dev
* Currently this project uses google ko at https://github.com/ko-build/ko to build the thv-operator.


# Rules

* Think, plan, and reason step by step.
* Do not jump to conclusions. 
* Once the plan is complete review each of the generated steps for accuracy. If a given step is found to be inaccurate then fix that step and any subsequent steps that transitively require updates.
* When Additional fields may exist do not omit them for brevity.
* Add a note at the beginning of any file that is created by AI indicating that the file was created in whole or in part by AI using Cursor and the model used.

# Current thv operator build process

* The OCI Container image for the kubernetes operator in the cmd/thv-operator directory is currently built by executing go-task tasks from TaskFile.yml. 
* In the TaskFile.yml from the cmd/thv-operator directory the task that builds the operator is "build-operator-image".
* The "build-operator-image" task uses google ko to Build the operator image.

# What to do

* It is required to add a new task in the TaskFile.yml in the cmd/thv-operator directory entitled "build-operator-image-from-Dockerfile".
* This new task "build-operator-image-from-Dockerfile" builds the golang project source code in cmd/thv-operator for the thv operator inside a container image via a new Dockerfile in the containers/thv-operator directory.


## Create the Dockerfile for the thv-operator

* The `Dockerfile` reference documentation is at https://docs.docker.com/reference/dockerfile/. The available instructions that can be used in the `Dockerfile` are listed in the Overview on that web page url.

* Write a `Dockerfile` to the containers/thv-operator directory of the project. If it is found to exist, overwrite the existing `Dockerfile` in the containers/thv-operator directory of the project.
** This `Dockerfile` will be a multi-stage build as described at https://docs.docker.com/get-started/docker-concepts/building-images/multi-stage-builds/ and at https://docs.docker.com/guides/golang/build-images/ for golang specifically.
  * There are two stages in this multi-stage `Dockerfile`:
    * Build the command line utility from golang source in the first `build-stage`.
    * Deploy the command line utility binary into a leaner image for later execution in the `build-release-stage`.
  * When building a multi-stage golang based OCI image via a `Dockerfile`:
    * You prefer to use registry.access.redhat.com/ubi9/go-toolset:1.23.9 as the golang source code `build-stage` builder image.
    * You prefer to use registry.access.redhat.com/ubi9/ubi-micro:9.4 as the image to Deploy the command line utility binary into during the `build-release-stage`.
  * During the `build-stage` the thv-operator should be compiled via the instructions in the "build-operator" task in the TaskFile.yml in the cmd/thv-operator directory.  

* Example multi-stage `Dockerfile` that builds from golang source, runs the tests in the container, and then copies the command line utility binary into a leaner image for later execution.
```
# Build the manager binary
FROM registry.access.redhat.com/ubi9/go-toolset:1.24.4-1752083840 AS builder

ARG TARGETOS
ARG TARGETARCH

ENV GO_MODULE=github.com/arkmq-org/activemq-artemis-operator

### BEGIN REMOTE SOURCE
# Use the COPY instruction only inside the REMOTE SOURCE block
# Use the COPY instruction only to copy files to the container path $REMOTE_SOURCE_DIR/app
ARG REMOTE_SOURCE_DIR=/tmp/remote_source
RUN mkdir -p $REMOTE_SOURCE_DIR/app
WORKDIR $REMOTE_SOURCE_DIR/app
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/
COPY entrypoint/ entrypoint/
COPY pkg/ pkg/
COPY version/ version/
### END REMOTE SOURCE

# Set up the workdir
WORKDIR /opt/app-root/src
RUN cp -r $REMOTE_SOURCE_DIR/app/* .

# Build
# the GOARCH has not a default value to allow the binary be built according to the host where the command
# was called. For example, if we call make docker-build in a local env which has the Apple Silicon M1 SO
# the docker BUILDPLATFORM arg will be linux/arm64 when for Apple x86 it will be linux/amd64. Therefore,
# by leaving it empty we can ensure that the container and binary shipped on it will have the same platform.
# CGO_ENABLED is set to 1 for dynamic linking to OpenSSL to use FIPS validated cryptographic modules
# when is executed on nodes that are booted into FIPS mode.
RUN CGO_ENABLED=1 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -a -ldflags="-X '${GO_MODULE}/version.BuildTimestamp=`date '+%Y-%m-%dT%H:%M:%S'`'" -o manager main.go

FROM registry.access.redhat.com/ubi9-minimal:9.6-1752069876 AS base-env

ENV BROKER_NAME=activemq-artemis
ENV USER_UID=1000
ENV USER_NAME=${BROKER_NAME}-operator
ENV USER_HOME=/home/${USER_NAME}
ENV OPERATOR=${USER_HOME}/bin/${BROKER_NAME}-operator

WORKDIR /

# Create operator bin
RUN mkdir -p ${USER_HOME}/bin

# Copy the manager binary
COPY --from=builder /opt/app-root/src/manager ${OPERATOR}

# Copy the entrypoint script
COPY --from=builder /opt/app-root/src/entrypoint/entrypoint ${USER_HOME}/bin/entrypoint

# Set operator bin owner and permissions
RUN chown -R `id -u`:0 ${USER_HOME}/bin && chmod -R 755 ${USER_HOME}/bin

# Upgrade packages
RUN microdnf update -y --setopt=install_weak_deps=0 && rm -rf /var/cache/yum

USER ${USER_UID}
ENTRYPOINT ["${USER_HOME}/bin/entrypoint"]

LABEL name="arkmq-org/activemq-artemis-operator"
LABEL description="ActiveMQ Artemis Broker Operator"
LABEL maintainer="Roddie Kieley <rkieley@redhat.com>"
LABEL version="2.0.3"
```
* For the build-stage in the `Dockerfile` utilize the registry.access.redhat.com/ubi9/go-toolset OCI image. Use the golang version that most closely matches that used in the Example multi-stage `Dockerfile`.

### Build the `Dockerfile` for the thv-operator into the OCI image

* Use the instructions in Using Containerfiles/Dockerfiles with Buildah at https://github.com/containers/buildah/blob/main/docs/tutorials/01-intro.md#using-containerfilesdockerfiles-with-buildah to build the `Dockerfile` from the previous 'Create the Dockerfile` section.
  * The name to be given to the created image is 'localhost/thv-operator' with tag `latest`

