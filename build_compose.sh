docker compose -f docker-compose.yml down -v

# docker rm -f $(docker ps -aq)

docker compose -f docker-compose.yml build --no-cache

docker compose -f docker-compose.yml up -d