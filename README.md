# avito-short-link-app
## Запуск
Запуск осуществуляется с помощью `docker-compose up` из корня приложения
## Описание проекта
Основные моменты:
- Проект запускается на порту `:9000`
- В качестве HTTP роутера используется `https://github.com/julienschmidt/httprouter`
- В качестве хранилища используется база MySQL (v. 8.0.21)
  + База данных содержит 1 таблицу `links` 
  + Структуру таблицы можно посмотреть в файле инициализации `_sql/db.sql`
- Для тестирования методов работы с БД использовал [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock)
  
## Описание API
- `localhost:9000/{:indentificator}` - получание полной ссылки
```
curl --request GET \
http://localhost:9000/test
```
- `localhost:9000/create` - запрос на создание сокращенной ссылки.
  + link - ссылка
  + name - пользовательское имя для сокращенной ссылки, по умолчанию генерируется рандмное имя длиной 6 символов
```
curl --header "Content-Type: application/x-www-form-urlencoded" \
--request POST \
--data "link=https://vk.com/feed&name=facebook" \
http://localhost:9000/create
```
## Структура проекта
Пытался структурировать проект в соответствии с [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
```
ShortLinkApp
│   README.md
│   Dockerfile
|   docker-compose.yml
|
└───bin
│   | shortlinkapp
|
└───_sql
│   | db.sql
│
└───cmd
│   └───short-link-app
│       │   main.go
│   
└───pkg
|   |
│   └───database
│   |   │   repo.go
│   |   │   repo_test.go
|   |
│   └───handlers
│   |   │   links.go
|   |
│   └───models
│   |   │   short_link.go
|   
└───templates
│   │   index.html
|   |   result.html  
|
└───script
    │   wait-for-it.sh
```

- `bin/shortlinkapp` - бинарник для запуска проекта в контейнере
- `_sql/db.sql` - файл инициализации базы данных с созданием нужных таблиц
- `cmd/short-link-app/main.go` - файл для запуска приложения
- `pkg/database` - реализация работы с БД по паттерну "Репозиторий"
- `pkg/handlers` - HTTP обработчики для запросов
- `pkg/models` - описания объектов
- `pkg/templates` - шаблоны для веб-страниц
- `script/wait-for-it.sh` - скрипт для ожидания доступности TCP хоста с портом [wait-for-it](https://github.com/vishnubob/wait-for-it)
  > Используется во время развертывания для ожидания запуска БД
