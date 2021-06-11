
build:
	bash ./scripts/build-all.sh

up-redis:
	docker-compose -f docker/docker-compose.yml -p redisbench up -d

down-redis:
	docker-compose -f docker/docker-compose.yml -p redisbench down

run:
	go run main.go -a localhost:6379 -c 10 -n 5000 -d 1000

run1:
	go run main.go -a localhost:6379 -c 10 -n 2000 -d 1000 -ma localhost:9001,localhost:9002 -mo 1

run2:
	go run main.go -a localhost:6379 -c 10 -n 2000 -d 1000 -ma localhost:9001,localhost:9002 -mo 2
