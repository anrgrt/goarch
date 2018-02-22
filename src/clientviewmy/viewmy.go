package clientviewmy

import
(
"fmt"
"bytes"
"net/http"
_"github.com/go-sql-driver/mysql"
"encoding/json"
"structure"
)

func ViewMy(){
	var u,u1 structure.User 
	url:="http://127.0.0.1:8080/viewmy"
	fmt.Println("Enter your Phonenumber and password")
	fmt.Scan(&u.Phone,&u.Pass)
	
	b:=new(bytes.Buffer)

	json.NewEncoder(b).Encode(u)

	

	resp,_:=http.Post(url,"application/json; charset=utf-8",b)
	
	json.NewDecoder(resp.Body).Decode(&u1)
	
	if u1.Exists{
		fmt.Println("\nYour Details are:\n\t",u1.First,"\n\t",u1.Last,"\n\t",u1.Phone,"\n\t",u1.Bal,"\n\t",u1.Pass)

	}else{
		fmt.Println("Invalid Phonenumber or password")
	}
}