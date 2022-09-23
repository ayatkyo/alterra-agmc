# Alterra AGMC (Day 7) - Dockerize Application

## Tasks

- Install docker & docker compose
- Create dockerfile
- Clone your code and integrate to your dockerfile
- Build your container
- Push the image to docker registry
- Deploy in yout local machine
- Try DockerCompose to run go app + mysql (opt.)
- try to Deploy in the Cloud (opt.)


## Image on Docker Registery

[https://hub.docker.com/r/ayatkyo/agmc-ayat](https://hub.docker.com/r/ayatkyo/agmc-ayat).

## How to run in local

Copy `.env.example` as `.env`, then fill to match your environment

### Building image

Run this command to build docker image.

```bash
docker build -t {IMAGE NAME} .
# for example
docker build -t ayatkyo/agmc-ayat .
```

### Running

Make sure the `.env` file has been created.

```bash
docker run --env-file .\.env -p 15000:15000 --name local-agmc-ayat {IMAGE NAME}
# for example
docker run --env-file .\.env -p 15000:15000 --name local-agmc-ayat ayatkyo/agmc-ayat:latest
```

Then visit `http://localhost:15000/status`.

## Using Docker Compose

### Configure env

To configure environment variable, open `docker-compose.yaml` and modify `environment` of `api` service.
### Docker Compose using local Dockerfile

Run following command:

```bash
docker compose up --build
```

Then visit `http://localhost:15000/status`.

### Docker Compose using image from docker hub

Run following command:

```bash
docker compose -f docker-compose.hub.yaml up
```

Then visit `http://localhost:15000/status`.