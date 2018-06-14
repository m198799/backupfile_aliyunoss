package file

import (
	"io/ioutil"
	"os"
	"fmt"
	"time"
)

var files []string

var PthSep string = string(os.PathSeparator)

func ListDir(dirPath string) (err error)  {
	dir,err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil
	}

	for _ ,fi := range dir {
		if fi.IsDir(){
			fmt.Println(fi)
			ListDir(dirPath+PthSep+fi.Name())
		} else{
			files = append(files,dirPath+PthSep+fi.Name())
		}
	}
	return nil
}

func ListChangeDir(dirPath string,expiration_time int64) (files []string,err error)  {
	dir,err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Println("not found file")
		return files, nil
	}

	for _ ,fi := range dir {
		if fi.IsDir(){
//			fmt.Println(dirPath+PthSep+fi.Name())
			fiinfo, err := os.Stat(dirPath+PthSep+fi.Name())
			if err != nil {
				panic(err)
			}
			fi_time := fiinfo.ModTime()
			timestamp := time.Now()
			lasthore :=GetHourDiffer(fi_time,timestamp)

			if lasthore > expiration_time {
				fmt.Printf("dir: \"%s\" not need to put \n",dirPath+PthSep+fi.Name())
				continue;
			} else {
				ListChangeDir(dirPath+PthSep+fi.Name(),expiration_time)
				fmt.Printf("dir: \"%s\" need to put \n",dirPath+PthSep+fi.Name())
			}

		} else{
			fiinfo, err := os.Stat(dirPath+PthSep+fi.Name())
			if err != nil {
				panic(err)
			}
			fi_time := fiinfo.ModTime()
			timestamp := time.Now()
			lasthore :=GetHourDiffer(fi_time,timestamp)

			if lasthore > expiration_time {
				fmt.Printf("file: \"%s\" not need to put \n",dirPath+PthSep+fi.Name())
				continue;
			} else {
				files = append(files,dirPath+PthSep+fi.Name())
				fmt.Printf("file: \"%s\" need to put \n",dirPath+PthSep+fi.Name())
			}

		}
	}
	return files, nil
}
