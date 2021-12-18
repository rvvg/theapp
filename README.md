# Project App
Simple web-application written in golang

# Description
This app only does two things, but it does them perfectly. It can calculate the square of an integer and it can also block the user. Yes, it can block you.

### Dependencies
* PC or Mac
* Go
* PostgreSQL
* Docker, docker-compose

### Using
How to run the App locally
```
$ docker-compose start
Starting db      ... done
Starting app     ... done
Starting adminer ... done
```
How to calculate square, e.g. 777
```
$ curl localhost:3000/?n=777
603729
```
How to get blocked? The App will response with 444 status code
```
$ curl -i localhost:3000/blacklisted
HTTP/1.1 444 Unknown Status Code
Date: Sat, 18 Dec 2021 05:41:34 GMT
Content-Type: text/plain; charset=utf-8
Content-Length: 11

Blacklisted
```

## Run the App without Docker
You should specify database credentials if you want to run the App outside the docker-compose stack
```
POSTGRES_HOST=MY_DB_HOST \
POSTGRES_PORT=5432 \
POSTGRES_PASSWORD=MY_PASSWORD \
POSTGRES_USER=MY_USER \
POSTGRES_DB=MY_DB \
make all
```
or put variables to file, e.g. ```my_vars.env``` and import them:
```
$ .  my_vars.env
$ make all
```