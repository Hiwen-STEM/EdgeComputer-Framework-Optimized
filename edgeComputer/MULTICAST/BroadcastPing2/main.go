package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"math"
	"github.com/dmichael/go-multicast/multicast"
	"github.com/urfave/cli"
)

const (
    defaultMulticastAddress = "239.0.0.0:9999"
)

func main() {
	app := cli.NewApp()

	app.Action = func(c *cli.Context) error {
		address := c.Args().Get(0)
		if address == "" {
			address = defaultMulticastAddress
		}
		fmt.Println("Broadcasting to %s\n", address)
		ping(address)
		return nil
	}

	app.Run(os.Args)
}

func ping(addr string) {
	conn, err := multicast.NewBroadcaster(addr)
	if err != nil {
		log.Fatal(err)
	}
	var cool int = 1
    	path := ""
    	path2 := ""
	for {
		path = "/go/mem1/"+strconv.Itoa(cool)+".jpg"




		//attempt to remove flag file GO-$Varialbe.txt
		path2 = "/go/mem1/GO-"+strconv.Itoa(cool)+".txt"
		for{
			_,err := os.Stat(path2)
			if err == nil{
			   for{
				err := os.Remove(path2)
				if err == nil{
				   break
				} else{
				  fmt.Println("Failed to remove file!...\n")
				}
			   }
			   break
			}
		}



		file, err := os.Open(path)
        	if err != nil{
            	   fmt.Println("RETRY!!!!!\n\n")
        	} else{
            	   data := make([]byte,math.MaxInt32)
		   transfer := make([]byte,65500)
            	   n, err := file.Read(data)
            	   if err != nil{
                      fmt.Println("File Reading error!\n\n")
            	   } else{

		      //Transmit the size of the file to be
		      //transmitted
		      conn.Write([]byte(strconv.Itoa(n)))
		      time.Sleep(1000 * time.Millisecond)
		      fmt.Println(n)




		      //Data will be transmitted in the form
		      //of 65500 byte blocks at a time.
		      //Variable u will be assigned the number of iterations that are to be made
		      //within the for loop.
		      var u int = (n / 65500)
		      if n % 65500 != 0{
		      	 u = u + 1
		      }
		      var indexG int = 0
		      for i := 1; i <= u; i++{
		      	  if (i != u) || (n % 65500 == 0){
			     for x:= 1; x <= 65500; x++{
			     	 transfer[x-1] = data[indexG]
				 n = n - 1
				 indexG = indexG + 1
			     }
			     conn.Write(transfer)
			     time.Sleep(80 * time.Millisecond)
			  } else {
			     sandbox := make([]byte,n)
			     fmt.Println(n)
			     for v := 1; v <= n; v++{
			     	 sandbox[v-1] = data[indexG]
				 indexG = indexG + 1
			     }
			     conn.Write(sandbox)
			     time.Sleep(80 * time.Millisecond)
			  }
		      }
		      cool = cool + 1
		   }
        	}
		file.Close()
		time.Sleep(10 * time.Second)
        }
}
