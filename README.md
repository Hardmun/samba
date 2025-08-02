# SAMBA + 1C API Documentation

This document outlines the API endpoints for interacting with the **SAMBA + 1C** system, a platform for managing advertising campaigns.

## Base URL

```
https://1cplatform.mitcoms.ru/Dev/U_MediaFinance_for_testing/hs/samba
```

## Endpoints 1

### 1. Final Advertisers

#### List Final Advertisers
- **Method**: `GET`
- **URL**: `/finaladv`
- **Parameters**: `inn`, `kpp`, `name`, `uuid`

**Example Response**:
```json
{
    "status": "ok",
    "message": "success",
    "result": [
        {
            "uuid": "301ca510-80b8-11ef-816c-0050568a4702",
            "name": "Карат Рус Медиа ООО FOR TEST",
            "fullName": "Карат Рус Медиа ООО FOR TEST",
            "inn": "7717707438",
            "kpp": "771701002"
        }
    ]
}
```

#### Create Final Advertiser
- **Method**: `POST`
- **URL**: `/finaladv`
- **Parameters**: `inn`, `kpp`, `name`, `fullName`

**Example Request**:
```json
{
    "inn": "7717707438",
    "kpp": "771701001",
    "name": "Карат Рус Медиа ООО",
    "fullName": "ООО \"Карат Рус Медиа\""
}
```

#### Update Final Advertiser
- **Method**: `PUT`
- **URL**: `/finaladv`
- **Parameters**: `uuid`, `inn`, `kpp`, `name`, `fullName`

---

### 2. Final Advertiser Contracts

#### List Final Advertiser Contracts
- **Method**: `GET`
- **URL**: `/finalcontracts`
- **Parameters**: `uuid`, `number`, `date`, `inn_finaladv`, `kpp_finaladv`, `uuid_finaladv`

#### Create Final Advertiser Contract
- **Method**: `POST`
- **URL**: `/finalcontracts`
- **Parameters**: `number`, `date`, `inn_finaladv`, `kpp_finaladv`, `uuid_finaladv`, `inn_contractor`, `kpp_contractor`, `uuid_contractor`

#### Update Final Advertiser Contract
- **Method**: `PUT`
- **URL**: `/finalcontracts`
- **Parameters**: `uuid`, `number`, `date`, `inn_finaladv`, `kpp_finaladv`, `uuid_finaladv`, `inn_contractor`, `kpp_contractor`, `uuid_contractor`

---

### 3. Legal Entities (Customers)

#### List Legal Entities
- **Method**: `GET`
- **URL**: `/customers`
- **Parameters**: `inn`, `kpp`, `name`, `uuid`

---

### 4. Brands

#### List Brands
- **Method**: `GET`
- **URL**: `/brands`
- **Parameters**: `uuid`, `name`, `inn`, `kpp`, `uuid_contractor`

#### Create Brand
- **Method**: `POST`
- **URL**: `/brands`
- **Parameters**: `name`, `inn_contractor`, `kpp_contractor`, `uuid_contractor`

#### Update Brand
- **Method**: `PUT`
- **URL**: `/brands`
- **Parameters**: `name`, `uuid`

---

### 5. Advertising Campaigns (RK)

#### List Advertising Campaigns
- **Method**: `GET`
- **URL**: `/rk`
- **Parameters**: `uuid`, `name`, `rk_id`, `numberMF`, `inn_contractor`, `kpp_contractor`, `uuid_contractor`

#### Create Advertising Campaign
- **Method**: `POST`
- **URL**: `/rk`
- **Parameters**: `rk_id`, `date_start`, `date_end`, `inn_contractor`, `kpp_contractor`, `uuid_contractor`, `inn_finaladv`, `kpp_finaladv`, `uuid_finaladv`, `number_contract`, `date_contract`, `uuid_contract`

#### Update Advertising Campaign
- **Method**: `PUT`
- **URL**: `/rk`
- **Parameters**: `uuid`, `rk_id`, `date_start`, `date_end`, `inn_contractor`, `kpp_contractor`, `uuid_contractor`, `inn_finaladv`, `kpp_finaladv`, `uuid_finaladv`, `number_contract`, `date_contract`, `uuid_contract`

---

### 6. Purchase Orders (PO)

#### List Purchase Orders
- **Method**: `GET`
- **URL**: `/po`
- **Parameters**: `date_start`, `date_end`, `date`, `uuid`, `number`, `uuid_rk`, `rk_id`, `numberMF`

#### Get PO Print Form
- **Method**: `GET`
- **URL**: `/files/po`
- **Parameters**: `number`, `date`, `uuid`

#### Create Purchase Order
- **Method**: `POST`
- **URL**: `/po`
- **Parameters**: `number`, `date`, `rk_id`, `rk_uuid`, `amount`

#### Update Purchase Order
- **Method**: `PUT`
- **URL**: `/po`
- **Parameters**: `uuid`, `number`, `date`, `rk_id`, `rk_uuid`, `amount`

---

## Authentication

The SAMBA + 1C API uses [authentication mechanism here, e.g., API keys, OAuth, etc.]. Make sure to include the necessary authentication information in your requests.