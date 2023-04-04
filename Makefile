start:
	go run main.go

stop:
	sudo kill -9 `sudo lsof -t -i:8080`

server:
	docker run --name gfiber -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -e MYSQL_ROOT_HOST=% -d mysql

.PHONY: start stop server