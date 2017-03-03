package main

import (
	"net"
	"time"
		
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/debug"
)


type myservice struct {}

func (m *myservice) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown | svc.AcceptPauseAndContinue
	changes <- svc.Status{State: svc.StartPending}
	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
	exitFromMain := make(chan uint32)
	
	go mainfunc(exitFromMain)
loop:
	for {
		select {
		case c := <-r:
			switch c.Cmd {
				case svc.Interrogate:
					changes <- c.CurrentStatus
					// Testing deadlock from https://code.google.com/p/winsvc/issues/detail?id=4
					time.Sleep(100 * time.Millisecond)
					changes <- c.CurrentStatus
				case svc.Stop, svc.Shutdown:
					break loop
				case svc.Pause:
					changes <- svc.Status{State: svc.Paused, Accepts: cmdsAccepted}
				case svc.Continue:
					changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
				default:
					Error.Printf("unexpected control request #%d", c)
			}
		case e := <- exitFromMain:
			return true, e
		}
	}
	changes <- svc.Status{State: svc.StopPending}
	return
}

func runService(name string, isDebug bool){
	var err error
	
	Info.Printf("starting %s service", name)
	run := svc.Run
	if isDebug {
		run = debug.Run
	}
	err = run(name, &myservice{})
	if err != nil {
		Error.Printf("%s service failed: %v", name, err)
		return
	}
	Info.Printf("%s service stopped", name)
}

func mainfunc(r chan<- uint32) {
	ln, err := net.Listen("tcp", ":8008")
	if err != nil{
	}//handle error
	
	for{
		conn, err := ln.Accept()
		if err != nil{
		  continue;
		}
		go handleConnection(r, conn)
	}
	r <- 0
}

func handleConnection(r chan<- uint32,conn net.Conn){
	//dec := gob.NewDecoder(conn)
	//p := &P{}
	//dec.Decode(p)
	//fmt.Printf("Received : %+v", p);
	//conn.Close();
}