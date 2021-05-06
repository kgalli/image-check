# Image Check

WORK IN PROGRESS

Continuously check for updates for a defined set of docker images on DockerHub.


## Development

```
# run postgres database
docker run -it --rm -e POSTGRES_PASSWORD=gorm -e POSTGRES_USER=gorm -e POSTGRES_DB=gorm -p 9970:5432 postgres
```

```
# connect to postgres database via console
docker run --rm -it --network=host postgres psql -h 127.0.0.1 -p9970 -U gorm
```

