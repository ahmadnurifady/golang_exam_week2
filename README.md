# PHINCON GOLANG EXAM WEEK 3

---

## API Spec

### Event

#### Create Event

Request :

- Method : `POST`
- Endpoint : `/api/v1/event/create`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
    "id": "EVENT-004",
    "event_name": "Concerto In The House",
    "location": "Jakarta",
    "date": "26 Februari 2026",
    "ticket": [
        {
            "type": "VIP",
            "price": 5000,
            "stock": 10
        },
        {
            "type": "CAT 1",
            "price": 250,
            "stock": 100
        }
    ]
}
```

Response :

```json
{
    "statuscode": 201,
    "message": "SUCCESS - CREATE - EVENT",
    "data": {
        "id": "EVENT-004",
        "event_name": "Concerto In The House",
        "location": "Jakarta",
        "date": "26 Februari 2026",
        "ticket": [
            {
                "id": "a91b594d-877f-45ec-9667-6912b1537e2b",
                "type": "VIP",
                "price": 5000,
                "stock": 10,
                "event_name": "Concerto In The House"
            },
            {
                "id": "1608251d-5d40-4509-8122-ffb1df318c0c",
                "type": "CAT 1",
                "price": 250,
                "stock": 100,
                "event_name": "Concerto In The House"
            }
        ]
    }
}
```




#### Find All Event

Request :

- Method : `GET`
- Endpoint : `/api/v1/event/find-all`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
```

Response :

```json
{
    "statuscode": 200,
    "message": "SUCCESS - FIND ALL - EVENT",
    "data": [
        {
            "id": "EVENT-001",
            "event_name": "Konser Music Van Priuck",
            "location": "Warakas, Tanjung Priuk, Jakarta Utara",
            "date": "25 Februari 2025",
            "ticket": [
                {
                    "id": "f12bfa2f-5285-4442-837a-9c62c6f8d7df",
                    "type": "VIP",
                    "price": 5000,
                    "stock": 4,
                    "event_name": "Konser Music Van Priuck"
                },
                {
                    "id": "9b54637e-07a4-4ab5-a92a-e0ee974b7bc3",
                    "type": "CAT 1",
                    "price": 250,
                    "stock": 95,
                    "event_name": "Konser Music Van Priuck"
                }
            ]
        },
        {
            "id": "EVENT-002",
            "event_name": "Music in The Cocass",
            "location": "Mall Cassablanca",
            "date": "26 Februari 2025",
            "ticket": [
                {
                    "id": "efbe78e7-db3d-4580-9d1d-4abba70d2e3f",
                    "type": "VIP",
                    "price": 5000,
                    "stock": 8,
                    "event_name": "Music in The Cocass"
                },
                {
                    "id": "411662f1-0862-4abf-b3f6-ecdaacf7d311",
                    "type": "CAT 1",
                    "price": 250,
                    "stock": 98,
                    "event_name": "Music in The Cocass"
                }
            ]
        },
        {
            "id": "EVENT-003",
            "event_name": "Music in The Cocasssssss",
            "location": "Mall Cassablancaaaaa",
            "date": "26 Februari 2025",
            "ticket": [
                {
                    "id": "7546f847-3291-4e10-b9e9-9a0cb90553dd",
                    "type": "VIP",
                    "price": 5000,
                    "stock": 6,
                    "event_name": "Music in The Cocasssssss"
                },
                {
                    "id": "1695d3de-7306-407f-bbef-a02d3549775a",
                    "type": "CAT 1",
                    "price": 250,
                    "stock": 96,
                    "event_name": "Music in The Cocasssssss"
                }
            ]
        },
        {
            "id": "EVENT-004",
            "event_name": "Concerto In The House",
            "location": "Jakarta",
            "date": "26 Februari 2026",
            "ticket": [
                {
                    "id": "a91b594d-877f-45ec-9667-6912b1537e2b",
                    "type": "VIP",
                    "price": 5000,
                    "stock": 10,
                    "event_name": "Concerto In The House"
                },
                {
                    "id": "1608251d-5d40-4509-8122-ffb1df318c0c",
                    "type": "CAT 1",
                    "price": 250,
                    "stock": 100,
                    "event_name": "Concerto In The House"
                }
            ]
        }
    ]
}
```


#### Find Event By Id

Request :

- Method : `GET`
- Endpoint : `/api/v1/event/find-by-id?id=`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
```


Response :

```json
{
    "statuscode": 200,
    "message": "SUCCESS - FIND BY ID - EVENT",
    "data": {
        "id": "EVENT-004",
        "event_name": "Concerto In The House",
        "location": "Jakarta",
        "date": "26 Februari 2026",
        "ticket": [
            {
                "id": "a91b594d-877f-45ec-9667-6912b1537e2b",
                "type": "VIP",
                "price": 5000,
                "stock": 10,
                "event_name": "Concerto In The House"
            },
            {
                "id": "1608251d-5d40-4509-8122-ffb1df318c0c",
                "type": "CAT 1",
                "price": 250,
                "stock": 100,
                "event_name": "Concerto In The House"
            }
        ]
    }
}
```




### Transaction

#### Create Transaction

- Method : `POST`
- Endpoint : `/api/v1/tx/create`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
    "user_id": "USR-002",
    "event_id": "EVENT-001",
    "ticket_type": "CAT 1"
}
```


Response : 

```json
{
    "statuscode": 201,
    "message": "SUCCESS-CREATE-TX",
    "data": {
        "id": "3f826ed1-5c78-460d-9d30-44bac7300b7a",
        "user": {
            "id": "USR-002",
            "name": "fandy"
        },
        "event": {
            "id": "EVENT-001",
            "event_name": "Konser Music Van Priuck",
            "location": "Warakas, Tanjung Priuk, Jakarta Utara",
            "date": "25 Februari 2025"
        },
        "ticket": {
            "id": "9b54637e-07a4-4ab5-a92a-e0ee974b7bc3",
            "type": "CAT 1",
            "price": 250,
            "stock": 95,
            "event_name": "Konser Music Van Priuck"
        },
        "create_at": "2024-08-03T21:53:07.5667189+07:00",
        "update_at": "2024-08-03T21:53:07.5667189+07:00"
    }
}
```


#### Create Many Transaction

- Method : `POST`
- Endpoint : `/api/v1/tx/create-many-tx`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
    "user_id": "USR-003",
    "request_event": [
        {
            "event_id": "EVENT-001",
            "ticket": [
                {
                    "ticket_type": "VIP",
                    "stock_request": 2
                },
                {
                    "ticket_type": "CAT 1",
                    "stock_request": 1
                }
            ]
        },
        {
            "event_id": "EVENT-002",
            "ticket": [
                {
                    "ticket_type": "VIP",
                    "stock_request": 2
                },
                {
                    "ticket_type": "CAT 1",
                    "stock_request": 2
                }
            ]
        }
    ]
}
```

Response : 

```json
{
    "statuscode": 201,
    "message": "SUCCESS-CREATE-TX",
    "data": [
        {
            "id": "041c59b5-a6d6-4df7-9fe3-10113fc397cb",
            "user": {
                "id": "USR-003",
                "name": "jhon doe"
            },
            "event": {
                "id": "EVENT-001",
                "event_name": "Konser Music Van Priuck",
                "location": "Warakas, Tanjung Priuk, Jakarta Utara",
                "date": "25 Februari 2025"
            },
            "ticket": {
                "id": "f12bfa2f-5285-4442-837a-9c62c6f8d7df",
                "type": "VIP",
                "price": 5000,
                "stock": 6,
                "event_name": "Konser Music Van Priuck"
            },
            "create_at": "2024-08-03T18:52:05.4968736+07:00",
            "update_at": "2024-08-03T18:52:05.4968736+07:00"
        },
        {
            "id": "c3ccab89-f265-48bb-a426-9b28cd25d27a",
            "user": {
                "id": "USR-003",
                "name": "jhon doe"
            },
            "event": {
                "id": "EVENT-001",
                "event_name": "Konser Music Van Priuck",
                "location": "Warakas, Tanjung Priuk, Jakarta Utara",
                "date": "25 Februari 2025"
            },
            "ticket": {
                "id": "f12bfa2f-5285-4442-837a-9c62c6f8d7df",
                "type": "VIP",
                "price": 5000,
                "stock": 6,
                "event_name": "Konser Music Van Priuck"
            },
            "create_at": "2024-08-03T18:52:05.5486317+07:00",
            "update_at": "2024-08-03T18:52:05.5486317+07:00"
        },
        {
            "id": "2774452f-170b-4268-a39d-49153f00583e",
            "user": {
                "id": "USR-003",
                "name": "jhon doe"
            },
            "event": {
                "id": "EVENT-001",
                "event_name": "Konser Music Van Priuck",
                "location": "Warakas, Tanjung Priuk, Jakarta Utara",
                "date": "25 Februari 2025"
            },
            "ticket": {
                "id": "9b54637e-07a4-4ab5-a92a-e0ee974b7bc3",
                "type": "CAT 1",
                "price": 250,
                "stock": 96,
                "event_name": "Konser Music Van Priuck"
            },
            "create_at": "2024-08-03T18:52:05.5550943+07:00",
            "update_at": "2024-08-03T18:52:05.5550943+07:00"
        },
        {
            "id": "6dda3b77-d6e9-45be-a118-eb4e78235a70",
            "user": {
                "id": "USR-003",
                "name": "jhon doe"
            },
            "event": {
                "id": "EVENT-002",
                "event_name": "Music in The Cocass",
                "location": "Mall Cassablanca",
                "date": "26 Februari 2025"
            },
            "ticket": {
                "id": "efbe78e7-db3d-4580-9d1d-4abba70d2e3f",
                "type": "VIP",
                "price": 5000,
                "stock": 10,
                "event_name": "Music in The Cocass"
            },
            "create_at": "2024-08-03T18:52:05.5596699+07:00",
            "update_at": "2024-08-03T18:52:05.5596699+07:00"
        },
        {
            "id": "30d8a0c2-96ff-4e96-b364-546d6da19af5",
            "user": {
                "id": "USR-003",
                "name": "jhon doe"
            },
            "event": {
                "id": "EVENT-002",
                "event_name": "Music in The Cocass",
                "location": "Mall Cassablanca",
                "date": "26 Februari 2025"
            },
            "ticket": {
                "id": "efbe78e7-db3d-4580-9d1d-4abba70d2e3f",
                "type": "VIP",
                "price": 5000,
                "stock": 10,
                "event_name": "Music in The Cocass"
            },
            "create_at": "2024-08-03T18:52:05.5610677+07:00",
            "update_at": "2024-08-03T18:52:05.5610677+07:00"
        },
        {
            "id": "86568d12-3cf4-4cac-90ab-df35c80e4cd3",
            "user": {
                "id": "USR-003",
                "name": "jhon doe"
            },
            "event": {
                "id": "EVENT-002",
                "event_name": "Music in The Cocass",
                "location": "Mall Cassablanca",
                "date": "26 Februari 2025"
            },
            "ticket": {
                "id": "411662f1-0862-4abf-b3f6-ecdaacf7d311",
                "type": "CAT 1",
                "price": 250,
                "stock": 100,
                "event_name": "Music in The Cocass"
            },
            "create_at": "2024-08-03T18:52:05.56717+07:00",
            "update_at": "2024-08-03T18:52:05.56717+07:00"
        },
        {
            "id": "1978119b-371e-4209-86ab-dafea9a903ed",
            "user": {
                "id": "USR-003",
                "name": "jhon doe"
            },
            "event": {
                "id": "EVENT-002",
                "event_name": "Music in The Cocass",
                "location": "Mall Cassablanca",
                "date": "26 Februari 2025"
            },
            "ticket": {
                "id": "411662f1-0862-4abf-b3f6-ecdaacf7d311",
                "type": "CAT 1",
                "price": 250,
                "stock": 100,
                "event_name": "Music in The Cocass"
            },
            "create_at": "2024-08-03T18:52:05.5700192+07:00",
            "update_at": "2024-08-03T18:52:05.5700192+07:00"
        }
    ]
}
```



#### Find All Transaction

- Method : `GET`
- Endpoint : `/api/v1/tx/find-all`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
```


Response:

```json
{
    "statuscode": 200,
    "message": "SUCCESS-FIND ALL-TX",
    "data": [
          {
            "id": "17545348-fb49-4a26-bad5-f3ec5825111e",
            "user": {
                "id": "USR-001",
                "name": "alexander"
            },
            "event": {
                "id": "EVENT-001",
                "event_name": "Konser Music Van Priuck",
                "location": "Warakas, Tanjung Priuk, Jakarta Utara",
                "date": "25 Februari 2025"
            },
            "ticket": {
                "id": "f12bfa2f-5285-4442-837a-9c62c6f8d7df",
                "type": "VIP",
                "price": 5000,
                "stock": 4,
                "event_name": "Konser Music Van Priuck"
            },
            "create_at": "2024-07-31T16:33:43.943837Z",
            "update_at": "2024-07-31T16:33:43.943837Z"
        },
        {
            "id": "a919555a-c6d3-4fd0-aa52-40401bbc41e4",
            "user": {
                "id": "USR-002",
                "name": "fandy"
            },
            "event": {
                "id": "EVENT-001",
                "event_name": "Konser Music Van Priuck",
                "location": "Warakas, Tanjung Priuk, Jakarta Utara",
                "date": "25 Februari 2025"
            },
            "ticket": {
                "id": "9b54637e-07a4-4ab5-a92a-e0ee974b7bc3",
                "type": "CAT 1",
                "price": 250,
                "stock": 94,
                "event_name": "Konser Music Van Priuck"
            },
            "create_at": "2024-08-01T11:00:19.864129Z",
            "update_at": "2024-08-01T11:00:19.864129Z"
        },
        {
            "id": "ee2c4d8d-b556-4e6a-a33e-797ddcba718f",
            "user": {
                "id": "USR-003",
                "name": "jhon doe"
            },
            "event": {
                "id": "EVENT-001",
                "event_name": "Konser Music Van Priuck",
                "location": "Warakas, Tanjung Priuk, Jakarta Utara",
                "date": "25 Februari 2025"
            },
            "ticket": {
                "id": "9b54637e-07a4-4ab5-a92a-e0ee974b7bc3",
                "type": "CAT 1",
                "price": 250,
                "stock": 94,
                "event_name": "Konser Music Van Priuck"
            },
            "create_at": "2024-08-01T14:05:59.354792Z",
            "update_at": "2024-08-01T14:05:59.354792Z"
        },
        {
            "id": "2622aa4c-394b-4280-8e24-e88335a19054",
            "user": {
                "id": "USR-003",
                "name": "jhon doe"
            },
            "event": {
                "id": "EVENT-001",
                "event_name": "Konser Music Van Priuck",
                "location": "Warakas, Tanjung Priuk, Jakarta Utara",
                "date": "25 Februari 2025"
            },
            "ticket": {
                "id": "9b54637e-07a4-4ab5-a92a-e0ee974b7bc3",
                "type": "CAT 1",
                "price": 250,
                "stock": 94,
                "event_name": "Konser Music Van Priuck"
            },
            "create_at": "2024-08-02T11:16:18.683507Z",
            "update_at": "2024-08-02T11:16:18.683507Z"
        }
    ]
}
```