package common

import (
	"fmt"
	"github.com/colinmarc/hdfs"
)

func ListDirs(hdfs_url, hdfs_dir string) error {

	client, err := hdfs.New(hdfs_url)
	if err != nil {
		fmt.Printf("Could not create client at %s\n", hdfs_url)
		return err
	}

	file, err := client.Open(hdfs_dir)
	if err != nil {
		fmt.Printf("file does not exist\n")
		return err
	}
	defer file.Close()

	names, _ := file.Readdirnames(0)

	for _, name := range names {
		fmt.Printf("directory name %s\n", name)
	}

	return nil

}
