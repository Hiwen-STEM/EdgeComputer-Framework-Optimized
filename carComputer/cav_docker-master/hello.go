package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
)



func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		dir := os.Args[1]
		fmt.Print("Serving ", dir)
		if err := http.ListenAndServe(":9000", http.FileServer(http.Dir(dir))); err != nil {
			panic(err)
		}
	}()

	go func() {
		l, err := net.Listen("tcp", ":3000")
		if err != nil {
			panic(err)
		}
		defer l.Close()

		for {
			c, err := l.Accept()
			if err != nil {
				continue
			}
			c.Write([]byte("EXITING"))
			fmt.Print("Exiting")
			c.Close()
			wg.Done()
		}
	}()

	wg.Wait()
}
