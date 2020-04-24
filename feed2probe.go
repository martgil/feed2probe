package main

import (
		"fmt"
		"bufio"
		"os"
		"sync"
		"net/http"
		"log"
		flags "github.com/jessevdk/go-flags"
		"strconv"
)

var opts struct {
	Filter string `short:"f" long:"filter" default:"*" description:"filter based on status Codes: 200,302,400,500 \ndefault: *"`
}

func main() {
	_,err := flags.ParseArgs(&opts, os.Args)
	if err != nil{
		os.Exit(1)
	}
	s := bufio.NewScanner(os.Stdin)
	urls := make(chan string)
	go func () {
		for s.Scan() {
			urls <- s.Text()

		}
		if err := s.Err(); err != nil {
			log.Println(err)
		}
		close(urls)
	}()

	worker := &sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		worker.Add(1)
		go Run(urls, worker)
	}
	worker.Wait()

}

func Run(urls chan string, worker *sync.WaitGroup) {
	for url := range urls {
		response,err := http.Get(url)
		if err != nil {
			return ;
		}

		if opts.Filter == "*" {
			fmt.Printf("[%d] %s \n", response.StatusCode, url)
		} else if opts.Filter == strconv.Itoa(response.StatusCode) {
			fmt.Printf("[%d] %s \n", response.StatusCode, url)
		}
		
	}
}
