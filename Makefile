
run:
	go build ./app/main.go && ./main

migrate:
	go build ./database/migration/migrate.go
	./migrate && rm ./migrate