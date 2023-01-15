# HTTP сервис для работы с записями клиента и магазина

## Для запуска необходимо:
1. Создать таблицы в PostgreSQL ("script.sql");
2. Изменить данные в "config.json";

## Запуск сервиса:
```shell
go run ./cmd/api
```

## Примеры запросов:
GET /client/list -- получить список клиентов по фамилии \
Request:
```json
{
  "last_name": "Sokolov"
}
```
Response:
```json
{
  "id": "b2d14bbd-94d5-11ed-a690-3aca73727d74",
  "last_name": "Sokolov",
  "first_name": "Petr",
  "patronymic": "Igorevich",
  "registration_date": "01-01-2012"
}
```

POST /client/create -- создать клиента \
Request:
```json
{
  "last_name": "Sokolov",
  "first_name": "Petr",
  "patronymic": "Igorevich",
  "registration_date": "01-01-2012"
}
```
Response:
```json
{
  "id": "b2d14bbd-94d5-11ed-a690-3aca73727d74"
}
```

PUT /client/update -- обновить клиента \
Request:
```json
{
  "id": "b2d14bbd-94d5-11ed-a690-3aca73727d74",
  "last_name": "Sokolov",
  "first_name": "Petr",
  "patronymic": "Igorevich",
  "age": 22,
  "registration_date": "01-01-2012"
}
```
Response:
```json
{"status": "success"}
```

DELETE /client/delete -- удалить клиента \
Request:
```json
{
  "id": "b2d14bbd-94d5-11ed-a690-3aca73727d74"
}
```
Response:
```json
{"status": "success"}
```

GET /market/list -- получить список магазинов по названию \
Request:
```json
{
  "name": "Magnit"
}
```
Response:
```json
{
  "id": "443e832c-94d6-11ed-a690-3aca73727d74",
  "name": "Magnit",
  "address": "Moscow",
  "active": true
}
```

POST /market/create -- создать магазин \
Request:
```json
{
  "name": "Magnit",
  "address": "Moscow",
  "active": true
}
```
Response:
```json
{
  "id": "443e832c-94d6-11ed-a690-3aca73727d74"
}
```

PUT /market/update -- обновить магазин \
Request:
```json
{
  "id": "443e832c-94d6-11ed-a690-3aca73727d74",
  "name": "Magnit",
  "address": "Moscow",
  "active": false,
  "owner": "Dude"
}
```
Response:
```json
{"status": "success"}
```

DELETE /market/delete -- удалить магазин \
Request:
```json
{
  "id": "443e832c-94d6-11ed-a690-3aca73727d74"
}
```
Response:
```json
{"status": "success"}
```
