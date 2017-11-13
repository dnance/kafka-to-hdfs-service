package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dnance/kafka-to-hdfs-service/pkg/common"
)

var (
	hdfs_url    = flag.String("hdfs-url", os.Getenv("KH_SERVICE_URL"), "The address to HDFS IP + PORT")
	hdfs_dir    = flag.String("hdfs-dir", os.Getenv("KH_SERVICE_DIR"), "The address to HDFS IP + PORT")
	hdfs_file   = flag.String("hdfs-file", os.Getenv("KH_SERVICE_FILE"), "The address to HDFS IP + PORT")
	hdfs_user   = flag.String("hdfs-user", os.Getenv("KH_SERVICE_USER"), "The address to HDFS IP + PORT")
	brokerList  = flag.String("brokers", os.Getenv("KH_SERVICE_BROKERS"), "The comma separated list of brokers in the Kafka cluster. You can also set the KAFKA_PEERS environment variable")
	topic       = flag.String("topic", os.Getenv("KH_SERVICE_TOPIC"), "REQUIRED: the topic to consume to")
	partitioner = flag.String("partitioner", "", "The partitioning scheme to use. Can be `hash`, `manual`, or `random`")
	partition   = flag.Int("partition", -1, "The partition to produce to.")
	partitions  = flag.String("partitions", "all", "The partitions to consume, can be 'all' or comma-separated numbers")
	offset      = flag.String("offset", os.Getenv("KH_SERVICE_OFFSET"), "The offset to start with. Can be `oldest`, `newest`")
	bufferSize  = flag.Int("buffer-size", 256, "The buffer size of the message channel.")
	verbose     = flag.Bool("verbose", false, "Turn on sarama logging to stderr")
	showMetrics = flag.Bool("metrics", false, "Output metrics on successful publish to stderr")
	silent      = flag.Bool("silent", false, "Turn off printing the message's topic, partition, and offset to stdout")

	logger = log.New(os.Stderr, "", log.LstdFlags)
)

func main() {

	flag.Parse()

	h, err := common.NewHdfsPublisher(*hdfs_url, *hdfs_dir, *hdfs_file, *hdfs_user)
	if err != nil {
		fmt.Printf("error creating client: %v", err)
		os.Exit(-1)
	}

	fmt.Printf("Create consumer...\n")

	err = common.ConsumeMessages(*brokerList, *topic, *bufferSize, *offset, *partitions, h)
	if err != nil {
		fmt.Printf("error consuming message, %v", err)
		os.Exit(-1)
	}
}

func printErrorAndExit(code int, format string, values ...interface{}) {
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", fmt.Sprintf(format, values...))
	fmt.Fprintln(os.Stderr)
	os.Exit(code)
}
