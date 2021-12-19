# Project App
Simple web-application written in Go

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
Starting tests   ... done
Starting adminer ... done
```

Also you can run tests
```
$ docker-compose up tests
Docker Compose is now in the Docker CLI, try `docker compose up`

app_db_1 is up-to-date
app_app_1 is up-to-date
Starting app_tests_1 ... done
Attaching to app_tests_1
tests_1    | PASS
tests_1    | ok         app     0.006s
app_tests_1 exited with code 0
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
## Deploy to Kubernetes
The App uses PostgreSQL database for storage. If you want to run the App locally in minikube you can use Kubegres operator to run local PostgreSQL cluster. Please find the example below.

Install Kubegres operator. The installation needs to be done once.
```
kubectl apply -f https://raw.githubusercontent.com/reactive-tech/kubegres/v1.15/kubegres.yaml
```
Create namespace for our App
```
$ kubectl create namespace app --dry-run=client -o yaml | kubectl apply -f -
```
Edit configuration files in ```kubegres``` directory according your needs and deploy your local PostgreSQL cluster:
```
$ kubectl apply -f kubegres/
```
Edit database section in ```helm/values.yaml``` if you have changed db, user or password in kubergres conf files above.

Deploy the App using Helm.
First of all we need to reuse Docker daemon from minikube
```
eval $(minikube docker-env)
```
Build App image
```
docker-compose build app
```
And finally deploy the App to minikube
```
$ helm install -n app app helm/
Release "app" has been upgraded. Happy Helming!
NAME: app
LAST DEPLOYED: Sun Dec 18 16:39:36 2021
NAMESPACE: app
STATUS: deployed
REVISION: 1
NOTES:
1. Get the application URL by running these commands:
  http://host.test/
```

## Accessing the App
In order to make the App reachable from your local curl or browser please follow the link ```https://minikube.sigs.k8s.io/docs/handbook/addons/ingress-dns/```