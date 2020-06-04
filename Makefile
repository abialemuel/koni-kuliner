
run:
	go build main.go && ./main

migrate:
	go build ./database/migration/migrate.go
	./migrate && rm ./migrate

deploy: 
	docker build -t gcr.io/koni-kuliner/koni-kuliner:v2 .
	docker push gcr.io/koni-kuliner/koni-kuliner:v2
	kubectl create deployment konikuliner-web --image=gcr.io/koni-kuliner/koni-kuliner:v2
	kubectl expose deployment konikuliner-web --type=LoadBalancer --port 5000 --target-port 5000
	docker pull gcr.io/koni-kuliner/koni-kuliner:v2