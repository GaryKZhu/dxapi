## DXAPI Server

### Quick Start
* go mod init doctorx
* go get -u github.com/gin-gonic/gin;go get -v golang.org/x/tools/gopls
* mkdir routers public pkg models docs conf
* cp the source e.g`cp -rf routers public pkg models docs conf main.go ../doctorx`
* modify the package to doctorx e.g. `find . -type f -exec sed -i -e 's/doctorx/<new project name>/g' {} \;`
* run the project `go run main.go`
* build the project `make`
