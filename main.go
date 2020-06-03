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
	Birthday string
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

//membuat routes
func routes() {
	http.HandleFunc("/register", register)
}

//main function

func main(){
	connect_db()
	routes()

	defer db.Close()

	fmt.Println("Server running on port :8080")
	http.ListenAndServe(":8000", nil)
}

//fungsi Query User yang berguna untuk mengambil data pengguna 
//berdasarkan username
func QueryUser(username string) user {
	var users = user{}
	err = db.QueryRow(`
		SELECT id, 
		first_name, 
		last_name, 
		email,
		bithday,
		username,
		password 
		FROM users WHERE username=?
		`, username).
		Scan(
			&users.ID,
			&users.FirstName,
			&users.LastName,
			&users.Email,
			&users.Birthday,
			&users.Username,
			&users.Password
		)
	return users
}

//func register 
func register(w http.ResponseWriter, r *http.Request){
	//kode ini saya buat untuk mengecek/memvalidasi apakah method yang digunakan post atau tidak
	//jika method yang digunakan bukan post maka akan terredirect/menampilkan halaman register.html yang
	//ada di folder views
	if r.method != "POST" {
		http.Serverfile(w, r, "views/register.html")
		return
	}

	first_name := r.formValue("first_name")
	last_name := r.formValue("last_name")
	email := r.formValue("email")
	birthday := r.formValue("birthday")
	username := r.formValue("username")
	password := r.formValue("password")
	
	users := QueryUser(username) 


}