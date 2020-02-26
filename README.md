# **Web Development using Golang**

## Pre-requisites:

- Work experience in any of the programming language (nice if it's Java, .Net or any other statically typed language) for web development.
- [Go Tour](https://tour.golang.org).
- [Go Playground](https://play.golang.org).


## Setup

- sudo snap install go --classic
- **go version**
- Reference: https://www.callicoder.com/golang-installation-setup-gopath-workspace/

**Go Commands**

- **go run main.go** [anotherfile.go]
- **go build** (To make it as executable file)
- **go install** (To run it from anywhere)
- **go help** [any command you need a help]

**Adding Dependencies for a project**

- **go get "github.com/narensundaram.dev/photo-blog"**
- **cd ~/projects/go/src/photo-blog**
- **go mod init**
- **cat go.mod**
- **go test**
- **cat go.mod** (To see all the dependencies noted down)
- **cat go.sum** (To see all the dependencies of dependencies noted down)
- **go mod download** (To download all the dependencies)

## Contents:

- **[Quick Introduction about Go](docs/1-quick-intro.md)**
- **[Server and MVC Architecture](docs/2-server-mvc-architecture.md)**
	- [Sqlite Database Connection (Model)](docs/2.1-model.md)
	- [Templates (View)](docs/2.2-view.md)
	- [Routers and Handlers (Controller)](docs/2.3-controller.md)
- **[Error Handling](docs/3-error-handling.md)**
- **[Middlewares](docs/4-middlewares.md)**
- **[File Handling & Serving Files](docs/5-file-handling-and-serving-files)**
- **[Sorting and Pagination](docs/6-sorting&pagination.md)**
- **[Responding JSON Data](docs/7-web-service-REST-APIs.md)** (Echo REST API)
