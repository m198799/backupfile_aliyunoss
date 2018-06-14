package file

import (
	"fmt"
	"os"
)

func Delete(dirPath string, delTime int64) error  {
	files, err := ListNotChangeDir(dirPath,delTime)
	if err != nil{
		fmt.Println(err)
		return err
	}
	for _, file := range files{
		fmt.Println("delete:",file)
		os.RemoveAll(file)
	}
	return nil
}