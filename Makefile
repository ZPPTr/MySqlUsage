up:
	docker-compose up -d
down:
	docker-compose down
restart:
	docker-compose down
	docker-compose up -d

run-init:
	cd src && go build -o ./../bin/app && ./bin/app -initmode

run:
	docker-compose up go
	#cd src && go build -o ./../bin/app && ./bin/app -ips=5000

init:
	docker-compose run go /dist/main -initmode

insert:
	docker-compose run go /dist/main