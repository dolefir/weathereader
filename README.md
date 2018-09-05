weathereader-docker-compose
=========================

Running weathereader development environment

## Prepare environment
Create a file `.env.docker` inside the file
```sh
SECRET_KEY=you secret key
APIKEY=you APIKEY openweathermap(for this example my key 5ce0bd41ca021e708f8907d2b04ae34e)
```

## Setup
```sh
glide install
docker-compose build
```


## Local run
```sh
docker-compose up
```

Look swagger API `http://localhost:8081`
