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

func main(){
	fmt.Println("test")
}