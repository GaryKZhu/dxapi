# This how we want to name the binary output
BINARY=doctorx-apiserver

# These are the values we want to pass for VERSION and BUILD
# git tag 1.0.1
# git commit -am "One more change after the tags"
##VERSION=`git describe --tags`
VERSION='1.0.0'
BUILD=`date +%FT%T%z`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-w -s -X doctorx/pkg/setting.Version=${VERSION} -X doctorx/pkg/setting.Build=${BUILD}"

# Builds the project
build:
	env GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY}

# Installs the project: copies binaries
install:
	go install ${LDFLAGS}

# Cleans the project: deletes binaries
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean install
