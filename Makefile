
kafka-hdfs-service: imports
	CGO_ENABLED=0

clean:
	go clean

clean-container:
	-docker stop kh-service
	-docker rm kh-service
		

imports:
	go get ./...

build: clean kafka-hdfs-service
	go build ./cmd/kafka2hdfs

container: clean-container
	docker build -t kafka-hdfs-service .
	

run: build
	./kafka2hdfs \
	-brokers="localhost:9092" \
	-offset="newest" \
	-topic="hdfs_test" \
	-hdfs-url="localhost:9000" \
	-hdfs-dir="/tmp" \
	-hdfs-file="output" \
	-hdfs-user="vmware"

docker-run: container
	docker run -d --name=kh-service --net=host \
	 -e KH_SERVICE_BROKERS="localhost:9092" \
	 -e KH_SERVICE_TOPIC="hdfs_test" \
	 -e KH_SERVICE_OFFSET="newest" \
	 -e KH_SERVICE_URL="localhost:9000" \
	 -e KH_SERVICE_DIR="/tmp" \
	 -e KH_SERVICE_FILE="output" \
	 -e KH_SERVICE_USER="vmware" \
	 kafka-hdfs-service

 