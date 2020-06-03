package main

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

var db *sql.DB
var err error

type user struct {
	ID        int
	FirstName  string
	LastName string
	Email  string
	Username  string
	Password  string
}

func main(){
	fmt.Println("test")
}