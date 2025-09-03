# Система управления пользователями с интерфейсами
Разработать систему управления пользователями с интерфейсами

1. Создать структуру `User` с полями:
```
ID string (UUID)
Name string
Email string
Role string ("admin", "user", "guest")
CreatedAt time.Time
```

2. Определить интерфейс `UserRepository` с методами:
* `Save(user User) error` — сохраняет пользователя
* `FindByID(id string) (User, error)` — находит пользователя по ID
* `FindAll() []User` — возвращает всех пользователей
* `DeleteByID(id string) error` — удаляет пользователя

3. Реализовать два типа хранилищ, которые соответствуют `UserRepository`:
* `InMemoryUserRepo` — хранение пользователей в `map[string]User`
* `MockUserRepo` — имитация работы с БД (например, просто логирует вызовы методов)

4. Добавить функцию `NewUserService(repo UserRepository) UserService`, которая возвращает объект с методами:
* `CreateUser(name, email, role string) (User, error)`
* `GetUser(id string) (User, error)`
* `ListUsers() []User`
* `RemoveUser(id string) error`

5. Дополнительно (если есть время)
Сделать метод `FindByRole(role string) []User` — поиск всех пользователей с определённой ролью.
