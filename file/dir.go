package file

import (
	"io/ioutil"
	"os"
	"fmt"
	"time"
)
var files []string

func ListDir(dirPath string) ([]string,error)  {
	var files []string
	var PthSep string = string(os.PathSeparator)
	dir,err := ioutil.ReadDir(dirPath)
	if err != nil {
		return files,err
	}

	for _ ,fi := range dir {
		if fi.IsDir(){
//			fmt.Println(fi)
			ListDir(dirPath+PthSep+fi.Name())
		} else{
			files = append(files,dirPath+PthSep+fi.Name())
		}
	}
	return files,nil
}

func ListChangeDir(dirPath string,expirationTime int64) ([]string,error)  {

	var PthSep string = string(os.PathSeparator)
	dir,err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Println("not found file")
		return files, err
	}

	for _ ,fi := range dir {
		fiInfo, err := os.Stat(dirPath+PthSep+fi.Name())
		if err != nil {
			panic(err)
		}
		fi_time := fiInfo.ModTime()
		timestamp := time.Now()
		lasthore :=GetHourDiffer(fi_time,timestamp)

		if fi.IsDir(){
//			fmt.Println(dirPath+PthSep+fi.Name())
			if lasthore > expirationTime {
				fmt.Printf("dir: \"%s\" not need to put \n",dirPath+PthSep+fi.Name())
			} else {
				ListChangeDir(dirPath+PthSep+fi.Name(),expirationTime)
				fmt.Printf("dir: \"%s\" need to put \n",dirPath+PthSep+fi.Name())
			}

		} else{
			if lasthore > expirationTime {
				fmt.Printf("file: \"%s\" not need to put \n",dirPath+PthSep+fi.Name())
			} else {
				files = append(files,dirPath+PthSep+fi.Name())
				fmt.Printf("file: \"%s\" need to put \n",dirPath+PthSep+fi.Name())
			}

		}
	}
	return files, nil
}

func ListNotChangeDir(dirPath string,expirationTime int64) ([]string,error)  {

	var PthSep string = string(os.PathSeparator)
	dir,err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Println("not found file")
		return files, err
	}

	for _ ,fi := range dir {
		fiInfo, err := os.Stat(dirPath+PthSep+fi.Name())
		if err != nil {
			panic(err)
		}
		fi_time := fiInfo.ModTime()
		timestamp := time.Now()
		lastTime :=GetHourDiffer(fi_time,timestamp)

		if fi.IsDir(){
			//			fmt.Println(dirPath+PthSep+fi.Name())


			if lastTime > expirationTime {
				files = append(files,dirPath+PthSep+fi.Name())
				fmt.Printf("dir: \"%s\" not changed in %d hours \n",dirPath+PthSep+fi.Name(),expirationTime)
			} else {
				ListNotChangeDir(dirPath+PthSep+fi.Name(),expirationTime)
				fmt.Printf("dir: \"%s\" is changed in %d hours \n",dirPath+PthSep+fi.Name(),expirationTime)
				continue;
			}

		} else{
			if lastTime > expirationTime {
				files = append(files,dirPath+PthSep+fi.Name())
				fmt.Printf("file: \"%s\" is not changed in %d hours \n",dirPath+PthSep+fi.Name(),expirationTime)

			} else {
				fmt.Printf("file: \"%s\" is changed in %d hours \n",dirPath+PthSep+fi.Name(),expirationTime)
				continue;
			}

		}
	}
	return files, nil
}
