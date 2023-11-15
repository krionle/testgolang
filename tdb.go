package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB
var qs string

func main() {
	var err error
	print("请输入：\n")
	fmt.Scanf("%s", &qs)
	db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/testdb")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}
	//用于设置闲置的连接数。如果 <= 0, 则没有空闲连接会被保留
	db.SetMaxIdleConns(0)
	//用于设置最大打开的连接数,默认值为0表示不限制。
	db.SetMaxOpenConns(30)
	if err = db.Ping(); err != nil {
		log.Fatalln(err)
	}
	print("===================\n")
	print(qs + "\n")
	tq(qs)
}
func showtb() {
	rows, err := db.Query("select * from testtb")
	defer rows.Close()
	if err != nil {
		print(err)
		return
	}
	for rows.Next() {
		var id int64
		var name string
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}
func showdb() {
	tbn, err := db.Query("show databases")
	if err != nil {
		return
	}
	for tbn.Next() {
		var name string
		tbn.Scan(&name)
		fmt.Println(name)
	}
	return
}
func tq(sq string) {
	print(sq + "============\n")
	rows, err := db.Query("show databases")

	if err != nil {
		print(err)
	}
	for rows.Next() {
		var name string
		rows.Scan(&name)
		fmt.Println(name)
	}
	return
}
