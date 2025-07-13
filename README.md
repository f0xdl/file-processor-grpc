
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
- [ ] Server:
  - [ ] Add `.proto` in gRPC-service
    - ProcessFiles(stream FilePath) returns (stream FileResult)
    - Using Server-side streaming
  - [ ] Accepts a list of file paths via gRPC
  - [ ] Processes files in parallel
    - [ ] Counting lines
    - [ ] Counting words
    - [ ] Limit: no more than 5 files can be processed simultaneously
    - [ ] Fan in processing files
      - [ ] Fan-out: dispatch processing to goroutines
      - [ ] Fan-in: collect results and send to client
  - [ ] Stream results back to the client
  - [ ] Support context cancellation on request
  - [ ] Stores processing history
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
