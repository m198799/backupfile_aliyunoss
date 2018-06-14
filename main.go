package main

import (
	"fmt"
	"os"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/widuu/goini"
	"flag"
	"github.com/pkevin0909/backupfile_aliyunoss/file"
	"strconv"
)

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

func file_exsit(file string){
	_, err := os.Stat(file)
	if os.IsNotExist(err){
		fmt.Println("file is not exist")
		os.Exit(-1)
	}
	fmt.Println("file path:",file)
}


func main() {

	var configfile string
	flag.StringVar(&configfile, "configfile", "./conf.ini", "config file path")
	flag.Parse()
	flag.Usage()
	_, err := os.Stat(configfile)
	file_exsit(configfile)
	config := goini.SetConfig(configfile)

	client, err := oss.New(config.GetValue("oss", "endpoint"), config.GetValue("oss", "accessKeyID"), config.GetValue("oss", "accessKeySecret"))
	if err != nil {
		handleError(err)
	}
	lsBks, err := client.ListBuckets()
	if err != nil {
		handleError(err)
	}
	for _, bucket := range lsBks.Buckets {
		fmt.Println("bucket:", bucket.Name)
	}

		bucket, err := client.Bucket(config.GetValue("oss","bucketName"))
		if err !=nil{
			handleError(err)
		}
	//	lsObs, err := bucket.ListObjects()

	//	if err != nil{
	//		handleError(err)
	//	}
	//	for _, object := range lsObs.Objects{
	//		fmt.Println("Object:", object.Key)
	//	}

	//	err = bucket.PutObjectFromFile("my-object","/Users/panxu/Downloads/office_mac.dmg")	/
	//  if err != nil{
	//		handleError(err)
	//	}

	// err = bucket.DeleteObject("my-object")
	// if err != nil {
	//    handleError(err)
	// }
	dirPath := config.GetValue("oss", "dirPath")
	file.ListDir(dirPath)
	file_exsit(dirPath)
	expiration_time := config.GetValue("oss","expiration")
	expiration, err := strconv.ParseInt(expiration_time,10,64)
	if err != nil{
		fmt.Println("strconv false")
		os.Exit(-1)
	}
	files, err:= file.ListChangeDir(dirPath,expiration)


	for _, table1 := range files {
		fmt.Println(table1)
		var table2 string = string([]rune(table1)[1:])
		fmt.Println("table2:",table2)
		err = bucket.PutObjectFromFile(table2,table1)
 		if err != nil{
			handleError(err)
		}
	}
}
