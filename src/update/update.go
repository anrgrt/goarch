package update

import (
"net/http"
_"github.com/go-sql-driver/mysql"
"encoding/json"
"structure"
"dbconnection"
)

func Update(w http.ResponseWriter, r *http.Request){
	
	var u1,u2 structure.User
	json.NewDecoder(r.Body).Decode(&u2)
	

	db,err1:= dbconnection.GetDb()
	checkErr(err1)
	rows,err2 := db.Exec("update wallet_holders set First_name=?,Last_name=?,Balance=? where Phone_number=?;",u2.First,u2.Last,u2.Bal,u2.Phone)
	checkErr(err2)
		
	//defer rows.Close()
	a,_:=rows.RowsAffected()

	if a>0{
			u1.Exists=true
			
	}else{
		u1.Exists=false
	}
	json.NewEncoder(w).Encode(u1)

}	
func checkErr(err error){
	if err!=nil{
		panic(err)
	}
}
