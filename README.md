# Тестовое задание

## запуск
```shell
docker compose build
docker compose up
```

## Выполнено
- Базовый функционал
- Поддержка долгих команд
- Сборка приложения (Dockerfile, compose.yml)

## Произвольный функционал
- Получение информации о выполнении команды по ее id.
- Обновление команды по id.

## Архитектура приложения:
Handlers Layer: Обрабатывает все запросы к API, соответствующие различным маршрутам.
- Проверка доступности сервера: ```GET /api/v1/ping```
- Получение списка всех команд: ```GET /api/v1/commands/```
- Получение информации о команде по её идентификатору: ```GET /api/v1/commands/:id```
- Получение вывода выполнения команды по её идентификатору: ```GET /api/v1/commands/:id/info```
- Создание новой команды: ```POST /api/v1/commands/```
- Остановка выполнения команды по её идентификатору: ```POST /api/v1/commands/:id/stop```
- Обновление информации о команде по её идентификатору: ```PUT /api/v1/commands/:id```

Service Layer: Включает в себя репозиторий для работы с базой данных и исполнитель для запуска команд.
- Runner: Отвечает за запуск и управление выполнением команд. Создаёт новый процесс для выполнения команды. Позволяет добавлять новые команды, запускать их, останавливать, а также получать вывод выполнения и статус завершения команды.

Repository Layer: Отвечает за взаимодействие с базой данных.

## Postgres
Таблица: commands 
| Поле              | Тип              | Обязательность   | Значение по умолчанию |
|-------------------|------------------|------------------|------------------|
| id                | UUID             | PRIMARY KEY      |gen_random_uuid() |
| name              | VARCHAR(64)      | NOT NULL         |                  |
| description       | TEXT             |                  |                  |
| script            | TEXT             | NOT NULL         |                  |

Таблица: commands_info
| Поле              | Тип              | Обязательность   | Значение по умолчанию |
|-------------------|------------------|------------------|------------------|
| id                | UUID             | PRIMARY KEY      | DEFAULT gen_random_uuid() |
| commands_id       | UUID             | REFERENCES commands(id) |           |
| start_time        | TIMESTAMP        | NOT NULL         | DEFAULT CURRENT_TIMESTAMP |
| end_time          | TIMESTAMP        |                  | DEFAULT NULL     |
| exitcode          | INTEGER          |                  | DEFAULT 0        |
| output            | TEXT             |                  | DEFAULT ''       |
