package main

import(
	"io"
	//"io/ioutil"
	"log"
	"os"
)

var (
	Trace 	*log.Logger
	Info 	*log.Logger
	Warning *log.Logger
	Error	*log.Logger
)

func Init(
	traceHandle   io.Writer,
	infoHandle    io.Writer,
	warningHandle io.Writer,
	errorHandle   io.Writer){
	
	Trace = log.New(traceHandle,
					"Trace: ",
					log.Ldate|log.Ltime|log.Lshortfile)
	
	Info = log.New(infoHandle,
					"INFO: ",
					log.Ldate|log.Ltime|log.Lshortfile)
	
	Warning = log.New(warningHandle,
					"WARNING: ",
					log.Ldate|log.Ltime|log.Lshortfile)
					
	Error = log.New(errorHandle,
					"ERROR: ",
					log.Ldate|log.Ltime|log.Lshortfile)
}

func main(){

	//Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	file, err := os.OpenFile("Tracefile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if(err != nil){
		log.Printf("cannot generate log file Tracefile")
	}
	
	file1, err := os.OpenFile("Infofile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if(err != nil){
		log.Printf("cannot generate log file Infofile")
	}
	
	file2, err := os.OpenFile("Warningfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if(err != nil){
		log.Printf("cannot generate log file Warningfile")
	}
	
	file3, err := os.OpenFile("Errorfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if(err != nil){
		log.Printf("cannot generate log file Errorfile")
	}
	
	Init(file, file1, file2, file3)
	
    Trace.Println("I have something standard to say")
    Info.Println("Special Information")
    Warning.Println("There is something you need to know about")
    Error.Println("Something has failed")
	
}

