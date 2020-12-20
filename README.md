# VkScrapper

API of a simple online shop

## Requirements
You need installed:


## Project setup 

1. `git clone https://github.com/Zaysevkun/DjangoECommerceAPI`
2.  `create POSTGRESQL db`
3. `cd DjangoECommerceAPI`
4. `python3 -m venv myvenv`
5. `source myvenv/bin/activate`
6. `pip install -r requirements.txt`
7. Add __.env__ file
##### EXAMPLE:
```
SECRET_KEY=qwerty123
DATABASE_URL=postgres://your_db_user_name:user_password@127.0.0.1:5432/your_db_name
ALLOWED_HOSTS=*
DEBUG=0
```
6. `python manage.py migrate`
7. `python manage.py collectstatic`
8. `python manage.py runserver 0.0.0.0:8000`

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
    "ID":198351038,
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

