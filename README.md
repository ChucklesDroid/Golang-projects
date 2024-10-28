# Basics

## Basic structure of project and file details 

* Every go file within the directory has the package name of the directory above it.
Eg-

--cmd
    --server
        --main.go
    --todocli
        --main.go
--db
    --store.go
--html
    --template.go
--web
    --middleware.go
    --server.go
--go.mod
--go.sum

In this example, cmd, server, todocli, db, html and web are directories and in turn called packages. So all the files within these directories will have the package name of the directory above it.

* Module is a collection of packages that live within the same project. A project typically contains 1 module with many packages.

* If the project has `go.mod` and `go.sum` file, this indicates that this is a go module. `go.mod` contains information about the name of the module, go version and its dependencies along with its version names.

* `go.sum` holds information about dependencies, their version along with the hash for each dependency which contains all file contents hashed into a single sum. The hash value is used to verify the contents the dependency

## Initialising a new project

1. $go mod init "<name of the module>"
 - The name of the module should be unique.
 - Others should be able to clone it for eg using `go get` command.

 Eg:- github.com/ChucklesDroid/test-project  //github.com can also be gitlab.com, bitbucket.com and so on.

2. There is no package distributor in go world like npm or packages for node and PHP. Packages and modules live in a distributed world in golang so there is no single point of failure.

3. You can change the name of module later on, by editing the `go.mod` file.

## Grabbing a dependency

1. $go get github.com/gorilla/mux
 - Once this command is executed within the project directory(inside a go module), it will automatically update `go.mod` and `go.sum` files with dependency information. (It can also create `go.sum` file if it doesn't already exist.) 

2. All the packages we get will be saved in `~/go/pkg/mod` directory. the modules in this directory are read-only. They are cached versions of external code that Go downloads when you run command like `go get`, `go mod tidy`, `go build`, `go install`.

3. `~/go/src` contains your source code.

## Download go dependencies in local cache- $go mod download
## Download missing dependencies and remove unused modules- $go mod tidy
## These commands can be checked out- $go mod
## Download, build and install binaries- $go install <github.com/example-project> 
