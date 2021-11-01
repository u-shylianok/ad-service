# ad-service
ad-service


## Описание
- Тестовый пользователь: u: test, p: test

//// - ошибка level=error msg="sql: no rows in result set" - при GET несуществующего id

### Структура объявления
- ID
- Название объявления
- Дата создания
- Цена
- Ссылка на главное фото
- Ссылки на все фото - (опционально)
- Описание - (опционально)
- Тэги - (опционально)

### Примечания
- Методы интерфейсов repository.go не содержат названия сущностей, к которым они относятся.
- При поиске объявления формат дат YYYY-MM-DD по умолчанию. Остальные форматы планируется добавить в дальнейших версиях


- buildAdFilterQuery (варианты)
- - SELECT ... FROM ... INNER JOIN ... INNER JOIN ... INNER JOIN ... WHERE (date statement)
- - SELECT ... FROM ... INNER JOIN ... AND (date statement) INNER JOIN ... INNER JOIN ...
- - *в примере выше (date statement) будет в первом доступном JOIN. Если JOIN нет, то WHERE в конце запроса*
- - SELECT ... FROM (SELECT WHERE date statement) INNER JOIN ... INNER JOIN ... INNER JOIN ...

- в итоге я получаю: ERROR: column \"ads.id\" must appear in the GROUP BY clause or be used in an aggregate function (SQLSTATE 42803)
- при запросе localhost:8000/ads/search?username=test&startdate=2021-10-12&enddate=2021-10-15&tags=ТЕСТ&tags=ЗЕЛЕНЫЙ