# WebCalc

---

Web-сервис для подсчета решения математических выражений
(финальное задание Спринт 1.)

## О проекте
Это простейший веб-сервер, который предоставляет 
API для вычисления математических выражений с 
использованием алгоритма **обратной польской нотации (RPN)**.

 Функционал:
 - Поддержка математических операций: 
   - сложение, вычитание
   - умножение, деление.
 - Логирование всех веб-запросов.
 - Тесты для алгоритма вычисления и веб-клиента.

---

## Как запустить
1. Клонируйте репозиторий
    ```bash
    git clone https://github.com/develslawer/webcalc.git
    cd webcalc
    ```
2. Установите зависимости
    ```bash
    go mod tidy
    ```
3. Запуск тестов (опционально)
    ```bash
    go test ./...
    ```
4. Запустите проект
    ```bash
    go run ./cmd/main.go
   ```

---

## Как пользоваться

конечные точки:
 - api/v1/calculate
### PowerShell
```bash
Invoke-WebRequest -Uri "http://localhost/api/v1/calculate?a=1" `
    -Method POST `
    -Headers @{ "Content-Type" = "application/json" } `
    -Body '{"expression": "YOUR EXPRESSION"}'
```

Метод: POST 

Тело запроса JSON:
```json
{
  "expression": "1+2+(1-3*2)/5+12"
}
```

Статус код: 200 |
Тело ответа JSON:
```json
{
  "result": 14.000000
}
```
### Другие случаи (ошибки):

Статус код: 400/500 |
Тело ответа JSON:
```json
{
   "error": "text of error"
}
```

---

## Дополнительные возможности

 - возможно задать порт на котором будет работать сервер:
   - windows
      ```bash
      SET VAR HOST=8080  
      ```
   - unix
      ```bash
      EXPORT HORT=8080
      ```