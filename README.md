# Go Web Example

An example project about creating web application and REST-API using Go and Gorilla framework. I created this project to serve as dictionary/example if i need to create web application or REST-API using Go (Gorilla).

## Using This Project

To use this project :

```bash
git pull {githubtothisproject}
```

create a .env in /.env and fill it with your preferred secret-key. That key used for Gorilla Session Management and JWT (JSON WEB TOKEN).

```bash
GORILLA_SESSION_SECRET=
JWT_SECRET=
```

To run the application :

```bash
go run main.go
```

This project should run at port 8081 by default.

## Project Utility

As a web application it can :

* Render / Serve HTML
* Process Form
* Login Process and Session Management

As a REST-API it can :

* Serve data with JSON
* Send POST Request
* Authentication With JWT

## Testing The Web Application Utility

To try the login process , after the application up and running , you can go to **login**.
You can find account and password of this project in : **/data-user.json**. You can modified it and create your own user data.

## Testing The REST API

Once application up and running , you can try a GET request:

```bash
curl --location 'http://localhost:8081/api/v1/product' \
--data ''
```

To try a POST request:

```bash
curl --location 'http://localhost:8081/api/v1/product' \
--header 'Content-Type: application/json' \
--data '{
    "newProductName": "Silverado 1700",
    "newProductPrice": 176320
}'
```

## Try The Authentication API

First you need to send a login reqeust (using data from data-user.json):

```bash
curl --location 'http://localhost:8081/api/v1/login' \
--header 'Content-Type: application/json' \
--data '{
    "username" : "osheere3",
    "password" : "1234"
}
'
```

If success , it will send back JWT as response : 

```bash
{
    "status": 200,
    "message": "OK",
    "payload": {
        "message": "success",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODQ4MDIxODYsImlhdCI6MTY4NDc5ODU4Niwic3ViIjoib3NoZWVyZTMifQ.sKRAW_l0DMUmzODUAqaCxHPrxt7DZ60PRDzVOSQ6bws"
    }
}
```

Now you can use the token as **Bearer** token authentication, example: 

```Bash
curl --location 'http://localhost:8081/api/v1/me' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODQ4MDIxODYsImlhdCI6MTY4NDc5ODU4Niwic3ViIjoib3NoZWVyZTMifQ.sKRAW_l0DMUmzODUAqaCxHPrxt7DZ60PRDzVOSQ6bws'
```

That request was send to API that will validate the token and send a payload of user that token refer to. Below is the response of the request:

```bash
{
    "status": 200,
    "message": "OK",
    "payload": {
        "username": "osheere3",
        "displayName": "Orbadiah Sheere",
        "email": "osheere3@stumbleupon.com",
        "password": "1234"
    }
}
```

