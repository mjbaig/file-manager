docker build -t my-golang-app .
docker run -p 8080:8080 -it --rm --name my-running-app my-golang-app