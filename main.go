package main

import (
	"fmt"
	"errors"
	"log"
	"io"
	"os"
)

func main(){
	file_name := "messages.txt"
	file, err := os.Open(file_name)
	if err != nil{
		log.Fatalf("couldn't open %s: Error %s",file_name, err)
	}
	defer file.Close()

	for {
		b := make([]byte,8,8)
		_, err := file.Read(b)
		if err != nil{
			if errors.Is(err, io.EOF){
				break
			}
			fmt.Printf("error: %s\n",err)
			break
		}
		printed := string(b[:])
		fmt.Printf("read: %s\n", printed)
	}
}
