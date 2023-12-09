# TP middleware example

## Run

Tidy / download modules :
```
go mod tidy
```
Build : 
```
go build -o middleware_collections cmd/main.go
```
Run : 
```
./middleware_collections
```
or
```
go run ./cmd/main.go
```

## Documentation

Documentation is visible in **api** directory ([here](api/swagger.json)). And at "/" when the server is started. 

Update swagger documentation:
```
go install github.com/swaggo/swag/cmd/swag@latest
```
If using Linux:
```
export PATH=$PATH:$HOME/go/bin
```
Generate documentation:
```
swag init -g ./internal/swagger_expose.go -o ./api
```