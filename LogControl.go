package main

import (
	"os"
	"bufio"
	"strings"
)
func checkError(e error){
	panic(e)
}
func readLog(url string) map[string]string{
	ret := make(map[string]string)
	//dat, err := ioutil.ReadFile(url)
	//checkError(err)
	file, err := os.Open(url)
	checkError(err)
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		pair := strings.Split(scanner.Text(), " was shorten to ")
		ret[pair[0]] = pair[1]
	}
	return ret
}

func writeLog(url string, log map[string]string){
	file, err := os.Open(url)
	checkError(err)
	writer := bufio.NewWriter(file)
	for key, value := range log{
		writer.WriteString(key + " was shorten to " + value)
	}
}