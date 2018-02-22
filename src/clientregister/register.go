package clientregister

import
(
"fmt"
"net/http"
_"github.com/go-sql-driver/mysql"
"encoding/json"
"structure"
"bytes"
)

func Register(){
	
	var u,u1 structure.User

	fmt.Println("To Register enter ur first name,last name, phone no, amount, password in order")
	fmt.Scan(&u.First,&u.Last,&u.Phone,&u.Bal,&u.Pass)

	b:=new(bytes.Buffer)

	json.NewEncoder(b).Encode(u)

	url:="http://127.0.0.1:8080"

	resp,_:=http.Post(url+"/register","application/json; charset=utf-8",b)
	
	json.NewDecoder(resp.Body).Decode(&u1)

	if u1.Exists{
		fmt.Println("Wallet Already Exists, for the number\n\t",u1.Phone)
	}else{

		//fmt.Println("Json Format\n")
		//io.Copy(os.Stdout, resp.Body)
		//json.NewDecoder(resp.Body).Decode(&u1)
		
		fmt.Println("\nWallet Created Succesfully, the details are:\n\t",u1.First,"\n\t",u1.Last,"\n\t",u1.Phone,"\n\t",u1.Bal,"\n\t",u1.Pass)

	
	}
	

}