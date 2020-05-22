REGISTRY = "viniciusramosdefaria"
RELEASE = "latest"
BINARY_NAME=configmap-reload

#go cmds
GOCMD=go
GOBUILD=$(GOCMD) build
CMD_PATH=./configmap-reload.go

#docker cmds
DOCKERCMD=docker
DOCKERBUILD=${DOCKERCMD} build
DOCKERPUSH=${DOCKERCMD} push

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GOBUILD} -o ./${BINARY_NAME} -v ${CMD_PATH}

publish: build-container
	${DOCKERPUSH} "${REGISTRY}/${BINARY_NAME}:${RELEASE}"

build-container: build
	$(DOCKERBUILD) -t "${REGISTRY}/${BINARY_NAME}:${RELEASE}" .

clean:
	rm ./${BINARY_NAME}