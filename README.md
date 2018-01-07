# learn-golang

Learning go language
====================

Beberapa file tentang belajar golang
-----------------------------

1. basic.go
***********
Berisi contoh tipe-tipe data di golang, bagaimana deklarasi variabel, konstanta, serta inisialisasi variabel. Juga mekanisme import di golang, alias import, dan mekanisme import dalam satu level 


2. collections.go
*****************
Berisi contoh tipe data yang berbentuk koleksi. deklarasi array, inisialisasi dan operasinya.
Kemudian ada tipe slice serta map beserta contohnya.


3. database.go
**************
Berisi contoh bagaimana berinteraksi dengan sql database driver, dalam contoh ini menggunakan mysql driver.


4. app.go
*********
Berisi contoh sederhana aplikasi berbasis http/web menggunakan golang net/http dengan dukungan beberapa library:
- database/sql, standard interface sql
- go-sql-driver/mysql, driver go untuk mysql
- go-jwt, golang json web token library
- gorilla/mux, router multiplexer dari gorilla
- gorilla/handlers, beberapa middleware berguna dari gorilla
- jwt-go/request, utility terkait http request dan jwt
- token.go, contoh sederhana package sendiri terkait jwt 
- beberapa library standar golang


5. token/token.go
*****************
Contoh membuat package sendiri yang digunakan di app.go di atas.


6. static/index.html
***********************
Contoh file yang digunakan app.go di atas, untuk bagaimana merender file html menggunakan fasilitas library html/template.
Termasuk integrasi dengan google polymer webcomponent library.


7. keys/*
*********
File-file public dan private rsa key sebagai contoh untuk signing key jwt payload yang digunakan di app.go yang digenerate dengan openssl command.


8. fileserver/main.go
*********************
Berisi contoh http file server dengan net/http


9. cleanarch/*
**************
Berisi contoh implementasi clean architecture dengan golang.
Lihat <===> terkait clean architecture.


10. main.go
***********
Entry point web app dengan clean architectures tersebut di atas.



===== semoga bermanfaat and happy learning =======

===== contact: green.shirt@yandex.com =====