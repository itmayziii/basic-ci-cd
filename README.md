# basic-ci-cd
Demonstrating basic CI/CD using Google Cloud Build and Google Cloud Deploy.

## Server
```shell
go run cmd/server/server.go
```

### Routes

* `/` - returns "hello world"
* `/echo` - returns the request body in the response body
    ```shell
    curl -X POST -d "tommy" localhost:8080/echo
    ```

## CI


