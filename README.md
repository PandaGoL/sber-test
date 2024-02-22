# sber-test
Сервис для расчета депозита



### Для запуска приложения:

```
make build && make run
```

Приложение стартует на 8000 порту


### Endpoints:

|Endpoint             |Метод|Ответ            |Структура запроса|Описание|
|---------------------|-----|----------------------------------------------------------------------------------------------------|---|---|
|/v1/deposit|POST  |Код 200 со структурой вида:<details><pre>{<p>    "31.01.2021":"10050",</p><p>    "28.02.2021":"10100.25",</p><p>    "31.03.2021":"10150.75",</p><p>...</p>}<p> </details>Код 400 со структурой вида: <details><pre>{<p>    "error":"описание ошибки"}</details>|date - дата заявки. обязательный формат даты: dd.mm.yyyy;<p>periods - количество месяцев по вкладу. Минимальное значение - 1, максимальное значение - 59<p>amount - сумма вклада. Минимальное значение - 10000, максимальное значение - 2999999<p>rate - процент по вкладу. Минимальное значение - 1, максимальное значение - 7 тело:<details><pre>{<p>    "date":"31.01.2021",</p><p>    "periods":3,</p><p>    "amount": 10000,</p><p>    "rate":6</p><p>}</pre></details>|Осуществляет расчет депозита.|