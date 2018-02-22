package register

import (
"fmt"
"net/http"
_"github.com/go-sql-driver/mysql"
"encoding/json"
"structure"
"dbconnection"
)
func Register(w http.ResponseWriter, r *http.Request){
	
	var u structure.User
	
	json.NewDecoder(r.Body).Decode(&u)
	

	db, err1:= dbconnection.GetDb()
	checkErr(err1)
	rows,err2 := db.Query("select * from wallet_holders where Phone_number=?", u.Phone)
	checkErr(err2)
		
	defer rows.Close()

	if rows.Next(){
			u.Exists=true
			json.NewEncoder(w).Encode(u)
			//fmt.Fprintf(w,"Wallet Already Exists\n")
			//fmt.Fprintf(w,"bhai kitne account banayga ek hi number se\n")
			fmt.Println("Wallet Already Exists")
}else{
	u.Exists=false
	stmt,err3:= db.Prepare("INSERT wallet_holders SET First_name=?,Last_name=?,Phone_number=?, Balance=?, password=?" )
	checkErr(err3)	
	
	_,err4:= stmt.Exec(u.First,u.Last,u.Phone,u.Bal,u.Pass)
	checkErr(err4)
	json.NewEncoder(w).Encode(u)
	fmt.Println("Wallet Created Successfully :")
}
}
func checkErr(err error){
	if err!=nil{
		panic(err)
	}
}
