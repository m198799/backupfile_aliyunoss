package main

import (
	"fmt"
	"os"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io/ioutil"
)
func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

var files []string
var PthSep string = string(os.PathSeparator)

func ListDir(dirPath string) (err error)  {

	dir,err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil
	}

	for _ ,fi := range dir {
		if fi.IsDir(){
			ListDir(dirPath+PthSep+fi.Name())
		} else{
			files = append(files,dirPath+PthSep+fi.Name())
		}
	}
	return nil
}


func main() {
	client, err := oss.New("http://oss-cn-beijing.aliyuncs.com", "LTAIzbTqnb3mdOiB", "Yb8yfBEdkHrBYx8bRI7ovyG0ifyg2Q")
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

	bucket, err := client.Bucket("pkevin")
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

	_ = ListDir("/Users/panxu/Downloads/123")

	for _, table1 := range files {
		fmt.Println(table1)
		var table2 string = string([]rune(table1)[:1])
		if table2 == string(os.PathSeparator){
			table2 = string([]rune(table1)[1:])
		}
		err = bucket.PutObjectFromFile(table2,table1)
		if err != nil{
			handleError(err)
		}
	}
}