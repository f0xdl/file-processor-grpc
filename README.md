
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![License](https://img.shields.io/github/license/Ileriayo/markdown-badges?style=for-the-badge)
![gRPC]()
![gin]()


# 📦FileProcessor gRPC `🛠 Practice`
## 📚Table of contents
🚧🚧🚧

## 🎯 Goals
The project is focused on practicing with gRPC in Golang:
### 📋 TODO
- [ ] Client:
  - [x] Accepts a list of file paths via HTTP API
  - [ ] Send paths to gRPC server
  - [ ] Cancel by timeout `context.WithTimeout`
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

- ⏳ Backlog
  - Authorization via gRPC metadata
  - Supports file uploads

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
## Makefile
Для использования команд make на Windows необходимо установить пакет:
```shell
choco install make
```

## Environment Variables
[Client](./doc/client-env.md#config),
[File Server](./doc/fileservice-env.md#config)
