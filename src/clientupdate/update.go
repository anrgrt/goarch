package clientupdate

import
(
"fmt"
"bytes"
"net/http"
_"github.com/go-sql-driver/mysql"
"encoding/json"
"structure"
)

func Update(){
	
var u,u1 structure.User 
	url:="http://127.0.0.1:8080/update"
	fmt.Println("Enter your all the deatils to be updated along with Phonenumber and password")
	fmt.Scan(&u.First,&u.Last,&u.Phone,&u.Bal,&u.Pass)
	
	b:=new(bytes.Buffer)

	json.NewEncoder(b).Encode(u)

	

	resp,_:=http.Post(url,"application/json; charset=utf-8",b)
	
	json.NewDecoder(resp.Body).Decode(&u1)
	
	if u1.Exists{
		fmt.Println("Ur record is up to date")

	}else{
		fmt.Println("Invalid Phonenumber or password")
	}
	
}