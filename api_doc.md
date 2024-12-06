# API Документация SAMBA + 1C

### Базовый URL

PRODUCTION:
`https://1cplatform.mitcoms.ru/Media/MF/hs/samba`

TEST:
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
| `inn`      | `string` | \*       | ИНН конечного рекламодателя                 |
| `kpp`      | `string` |          | КПП конечного рекламодателя                 |
| `name`     | `string` | \*       | Наименование конечного рекламодателя        |
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
| `uuid`     | `string` | \*(or)   | Уникальный идентификатор конечного рекламодателя |
| `inn`      | `string` | \*(or)   | ИНН конечного рекламодателя                      |
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
| `inn_finaladv`  | `string` |          | ИНН конечного рекламодателя                      |
| `kpp_finaladv`  | `string` |          | КПП конечного рекламодателя                      |
| `uuid_finaladv` | `string` |          | Уникальный идентификатор конечного рекламодателя |

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
| `number`          | `string` | \*       | Номер конечного договора рекламодателя           |
| `date`            | `string` | \*       | Дата конечного договора рекламодателя            |
| `inn_finaladv`    | `string` | \*(or)   | ИНН конечного рекламодателя                      |
| `kpp_finaladv`    | `string` |          | КПП конечного рекламодателя                      |
| `uuid_finaladv`   | `string` | \*(or)   | Уникальный идентификатор конечного рекламодателя |
| `inn_contractor`  | `string` | \*\*(or) | ИНН юр. лица                                     |
| `kpp_contractor`  | `string` |          | КПП юр. лица                                     |
| `uuid_contractor` | `string` | \*\*(or) | Уникальный идентификатор юр. лица                |

**Пример запроса JSON:**

```json
{
  "number": "1-Test",
  "date": "20220721",
  "inn_finaladv": "7717707438",
  "kpp_finaladv": "771701004",
  "inn_contractor": "7726472670"
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

| Parameter         | Type     | Required   | Description                                       |
| ----------------- | -------- | ---------- | ------------------------------------------------- |
| `uuid`            | `string` | \* (or)    | Уникальный номер конечного договора рекламодателя |
| `number`          | `string` | \* (or)    | Номер конечного договора рекламодателя            |
| `date`            | `string` | \* (or)    | Дата конечного договора рекламодателя             |
| `inn_finaladv`    | `string` | \*\*(or)   | ИНН конечного рекламодателя                       |
| `kpp_finaladv`    | `string` |            | КПП конечного рекламодателя                       |
| `uuid_finaladv`   | `string` | \*\*(or)   | Уникальный идентификатор конечного рекламодателя  |
| `inn_contractor`  | `string` | \*\*\*(or) | ИНН юр. лица                                      |
| `kpp_contractor`  | `string` |            | КПП юр. лица                                      |
| `uuid_contractor` | `string` | \*\*\*(or) | Уникальный идентификатор юр. лица                 |

**Пример запроса JSON:**

```json
{
  "uuid": "cd2539af-80f1-11ef-816c-0050568a4702",
  "number": "44-Test",
  "date": "20220721",
  "uuid_finaladv": "06c71ae6-80c2-11ef-816c-0050568a4702",
  "inn_contractor": "2309115177"
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

| Parameter | Type     | Required | Description                       |
| --------- | -------- | -------- | --------------------------------- |
| `inn`     | `string` |          | ИНН юр. лица                      |
| `kpp`     | `string` |          | КПП юр. лица                      |
| `name`    | `string` | (%like%) | Наименование юр. лица             |
| `uuid`    | `string` |          | Уникальный идентификатор юр. лица |

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
| `name`            | `string` | \*       | Наименование бренда                      |
| `inn_contractor`  | `string` | \* (or)  | ИНН юр.лица бренда                       |
| `kpp_contractor`  | `string` |          | КПП юр.лица бренда                       |
| `uuid_contractor` | `string` | \* (or)  | Уникальный идентификатор юр. лица бренда |

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
| `name`    | `string` | \*       | Наименование бренда             |
| `uuid`    | `string` | \*       | Уникальный идентификатор бренда |

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

| Parameter         | Type      | Required  | Description                                      |
| ----------------- | --------- | --------- | ------------------------------------------------ |
| `rk_id`           | `string`  | \*        | Номер РК в SAMBA                                 |
| `date_start`      | `string`  |           | Дата начала РК                                   |
| `date_end`        | `string`  |           | Дата окончания РК                                |
| `advance`         | `boolean` |           | Признак аванса                                   |
| `inn_contractor`  | `string`  | \* (or_1) | ИНН юр.лица РК                                   |
| `kpp_contractor`  | `string`  |           | КПП юр.лица РК                                   |
| `uuid_contractor` | `string`  | \* (or_1) | Уникальный идентификатор юр.лица РК              |
| `inn_finaladv`    | `string`  | (or_2)    | ИНН конечного рекламодателя                      |
| `kpp_finaladv`    | `string`  |           | КПП конечного рекламодателя                      |
| `uuid_finaladv`   | `string`  | (or_2)    | Уникальный идентификатор конечного рекламодателя |
| `number_contract` | `string`  | (or_3)    | Номер договора рекламодателя                     |
| `date_contract`   | `string`  | (or_3)    | Дата договора рекламодателя                      |
| `uuid_contract`   | `string`  | (or_3)    | Уникальный идентификатор договора рекламодателя  |

**Пример запроса JSON:**

```json
{
  "rk_id": "У191994",
  "inn_contractor": "3526019521",
  "kpp_contractor": "352601001",
  "advance": false,
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

| Parameter         | Type      | Required | Description                                      |
| ----------------- | --------- | -------- | ------------------------------------------------ |
| `uuid`            | `string`  | \*(or)   | Уникальный номер РК                              |
| `rk_id`           | `string`  | \*(or)   | Номер РК в SAMBA                                 |
| `date_start`      | `string`  |          | Дата начала РК                                   |
| `date_end`        | `string`  |          | Дата окончания РК                                |
| `advance`         | `boolean` |          | Признак аванса                                   |
| `inn_contractor`  | `string`  | (or_1)   | ИНН юр.лица РК                                   |
| `kpp_contractor`  | `string`  |          | КПП юр.лица РК                                   |
| `uuid_contractor` | `string`  | (or_1)   | Уникальный идентификатор юр.лица РК              |
| `inn_finaladv`    | `string`  | (or_2)   | ИНН конечного рекламодателя                      |
| `kpp_finaladv`    | `string`  |          | КПП конечного рекламодателя                      |
| `uuid_finaladv`   | `string`  | (or_2)   | Уникальный идентификатор конечного рекламодателя |
| `number_contract` | `string`  | (or_3)   | Номер договора рекламодателя                     |
| `date_contract`   | `string`  | (or_3)   | Дата договора рекламодателя                      |
| `uuid_contract`   | `string`  | (or_3)   | Уникальный идентификатор договора рекламодателя  |

**Пример запроса JSON:**

```json
{
  "uuid": "c833b7d2-831a-11ef-816c-0050568a4702",
  "rk_id": "У191997",
  "inn_contractor": "3526019521",
  "kpp_contractor": "352601001",
  "advance": false,
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

| Parameter    | Type      | Required | Description                    |
| ------------ | --------- | -------- | ------------------------------ |
| `date_start` | `string`  |          | Период начала                  |
| `date_end`   | `string`  |          | Период окончания               |
| `date`       | `string`  |          | Дата счета                     |
| `advance`    | `boolean` |          | Счет на аванс                  |
| `uuid`       | `string`  |          | Уникальный идентификатор счета |
| `number`     | `string`  |          | Номер счета в 1С               |
| `uuid_rk`    | `string`  |          | Уникальный идентификатор РК    |
| `rk_id`      | `string`  |          | Номер РК в SAMBA               |
| `numberMF`   | `string`  |          | Внутренний номер РК в 1С       |

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
      "advance": true,
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
| `number`  | `string` | \* (or)  | Номер счета в 1С               |
| `date`    | `string` | \* (or)  | Дата счета                     |
| `uuid`    | `string` | \* (or)  | Уникальный идентификатор счета |
| `ext`     | `string` |          | pdf, docx, xlsx (pdf-default)  |

**Параметры ответа JSON:**

| Parameter | Type     | Required | Description          |
| --------- | -------- | -------- | -------------------- |
| `status`  | `string` |          | Статус ответа 1С     |
| `message` | `string` |          | Сообщение            |
| `result`  | 'struct' |          |                      |
| `name`    | `string` |          | Имя файла            |
| `ext`     | `string` |          | Расширение           |
| `data`    | `string` |          | Binary data (base64) |

**Пример ответа JSON:**

```json
{
  "status": "ok",
  "message": null,
  "result": {
    "name": "Счет УСУ082753/10 от 09.10.2024.pdf",
    "ext": "pdf",
    "data": "JVBERi0xLjcKJeLjz9MKMSAwIG9iago8PAovRmlsdGVyIC9GbGF0ZURlY29k"
  }
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
| `rk_id`   | `string` | \* (or)  | Номер РК в SAMBA            |
| `rk_uuid` | `string` | \* (or)  | Уникальный идентификатор РК |
| `amount`  | `string` | \*       | Сумма счета                 |

**Пример запроса JSON:**

```json
{
  "number": "TEST PO #11",
  "date": "20241008",
  "rk_id": "У191994",
  "amount": 60000
}
```

**Пример ответа JSON:**

```json
{
  "status": "ok",
  "message": "Счет на оплату покупателю успешно создан",
  "result": {
    "uuid": "59fed96e-8646-11ef-816c-0050568a4702",
    "number": "УЛУ085454/05",
    "date": "2024-10-09T16:57:01",
    "amount": 60000,
    "vat": 10000
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
| `rk_id`   | `string` | \* (or)  | Номер РК в SAMBA                    |
| `rk_uuid` | `string` | \* (or)  | Уникальный идентификатор РК         |
| `amount`  | `string` | \*       | Сумма счета                         |

**Пример запроса JSON:**

```json
{
  "uuid": "",
  "number": "TEST PO #11",
  "date": "20241008",
  "rk_id": "У191994",
  "amount": 60000
}
```

**Пример ответа JSON:**

```json
{
  "status": "ok",
  "message": "Счет на оплату покупателю успешно создан",
  "result": {
    "uuid": "59fed96e-8646-11ef-816c-0050568a4702",
    "number": "УЛУ085454/05",
    "date": "2024-10-09T16:57:01",
    "amount": 60000,
    "vat": 10000
  }
}
```

### 6.5 **Удалить счет на оплату покупателю**

**Метод:** `DELETE`

**URL:** `/po`

**Параметры JSON:**

| Parameter | Type     | Required | Description                         |
| --------- | -------- | -------- | ----------------------------------- |
| `uuid`    | `string` | (or)     | Уникальный идентификатор счета в 1С |
| `number`  | `string` | (or)     | Номер счета в 1С                    |
| `date`    | `string` | (or)     | Дата счета                          |

**Пример запроса JSON:**

```json
{
  "number": "УМУ083765/01",
  "date": "20240913"
}
```

**Пример ответа JSON:**

```json
{
  "status": "ok",
  "message": "Найдены документы по счету. Удаление невозможно",
  "result": [
    {
      "uuid": "c0fa0861-293a-47de-8d29-cb9f3fd69fbe",
      "number": "УМУ083765/01",
      "date": "2024-09-13T15:55:26",
      "uuid_doc": "70ba84ca-7599-11ef-8b6d-0050569f1124",
      "name_doc": "Платежное поручение входящее БМ000003417 от 17.09.2024 23:53:18"
    }
  ]
}
```

### 7. Оплаты и акты

### 7.1 **Оплата по счету покупателю**

**Метод:** `GET`

**URL:** `/payment/po`

**Параметры JSON:**

| Parameter | Type     | Required | Description                         |
| --------- | -------- | -------- | ----------------------------------- |
| `uuid`    | `string` | (or)     | Уникальный идентификатор счета в 1С |
| `number`  | `string` | (or)     | Номер счета в 1С                    |
| `date`    | `string` | (or)     | Дата счета                          |

**Параметры ответа JSON:**

| Parameter         | Type     | Required | Description                                                             |
| ----------------- | -------- | -------- | ----------------------------------------------------------------------- |
| `uuid`            | `string` |          | Уникальный идентификатор оплаты в 1С                                    |
| `doctype`         | `string` |          | Тип документа оплаты                                                    |
| `numberMF`        | `string` |          | Номер документа в МФ 1С                                                 |
| `status`          | `string` |          | Стутус оплаты в 1С(Оплачено, Не оплачено, Частично оплачено, Переплата) |
| `number`          | `string` |          | Номер платежного документа                                              |
| `date`            | `date`   |          | Дата платежного документа                                               |
| `paymentdate`     | `date`   |          | Фактическая дата оплаты                                                 |
| `amount`          | `number` |          | Сумма оплаты                                                            |
| `payment_details` | `string` |          | Назначение платежа                                                      |
| `uuid_po`         | `string` |          | Уникальный идентификатор счета в 1С                                     |
| `number_po`       | `string` |          | Номер счета в 1С                                                        |
| `date_po`         | `date`   |          | Дата счета в 1С                                                         |

**Пример ответа JSON:**

```json
{
  "status": "ok",
  "message": null,
  "result": [
    {
      "uuid": "d2844684-e040-11ee-bb0e-005056995bd0",
      "doctype": "Платежное поручение входящее",
      "numberMF": "0Л000000023",
      "status": "Частично оплачено",
      "number": "145",
      "date": "2024-03-11T00:00:00",
      "paymentdate": "2024-03-11T00:00:00",
      "amount": 68714.65,
      "payment_details": "ОПЛАТА ПО СЧЕТУ №УЛУ080367/01 ОТ 07.03.2024 Г. ЗА РАСХОДЫ ПО ВЫПОЛ. ПОРУЧ. ПО РАЗМЕЩ. РЕКЛАМНЫХ МАТЕРИАЛОВ ПО ДОГ.№3/2024/ОФ ОТ 07.03.2024 Г СУММА 68714-65 РУБ. В Т.Ч. НДС (20%) 11452-45",
      "uuid_po": "e43a0b45-080c-4f59-8aae-10b03f8cf776",
      "number_po": "УЛУ080367/01",
      "date_po": "2024-03-07T14:28:13"
    }
  ]
}
```

### 7.2 **Акты по рекламным кампаниям**

**Метод:** `GET`

**URL:** `/acts/rk`

**Параметры JSON:**

| Parameter | Type     | Required | Description                 |
| --------- | -------- | -------- | --------------------------- |
| `uuid`    | `string` | (or)     | Уникальный идентификатор РК |
| `rk_id`   | `string` | (or)     | Идентификатор РК в SAMBA    |

**Параметры ответа JSON:**

| Parameter           | Type     | Required | Description                   |
| ------------------- | -------- | -------- | ----------------------------- |
| `uuid_rk`           | `string` |          | Уникальный идентификатор РК   |
| `placements_amount` | `number` |          | Сумма размещений РК           |
| `acts_amount`       | `number` |          | Сумма актов по размещениям РК |

**Пример ответа JSON:**

```json
{
  "status": "ok",
  "message": null,
  "result": [
    {
      "uuid_rk": "62d2dda1-69fe-11ef-816a-005056b27e80",
      "placements_amount": 436704.38,
      "acts_amount": 218352.19
    }
  ]
}
```

### 7.3 **Детализация актов по рекламным кампаниям**

**Метод:** `GET`

**URL:** `/acts/rk/details`

**Параметры JSON:**

| Parameter | Type     | Required | Description                 |
| --------- | -------- | -------- | --------------------------- |
| `uuid`    | `string` | (or)     | Уникальный идентификатор РК |
| `rk_id`   | `string` | (or)     | Идентификатор РК в SAMBA    |

**Параметры ответа JSON:**

| Parameter           | Type     | Required | Description                 |
| ------------------- | -------- | -------- | --------------------------- |
| `uuid_rk`           | `string` |          | Уникальный идентификатор РК |
| `placements`        | `string` |          | Размещение                  |
| `placements_amount` | `number` |          | Сумма по размещению         |

| Parameter     | Type     | Required | Description                 |
| ------------- | -------- | -------- | --------------------------- |
| `uuid_rk`     | `string` |          | Уникальный идентификатор РК |
| `placements`  | `string` |          | Размещение                  |
| `document`    | `string` |          | Документ                    |
| `acts_amount` | `number` |          | Сумма акта по размещению    |

**Пример ответа JSON:**

```json
{
  "status": "ok",
  "message": null,
  "result": [
    [
      {
        "uuid_rk": "62d2dda1-69fe-11ef-816a-005056b27e80",
        "placement": "{8e9a83af-fe7e-4d4d-8642-ca6f01a395be}",
        "placement_amount": 358709.16
      },
      {
        "uuid_rk": "62d2dda1-69fe-11ef-816a-005056b27e80",
        "placement": "{081f755a-25ed-42cb-bdb2-a9faec51a725}",
        "placement_amount": 77995.22
      }
    ],
    [
      {
        "uuid_rk": "62d2dda1-69fe-11ef-816a-005056b27e80",
        "placement": "{8e9a83af-fe7e-4d4d-8642-ca6f01a395be}",
        "document": "Реализация товаров и услуг УМ0711/0001 от 07.11.2024 14:55:08",
        "act_amount": 179354.58
      },
      {
        "uuid_rk": "62d2dda1-69fe-11ef-816a-005056b27e80",
        "placement": "{8e9a83af-fe7e-4d4d-8642-ca6f01a395be}",
        "document": "Реализация товаров и услуг УМ3009/0204 от 30.09.2024 0:00:00",
        "act_amount": 179354.58
      },
      {
        "uuid_rk": "62d2dda1-69fe-11ef-816a-005056b27e80",
        "placement": "{081f755a-25ed-42cb-bdb2-a9faec51a725}",
        "document": "Реализация товаров и услуг УМ3108/0329 от 31.08.2024 0:00:00",
        "act_amount": 38997.61
      }
    ]
  ]
}
```
