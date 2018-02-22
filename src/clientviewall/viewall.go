package clientviewall

import 
(
"fmt"
"net/http"
_"github.com/go-sql-driver/mysql"
"encoding/json"
"structure"
)

func ViewAll(){
	url:="http://127.0.0.1:8080/viewall"
	var u2 []structure.User
	resp,_:=http.Get(url)
	json.NewDecoder(resp.Body).Decode(&u2)
	fmt.Println("All wallet holder in Database")
	for i:=0;i<len(u2);i++{
		fmt.Println("\t\t",u2[i])
	}
}