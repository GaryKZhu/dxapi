## Medi Server

### Quick Start
* go mod init doctorx
* go get -u github.com/gin-gonic/gin;go get -v golang.org/x/tools/gopls
* mkdir routers public pkg models docs conf
* cp the source e.g`cp -rf routers public pkg models docs conf main.go ../doctorx`
* modify the package to doctorx e.g. `find . -type f -exec sed -i -e 's/doctorx/<new project name>/g' {} \;`
* run the project `go run main.go`
* build the project `make`

### token generator 
* base64 (two ways) - `echo "doctorx-gary" | base64"` , `echo "ZG9jdG9yeC1nYXJ5Cg==" | base64 -d`
* md5 (one way) - `md5 -s "doctorx-gary"`

### Project Structure:
.
├── Makefile              ## make file
├── README.md             ## README 
├── conf
│   └── app.conf          ## app config file
├── docs                  ## swagger documents
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod                ## go module
├── go.sum                ## go summary
├── main.go               ## main 
├── pkg                   ## package folder
│   ├── jwt
│   │   └── jwt.go        ## token control
│   └── setting
│       └── setting.go    ## setting management to process app.conf
├── public                ## web folder
│   ├── build
│   └── html
│       ├── 404.html      
│       └── index.html
└── routers
    ├── api
    │   └── v1
    │       └── ping.go    ## ping 
    └── routers.go         ## routing management

