package main

import (
	"fmt"
	"database/sql"
	"github.com/didi/gendry/manager"
	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/scanner"
	_ "github.com/go-sql-driver/mysql"
)

const (
	defaultDriver = "mysql"
	user = "root"
	password = "root"
	host = "localhost"
	port = 3306
	dbName ="godemo"
	setstr = "charset=utf8"

	format = "%s:%s@tcp(%s:%d)/%s?%s"
)

func main()  {
	fmt.Println("hello database! datasource test")
	var db *sql.DB
	db,err := open(defaultDriver,createRealDSN())
	if err != nil{
		fmt.Println("db open error err:",err)
		return
	}

	err = testPing(db)
	if err != nil{
		fmt.Println("db connection error err:",err)
	}
	queryUser(db)
	testapi()


}
func testPing(db *sql.DB)  error{
	err := db.Ping();
	fmt.Println("err",err,"db",db)
	return err
}
func createRealDSN()string{
	realDSN := fmt.Sprintf(format, user, password, host, port, dbName, setstr)
	fmt.Println("realDSN",realDSN)
	return realDSN
}

func testapi()  {
	var db2 *sql.DB
	var err2 error
	db2, err2 =manager.New(dbName,user,password,host).Set(
		manager.SetCharset("utf8"),
	).Port(port).Open(true)

	fmt.Println("err2",err2,"db2",db2)
}
func open(driver,rdsn string) (*sql.DB, error) {
	return sql.Open(driver, rdsn)
}

func queryUser(db *sql.DB) error {
	where := map[string]interface{}{
		"account":"chen",
	}
	table := "sys_user"
	field :=[]string{"account","address"}
	con,values,err := builder.BuildSelect(table,where,field)
	fmt.Println("con:",con,"values:",values,"err:",err)

	var account,address string
	err = db.QueryRow(con,"chen").Scan(&account,&address)
	fmt.Println("account:",account,"address",address,"err:",err)



	//field = []string{"id","account","password","name","address"}
	con,values,err = builder.BuildSelect(table,where,nil)
	fmt.Println("con:",con)

	rows ,err := db.Query(con,values...)

	var user []SysUser
	scanner.Scan(rows,&user)

	fmt.Println("user:",user)
	return err
}

type SysUser struct  {
	Id int `ddb:"id"`
	Account string `ddb:"account"`
	Password string `ddb:"password"`
	Name string `ddb:"name"`
	Address string `ddb:"address"`
}