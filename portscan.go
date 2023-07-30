package main 

import (
    "log"
    "net"
    "fmt"
    "sync"
)

var wg sync.WaitGroup

func scanPort(i int){
   defer wg.Done()
   address := fmt.Sprintf("127.0.0.1:%d", i)
   conn, e := net.Dial("tcp", address)

   if e != nil {
     return
   } 
   defer conn.Close()  
 
   log.Printf("%d is open\n", i)
   
}

func main(){
   for i := 0; i < 65535; i++{
     wg.Add(1)
     go scanPort(i)
   }

 wg.Wait()
}
