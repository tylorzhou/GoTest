package main

import (
	"fmt"
	"log"
	"net"
	"encoding/gob"
)
type P struct{
	M, N int64
}

func main(){
	fmt.Println("start client")
	conn, err := net.Dial("tcp", "localhost:8008")
	if err != nil {
		log.Fatal("Connection error", err)
	}
	encoder := gob.NewEncoder(conn)
	p := &P{3, 5}
	encoder.Encode(p)
	conn.Close()
	fmt.Println("done")
}