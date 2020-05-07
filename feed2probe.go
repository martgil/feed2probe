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
                "strings"
)

var opts struct {
        Filter string `short:"f" long:"filter" default:"*" description:"filter based on status Codes: 200,302,400,500"`
        Probe bool `short:"p" long:"probe" description:"perform domain check to extract alive domains"`
}

func main() {
        _,err := flags.ParseArgs(&opts, os.Args)
        if err != nil{
                fmt.Println(err)
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

        for i := 0; i < 8; i++ {
                worker.Add(1)
                go Run(urls, worker)
        }
        worker.Wait()
}

func Run(urls chan string, worker *sync.WaitGroup) {
        for url := range urls {
                scheme := "https://"
                var furl string
                if opts.Probe { furl = scheme + url } else { furl = url }
                response,err := http.Get(furl)
                if err != nil {
                        continue
                }

                if opts.Filter == "*" {
                        fmt.Printf("[%d] %s \n", response.StatusCode , url)
                } else {
                        i := 0
                        for ; i < len(strings.Split(opts.Filter,",")); {
                                if strings.Split(opts.Filter,",")[i] == strconv.Itoa(response.StatusCode) {
									if response.StatusCode == 200 && opts.Probe  {
										fmt.Printf("%s \n", url)
									} else {
										fmt.Printf("[%d] %s \n", response.StatusCode, url)
									}
                                }
                                i += 1
                        }
                }
        }
        worker.Done()
}
