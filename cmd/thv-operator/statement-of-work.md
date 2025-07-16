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

# Rules

* Think, plan, and reason step by step.
* Do not jump to conclusions. 
* Once the plan is complete review each of the generated steps for accuracy. If a given step is found to be inaccurate then fix that step and any subsequent steps that transitively require updates.
* When Additional fields may exist do not omit them for brevity.
* Add a note at the beginning of any file that is created by AI indicating that the file was created in whole or in part by AI using Cursor and the model used.

# Data provided

* The data provided is in the YAML 1.2 data language as per the definition of the YAML specification v1.2.2
* The data conforms to the OpenAPI Specification v3.1.0
* When the data contains the `url` or `URL` identifiers the meaning is defined as per 'Uniform Resource Locators (URL)'
* The OpenAPI Specification v3.1.0 compliant data structure is specified in the `customresourcedefinition-mcpservers.mcp.opendatahub.io.yaml` file in the `ref` directory in this project.
* In the Example provided YAML data that follows there will be multiple entries in the Example provided YAML data in the form of a numbered list, each of which begins with `spec:`.
* Example provided YAML data is as follows:
1.
```
spec:
  server_detail:
    description: Awesome MCP Servers - A curated list of Model Context Protocol servers
    id: 0007544a-3948-4934-866b-b4a96fe53b55
    name: io.github.appcypher/awesome-mcp-servers
    packages:
      - name: appcypher/awesome-mcp-servers
        registry_name: unknown
        version: ''
    repository:
      id: '895801050'
      source: github
      url: 'https://github.com/appcypher/awesome-mcp-servers'
    version_detail:
      is_latest: true
      release_date: '2025-05-16T19:16:40Z'
      version: 0.0.1-seed
```
2.  
```
spec:
  server_detail:
    description: ''
    id: 010904ec-6e39-4bdc-878a-75a6e79d0500
    name: io.github.kyrietangsheng/mcp-server-nationalparks
    packages:
      - environment_variables:
          - description: YOUR_NPS_API_KEY
            name: NPS_API_KEY
        name: mcp-server-nationalparks
        registry_name: npm
        version: 1.0.0
    repository:
      id: '951713109'
      source: github
      url: 'https://github.com/KyrieTangSheng/mcp-server-nationalparks'
    version_detail:
      is_latest: true
      release_date: '2025-05-16T19:11:05Z'
      version: 0.0.1-seed
```

# What to do

* Write each of the provided Example provided YAML data entries from the numbered list to a file `example-data-#.yaml` where the `#` is replaced with the numbered list number itself. Write these files with numbered filenames in the `data` directory within the project. If one of the numbered list number files is found to already exist in the project then overwrite the existing file with the Example provided YAML data from the numbered list entry previously provided.

## Create the command line utility

### Process the data into golang types

* From the `customresourcedefinition-mcpservers.mcp.opendatahub.io.yaml` file in the `ref` directory in this project the data structure in the `v1alpha1` named version in the `spec:` section of the `openAPIV3Schema` schema of `server_detail` is to be processed into golang struct types.
* Analyse the `customresourcedefinition-mcpservers.mcp.opendatahub.io.yaml` file in the `ref` directory in this project, creating matching valid golang structs with json: annotations and yaml: annotations.
  * In particular the command line utility needs to read and process the `server_detail` object data structure. Ensure all properties of the `server_detail` object have matching valid golang struct types created.
  * Further more any subsequent related golang struct types that are required should also be used directly from https://github.com/RHEcosystemAppEng/mcp-catalog-operator/blob/main/api/v1alpha1/mcpserver_types.go where possible.
  * Write these required golang struct types to a file `mcpserver_types.go` in the `types` directory.
  * Example golang struct type with json: annotation for the top level `spec:` and child `server_detail:` sections from the Example provided YAML data previously provided:
```
// McpServerSpec defines the desired state of McpServer.
type McpServerSpec struct {
	ServerDetail ServerDetail `json:"server_detail"`
}

// ServerDetail represents detailed server information as defined in the spec
type ServerDetail struct {
	Server   `json:",inline"`
	Packages []Package `json:"packages,omitempty"`
	Remotes  []Remote  `json:"remotes,omitempty"`
}
```
  * Where the golang struct types differ between the Example golang struct type and those in the https://github.com/RHEcosystemAppEng/mcp-catalog-operator/blob/main/api/v1alpha1/mcpserver_types.go file differ, prefer those in the mcpserver_types.go file where possible.


### Write the golang source code for the command line utility

* Write a command line utility in golang that utilizes the valid golang structs previously written to the `mcpserver_types.go` file in the `types` directory.
* This command line utility in golang must:
  * Examine the `data` directory to find all `*.yaml` files.
  * For each of the `*.yaml` files found in the `data` directory, read and parse the YAML data from the individual file into the valid golang structs found in the `mcpserver_types.go` file.
  * When the read and parse of an individual file is complete write a message to stdout indicating success or failure for that file as well as a newline.
  * Evaluate each of the read and parsed golang struct instances representing server_detail. Check to see if the server_detail repository source has a value of `github`. If the server_detail repository source is `github` then for that server_detail instance check the server_detail repository url field. If the server_detail repository url field contains a valid Uniform Resource Locator it will refer to a `git` repository. Check this URL's git repository file list for a valid `Dockerfile`. If the git repository at this URL's has a valid `Dockerfile` print the server_detail name with " has Dockerfile" appended to stdout.


### Write the golang test source code for the command line utility

* You prefer to write the golang test source code for the command line utility using the ginkgo Modern Testing Framework from https://github.com/onsi/ginkgo and the gomega Preferred Matcher Library from https://github.com/onsi/gomega. The ginkgo Modern Testing Framework focuses on Behaviour-driven development as documented at https://en.wikipedia.org/wiki/Behavior-driven_development
* Write matching tests for this golang command line utility using the ginkgo Modern Testing Framework and gomega Preferred Matcher Library using a Behaviour-driven development approach.
  * In particular ensure the tests cover the functionality reading, parsing, and marshalling the YAML into golang struct type instances from `mcpserver_types.go` by using the original Example provided YAML data files from the `example-data-#.yaml` files in the `data` directory.
  * A test must be provided that takes the each of the Example provided YAML data files of the name `example-data-#.yaml` in the `data` directory in turn, load it successfully into the golang struct types in `mcpserver_types.go`, then writes the data back to a new file named `example-data-rewritten-#.yaml` in the `data` directory in correct YAML format such that it is identical and compares successfully against the `example-data-#.yaml` in the `data` directory.
* Execute the tests created using the `go test ./...` shell command. If all tests are Passed then proceed to the `Build the OCI image` section.

## Build the OCI image for the command line utility

* This command line utility should build correctly via the `go build ./...` shell command. If not return to the `Create the command line utility` section.

### Create the Dockerfile for the command line utility

* The `Dockerfile` reference documentation is at https://docs.docker.com/reference/dockerfile/. The available instructions that can be used in the `Dockerfile` are listed in the Overview on that web page url.

* Write a `Dockerfile` to the root directory of the project. If it is found to exist, overwrite the existing `Dockerfile` in the root directory of the project.
** This `Dockerfile` will be a multi-stage build as described at https://docs.docker.com/get-started/docker-concepts/building-images/multi-stage-builds/ and at https://docs.docker.com/guides/golang/build-images/ for golang specifically.
  * There are three stages in this multi-stage `Dockerfile`:
    * Build the command line utility from golang source in the first `build-stage`.
    * Run the tests in the container in the second `run-test-stage`.
    * Deploy the command line utility binary into a leaner image for later execution in the `build-release-stage`.
  * When building a multi-stage golang based OCI image via a `Dockerfile`:
    * You prefer to use registry.access.redhat.com/ubi9/go-toolset:1.23.9 as the golang source code `build-stage` builder image.
    * You prefer to use registry.access.redhat.com/ubi9/ubi-micro:9.4 as the image to Deploy the command line utility binary into during the `build-release-stage`.
  * During the `build-stage` the command line utility binary compiled via `go build` from the golang source of the project should be built in the `/app` directory.  

* Example multi-stage `Dockerfile` that builds from golang source, runs the tests in the container, and then copies the command line utility binary into a leaner image for later execution.
```
# syntax=docker/dockerfile:1

# Build the command line utility from golang source
FROM registry.access.redhat.com/ubi9/go-toolset:1.23.9 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the command line utility binary into a leaner image for later execution
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /docker-gs-ping /docker-gs-ping

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/docker-gs-ping"]
```
* For the build-stage in the `Dockerfile` utilize the registry.access.redhat.com/ubi9/go-toolset OCI image. Use the golang version that most closely matches that used in the Example multi-stage `Dockerfile`.

### Build the `Dockerfile` for the command line utility into the OCI image

* Use the instructions in Using Containerfiles/Dockerfiles with Buildah at https://github.com/containers/buildah/blob/main/docs/tutorials/01-intro.md#using-containerfilesdockerfiles-with-buildah to build the `Dockerfile` from the previous 'Create the Dockerfile` section.
  * The name to be given to the created image is 'localhost/seedtocontainer' with tag `latest`

### Test the execution of the OCI image built from the `Dockerfile` using the following command:

go test ./...                         # confirm green
buildah bud -t localhost/seedtocontainer:latest .   # build image
podman run --rm localhost/seedtocontainer:latest    # parse YAML(s)
