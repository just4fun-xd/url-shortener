# URL Shortener

[![Go](https://img.shields.io/badge/-Go-00ADD8?style=flat&logo=Go&logoColor=ffffff)](https://golang.org/)
[![SQLite](https://img.shields.io/badge/-SQLite-003B57?style=flat&logo=SQLite&logoColor=ffffff)](https://www.sqlite.org/)
[![GORM](https://img.shields.io/badge/-GORM-374151?style=flat&logo=data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMTAwMCIgaGVpZ2h0PSI1MDAiIHZpZXdCb3g9IjAgMCAxMDAwIDUwMCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48cGF0aCBkPSJNMCAyNSBIMTAwMFYyNUgwVjI1Wk0wIDQ3NSBIMTAwMFY0NzVIMFY0NzVaIiBmaWxsPSIjMzc0MTUxIi8+PHBhdGggZD0iTTUwMCAwQzI3OS43NTggMCAxMCAyNzkuNzU4IDEwIDUwMEMxMCA3MjAuMjQyIDI3OS43NTggMTAwMCA1MDAgMTAwMUM3MjAuMjQyIDEwMDAgMTAwMCA3MjAuMjQyIDEwMDAgNTAwQzEwMDAgMjc5Ljc1OCA3MjAuMjQyIDEwIDUwMCAwWk01MDAgODU0LjM1NEMzMDEuNjk5IDg1NC4zNTQgMTQ1LjY0NiA2OTguMzAxIDE0NS42NDYgNTAwQzE0NS42NDYgMzAxLjcwMiAzMDEuNzAzIDE0NS42NDYgNTAwIDE0NS42NDZDNzg4LjM1NCAxNDUuNjQ2IDk1NC4zNTQgMzAxLjY5OSA5NTQuMzU0IDUwMUM5NTQuMzU0IDY5OC4zMDEgNzg4LjMwMSA4NTQuMzU0IDUwMCA4NTQuMzU0WiIgZmlsbD0id2hpdGUiLz48L3N2Zz4=)](https://gorm.io/)
[![Gorilla Mux](https://img.shields.io/badge/-Gorilla%20Mux-74D269?style=flat&logo=data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMTAwMCIgaGVpZ2h0PSI1MDAiIHZpZXdCb3g9IjAgMCAxMDAwIDUwMCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48cGF0aCBkPSJNMCAyNSA0MDAgMjUgNDAwIDI1SDBWMjVaTTAgNDc1IDQwMCA0NzUgNDAwIDQ3NUgwVjQ3NVpNNTI4IDI1SDYwMFY1MDBINTQyVjI1WiIgZmlsbD0iIzQ4QkQ2QyIvPjxwYXRoIGQ9Ik02MDAgMEgzMDBDMTQwIDAgMCAxNDAgMCAzMDBWMjAwSDEwMEwxMDAgMzAwSDBWMzAwQzAgNDQwIDE0MCA1ODAgMzAwIDYwMEgyMDAgNDgwSDMwMFM0MjAwIDUyMDAgNTIwMCA0MzAwTDYwMCA0MjAwVjI4MDAgTDAgMjQwMFYzNzAwTDMwMCAzODAwVjI2MDBMMzYwMCAyNTgwbDAgMCIgZmlsbD0id2hpdGUiLz48L3N2Zz4=)](https://www.gorillatoolkit.org/)

URL Shortener — это простой сервис для сокращения длинных URL и перенаправления по коротким ссылкам. Проект написан на Go с использованием GORM и Gorilla Mux.

---

## **Особенности проекта**

- Генерация уникальных коротких ссылок.
- Перенаправление пользователей по коротким ссылкам на оригинальные URL.
- Хранение данных в базе SQLite.
- Учёт количества переходов по короткой ссылке.
- Простая архитектура и понятный код.

---

## **Технологии**

- **Язык:** Go
- **Фреймворк:** Gorilla Mux
- **ORM:** GORM
- **База данных:** SQLite

---

## **Как запустить проект**

### **1. Клонируйте репозиторий**

```bash
git clone https://github.com/just4fun-xd/url-shortener
cd url-shortener
```

### **2. Установите зависимости**

Убедитесь, что у вас установлен Go версии 1.18 или выше. Установите необходимые зависимости:

```bash
go mod tidy
```

### **3. Подготовьте базу данных**

Создайте файл базы данных и выполните миграцию:

```bash
sqlite3 url_shortener.db < migrations/init.sql
```

### **4. Запустите сервер**

Запустите сервер с помощью команды:

```bash
go run cmd/main.go
```

Сервер запустится на `http://localhost:8080`.

---

## **API Эндпоинты**

### **1. Создание короткой ссылки**

**POST /shorten**

- **Описание:** Создаёт короткую ссылку для переданного оригинального URL.
- **Пример запроса:**
  ```json
  {
      "original_url": "https://example.com"
  }
  ```
- **Пример ответа:**
  ```json
  {
      "ID": 1,
      "ShortURL": "abc123",
      "OriginalURL": "https://example.com",
      "CreatedAt": "2024-12-25T12:00:00Z",
      "VisitCount": 0
  }
  ```

### **2. Перенаправление по короткой ссылке**

**GET /{short_url}**

- **Описание:** Перенаправляет пользователя на оригинальный URL.
- **Пример запроса:**
  ```bash
  curl -X GET http://localhost:8080/abc123
  ```

- **Ответ:** Перенаправление на `https://example.com`.

---

## **Архитектура проекта**

```
url-shortener/
├── cmd/              // Главная точка входа
│   └── main.go       // Запуск приложения
├── pkg/              // Основная логика
│   ├── api/          // Обработчики HTTP
│   │   └── router.go // Настройка маршрутов
│   ├── storage/      // Работа с базой данных
│   │   └── database.go
│   └── migrations/   // SQL-скрипты для миграции
├── go.mod            // Модуль Go
└── README.md         // Описание проекта
```

---

## **Как добавить новые функции**

1. **Добавление нового маршрута:**
   - Добавьте новый маршрут в `pkg/api/router.go`.
   - Реализуйте обработчик для маршрута.

2. **Добавление полей в таблицу:**
   - Измените структуру `URL` в `pkg/storage/database.go`.
   - Добавьте соответствующую миграцию в `migrations/`.

3. **Подключение к другой базе данных:**
   - Измените `gorm.Open(sqlite.Open(...))` на нужный драйвер (например, PostgreSQL).

---

## **Идеи для улучшения**

- **Статистика:** Добавьте эндпоинт для вывода статистики по коротким ссылкам.
- **Аутентификация:** Реализуйте поддержку пользователей и авторизацию.
- **Истечение срока действия ссылки:** Добавьте возможность задавать срок действия для короткой ссылки.
- **Кэширование:** Используйте Redis для ускорения перенаправлений.

---

## **Контакты**

Если у вас есть вопросы или предложения, свяжитесь со мной:

- **Email:** k.shalygin@yandex.ru
- **GitHub:** [github.com/just4fun-xd](https://github.com/just4fun-xd)
