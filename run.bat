kafka2hdfs -brokers="localhost:9092" -offset="oldest" -topic="hdfs_test" -hdfs-url="localhost:9000" -hdfs-dir="/tmp" -hdfs-file="output.txt" -hdfs-user="vmware"

