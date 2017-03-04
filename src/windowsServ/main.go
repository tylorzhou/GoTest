package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	
	"golang.org/x/sys/windows/svc"
)

func usage(errmsg string) {
	fmt.Fprintf(os.Stderr,
	"%s\n\n"+
		"usage: %s <command>\n"+
		"		where <command> is one of\n"+
		"		install, remove, debug, start, stop, pause or continue.\n",
	errmsg, os.Args[0])
	os.Exit(2)
}

func main() {
	const svcName = "myservice"
	
	isIntSess, err := svc.IsAnInteractiveSession()
	if err != nil {
		log.Fatalf("failed to determine if we are running in an interactive session: %v", err)
	}	
	initLogFile(svcName)
	
	if !isIntSess {
		runService(svcName, false)
		return
	}
	
	if len(os.Args) < 2 {
		usage("no command specified")
	}
	
	cmd := strings.ToLower(os.Args[1])
	switch cmd {
		case "debug":
			Info.Printf("debug mode server: " + svcName)
			runService(svcName, true)
			return
		case "install":
			Info.Printf("install server: " + svcName)
			err = installService(svcName, "My service display")
		case "remove":
			Info.Printf("remove server: " + svcName)
			err = removeService(svcName)
		case "start":
			Info.Printf("start server: " + svcName)
			err = startService(svcName)
		case "stop":
			Info.Printf("stop server: " + svcName)
			err = controlService(svcName, svc.Stop, svc.Stopped)
		case "pause":
			Info.Printf("pause server: " + svcName)
			err = controlService(svcName, svc.Pause, svc.Paused)
		case "continue":
			Info.Printf("Continue server: " + svcName)
			err = controlService(svcName, svc.Continue, svc.Running)
		default:
			usage(fmt.Sprintf("invalid command  %s", cmd))	
	}
	if err != nil {
		log.Fatalf("failed to %s %s: %v", cmd, svcName, err)
	}
	return	
}