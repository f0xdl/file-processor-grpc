
![License](https://img.shields.io/github/license/Ileriayo/markdown-badges?style=flat-square)
![Go](https://img.shields.io/badge/go-1.25-%2300ADD8.svg?style=flat-square&logo=go&logoColor=white)
![gRPC](https://img.shields.io/badge/gRPC-1.73-blue?style=flat-square&logo=go&logoColor=white)
![gin](https://img.shields.io/badge/Gin-1.10-00C397?style=flat-square&logo=go&logoColor=white)
![zerolog](https://img.shields.io/badge/zerolog-f33?style=flat-square)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=flat-square&logo=docker&logoColor=white)

# 📦FileProcessor gRPC `🛠 Practice`

> [!note]
> The project was written for educational purposes.
> 🚧🚧🚧

## 📚Table of contents
- Features
- Todo

## 🛠️ Features
+ 🔗 Accepts a list of virtual file names
+ ⚙️ Processes files in parallel (counting lines and words)
+ 🔄 Streams results back to the client
+ ⏹️ Support context cancellation
+ 🕓 Stores processing history
+ 📤 Supports file uploads
- 🧩 Using middleware (rate-limit and logger)
- 🛡️ Authorization via gRPC metadata

##  📋 TODO
The project is focused on practicing with gRPC in Golang:
### Http Client
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
  - [x] Supports file uploads to store
- [x] Lifecycle organization:
  -  [x] Graceful shutdown
  -  [x] Healthcheck
- [ ] Testing
  - [ ] Unit testing:
    - [ ] Counting lines
    - [ ] Counting words 
    - [ ] Processing raise a panic 
  - [ ] Integrated testing:
    - [ ] Client: send path -> Server: calculate -> Client: return result
- [ ] Deployment
  - [x] Create dockerfiles
  - [x] Create docker-compose
  - [ ] Write ReadMe
  - [x] Add Taskfile.yml in project



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