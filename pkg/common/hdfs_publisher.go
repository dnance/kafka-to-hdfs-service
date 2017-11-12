package common

import (
	"fmt"
	"github.com/colinmarc/hdfs"
)

type HdfsPublisher struct {
	client *hdfs.Client
	url    string
	dir    string
	file   string
	user   string
}

func NewHdfsPublisher(hdfs_url, hdfs_dir, hdfs_file, hdfs_user string) (*HdfsPublisher, error) {

	client, err := hdfs.NewForUser(hdfs_url, hdfs_user)
	if err != nil {
		fmt.Printf("Could not create client at %s\n", hdfs_url)
		return nil, err
	}

	hp := &HdfsPublisher{
		client: client,
		url:    hdfs_url,
		dir:    hdfs_dir,
		file:   hdfs_file,
		user:   hdfs_user,
	}

	return hp, nil
}

func (h *HdfsPublisher) WriteData(suf string, data []byte) error {

	// check for existing directory
	_, err := h.client.Stat(h.dir)
	if err != nil {
		// TODO figure out the correct permission value
		err = h.client.Mkdir(h.dir, 0777)
		if err != nil {
			fmt.Printf("Can't create directory...%s\n", h.dir)
			return err
		}
	}

	abs_path_to_file := fmt.Sprintf("%s/%s.%s", h.dir, h.file, suf)
	fw, err := h.client.Create(abs_path_to_file)
	if err != nil {
		fmt.Printf("Can't create file...%s\n", abs_path_to_file)
		return err
	}
	defer fw.Close()

	_, err = fw.Write(data)
	if err != nil {
		fmt.Printf("error writing data\n")
		return err
	}
	fmt.Printf("successfully wrote data %s...", data)

	return nil
}

func WriteData(hdfs_url, hdfs_dir, hdfs_file string, data []byte, hdfs_user string) error {

	client, err := hdfs.NewForUser(hdfs_url, hdfs_user)
	if err != nil {
		fmt.Printf("Could not create client at %s\n", hdfs_url)
		return err
	}

	// check for existing directory
	_, err = client.Stat(hdfs_dir)
	if err != nil {
		// TODO figure out the correct permission value
		err = client.Mkdir(hdfs_dir, 0777)
		if err != nil {
			fmt.Printf("Can't create directory...%s\n", hdfs_dir)
			return err
		}
	}

	abs_path_to_file := fmt.Sprintf("%s/%s", hdfs_dir, hdfs_file)

	// TODO issue that created file from here cannot be appended to, but
	// if use hdfs dfs -touchz [file] then
	// append works fine
	// maybe need to create file with different permissions from here

	_, err = client.Stat(abs_path_to_file)
	var fw *hdfs.FileWriter

	if err != nil { // file does not exist
		fw, err = client.Create(abs_path_to_file)
		if err != nil {
			fmt.Printf("Can't create file...%s\n", abs_path_to_file)
			return err
		}
	} else {
		fw, err = client.Append(abs_path_to_file)
		if err != nil {
			fmt.Printf("Can't open file...%s\n", abs_path_to_file)
			return err
		}
	}
	defer fw.Close()

	_, err = fw.Write(data)
	if err != nil {
		fmt.Printf("error writing data\n")
		return err
	}
	fmt.Printf("successfully wrote data %s...", data)

	return nil
}

func ListDirs(hdfs_url, hdfs_dir string, hdfs_user string) error {

	client, err := hdfs.NewForUser(hdfs_url, hdfs_user)
	if err != nil {
		fmt.Printf("Could not create client at %s\n", hdfs_url)
		return err
	}

	file, err := client.Open(hdfs_dir)
	if err != nil {
		fmt.Printf("file does not exist...%s\n", hdfs_dir)
		return err
	}
	defer file.Close()

	names, _ := file.Readdirnames(0)

	for _, name := range names {
		fmt.Printf("item name %s\n", name)
	}

	return nil

}
