package main

import "bp_api/json_driver"

func main() {
	db, ok := json_driver.Init()
	if !ok {
		panic("DB error")
	}
	db.Test()
}
