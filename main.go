package main

//import pacakge dan library 

import (

	 "database/sql"
	 "fmt"
	 "html/template"
	 "log"
	 "net/http"
	 "golang.org/x/crypto/bcrypt"
	 _ "github.com/go-sql-driver/mysql"
	"github.com/kataras/go-sessions"
)

//membuat variabel db dan err

var db *sql.DB
var err error

//membuat struct user 
type user struct {
	ID        int
	FirstName  string
	LastName string
	Email  string
	Username  string
	Password  string
}

//membuat fungsi untuk melakukan koneksi kedalam database
func connect_db(){
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1)/taptalk")

	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err !=nil {
		log.Fatalln(err)
	}
}

func main(){
	fmt.Println("test")
}