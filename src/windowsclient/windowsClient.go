package main

import (
    "net"
    "os"
    "fmt"
	"bufio"
	"strings"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
        os.Exit(1)
    }
    service := os.Args[1]

    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)

    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    checkError(err)

	
	reader := bufio.NewReader(os.Stdin)
	Loop:
		for{
			fmt.Print("Enter text: \n")			
			text, _ := reader.ReadString('\n')
			fmt.Printf("before trim %s size %d \n", text, len(text)) 
			text = strings.TrimRight(text, "\r\n")
			fmt.Printf("after trim %s size %d \n", text, len(text)) 
			_, err = conn.Write([]byte(text))
			checkError(err)
			if(text == "exit") {
				break Loop
			}				
		}  
    os.Exit(0)
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}