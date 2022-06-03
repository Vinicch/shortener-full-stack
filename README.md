# Shortener
Solution for shortening URLs using aliases. With this it's possible to:
- Shorten URLs
- Use a shortened URL to redirect to the original one
- Check the top 10 most visited URLs

## How It Works
![](docs/Shorten%20URL.png)

![](docs/Retrieve%20URL.png)

![](docs/Get%20Most%20Visited%20URLs.png)

## Project Setup
### Local
You need to create a .env file. All the environment variables are available in the .env.example
file located in the root of the project. You can also use its current values to ease the process.

The database is inside a Docker container. There's a couple of ways to run it using the `docker-compose.yaml` provided:
#### Docker Compose (for Compose V1 use "docker-compose" instead)
```shell
docker compose up -d
```
#### GNU Make
```shell
make up
```

You also need to run the database migration scripts. To do that you can either run the `db/schema.sql`
file manually or use [dbmate](https://github.com/amacneil/dbmate) to run the scritps inside `db/migrations/`
from the root folder (which is why there's a DBMATE_MIGRATIONS_TABLE environment variable).

Finally, you can run the API:

#### Development build
```shell
go run main.go
```
#### Production build
```shell
go build
./shortener-go
```

### Docker
You can also run the API from within a Docker container:
#### Docker Compose
```shell
docker compose --profile release up
```
#### GNU Make
```shell
make deploy
```

To check the logs:
#### Docker Compose
```shell
docker compose logs -tf
```
#### GNU Make
```shell
make logs
```

To stop the containers:
#### Docker Compose
```shell
docker compose down --remove-orphans
```
#### GNU Make
```shell
make down
```

## Tests
To run both unit and integration tests run:
```shell
go test ./...
```

## Client
The solution comes with a client to consume the API. To use it, run the following commands
inside `client/`:
#### Development build
```shell
npm install
npm run dev
```
#### Production build
```shell
npm install
npm run build
npm run preview
```

You may also use Yarn or PNPM, if desired