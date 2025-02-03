# Online Music Library API

![Go Version](https://img.shields.io/badge/Go-1.21-blue) ![License](https://img.shields.io/badge/License-MIT-green)

## Описание

Тестовое задание для Effective Mobile. Реализация онлайн библиотеки песен.

## Возможности

- Получение данных библиотеки с фильтрацией по всем полям и пагинацией
- Получение текста песни с пагинацией по куплетам
- Удаление песни
- Изменение данных песни
- Добавление новой песни в формате

## Стек технологий

- **Язык**: Go (Golang)
- **Фреймворк**: Chi
- **База данных**: PostgreSQL
- **Документация API**: Swagger (OpenAPI 3.0)
- **Логирование**: Zap

## Установка и запуск

### Требования:

- Go 1.21+
- PostgreSQL

### Клонирование репозитория

```sh
$ git clone https://github.com/plasmatrip/music-library-api.git
$ cd music-library-api
```

### Настройка окружения

Создайте `.env` файл в директории c исполняемым файлом и добавьте конфигурацию:

```env
RUN_ADDRESS"`          //адрес веб-сервера
DATABASE_URI"`         //DSN базы данных
INFO_SERVICE_ADDRESS"` //адрес внешнего сервиса
LOG_LEVEL"`            //уровень логирования
```

### Установка зависимостей

```sh
$ go mod tidy
```

### Компиляция сервиса

```sh
$ go build -o ./cmd/muslib ./cmd/main.go
```

### Запуск сервиса

```sh
$ ./cmd/muslib
```

## Автор

[plasmatrip](https://github.com/plasmatrip)
