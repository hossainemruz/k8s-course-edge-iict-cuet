# Sample Web Server

## Run Locally

You must have Go installed to run the sample server locally.

```bash
go run main.go
```

## Run using Docker

Build the docker image:

```bash
docker build -t sample-web-server:v1 . 
```

Run the docker image:
```bash
docker run -it sample-web-server:v1 -p 8080:8080
```


## Access the Server
```bash
curl 172.17.0.2:8080/hello
```
