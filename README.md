# VkScrapper

API of a simple online shop

## Requirements
You need installed:

1. vowel wabbit
2. postgreSQL
3. go 1.15

## Project setup 

1. `git clone https://github.com/Limbus-hack/vk-scrapper.git`
2.  `create POSTGRESQL db`
3. `cd vk-scrapper`
4. `psql -U username -d myDataBase -a -f init.sql`
5. `make`
7. Add __.env__ file
##### EXAMPLE:
```
POSTGRES_DB_STR=postgresql://postgres:postgres@127.0.0.1:5432/your_db_name
VK_APP_ID=example
VK_PRIVATE_KEY=example
VK_SERVICE_TOKEN=example
VK_CLIENT_TOKEN=example
```
8. `go run github.com/code7unner/vk-scrapper`

## END POINTS

| *URL* | *Method*|*Description*|
|-------|---------|-------------|
| `prediction/` | `GET` | Retrieve Cached Predictions|

Requiered query params:
```
city=<integer>
school=<string>
```

GET response body:
```
[
  {
    "id":198351038,
    "Name":"Alexey",
    "LastName":"Zakirov",
    "Region":60,
    "Prediction":2,
    "Status":"new"
  },
]
```
| *URL* | *Method*|*Description*|
|-------|---------|-------------|
| `status/` | `PATCH` | endpoint for updating user status|

PATCH request body:
```
{"id", "status"}
```

