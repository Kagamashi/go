Go modules are standard way to manage dependencies in Go projects.
from version 1.14 modules replaced older GOPATH-based dependency management system 

Navigate to project directory and run:
go mod init [module]
this creates go.mod file, which defines module's name and dependencies


go.mod file keeps track of modules dependencies and contains:
- module name (usually repository URL for shared packages)
- Go version used for project
- dependency requirements

module github.com/user/mymodule

go 1.20

require (
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0
)




go mod init

go.mod 


go.sum 


versioning, upgrading depndencies



go get -u 



go mod tidy 


