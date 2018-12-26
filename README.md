# Как использовать

### Создание базы данных и таблиц PostgreSQL
Из корневой папки проекта ввести следующие команды:
```
sh parser/scripts/create_db.sh
```

### Запуск парсера
```
cd parser && sh scripts/create_db.sh && go build . && ./parser --config="config/config.json" && cd ..
```

### Запуск сервера с API
```
cd server && go build . && ./server && cd ..
```

### Очистка
```
sh parser/scripts/drop_db.sh
```
