package main

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(42.194.141.33:3306)/testdb?charset=utf8mb4&parseTime=true")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}
	//用于设置闲置的连接数。如果 <= 0, 则没有空闲连接会被保留
	db.SetMaxIdleConns(0)
	//用于设置最大打开的连接数,默认值为0表示不限制。
	db.SetMaxOpenConns(30)
	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}
	showdbs()
}
func showdbs() {
	rows, err := db.Query("select * from testtb")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int64
		var name string
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
}
