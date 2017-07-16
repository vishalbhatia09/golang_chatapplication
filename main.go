package main

import "fmt"
import "net"
import "bufio"

func main(){
	conns := make(chan net.Conn)
	msgs := make(chan string)
	tc,err := net.Listen("tcp", "127.0.0.1:3030")
	if err != nil{
		fmt.Println(err)
	}


	go func(){
		for{
			conn,err := tc.Accept()
			if err != nil{
				fmt.Println(err)
			}
			conns <- conn
		}
	}()

	for{
		select{
		case conn := <- conns:
			go func(conn net.Conn){
				rdmsg := bufio.NewReader(conn)
				for{
					m,err := rdmsg.ReadString('\n')
					if err!= nil{
						fmt.Println(err)
						break
					}
					msgs <- m 
				}
			}(conn)

		case msg := <-msgs :
			for conn := range conns{
				conn.Write([]byte(msg))
			} 
		}
	}

}