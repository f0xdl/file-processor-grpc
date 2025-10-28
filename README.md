![License](https://img.shields.io/github/license/Ileriayo/markdown-badges?style=flat-square)
![Go](https://img.shields.io/badge/go-1.25-%2300ADD8.svg?style=flat-square&logo=go&logoColor=white)
![gRPC](https://img.shields.io/badge/gRPC-1.73-blue?style=flat-square&logo=go&logoColor=white)
![gin](https://img.shields.io/badge/Gin-1.10-00C397?style=flat-square&logo=go&logoColor=white)
![zerolog](https://img.shields.io/badge/zerolog-f33?style=flat-square)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=flat-square&logo=docker&logoColor=white)

# 📦FileProcessor gRPC `🛠 Practice`

> [!note]
> This project is intended for research purposes only.
> Its functionality on your device is not guaranteed.
> The concepts used in this project are not absolute;
> they were chosen solely for the purpose of testing gRPC stream methods.

> [!note]
> Данный проект выполнен исключительно в исследовательских целях. 
> Его работоспособность на вашем устройстве не гарантирована.
> Концепции использованные в данном проекте не являются абсолютом,
> они выбраны исключительно их желания  протестировать пакетную передачу gRPC.    

## 📚Table of contents

🚧🚧🚧

## 📝Description

🚧🚧🚧

## 🛠️ Features

+ 🔗 Accepts a list of virtual file names
+ ⚙️ Processes files in parallel (counting lines and words)
+ 🔄 Streams results back to the client
+ ⏹️ Support context cancellation
+ 🕓 Stores processing history
+ 📤 Supports file uploads
+ 🧩 Using middleware (rate-limit and logger)

## 🚀Getting Started

### Prerequisites

- Docker
- go 1.25

### Environment Variables

- Client
    - `GRPC_SERVER_ADDRESS` - Address to gRPC file service, [host]:port
    - `HTTP_ADDRESS` - HTTP gateway, [host]:port
- FileServer
    - `STORAGE_DIR` - Storage directory for processing files
    - `GRPC_ADDRESS` - Address to gRPC file service, [host]:port

### Install

```sh
task install
```

### Launch with Docker

### Deployment via Docker

🚧🚧🚧

## 📋 TODO

### Http Client

- [x] Accepts a list of file paths via HTTP API
- [x] Send paths to gRPC server
- [x] Supports file uploads to gRPC
- [x] Cancel by timeout `context.WithTimeout`
- [x] Middlewares:
    - [x] Logging method, duration and errors
    - [x] Panic processing
    - [x] Rate limiting

### File Server

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
- [ ] Deployment
    - [x] Create dockerfiles
    - [x] Create docker-compose
    - [ ] Add install bash in taskfile.yml
    - [ ] Write ReadMe
    - [x] Add Taskfile.yml in project

## 📁 Folder structure

```shell
cmd/            # Entrypoint
internal/       
  client/     
  domain/   
  fileservice/
pkg/
  logger/
  safe_service/
```

