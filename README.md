# taptalk-test for authentication
Repo ini berisikan test untuk posisi backend developer di taptalk.io 

1. Langkah pertama yang saya lakukan adalah melakukan inisialisasi project
dengan menggunakan command 
mkdir taptalk.io
cd taptalk.io
go mod init taptalk

2. Langkah kedua yang saya lakukan adalah membuat database yang akan menyimpan data user 
disini saya menggunakan SQL Database yaitu MySQL
Disini saya memodifikasi sedikit dari requirement yang diberikan , dimana saya mengganti atribut user yang semula fullname
berganti menjadi firt_name dan last_name

3. Langkah ketiga yang saya lakukan adalah membuat folder views dimana nanti nya akan berisi template engine yang akan mempercantik tampilan website dalam project test kali ini

4. Langkah ke 4 saya membuat file register.html, login.html dan home.html didalam folder views

5. Langkah ke 5 saya menginstall library yang dibutuhkan seperti sql driver, cryptto/bcrypt dan kataras/go-sessions karena dalam requirement yang diberikan oleh taptalk.io user can only have 1 active session at one time

6. Langkah ke 6 saya mengimport beberapa package yang akan kita gunakan sepertii net/http dan lainnya. beserta library yang 
sudah saya download sebelumnya 

7. Langkah selanjutnya saya membuat 2 buah variabel yaitu var err dan  var db