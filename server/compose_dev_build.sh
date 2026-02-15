docker compose -f docker-compose-dev.yml down -v

docker rm -f $(docker ps -aq)

go mod tidy

docker compose -f docker-compose-dev.yml build --no-cache

docker compose -f docker-compose-dev.yml up -d