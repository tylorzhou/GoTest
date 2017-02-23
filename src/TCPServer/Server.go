package main
import(
	"fmt"
	"net"
	"encoding/gob"
)

type P struct{
	M, N int64
}

func handleConnection(conn net.Conn){
	dec := gob.NewDecoder(conn)
	p := &P{}
	dec.Decode(p)
	fmt.Printf("Received : %+v", p);
	conn.Close();
}

func main(){
	fmt.Println("start")
	ln, err := net.Listen("tcp", ":8008")
	if err != nil{
	}//handle error
	
	for{
		conn, err := ln.Accept()
		if err != nil{
		  continue;
		}
		go handleConnection(conn)
	}
}	