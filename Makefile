.PHONY: mongo

mongo:
	docker exec -it go-election-mongodb mongo -u root -p root --authenticationDatabase admin

http-api:
	go run cmd/http/main.go