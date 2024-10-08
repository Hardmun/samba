# API Документация

### Базовый URL

`https://1cplatform.mitcoms.ru/Dev/U_MediaFinance_for_testing/hs/samba`

## Точки доступа (Endpoints)

### 1. Конечные рекламодатели

### 1.1 **Cписок конечных рекламодателей.**

**Метод:** `GET`

**URL:** `/finaladv`

**Параметры запроса:**

| Parameter | Type              | Required | Description                                      |
| --------- | ----------------- | -------- | ------------------------------------------------ |
| `inn`     | `string`          |          | ИНН конечного рекламодателя                      |
| `kpp`     | `string`          |          | КПП конечного рекламодателя                      |
| `name`    | `string` (%like%) |          | Наименование конечного рекламодателя             |
| `uuid`    | `string`          |          | Уникальный идентификатор конечного рекламодателя |

**Пример ответа JSON:**

```json
{
    "status": "ok",
    "message": "seccess",
    "result": [
        {
            "uuid": "301ca510-80b8-11ef-816c-0050568a4702",
            "name": "Карат Рус Медиа ООО FOR TEST",
            "fullName": "Карат Рус Медиа ООО FOR TEST",
            "inn": "7717707438",
            "kpp": "771701002"
        },
        {
            "uuid": "9e188360-56d2-11ed-b75a-005056995bef",
            "name": "Карат Рус Медиа ООО",
            "fullName": "Карат Рус Медиа ООО",
            "inn": "7717707438",
            "kpp": "771701001"
        }
    ]
}
```

### 1.2 **Создать конечного рекламодателя.**

**Метод:** `POST`

**URL:** `/finaladv`

**Параметры JSON:**

| Parameter  | Type     | Required | Description                                 |
| ---------- | -------- | -------- | ------------------------------------------- |
| `inn`      | `string` | *        | ИНН конечного рекламодателя                 |
| `kpp`      | `string` |          | КПП конечного рекламодателя                 |
| `name`     | `string` | *        | Наименование конечного рекламодателя        |
| `fullName` | `string` |          | Полное наименование конечного рекламодателя |

**Пример запроса JSON:**

```json
{
    "inn": "7717707438",
    "kpp": "771701001",
    "name": "Карат Рус Медиа ООО",
    "fullName": "ООО \"Карат Рус Медиа\""
}
```

**Пример ответа JSON:**

```json
{
    "status": "error",
    "message": "Конечный рекламодатель с таким ИНН уже существует: Карат Рус Медиа ООО",
    "result": {
        "name": "Карат Рус Медиа ООО",
        "uuid": "9e188360-56d2-11ed-b75a-005056995bef"
    }
}
```

### 1.3 **Обновить конечного рекламодателя.**

**Метод:** `PUT`

**URL:** `/finaladv`

**Параметры JSON:**

| Parameter  | Type     | Required | Description                                      |
| ---------- | -------- | -------- | ------------------------------------------------ |
| `uuid`     | `string` |          | Уникальный идентификатор конечного рекламодателя |
| `inn`      | `string` |          | ИНН конечного рекламодателя                      |
| `kpp`      | `string` |          | КПП конечного рекламодателя                      |
| `name`     | `string` |          | Наименование конечного рекламодателя             |
| `fullName` | `string` |          | Полное наименование конечного рекламодателя      |

**Пример запроса JSON:**

```json
{
    "inn": "7717707438",
    "kpp": "771701001",
    "name": "Карат Рус Медиа ООО",
    "fullName": "ООО \"Карат Рус Медиа\""
}
```

**Пример ответа JSON:**

```json
{
    "status": "ok",
    "message": "Конечный рекламодатель обновлен.",
    "result": {
        "uuid": "8a378302-80c0-11ef-816c-0050568a4702",
        "inn": "7717707438",
        "kpp": "771701003",
        "name": "Карат Рус Медиа ООО FOR TEST"
    }
}
```

## 2. Конечные договоры рекламодателей

### 2.1 **Список конечных договоров рекламодателей**

**Метод:** `GET`

**URL:** `/finalcontracts`

**Параметры запроса:**

| Parameter       | Type     | Required | Description                                      |
| --------------- | -------- | -------- | ------------------------------------------------ |
| `uuid`          | `string` |          | Уникальный идентификатор конечного договора      |
| `number`        | `string` |          | Номер конечного договора                         |
| `date`          | `string` |          | Дата конечного договора                          |
| `inn_finaladv`  | `string` | *(or)    | ИНН конечного рекламодателя                      |
| `kpp_finaladv`  | `string` |          | КПП конечного рекламодателя                      |
| `uuid_finaladv` | `string` | *(or)    | Уникальный идентификатор конечного рекламодателя |

**Пример ответа JSON:**

```json
{
    "status": "ok",
    "message": "seccess",
    "result": [
        {
            "uuid": "cd2539af-80f1-11ef-816c-0050568a4702",
            "number": "1-Test",
            "date": "2022-07-21T00:00:00",
            "inn_finaladv": "7717707438",
            "kpp_finaladv": "771701004",
            "name_finaladv": "Карат Рус Медиа ООО FOR TEST 22",
            "uuid_finaladv": "301ca510-80b8-11ef-816c-0050568a4702",
            "inn_contractor": "7726472670",
            "kpp_contractor": "772601001",
            "name_contractor": "1001101 ООО ( до 31.01.2023 АЙТАРГЕТ ТЕХНОЛОГИИ)",
            "uuid_contractor": "9b955c16-80ef-11ef-816c-0050568a4702"
        }
    ]
}
```

### 2.2 **Создать договор конечного рекламодателя.**

**Метод:** `POST`

**URL:** `/finalcontracts`

**Параметры JSON:**

| Parameter         | Type     | Required | Description                                      |
| ----------------- | -------- | -------- | ------------------------------------------------ |
| `number`          | `string` | *        | Номер конечного договора рекламодателя           |
| `date`            | `string` | *        | Дата конечного договора рекламодателя            |
| `inn_finaladv`    | `string` | *(or)    | ИНН конечного рекламодателя                      |
| `kpp_finaladv`    | `string` |          | КПП конечного рекламодателя                      |
| `uuid_finaladv`   | `string` | *(or)    | Уникальный идентификатор конечного рекламодателя |
| `inn_contractor`  | `string` | **(or)   | ИНН юр. лица                                     |
| `kpp_contractor`  | `string` |          | КПП юр. лица                                     |
| `uuid_contractor` | `string` | **(or)   | Уникальный идентификатор юр. лица                |

**Пример запроса JSON:**

```json
{
    "number": "1-Test",
    "date": "20220721",
    "inn_finaladv": "7717707438",
    "kpp_finaladv": "771701004",
    "inn_contractor":"7726472670"
}
```

**Пример ответа JSON:**

```json
{
    "status": "ok",
    "message": "Конечный договор рекламодателя успешно создан",
    "result": {
        "uuid": "569d3eee-80f2-11ef-816c-0050568a4702",
        "number": "2-Test",
        "date": "2022-07-21T00:00:00"
    }
}
```

### 2.3 **Обновить договор конечного рекламодателя.**

**Метод:** `PUT`

**URL:** `/finalcontracts`

**Параметры JSON:**

| Parameter         | Type     | Required | Description                                       |
| ----------------- | -------- | -------- | ------------------------------------------------- |
| `uuid`            | `string` | * (or)   | Уникальный номер конечного договора рекламодателя |
| `number`          | `string` | ** (or)  | Номер конечного договора рекламодателя            |
| `date`            | `string` | ** (or)  | Дата конечного договора рекламодателя             |
| `inn_finaladv`    | `string` | **(or)   | ИНН конечного рекламодателя                       |
| `kpp_finaladv`    | `string` |          | КПП конечного рекламодателя                       |
| `uuid_finaladv`   | `string` | **(or)   | Уникальный идентификатор конечного рекламодателя  |
| `inn_contractor`  | `string` |          | ИНН юр. лица                                      |
| `kpp_contractor`  | `string` |          | КПП юр. лица                                      |
| `uuid_contractor` | `string` |          | Уникальный идентификатор юр. лица                 |

**Пример запроса JSON:**

```json
{
    "uuid": "cd2539af-80f1-11ef-816c-0050568a4702",
    "number": "44-Test",
    "date": "20220721",
    "uuid_finaladv": "06c71ae6-80c2-11ef-816c-0050568a4702",
    "inn_contractor":"2309115177"
}
```

**Пример ответа JSON:**

```json
{
    "status": "ok",
    "message": "Конечный договор рекламодателя обновлен",
    "result": {
        "uuid": "cd2539af-80f1-11ef-816c-0050568a4702",
        "number": "44-Test",
        "date": "2022-07-21T00:00:00"
    }
}
```

---

### 3. Юридические лица (Customers)

### 3.1 **Cписок юр. лиц**

**Метод:** `GET`

**URL:** `/customers`

**Параметры запроса:**

| Parameter | Type              | Required | Description                       |
| --------- | ----------------- | -------- | --------------------------------- |
| `inn`     | `string`          |          | ИНН юр. лица                      |
| `kpp`     | `string`          |          | КПП юр. лица                      |
| `name`    | `string` (%like%) |          | Наименование юр. лица             |
| `uuid`    | `string`          |          | Уникальный идентификатор юр. лица |

**Пример ответа JSON:**

```json
{
    "status": "ok",
    "message": "seccess",
    "result": [
        {
            "uuid": "d241d1cf-dbc0-46f7-80c5-0a043c390827",
            "name": "АЛД Автомотив ООО ",
            "fullName": "ООО “АЛД Автомотив”",
            "inn": "7725514969",
            "kpp": "772501001"
        }
    ]
}
```

---

### 4. Бренды (Brands)

### 4.1 **Cписок брендов**

**Метод:** `GET`

**URL:** `/brands`

**Параметры запроса:**

| Parameter         | Type     | Required | Description                              |
| ----------------- | -------- | -------- | ---------------------------------------- |
| `uuid`            | `string` |          | Уникальный идентификатор бренда          |
| `name`            | `string` |          | Наименование бренда                      |
| `inn`             | `string` |          | ИНН юр. лица бренда                      |
| `kpp`             | `string` |          | КПП юр. лица бренда                      |
| `uuid_contractor` | `string` |          | Уникальный идентификатор юр. лица бренда |

**Пример ответа JSON:**

```json
{
    "status": "ok",
    "message": "seccess",
    "result": [
        {
            "uuid": "1dff21fa-6a02-11ef-816a-005056b27e80",
            "name": "SAMBA (ЛТ)",
            "name_contractor": "ТРАДИЦИЯ ООО",
            "inn_contractor": "7810751880",
            "kpp_contractor": "781001001",
            "uuid_contractor": "a787dc66-52fb-11ef-8168-005056b27e80"
        },
        {
            "uuid": "04e19a7f-52ff-11ef-8168-005056b27e80",
            "name": "Русские традиции ",
            "name_contractor": "ТРАДИЦИЯ ООО",
            "inn_contractor": "7810751880",
            "kpp_contractor": "781001001",
            "uuid_contractor": "a787dc66-52fb-11ef-8168-005056b27e80"
        }
    ]
}
```

### 4.2 **Создать бренд**

**Метод:** `POST`

**URL:** `/brands`

**Параметры JSON:**

| Parameter         | Type     | Required | Description                              |
| ----------------- | -------- | -------- | ---------------------------------------- |
| `name`            | `string` | *        | Наименование бренда                      |
| `inn_contractor`  | `string` | * (or)   | ИНН юр.лица бренда                       |
| `kpp_contractor`  | `string` |          | КПП юр.лица бренда                       |
| `uuid_contractor` | `string` | * (or)   | Уникальный идентификатор юр. лица бренда |

**Пример запроса JSON:**

```json
{
    "name": "My brand NIKE",
    "inn_contractor": "3526019521"
}
```

**Пример ответа JSON:**

```json
{
    "status": "ok",
    "message": "Бренд успешно создан.",
    "result": {
        "uuid": "da7d0cd6-6fb8-4bd2-b93b-096be4e1c11f",
        "name": "My brand NIKE"
    }
}
```

### 4.3 **Обновить бренд**

**Метод:** `PUT`

**URL:** `/brands`

**Параметры JSON:**

| Parameter | Type     | Required | Description                     |
| --------- | -------- | -------- | ------------------------------- |
| `name`    | `string` | *        | Наименование бренда             |
| `uuid`    | `string` | *        | Уникальный идентификатор бренда |

**Пример запроса JSON:**

```json
{
    "name": "NIKE NIKE",
    "uuid": "da7d0cd6-6fb8-4bd2-b93b-096be4e1c11f"
}
```

**Пример ответа JSON:**

```json
{
    "status": "ok",
    "message": "Бренд обновлен",
    "result": {
        "uuid": "da7d0cd6-6fb8-4bd2-b93b-096be4e1c11f",
        "name": "NIKE NIKE"
    }
}
```

---

### 5. Рекламные кампании (RK)

### 5.1 **Cписок рекламных кампаний**

**Метод:** `GET`

**URL:** `/rk`

**Параметры запроса:**

| Parameter         | Type     | Required | Description                          |
| ----------------- | -------- | -------- | ------------------------------------ |
| `uuid`            | `string` |          | Уникальный идентификатор РК          |
| `name`            | `string` | (%like%) | Наименование РК                      |
| `rk_id`           | `string` |          | Идентификатор РК в SAMBA             |
| `numberMF`        | `string` |          | Внутренний номер РК в 1С             |
| `inn_contractor`  | `string` |          | ИНН юр. лица РК                      |
| `kpp_contractor`  | `string` |          | КПП юр. лица РК                      |
| `uuid_contractor` | `string` |          | Уникальный идентификатор юр. лица РК |

**Пример ответа JSON:**

```json
{
    "status": "ok",
    "message": "seccess",
    "result": [
        {
            "uuid": "95e85794-8285-11ef-816c-0050568a4702",
            "name": "SAMBA|TORNADO|Oct - 2024",
            "rk_id": "У191994",
            "numberMF": "У085450",
            "name_brand": "NIKE JUST DO IT",
            "uuid_brand": "8e73927f-8c97-4737-b99b-f4139c3f9a80",
            "name_contractor": "ТОРНАДО ООО",
            "inn_contractor": "3526019521",
            "kpp_contractor": "352601001",
            "uuid_contractor": "69f2faf8-7a57-11ef-816b-005056b27e80",
            "name_finaladv": "«ЛИДКОМ ИНВЕСТМЕНТС ЛИМИТЕД»",
            "inn_finaladv": "9909369754",
            "kpp_finaladv": "774751001",
            "uuid_finaladv": "8e6796de-585d-11ed-b75a-005056995bef",
            "number_contract": "№ 135/19",
            "date_contract": "2019-10-01T00:00:00",
            "uuid_contract": "e93cfac2-585d-11ed-b75a-005056995bef"
        }
    ]
}
```
### 5.2 **Создать рекламную кампанию**

**Метод:** `POST`

**URL:** `/rk`

**Параметры JSON:**

| Parameter         | Type     | Required | Description                                      |
| ----------------- | -------- | -------- | ------------------------------------------------ |
| `rk_id`           | `string` | *        | Номер РК в SAMBA                                 |
| `date_start`      | `string` | *        | Дата начала РК                                   |
| `date_end`        | `string` | *        | Дата окончания РК                                |
| `inn_contractor`  | `string` | * (or_1) | ИНН юр.лица РК                                   |
| `kpp_contractor`  | `string` |          | КПП юр.лица РК                                   |
| `uuid_contractor` | `string` | * (or_1) | Уникальный идентификатор юр.лица РК              |
| `inn_finaladv`    | `string` | (or_2)   | ИНН конечного рекламодателя                      |
| `kpp_finaladv`    | `string` |          | КПП конечного рекламодателя                      |
| `uuid_finaladv`   | `string` | (or_2)   | Уникальный идентификатор конечного рекламодателя |
| `number_contract` | `string` | (or_3)   | Номер договора рекламодателя                     |
| `date_contract`   | `string` | (or_3)   | Дата договора рекламодателя                      |
| `uuid_contract`   | `string` | (or_3)   | Уникальный идентификатор договора рекламодателя  |

**Пример запроса JSON:**

```json
{
    "rk_id": "У191994",
    "inn_contractor": "3526019521",
    "kpp_contractor": "352601001",
    "date_start": "20241001",
    "date_end": "20241031",
    "inn_finaladv": "9909369754",
    "kpp_finaladv": "",
    "number_contract": "№ 135/19",
    "date_contract": "01.10.2019"
}
```

**Пример ответа JSON:**

```json
{
    "status": "error",
    "message": "РК уже существует",
    "result": {
        "rk_id": "У191994",
        "uuid": "95e85794-8285-11ef-816c-0050568a4702",
        "uuid_brand": "8e73927f-8c97-4737-b99b-f4139c3f9a80",
        "name_brand": "My test brand NIKE",
        "inn_contractor": "3526019521",
        "kpp_contractor": "352601001",
        "name_contractor": "ТОРНАДО ООО",
        "inn_finaladv": "9909369754",
        "kpp_finaladv": "774751001",
        "name_finaladv": "«ЛИДКОМ ИНВЕСТМЕНТС ЛИМИТЕД»",
        "number_contract": "№ 135/19",
        "date_contract": "2019-10-01T00:00:00"
    }
}
```

### 5.3 **Обновить рекламную кампанию**

**Метод:** `PUT`

**URL:** `/rk`

**Параметры JSON:**

| Parameter         | Type     | Required | Description                                      |
| ----------------- | -------- | -------- | ------------------------------------------------ |
| `uuid`            | `string` | *(or)    | Уникальный номер РК                              |
| `rk_id`           | `string` | *(or)    | Номер РК в SAMBA                                 |
| `date_start`      | `string` |          | Дата начала РК                                   |
| `date_end`        | `string` |          | Дата окончания РК                                |
| `inn_contractor`  | `string` | (or_1)   | ИНН юр.лица РК                                   |
| `kpp_contractor`  | `string` |          | КПП юр.лица РК                                   |
| `uuid_contractor` | `string` | (or_1)   | Уникальный идентификатор юр.лица РК              |
| `inn_finaladv`    | `string` | (or_2)   | ИНН конечного рекламодателя                      |
| `kpp_finaladv`    | `string` |          | КПП конечного рекламодателя                      |
| `uuid_finaladv`   | `string` | (or_2)   | Уникальный идентификатор конечного рекламодателя |
| `number_contract` | `string` | (or_3)   | Номер договора рекламодателя                     |
| `date_contract`   | `string` | (or_3)   | Дата договора рекламодателя                      |
| `uuid_contract`   | `string` | (or_3)   | Уникальный идентификатор договора рекламодателя  |

**Пример запроса JSON:**

```json
{
    "uuid": "c833b7d2-831a-11ef-816c-0050568a4702",
    "rk_id": "У191997",
    "inn_contractor": "3526019521",
    "kpp_contractor": "352601001",
    "date_start": "20241001",
    "date_end": "20241031",
    "inn_finaladv": "9909369754",
    "kpp_finaladv": "",
    "number_contract": "№ 135/19",
    "date_contract": "01.10.2019"
}
```

**Пример ответа JSON:**

```json
{
    "status": "ok",
    "message": "РК обновлена",
    "result": {
        "rk_id": "У191997",
        "uuid": "c833b7d2-831a-11ef-816c-0050568a4702",
        "uuid_brand": "8e73927f-8c97-4737-b99b-f4139c3f9a80",
        "name_brand": "NIKE JUST DO IT",
        "inn_contractor": "3526019521",
        "kpp_contractor": "352601001",
        "name_contractor": "ТОРНАДО ООО",
        "inn_finaladv": "9909369754",
        "kpp_finaladv": "774751001",
        "name_finaladv": "«ЛИДКОМ ИНВЕСТМЕНТС ЛИМИТЕД»",
        "number_contract": "№ 135/19",
        "date_contract": "2019-10-01T00:00:00"
    }
}
```

---

### 6. Счет на оплату покупателю (PO)

### 6.1 **Cписок счетов на оплату покупателю**

**Метод:** `GET`

**URL:** `/po`

**Параметры запроса:**

| Parameter    | Type     | Required | Description                    |
| ------------ | -------- | -------- | ------------------------------ |
| `date_start` | `string` |          | Период начала                  |
| `date_end`   | `string` |          | Период окончания               |
| `date`       | `string` |          | Дата счета                     |
| `uuid`       | `string` |          | Уникальный идентификатор счета |
| `number`     | `string` |          | Номер счета в 1С               |
| `uuid_rk`    | `string` |          | Уникальный идентификатор РК    |
| `rk_id`      | `string` |          | Номер РК в SAMBA               |
| `numberMF`   | `string` |          | Внутренний номер РК в 1С       |

**Пример ответа JSON:**

```json
{
    "status": "ok",
    "message": null,
    "result": [
        {
            "uuid": "75cfdd86-c090-4ed8-b7f5-ee05a51b7404",
            "number": "УЛУ082118/02",
            "date": "2024-05-29T16:26:29",
            "amount": 19729.97,
            "uuid_rk": "1fb38b2a-1daa-11ef-8168-005056b27e80",
            "rk_id": "",
            "numberMF": "У082118"
        }
    ]
}
```

### 6.2 **Печатная форма существующего счета на оплату покупателю**

**Метод:** `GET`

**URL:** `/files/po`

**Параметры запроса:**

| Parameter | Type     | Required | Description                    |
| --------- | -------- | -------- | ------------------------------ |
| `number`  | `string` | * (or)   | Номер счета в 1С               |
| `date`    | `string` | * (or)   | Дата счета                     |
| `uuid`    | `string` | * (or)   | Уникальный идентификатор счета |

**Параметры ответа JSON:**

| Parameter | Type     | Required | Description                         |
| --------- | -------- | -------- | ----------------------------------- |
| `status`  | `string` |          | Статус ответа 1С                    |
| `message` | `string` |          | Сообщение                           |
| `result`  | `string` |          | Binary data encoded in Base64 (PDF) |

**Пример ответа JSON:**

```json
{
    "status": "ok",
    "message": null,
    "result": "JVBERi0xLjcKJeLjz9MKMSAwIG9iago8PAovRmlsdGVyIC9GbGF0ZURlY29kZQov\r\nTGVuZ3RoIDIgMCBSCi9MZW5ndGgxIDQ2MjQwCi9MZW5ndGgyIDAKL0xlbmd0aDMg\r\nMAo+PgpzdHJlYW0KeJztvQt8VMX1OD4z97F333ffm+xm9yabLI8NBBIgBKLZQMJD\r\nHokhBIJECEmAQCAhCSK+iFYEIxVqq63aKlrfj7qEAAG0UqS21VL9VmtbW4W21Gor\r\nSltKWyWb/5m59242qP339/t//v/P7//5yGbuPTN37jzOOXPOmTMzF4QRQi"
}
```

### 6.3 **Создать счет на оплату покупателю** 

**Метод:** `POST`

**URL:** `/po`

**Параметры JSON:**

| Parameter | Type     | Required | Description                 |
| --------- | -------- | -------- | --------------------------- |
| `number`  | `string` |          | Номер счета в 1С            |
| `date`    | `string` |          | Дата счета                  |
| `rk_id`   | `string` | * (or)   | Номер РК в SAMBA            |
| `rk_uuid` | `string` | * (or)   | Уникальный идентификатор РК |
| `amount`  | `string` | *        | Сумма счета                 |

**Пример запроса JSON:**

```json
{
    "number": "TEST PO #11",
    "date": "20241008",
    "rk_id": "У191994",
    "amount": 1000
}
```

**Пример ответа JSON:**

```json
{
    "status": "ok",
    "message": "Счет на оплату покупателю успешно создан",
    "result": {
        "uuid": "860c7845-8566-11ef-816c-0050568a4702",
        "number": "УЛУ085450/02",
        "date": "2024-10-08T14:14:47",
        "po_file": ""
        }
}
```

### 6.4 **Обновить счет на оплату покупателю**

**Метод:** `PUT`

**URL:** `/po`

**Параметры JSON:**

| Parameter | Type     | Required | Description                         |
| --------- | -------- | -------- | ----------------------------------- |
| `uuid`    | `string` |          | Уникальный идентификатор счета в 1С |
| `number`  | `string` |          | Номер счета в 1С                    |
| `date`    | `string` |          | Дата счета                          |
| `rk_id`   | `string` | * (or)   | Номер РК в SAMBA                    |
| `rk_uuid` | `string` | * (or)   | Уникальный идентификатор РК         |
| `amount`  | `string` | *        | Сумма счета                         |

**Пример запроса JSON:**

```json
{
    "uuid": "",
    "number": "TEST PO #11",
    "date": "20241008",
    "rk_id": "У191994",
    "amount": 1000
}
```

**Пример ответа JSON:**

```json
{
    "status": "ok",
    "message": "Счет на оплату покупателю обновлен",
    "result": {
        "uuid": "860c7845-8566-11ef-816c-0050568a4702",
        "number": "УЛУ085450/02",
        "date": "2024-10-08T14:14:47",
        "po_file": ""
        }
}
```