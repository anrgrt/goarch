package viewall

import (
"net/http"
_"github.com/go-sql-driver/mysql"
"encoding/json"
"structure"
"dbconnection"
)
func Viewall(w http.ResponseWriter, r *http.Request){
	var u2 []structure.User
	var u1 structure.User
	
	db, err1:= dbconnection.GetDb()
	checkErr(err1)
	rows,err2 := db.Query("select * from wallet_holders")
	checkErr(err2)
	
	for rows.Next(){
		rows.Scan(&u1.First,&u1.Last,&u1.Phone,&u1.Bal,&u1.Pass)
		u2=append(u2,u1)
	}		
	defer rows.Close()
	json.NewEncoder(w).Encode(u2)
	
}
func checkErr(err error){
	if err!=nil{
		panic(err)
	}
}
