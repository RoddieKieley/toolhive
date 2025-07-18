# This Dockerfile was created in whole or in part by AI using Cursor and the o3 model.
# Multi-stage build for the ToolHive Kubernetes operator (thv-operator)

# ---------------------------
# Build stage
# ---------------------------
FROM registry.access.redhat.com/ubi9/go-toolset:1.24.4-1752083840 AS builder

ARG TARGETOS
ARG TARGETARCH

# Enable Go modules and set build environment
ENV CGO_ENABLED=1 \
    GO111MODULE=on \
    GOOS=${TARGETOS:-linux} \
    GOARCH=${TARGETARCH:-amd64} \
    WORK_DIR=/opt/app-root/src

WORKDIR ${WORK_DIR}

# Copy Go module manifests first to leverage layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source tree
COPY . .

# Build the operator binary
RUN mkdir -p /opt/app-root/src/bin && go build -buildvcs=false -ldflags="-s -w" -o /opt/app-root/src/manager ./cmd/thv-operator

# ---------------------------
# Runtime stage
# ---------------------------
FROM registry.access.redhat.com/ubi9/ubi-micro:9.4 AS runtime

LABEL name="stacklok/toolhive-operator"
LABEL description="ToolHive Kubernetes operator"
LABEL source="https://github.com/stacklok/toolhive"

# Environment and user setup
ENV USER_UID=1000
ENV USER_NAME=thv-operator
ENV USER_HOME=/home/${USER_NAME}
ENV OPERATOR=${USER_HOME}/bin/thv-operator

WORKDIR /

# Create operator bin
RUN mkdir -p ${USER_HOME}/bin

# Copy the manager binary
COPY --from=builder /opt/app-root/src/manager ${OPERATOR}

# Copy the entrypoint script
COPY --from=builder /opt/app-root/src/cmd/thv-operator/entrypoint ${USER_HOME}/bin/entrypoint

# Set operator bin owner and permissions
RUN chown -R `id -u`:0 ${USER_HOME}/bin && chmod -R 755 ${USER_HOME}/bin

USER ${USER_UID}
ENTRYPOINT ["/home/thv-operator/bin/entrypoint"]
