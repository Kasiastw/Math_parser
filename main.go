package main

import (
	"bufio"
	"flag"
	"log"
	"math_parser/calculate"
	"os"
)

func main()  {
	filename := flag.String("filename", "input.txt", "The file to parse")
	flag.Parse()

	if *filename == "" {
		log.Fatal("Provide a file to parse")
	}

	log.Println("Getting file")

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	ch:=make(chan string)
	for scanner.Scan() {
		str:= scanner.Text()
		go func() {
			output:=calculate.Result(str)
			ch<-output
		}()
		log.Println(<-ch)
	}
}

