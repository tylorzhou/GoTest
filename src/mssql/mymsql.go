package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var debug = flag.Bool("debug", false, "enable debugging")
var password = flag.String("password", "", "the database password")
var port = flag.Int("port", 1433, "the database port")
var server = flag.String("server", "", "the database server")
var user = flag.String("user", "", "the databse user")

func main() {
	flag.Parse() //parse the command line args

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", *server, *user, *password, *port)
	*debug = true
	//connString := fmt.Sprintf("server=%s\\SQLExpress;user id=%s;password=%s;port=%d", "10.40.1.45", "sa", "Altigen1234", 1433)

	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("open connection failed:", err.Error())
	}
	defer conn.Close()

	str1 := `SELECT [systemId] ,[ServerAddress] FROM [CWSDB1].[dbo].[MaxACDAS]`
	stmt, err := conn.Prepare(str1)
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	var somenumber int64
	var somechars string
	err = row.Scan(&somenumber, &somechars)
	if err != nil {
		log.Fatal("scan failed:", err.Error())
	}
	fmt.Printf("somenumber:%d\n", somenumber)
	fmt.Printf("somechars:%s\n", somechars)

}
