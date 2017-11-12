
clean:
	go clean

build: clean
	go build ./cmd/kafka2hdfs

run: build
kafka2hdfs \
-brokers="localhost:9092"\
-offsets="oldest"\
-topic="hdfs_test"\
-hdfs_url="localhost:9000" \
hdfs_dir="/tmp" \
-hdfs_file="output.txt" \
-user="vmware"


