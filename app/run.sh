if docker stop prime-image; then docker rm prime-image; fi
docker build -t prime-image .
docker run -d --rm --name prime-instance -p 8080:8080 prime-image