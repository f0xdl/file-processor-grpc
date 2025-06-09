# 📦 FileProcessor gRPC
## 🎯 Цель проекта

Создать gRPC-сервис, который:

* Принимает список файловых путей (или имена виртуальных файлов).
* Параллельно обрабатывает файлы (подсчёт строк и слов).
* Стримит результаты обратно клиенту.
* Использует middleware (interceptor) и поддержку `context cancellation`.


## 📐 Архитектура

### gRPC API (protobuf)

```proto
syntax = "proto3";

package fileprocessor;

service FileService {
  rpc ProcessFiles (FileList) returns (stream FileStats);
}

message FileList {
  repeated string paths = 1;
}

message FileStats {
  string path = 1;
  int32 lines = 2;
  int32 words = 3;
  string error = 4; // если ошибка при обработке
}
```



## ✅ Функциональность

### Клиент

* Отправляет список путей (`paths`) на сервер.
* Может отменить выполнение через `context.WithTimeout`.

### Сервер

* Обрабатывает каждый файл в отдельной горутине (**fan-out**).
* Подсчитывает количество строк и слов (можно использовать `bufio.Scanner`).
* Отправляет результат по мере готовности в виде стрима.
* Завершает выполнение по `ctx.Done()` (таймаут или отмена).



## 🔄 Конкурентность и ограничение ресурсов

* **Fan-out/fan-in**:

  * fan-out: воркеры получают задания.
  * fan-in: собираем и отправляем клиенту результаты.
* Ограничение: одновременно обрабатывается не более **5 файлов**.

  * Можно использовать `semaphore` или буферизированный `channel`.



## 🔒 Middleware

### Interceptors
* Логируют метод, длительность, ошибки.
* Обрабатывают panic.



## 🧪 Тестирование

* Юнит-тесты:

  * Подсчёт строк/слов.
  * Ошибка чтения файла.
  * Завершение по таймауту (`context.Cancel()`).
* Интеграционный тест:

  * Клиент ↔ Сервер ↔ Результат.



## ⚙️ Используемые библиотеки

* [`google.golang.org/grpc`](https://pkg.go.dev/google.golang.org/grpc)
* [`google.golang.org/protobuf`](https://pkg.go.dev/google.golang.org/protobuf)
* `context`, `os`, `sync`, `bufio`, `time`
* (опционально) [`golang.org/x/sync/semaphore`](https://pkg.go.dev/golang.org/x/sync/semaphore)



## 🚀 Возможные улучшения

* Авторизация с помощью gRPC metadata.
* Хранение истории обработки.
* Поддержка загрузки файлов (а не только по имени).
