# Fampay Backend Assignment
This is a backend assignment for Fampay. The assignment is to create a REST API that fetches videos from Youtube and stores them in a database. The API should also provide a way to search the stored videos.

# Table of Contents
- [Fampay Backend Assignment](#fampay-backend-assignment)
- [Table of Contents](#table-of-contents)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Linux/MacOS](#linuxmacos)
    - [Docker](#docker)
      - [Troubleshooting](#troubleshooting)
  - [API Documentation](#api-documentation)
  - [Explanation](#explanation)
      - [Database](#database)
    - [API](#api)
      - [Testing](#testing)
    - [Youtube Fetcher](#youtube-fetcher)
    - [DbApp](#dbapp)

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.
### Prerequisites
In case of running using docker, you need to have the following installed on your machine:
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Google Cloud Platform](https://cloud.google.com/) account with Youtube Data API enabled

In case of running the application directly, you need to have the following installed on your machine:
- [Go](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)
- [Google Cloud Platform](https://cloud.google.com/) account with Youtube Data API enabled

### Linux/MacOS
1. Clone the repository
```bash
git clone https://github.com/bearTS/fampay-backend-assignment.git fb_assignment
cd fb_assignment
```

2. Copy the environment variables file
```bash
cp .env.example .env
```
3. Fill the appropriate values in the environment variables file
4. Run the following commands in separate terminal windows
```bash
go run main.go dbapp migrate && go run main.go api
```
```bash
go run main.go youtube-fetcher
```
5. The application should be running on `http://localhost:3000` by default

### Docker
1. Clone the repository
```bash
git clone https://github.com/bearTS/fampay-backend-assignment.git fb_assignment
cd fb_assignment
```

2. Fill the appropriate values in the environment variables in docker-compose
For `YOUTUBE_API_KEYS`, you can add multiple keys separated by a comma. The application will switch to the next key if the quota for the current key is exhausted.

3.. Run the postgres container
```bash
docker-compose up -d psql
```

4. Run the following command to run the application
```bash
docker-compose up -d
```


The application should be running on `http://localhost:3000` by default.

Note: The setup is used for local or development purposes. For production, the environment variables should be set in a more secure way and additional measures should be taken.

#### Troubleshooting

In case you get the following error:
```bash
[error] failed to initialize database, got error failed to connect to
```

It could be because the database is not ready yet. You can run the following command to check the logs of the database container:
```bash
docker-compose logs -f psql
```

You can check the logs to see if the database is ready. If the database is ready, you can run the following command to run the application:
```bash
docker-compose up -d migrate
```
The above should resolve the issue. If the issue still persists, kindly check the environment variables.

## API Documentation
<!-- api/pkg/routes/openapi-spec.yaml -->
You can get the API documentation in the OpenAPI format from this [link](/api/pkg/routes/openapi-spec.yaml).
You can directly import this file to Postman or Swagger to get the API documentation.

## Explanation
The application is divided into 2 separate services:
1. [API](#api)
2. [Youtube Fetcher](#youtube-fetcher)
   
We use our own command line tool, [DbApp](#dbapp) to handle one-time database operations.

#### Database
We are using PostgreSQL as the database. The database is used to store the videos fetched from Youtube. The database is managed using the [GORM](https://gorm.io/) library. The database schema is defined in the `pkg/db` directory. The database schema is defined using the GORM tags.

### API
The API is the main part of the application. It is responsible for handling the requests from the client and returning the appropriate response. The API is built using the [Go Echo](https://echo.labstack.com/) framework. The API is divided into 3 main parts:
1. **Routes**: The routes are defined in the `api/pkg/routes` directory. The routes are generated using openapi-spec with the help of [oapi-codegen](https://github.com/deepmap/oapi-codegen). The routes are defined in the `openapi-spec.yaml` file. The routes are then generated using the following command:
If you have [task](https://taskfile.dev/) installed, use the following command:
```bash
task gen
```
else use the following command:
```bash
cd ./api/pkg/routes
go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=oapi-codegen.yaml ./openapi-spec.yaml
```
2. **Handlers**: The handlers are defined in the `api/pkg/handlers` directory. The handlers are responsible for handling the requests from the client and returning the appropriate response. 
    We use dependency injection to inject the required services into the handlers. This makes the handlers more testable and modular.
3. **Services**: The services are defined in the `api/pkg/services` directory. The services are responsible for handling the business logic.

The service can be run using the following command:
```bash
go run main.go api
```

#### Testing
You can hit the API using the following command:
```bash
curl -X 'GET' \
  'http://localhost:3000/v1/videos' \
  -H 'accept: application/json'
```
You can also use the Postman collection provided in the repository to test the API, or you can directly import the OpenAPI spec to Postman or Swagger to get the API documentation.


### Youtube Fetcher
The Youtube Fetcher is responsible for fetching the videos from Youtube and storing them in the database. The Youtube Fetcher is built using the [Go](https://golang.org/) programming language. 

The Youtube Fetcher can be run using the following command:
```bash
go run main.go youtube-fetcher
```

### DbApp
The DbApp is a command line tool that is responsible for handling one-time database operations. The DbApp is built using the [Go](https://golang.org/) programming language.
This is **not a service**, but a command line tool.

The service can be run using the following command:
```bash
go run main.go dbapp migrate
```
Currently, the DbApp only handles the database migrations. In the future, it can be extended to handle other database operations.


