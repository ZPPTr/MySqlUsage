up:
	docker-compose up -d
down:
	docker-compose down
restart:
	docker-compose down
	docker-compose up -d

run-init:
	go build -o ./bin/app && ./bin/app -initmode

run:
	./bin/app