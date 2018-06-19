package main

import (
	"fmt"
	"os"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/widuu/goini"
	"flag"
	"github.com/pkevin0909/backupfile_aliyunoss/file"
	"strconv"
	"time"
	"strings"
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
	flag.StringVar(&configfile, "configfile", "./online.ini", "config file path")
	flag.Parse()
//	flag.Usage()

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
	expirationTime := config.GetValue("oss","expirationTime")
	expTime, err := strconv.ParseInt(expirationTime,10,64)
	runTimeOne := config.GetValue("oss","runTime")
	runTime, err := strconv.Atoi(runTimeOne)
	if err != nil{
		fmt.Println("strconv false\n")
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	if runTime <= 0 {
		runTime = 1
	}

	dirPath := config.GetValue("oss", "dirPath")
	fmt.Println(dirPath)
//	filess := map[int]string{}
	for _, v := range strings.Split(dirPath, ",") {

		file_exsit(v)
		file.ListDir(v)
		for {
			files, err := file.ListChangeDir(v, expTime)

			for _, table1 := range files {
				fmt.Println("table1:", table1)
				var table2 string = string([]rune(table1)[1:])
				fmt.Println("table2:", table2)
				err = bucket.PutObjectFromFile(table2, table1)
				if err != nil {
					handleError(err)
				}
			}
			delTimeStr := config.GetValue("oss", "delTime")
			delTime, err := strconv.ParseInt(delTimeStr, 10, 64)
			if err != nil {
				fmt.Println("strconv false\n")
				fmt.Println("Error:", err)
				os.Exit(-1)
			}
			if delTime == 0 {
				fmt.Println("do not delete files")
			} else {
				if delTime > expTime {
					file.Delete(v, delTime)
				} else {
					fmt.Println("delTime need > expirationTime")
				}
			}
			time.Sleep(time.Hour * time.Duration(runTime))
		}
	}
}
