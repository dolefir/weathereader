weathereader-docker-compose
=========================

## Running application
- You need to ensure that you have correct one env file.
If you don't, then you need to create a file `.env.docker` with the next content:
```sh
SECRET_KEY=your secret key
APIKEY=your APIKEY openweathermap (for this example my key 5ce0bd41ca021e708f8907d2b04ae34e)
```
- Use docker-compose: `$ docker-compose up `
- Now you can use swagger: http://localhost:8081

## Development
For development, your IDE probably would need every dependency that was used inside this project.
You can download them with glide. Run the next command:`$ glide install.`