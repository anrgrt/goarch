package main

import(
"net/http"
_"github.com/go-sql-driver/mysql"
"update"
"register"
"viewall"
"viewmy"
)



func main(){
	http.HandleFunc("/update",update.Update)
	http.HandleFunc("/register",register.Register)
	http.HandleFunc("/viewall",viewall.Viewall)
	http.HandleFunc("/viewmy",viewmy.Viewmy)
	http.ListenAndServe(":8080",nil)
}
