
## kafka-to-hdfs-service
### Description
Service that listens on a given kafka topic and writes the value to hdfs as the given [hdfs_dir]+[hdfs_file]+ the kafka offset.

### Example
./kafka2hdfs -brokers="localhost:9092" -offset="newest" -topic="hdfs_test" -hdfs-url="localhost:9000" -hdfs-dir="/tmp" -hdfs-file="output" -hdfs-user="vmware"  

Output file:  
/tmp/output.0  
/tmp/output.1  

