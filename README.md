
![License](https://img.shields.io/github/license/Ileriayo/markdown-badges?style=flat-square)
![Go](https://img.shields.io/badge/go-1.25-%2300ADD8.svg?style=flat-square&logo=go&logoColor=white)
![gRPC](https://img.shields.io/badge/gRPC-1.73-blue?style=flat-square&logo=go&logoColor=white)
![gin](https://img.shields.io/badge/Gin-1.10-00C397?style=flat-square&logo=go&logoColor=white)
![zerolog](https://img.shields.io/badge/zerolog-f33?style=flat-square)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=flat-square&logo=docker&logoColor=white)

>[!tldr] Abstract
> The project was written for educational purposes.
> Windows platform, vscode editor. Support for other operating systems is not guaranteed.

# 📦FileProcessor gRPC `🛠 Practice`
## 📚Table of contents
🚧🚧🚧

## 🎯 Goals
The project is focused on practicing with gRPC in Golang:
### 📋 TODO
- [ ] Client:
  - [x] Accepts a list of file paths via HTTP API
  - [x] Send paths to gRPC server
  - [ ] Supports file uploads to gRPC
  - [x] Cancel by timeout `context.WithTimeout`
  - [ ] Middlewares:
    - [ ] Logging method, duration and errors
    - [x] Panic processing
    - [ ] Rate limiting
- [x] Server:
  - [x] Panic middleware
  - [x] Add `.proto` in gRPC-service
    - ProcessFiles(FileList) returns (stream FileResult)
    - Using Server-side streaming
  - [x] build proto files
  - [x] Accepts a list of file paths via gRPC
  - [x] Processes files in parallel
  - [x] Counting lines
  - [x] Counting words
  - [x] Limit: no more than 5 files can be processed simultaneously
  - [x] Fan in processing files
    - [x] Fan-out: dispatch processing to goroutines
    - [x] Fan-in: collect results and send to client
  - [x] Stream results back to the client
  - [x] Support context cancellation on request
  - [x] Stores processing history
  - [ ] Supports file uploads to store
- [ ] Lifecycle organization:
  -  [x] Graceful shutdown
  -  [ ] Healthcheck
- [ ] Testing
  - [ ] Unit testing:
    - [ ] Counting lines
    - [ ] Counting words 
    - [ ] Processing raise a panic 
  - [ ] Integrated testing:
    - [ ] Client: send path -> Server: calculate -> Client: return result
- [ ] Deployment
  - [x] Create dockerfiles
  - [ ] Create docker-compose
  - [ ] Write ReadMe
  - [x] Add Taskfile.yml in project

- ⏳ Backlog
  - Authorization via gRPC metadata

## 🛠️ Features
- 🔗 Accepts a list of file paths (or virtual file names)
- ⚙️ Processes files in parallel (counting lines and words)
- 🔄 Streams results back to the client
- 🧩 Using middleware (interceptor, rate-limit, etc.)
- ⏹️ Support context cancellation
- 🛡️ Authorization via gRPC metadata
- 🕓 Stores processing history
- 📤 Supports file uploads

## 📐 Architecture
🚧🚧🚧

# DEBUG
🚧🚧🚧
## Taskfile
```sh
task install
```


## Environment Variables
[Client](./doc/client-env.md#config),
[File Server](./doc/fileservice-env.md#config)


# TODO
- https://protobuf.dev/installation/