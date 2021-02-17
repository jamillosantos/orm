# Library example

## Instructions

From the root directory:

1. Fire the environment up

```
docker-compose -f samples/library/docker-compose.yml up  -d
```

2. Generate the modules

```
go generate ./...
```

3. Run the sample itself

```
go run samples/library/code/main.go
```
