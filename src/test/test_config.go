package main

import (
	"fmt"
	"goApi/src/config_log"
)

func testConf() {
	r := config_log.GetConfigValue("mysql", "host")
	fmt.Println(r)
}

func main()  {
	testConf()
}