# InOrder

### Quickstart

Clone the repo and create a `config.yaml`


The repository comes with dummy data to test with.



### Setup and Configuration

Make a `config.yaml` in the root directory of the project based on the [sample config](https://github.com/TanmayArya-1p/MVCAssignment/blob/main/sample.config.yaml).

### Run


If you have [go](go.dev) installed, run the following command to start the service
```bash
make run
```

If you would like to run it as a docker container then make sure to edit the `config.yaml` accordingly and run

```bash
docker compose up -d
```
### Database Migrations

Migration files are located in the `/database/migrations` directory.

**You can also automate database migrations using [golang-migrate](https://github.com/golang-migrate/migrate).**

For convenience, you can also use the provided Makefile commands that are basically aliases for [golang-migrate](https://github.com/golang-migrate/migrate). To do this put your DSN in a `.dsn` file located at the root of the project.

- To apply migrations:
    ```bash
    make db-up
    ```
- To rollback migrations:
    ```bash
    make db-down
    ```

### Mocking Requests

This project comes with an [OpenAPI Spec](https://github.com/TanmayArya-1p/MVCAssignment/blob/main/openapi.json) that can be imported into any request mocking client.
