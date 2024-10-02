# API Документация

### Базовый URL
`https://1cplatform.mitcoms.ru/Dev/U_MediaFinance_for_testing/hs/samba`

## Точки доступа (Endpoints)

### 1. Конечные рекламодатели

### 1.1 **Получить список конечных рекламодателей.**

**Метод:** `GET`

**URL:** `/finaladv`

**Параметры:**

| Parameter          | Type      | Required | Description                               |
|--------------------|-----------|----------|-------------------------------------------|
| `inn`              | `string`  |         | ИНН конечного рекламодателя              |
| `kpp`              | `string`  |        | КПП конечного рекламодателя              |
| `name`              | `string`  |        | Наименование конечного рекламодателя              |
| `uuid`              | `string`  |       | Уникальный идентификатор конечного рекламодателя              |

**Пример ответа:**
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

| Parameter          | Type      | Required | Description                               |
|--------------------|-----------|----------|-------------------------------------------|
| `inn`              | `string`  | *        | ИНН конечного рекламодателя              |
| `kpp`              | `string`  |          | КПП конечного рекламодателя              |
| `name`             | `string`  | *        | Наименование конечного рекламодателя    |
| `fullName`         | `string`  |          | Полное наименование конечного рекламодателя |


**Пример запроса:**
```json
{
    "inn": "7717707438",
    "kpp": "771701001",
    "name": "Карат Рус Медиа ООО",
    "fullName": "ООО \"Карат Рус Медиа\""
}
```

**Пример ответа:**
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

| Parameter          | Type      | Required | Description                               |
|--------------------|-----------|----------|-------------------------------------------|
| `uuid`              | `string`  |        | ИНН конечного рекламодателя              |
| `inn`              | `string`  |         | ИНН конечного рекламодателя              |
| `kpp`              | `string`  |          | КПП конечного рекламодателя              |
| `name`             | `string`  |            | Наименование конечного рекламодателя     |
| `fullName`         | `string`  |          | Полное наименование конечного рекламодателя |


**Пример запроса:**
```json
{
    "inn": "7717707438",
    "kpp": "771701001",
    "name": "Карат Рус Медиа ООО",
    "fullName": "ООО \"Карат Рус Медиа\""
}
```

**Пример ответа:**
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

---

### 2. Customers (Юридические лица)

### 2.1 **Получить список юр. лиц**

**Метод:** `GET`

**URL:** `/customers`

**Параметры:**

| Parameter          | Type      | Required | Description                               |
|--------------------|-----------|----------|-------------------------------------------|
| `inn`              | `string`  |        | ИНН юр. лица              |
| `kpp`              | `string`  |        | КПП юр. лица              |
| `name`              | `string`  |        | Наименование юр. лица              |
| `uuid`              | `string`  |        | Уникальный идентификатор юр. лица              |

**Пример ответа:**
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

### 3. Brands

**Назначение:** Получить список брендов.

**Метод:** `GET`

**URL:** `/brands`

**Параметры JSON:**
- `uuid` (string): Уникальный идентификатор.
- `name` (string)*: Наименование бренда.
- `inn` (string): ИНН конечного рекламодателя.
- `kpp` (string): КПП конечного рекламодателя.
- `uuid_contractor` (string): Уникальный идентификатор юр. лица.

**Пример ответа:**
```json
{
    "status": "ok",
    "message": "success",
    "result": [
        {
            "uuid": "1dff21fa-6a02-11ef-816a-005056b27e80",
            "name": "SAMBA (ЛТ)",
            "fullName": "ТРАДИЦИЯ ООО",
            "inn": "7810751880",
            "kpp": "781001001",
            "uuid_contractor": "a787dc66-52fb-11ef-8168-005056b27e80"
        }
    ]
}
```

**Назначение:** Создать новый бренд.

**Метод:** `POST`

**URL:** `/brands`

**Параметры JSON:**
- `name` (string): Наименование бренда.
- `inn` (string): ИНН юридического лица.
- `kpp` (string): КПП юридического лица.
- `uuid_contractor` (string): Уникальный идентификатор юр. лица.

**Пример запроса:**
```json
{
    "name": "My test brand",
    "inn": "381508339299"
}
```

**Пример ответа:**
```json
{
    "status": "ok",
    "message": "Бренд успешно создан.",
    "result": {
        "uuid": "2be267b3-a6bf-408b-b385-3bda0a226725",
        "name": "My test brand"
    }
}
```

---

### 4. RK

**Назначение:** Создание рекламной кампании.

**Метод:** `POST`

**URL:** `/rk`

**Параметры JSON:**
- `RK_ID` (string)*: Номер рекламной кампании.
- `AdvName` (string)*: Наименование рекламодателя.
- `AdvINN` (string)*: ИНН рекламодателя.
- `AdvKPP` (string)*: КПП рекламодателя.
- `BegDate` (string)*: Дата начала РК (формат: `dd.mm.yyyy`).
- `EndDate` (string)*: Дата окончания РК (формат: `dd.mm.yyyy`).
- `LastAdvName` (string): Наименование конечного рекламодателя.
- `LastAdvINN` (string): ИНН конечного рекламодателя.
- `LastAdvKPP` (string): КПП конечного рекламодателя.
- `LastAdvAgrNumber` (string): Номер договора конечного рекламодателя.
- `LastAdvAgrDate` (string): Дата договора конечного рекламодателя.

**Пример запроса:**
```json
{
    "RK_ID": "12345",
    "AdvName": "Sample Advertiser",
    "AdvINN": "7701234567",
    "AdvKPP": "770101001",
    "BegDate": "01.01.2024",
    "EndDate": "31.12.2024",
    "LastAdvName": "ООО Конечный Рекламодатель",
    "LastAdvINN": "7707654321",
    "LastAdvKPP": "770701001",
    "LastAdvAgrNumber": "Д-12345",
    "LastAdvAgrDate": "01.01.2024"
}
```

**Пример ответа:**
```json
{
    "status": "ok",
    "message": "",
    "result": null
}
```

---

### 5. PO

**Назначение:** Получить данные о PO (заказах).

**Метод:** `GET`

**URL:** `/po`

**Параметры JSON:** Нет.

**Пример ответа:**
```json
{
    "status": "ok",
    "message": null,
    "result": [
        {
            "uuid": "7399e00d-37e7-4eef-b87a-d16b7d421155",
            "number": "УЛУ080366/01",
            "date": "2024-03-07T14:37:55",
            "amount": 348183.61,
            "uuid_rk": "c10d0693-291d-4542-b1c8-293458b6fa98",
            "rk_id": "SAMBA|Nabiullin|Mar - 2024",
            "numberMF": "У081095"
        }
    ]
}
```

## Примечания

Все запросы должны отправляться с использованием заголовка `Authorization` с токеном авторизации для аутентификации.