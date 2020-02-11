package main
import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)
var MAX int = 1000
var SHARED_DIR string = "/shared"
func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 2 {
		var err error
		MAX, err = strconv.Atoi(argsWithoutProg[0])
		if err != nil {
			panic(err)
		}
	} else {
		var err error
		MAX, err = strconv.Atoi(argsWithoutProg[0])
		if err != nil {
			panic(err)
		}
		SHARED_DIR = argsWithoutProg[1]
	}

	var files = make([]*os.File, MAX)
	defer closeFiles(files)
	for i := 0; i < MAX; i++ {
		file, err := ioutil.TempFile(SHARED_DIR, "ulimit-test-")
		if err != nil {
			fmt.Println(err)
			fmt.Printf("Failed to open file %d, see error above\n", i)
			return
		}
		files[i] = file
	}
}
func closeFiles(files []*os.File) {
	for i := 0; i < len(files); i++ {
		if files[i] != nil {
			err := files[i].Close()
			if err != nil {
				fmt.Println(err)
				fmt.Println("Failed to close file, see error above")
			}
			err = os.Remove(files[i].Name())
			if err != nil {
				fmt.Println(err)
				fmt.Println("Failed to remove file, see error above")
			}
		} else {
			return
		}
	}
}