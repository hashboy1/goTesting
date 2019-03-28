package main
import (
	//"crypto/md5"
    "database/sql"
    //"encoding/hex"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    //"strconv"
    //"time"
 )
//CheckErr Function
 func CheckErr(err error) {
    if err != nil {
        panic(err)
        //fmt.Println("err:", err)
    }
}
//TextPrint function
func TextPrint(content string){
	fmt.Println(content)
}

func main() {
	var DBDriver string
	DBDriver = "vrkb:3dms@tcp(192.168.0.249:3306)/testdb?charset=utf8"
	var isOpen bool
	db, err := sql.Open("mysql", DBDriver)
	if err != nil {
			isOpen = false
	} else {
			isOpen = true
	}
	CheckErr(err)
	fmt.Println(isOpen)
	//fmt.Println(db)
	//return isOpen, db

    rows, err := db.Query("SELECT * FROM employee")
    CheckErr(err)
    if err != nil {
        fmt.Println("error:", err)
    } else {
    }
    for rows.Next() {
        var ID string
        var EmployeeNO string
        var RealName string
        var Password string
        var Role string
        var Department string
        CheckErr(err)
        err = rows.Scan(&ID, &EmployeeNO, &RealName, &Password, &Role, &Department)
        //fmt.Println(ID)
        //fmt.Println(EmployeeNO)
        //fmt.Println(RealName)
        //fmt.Println(Password)
        //fmt.Println(Role)
		//fmt.Println(Department)
		//fmt.Println(ID+EmployeeNO+RealName+Password+Role+Department)
		go TextPrint(ID+EmployeeNO+RealName+Password+Role+Department)
    }

}
  