
create-cluster:
	docker-compose exec redis1 redis-cli --cluster create \
		127.0.0.1:6001 127.0.0.1:6002 127.0.0.1:6003 \
		127.0.0.1:6004 127.0.0.1:6005 127.0.0.1:6006 \
		--cluster-replicas 1

up:
	docker-compose up --force-recreate -d

down:
	docker-compose down

cli:
	docker-compose exec redis1 redis-cli -c -p 6001

nodes:
	docker-compose exec redis1 redis-cli -p 6001 cluster nodes