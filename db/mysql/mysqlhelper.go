package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var (
	db *sql.DB
	dataSourceName string
)


/* 连接数据库*/
func Open(uname string, upwd string, host string, port string, dbname string) (error) {
	//数据库连接
	//var dataSourceName string = "root:mustang@tcp(localhost:3306)/test?timeout=90s&charset=utf8"
	dataSourceName = uname + ":" + upwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?timeout=90s&charset=utf8"
	fmt.Printf(dataSourceName + "\n")

	checkConnect()
	return nil
}

func checkConnect()  {
	if db == nil {
		connect()
	}

}

/*连接数据库*/
func connect()  {
	db, _ = sql.Open("mysql", dataSourceName)
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
}

/*数据插入操作*/
func Insert(sqlcmd string, args ...interface{}) {
	checkConnect()

	//fmt.Println(sqlcmd + "\n")
	stmt, err := db.Prepare(sqlcmd)//`INSERT user (user_name,user_age,user_sex) values (?,?,?)`)
	checkErr(err)
	res, err := stmt.Exec(args...)//"tony", 20, 1)

	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id, "\n")
}

/*数据查询操作*/
func Query(sqlcmd string, args ...interface{}) ([]map[string]string){
	checkConnect()

	rows, err := db.Query(sqlcmd)//"SELECT * FROM user")
	checkErr(err)

	//字典类型
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	res := []map[string]string{}
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		res = append(res, record)
		//fmt.Println(record)
	}

	return res
}

/*数据更新操作*/
func Update(sqlcmd string, args ...interface{}) {
	checkConnect()

	stmt, err := db.Prepare(sqlcmd)//`UPDATE user SET user_age=?,user_sex=? WHERE user_id=?`)
	checkErr(err)
	res, err := stmt.Exec(21, 2, 1)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

/*删除数据*/
func Remove(sqlcmd string, args ...interface{}) {
	checkConnect()

	stmt, err := db.Prepare(sqlcmd)//`DELETE FROM user WHERE user_id=?`)
	checkErr(err)
	res, err := stmt.Exec(1)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

/*检查错误信息*/
func checkErr(err error) {
	if err != nil {
		println("mysql checkErr: " + err.Error())
		panic(err)
	}
}

