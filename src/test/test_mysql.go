package main

import (
	"fmt"
	"goApi/src/db"
)

func main() {
	sql := "select * from qa_xiaomu limit 10;"
	error := db.GetDataByQuery("mysql", "qa_log", sql)
	fmt.Println(error)
}
