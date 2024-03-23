# yos-server

## Running

### .env file

```
PB_ADMIN_EMAIL="admin@example.com"
PB_ADMIN_PW=

YOS_USERNAME=
YOS_PASSWORD=
YOS_NAME=
YOS_EMAIL="user@example.com"
```

> \[!TIP\]
> Use a non-trivial password or otherwise your browser might scream at you

## Development

* Start Server `go run main.go serve`

### Testing

* run `go test yos/controller`

### Benchmarking

* run `go test -v yos/controller -bench=.`