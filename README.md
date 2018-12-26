# Как использовать

### Создание базы данных и таблиц PostgreSQL
Из корневой папки проекта ввести следующие команды:
```
sh parser/scripts/create_db.sh
```

### Запуск парсера
```
go build parser/ && ./parser/parser --config="parser/config/config.json"
```

### Запуск сервера с API
```
go build server/ && ./server/server
```

### Очистка
```
sh parser/scripts/drop_db.sh
```
