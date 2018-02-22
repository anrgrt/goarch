package dbconnection

import
(
"database/sql"
)

func GetDb() (*sql.DB,error) {
return sql.Open("mysql", "root:@tcp(127.0.0.1:3307)/mydb_test")
}
