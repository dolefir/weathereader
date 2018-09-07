weathereader-docker-compose
=========================

## Running application
- You need to ensure that you have correct one env file.
If you don't, then you need to create a file `.env.docker` with the next content:
```sh
SECRET_KEY=your secret key
APIKEY=your APIKEY openweathermap (for this example my key 5ce0bd41ca021e708f8907d2b04ae34e)
```
- Run the next command:
```sh
$ glide install
$ docker-compose build
$ docker-compose up
```
- Now you can use swagger: http://localhost:8081
