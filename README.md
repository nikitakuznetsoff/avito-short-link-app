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
- `localhost:9000/balance` - получение баланса пользователя
```
curl --header "Content-Type: application/json" \
--request POST \
--data '{"userID": 2}' \
http://localhost:9000/balance
```
- `localhost:9000/balance?currency=USD` - получение баланса пользователя в заданной валюте
```
curl --header "Content-Type: application/json" \
--request POST \
--data '{"userID": 2}' \
http://localhost:9000/balance?currency=USD
```
- `localhost:9000/balance/withdraw` - снятие денег с баланса пользователя
```
curl --header "Content-Type: application/json" \
--request POST \
--data '{"userID": 2, "value": 50.44}' \
http://localhost:9000/balance/withdraw
```
## Структура проекта
Пытался структурировать проект в соответствии с [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
```
BalanceApp
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
|   |
│   └───templates
│       │   index.html
|       |   result.html  
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
