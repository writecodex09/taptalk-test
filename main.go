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
	Birthday  string
	Email string
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
		birthday,
		email,
		username,
		password 
		FROM users WHERE username=?
		`, username).
		Scan(
			&users.ID,
			&users.FirstName,
			&users.LastName,
			&users.Birthday,
			&users.Email,
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
	birthday := r.formValue("email")
	email := r.formValue("birthday")
	username := r.formValue("username")
	password := r.formValue("password")
	
	users := QueryUser(username) 

	if(user{}) == users {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		if len(hashedPassword) != 0 && checkErr(w, r, err){
			//code ini saya gunakan untuk mengecek apakah username yang dimasukan sudah ada atau belum didalam database,
			//jika tidak ada maka proses akan dilanjutkakn 
			stmt, err := db.Prepare("INSERT INTO users SET first_name=?, last_name=?, birthday=?, email=?, username=?, password=?")
			if err == nil {
				_, err := stmt.Exec(&first_name, &hashedPassword, &last_name, &username)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
		}
	}else{
			http.Redirect(w, r, "/register", 302)
	}
