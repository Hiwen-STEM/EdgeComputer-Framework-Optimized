
package main

import (
       "fmt"
       "net"
       "os"
       "strconv"
       "sync"
       "log"
       "github.com/urfave/cli"
)       

const (
      defaultMulticastAddress = "239.0.0.0:9999"
      maxDatagramSize = 65500
)


//credit for this function (msgHandler) goes to user dmichael on github
//Though it is not being used for anything within this framework
//Feel free to uncomment the code within and add more to suit your
//needs!
func msgHandler(src *net.UDPAddr, n int, b []byte) {
     //fmt.Println(n, "bytes read from", src)
     //fmt.Println(hex.Dump(b[:n]))
}


//Credit for the basis of the Listen function goes to user dmichael on github
// Listen binds to the UDP address and port given and writes packets received
// from that address to a buffer which is passed to a hander
func Listen2(address string, handler func(*net.UDPAddr, int, []byte)) {
     // Parse the string address and
     //set variables that will
     //assist in the creation of new files
     //and prvent data from being read to a file
     //pre-maturely
     var upper int = 0
     var cooler int = 34
     var popper int = 1
     path := "/go/mydata/"+strconv.Itoa(popper)+".jpg"

     //bind to udp address
     addr, err := net.ResolveUDPAddr("udp", address)
     if err != nil {
     	log.Fatal(err)
     }
     
     // Open up a connection
     conn, err := net.ListenMulticastUDP("udp", nil, addr)
     if err != nil {
	   log.Fatal(err)
     }

    //set the read buffer
     conn.SetReadBuffer(maxDatagramSize)
     
     // Loop forever reading from the socket
     file1, err := os.OpenFile(path,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
     if err != nil {
            fmt.Println("File Creation error!!\n")
     }

     var checker int = 0
     var mimicker int = 0
     var flag1 int = 1
 
     for {
     	 
	 //make the read buffer
	 buffer := make([]byte, maxDatagramSize)
	 numBytes, _, _ := conn.ReadFromUDP(buffer)

	 //This chunk of code prevents the creation reading of
	 //data into a new file pre-maturely.
	 //As some files may contain useless data at the end.
	 if flag1 != 1{
	    mimicker,_=strconv.Atoi(string(buffer[:numBytes]))
	    if mimicker != 0{
	       flag1 = 1
	    }
	 }
	 if (numBytes != 0) && (flag1 == 1){
	    if cooler == 34{
	       cooler = 1
	       checker,_ = strconv.Atoi(string(buffer[:numBytes]))
	    } else{

		//keep track of how many bytes have are being
		//read from the buffer
	        upper = upper + numBytes



	   
		n, err := file1.Write(buffer)
	        _ = n
	        if err != nil{
	            fmt.Println("File writing error!\n\n")
	            os.Exit(1)
	        }
	    	fmt.Println((checker / upper) * 100)



        	//This is where the commented out msgHandler function
        	//was originally called. Feel free to uncomment in order
        	//to suit your needs.
	    	//handler(src, numBytes, buffer)



		//This chunk of code creates the next file
		//data will be transferred to
	    	if checker == upper{
	       	   upper = 0
	       	   flag1 = 5
	       	   cooler = 34
	       	   popper = popper + 1
	       	   path := "/go/mydata/"+strconv.Itoa(popper)+".jpg"
	       	   fmt.Println(path)
	       	   file1.Close()
	       	   file1, err = os.OpenFile(path,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	       	   if err != nil{
	       	      fmt.Println("File creation error!!!\n")
	       	   }
	    	}
            }
	 }
     }
}






//credit for code listening on port 3001 goes to user dmichael on github
func main() {
     var wg sync.WaitGroup

     wg.Add(1)

     //This function will act as the listener of the car
     go func(){
	    

	    app := cli.NewApp()
	    fmt.Println("asdflasdkjf;alskdfja;sdlkfj\n\n")
    	    app.Action = func(c *cli.Context) error {
    	       	       address := c.Args().Get(0)
	       	       if address == "" {
		       	  	  address = defaultMulticastAddress
		       }
		       fmt.Printf("Listening on %s\n", address)
		       Listen2(address, msgHandler)
		       return nil
	    }

	    app.Run(os.Args)


     }()

     wg.Wait()
}
