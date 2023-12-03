package main
import (
    "bufio"
    "os"
)

func readFile(filename string){
	readFile, err := os.Open(filename)
	
    if err != nil {
		panic(err)
    }
    fileScanner := bufio.NewScanner(readFile)
	
    fileScanner.Split(bufio.ScanLines)
	
	
	var input []string
    for fileScanner.Scan() {
		input = append(input, fileScanner.Text())
    }
	
    readFile.Close()
}