package viewmy

import (
"net/http"
_"github.com/go-sql-driver/mysql"
"encoding/json"
"structure"
"dbconnection"
)
func Viewmy(w http.ResponseWriter, r *http.Request){
	var u1,u2 structure.User

	json.NewDecoder(r.Body).Decode(&u2)
	

	db, err1:= dbconnection.GetDb()
	checkErr(err1)
	rows,err2 := db.Query("select * from wallet_holders where Phone_number=?", u2.Phone)
	checkErr(err2)
		
	defer rows.Close()

	if rows.Next(){
			u1.Exists=true
			rows.Scan(&u1.First,&u1.Last,&u1.Phone,&u1.Bal,&u1.Pass)

			json.NewEncoder(w).Encode(u1)			
}else{
	u1.Exists=false
	json.NewEncoder(w).Encode(u1)

}
}
func checkErr(err error){
	if err!=nil{
		panic(err)
	}
}
