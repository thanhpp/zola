export DOCKER_BUILDKIT=1
docker-compose -f docker-compose.auco.yml down
docker-compose -f docker-compose.auco.yml up --build --force-recreate