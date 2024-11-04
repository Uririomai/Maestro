swagger:
	swag init -g .\cmd\main.go

build: swagger
	go build -o build.exe .\cmd\main.go

run: swagger
	go run .\cmd\main.go