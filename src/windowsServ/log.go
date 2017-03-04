package main

import(
	"io"
	"path/filepath"
	"log"
	"os"
)

var (
	Trace 	*log.Logger
	Info 	*log.Logger
	Warning *log.Logger
	Error	*log.Logger
)

var ServDir string

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



func initLogFile(name string){

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
            log.Fatal(err)
    }	
	_ = os.Mkdir(dir + "\\log", os.ModeDir)

	ServDir = dir + "\\log"

	file, err := os.OpenFile(ServDir + "\\Tracefile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if(err != nil){
		log.Printf("cannot generate log file Tracefile: %v", err)
	}
	
	file1, err := os.OpenFile(ServDir + "\\Infofile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if(err != nil){
		log.Printf("cannot generate log file Infofile: %v", err)
	}
	
	file2, err := os.OpenFile(ServDir + "\\Warningfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if(err != nil){
		log.Printf("cannot generate log file Warningfile: %v", err)
	}
	
	file3, err := os.OpenFile(ServDir + "\\Errorfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if(err != nil){
		log.Printf("cannot generate log file Errorfile: %v", err)
	}
	
	Init(file, file1, file2, file3)
}

