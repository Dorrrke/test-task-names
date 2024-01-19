### Тестовое задание "Имена"

## Endpoints
Сервис обрабатывает следующие ендпоинты: 
1. /new - добавление новых имен
2. /users - получение всех имен
3. /user/{id} - получение имени по id
4. /user/{id}/update - обновление имени по id
5. /user/{id}/delete - удаление имени по id
6. /users/filter/surmane - вывод имен по фамилии
7. /users/filter/patronymic - вывод имен по отчеству
8. /users/filter/age - вывод имен по возрасту
9. /users/filter/gender - вывод имен по гендеру
10. /users/filter/national - вывод имен по национальности

### Запуск
## Запуск миграций
Для запуска миграций требуется запустить мигратор с соответствующими флагами:
`go run .\cmd\migrator\main.go --storage-path=<адрес базы данных> --migration-path=./migrations `
## Запуск сервиса
Сервис конфигурируется через .env файл, который лежит в корневой директории (values.env) либо при помощи флагов:
1. -a addres and port to rin server
2. -d data base addres
3. -env application env
