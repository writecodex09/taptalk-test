package main

//import pacakge dan library 

import (
	 "database/sql"
	 "fmt"
	 "log"
	 "net/http"
	 "golang.org/x/crypto/bcrypt"
	 _"github.com/go-sql-driver/mysql"
	 "github.com/kataras/go-sessions"
	// "os"
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

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
//fungsi login
func login(w http.ResponseWriter, r *http.Request){
	//disini kita akan mengecek session 
	//code yang ada dibawah ini saya gunakan untuk membuat session
	session := sessions.Start(w, r)
	if len(session.GetString(username)) != 0 && checkErr(w, r, err){
		http.Redirect(w, r, "/", 302)
	}
	//disini saya mencek apakah bila ada session , dan session itu benar bermethod post atau tidak 
	if r.Method != "POST" {
		//jika tidak maka akan diredirect kehalaman login
		http.ServeFile(w, r, "views/login.html")
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	//kita akan mengambil data pengguna berdasarkan username 
	users := QueryUser(username)

	//disini saya melakukan pengecekan 
	//dimana saya akan melakukan perbandingan password yang ada didatabase
	var password_tes = bcrypt.CompareHashAndPassword([]byte(users.Password),[]byte(password))

	if password_tes == nil {
		session := sessions.Start(w, r)
		session.Set("username", users.Username)
		session.Set("password", users.Password)
		http.Redirect(w, r, "/", 302))
	}else{
		http.Redirect(w, r, "/login", 302)
	}
}
//func checkErr
func checkErr(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {

		fmt.Println(r.Host + r.URL.Path)

		http.Redirect(w, r, r.Host+r.URL.Path, 301)
		return false
	}

	return true
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
		FROM users WHERE username=?`, username).
		Scan(
			&users.ID,
			&users.FirstName,
			&users.LastName,
			&users.Birthday,
			&users.Email,
			&users.Username,
			&users.Password,
		)
	return users
}

//func register 
func register(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		http.ServeFile(w, r, "views/register.html")
		return
	}

	first_name := r.FormValue("first_name")
	last_name := r.FormValue("last_name")
	birthday := r.FormValue("birthday")
	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")
	
	users := QueryUser(username) 

	if(user{}) == users {
		//disini saya akan mengenkripsi password yang dimasukan oleh user untuk menunjang keamanan data
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		if len(hashedPassword) != 0 && checkErr(w, r, err) {
			//pada bagian ini saya akan cek apakah username yang ingin di daftarkan dalam sistem/aplikasi
			//sudah ada atau belum, jika belum ada maka proses akan dilanjutkan 
			stmt, err := db.Prepare("INSERT INTO users SET first_name=?, last_name=?, birthday=?, email=?, username=?, password=?")
			if err == nil {
				_, err := stmt.Exec(&first_name, &last_name, &birthday, &email, &username, &hashedPassword)
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
}