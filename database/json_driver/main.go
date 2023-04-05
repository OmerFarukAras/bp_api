package json_driver

import (
	"bp_api/database/model"
	"bp_api/database/util"
	"github.com/asdine/storm/v3"
	"time"
)

func Init() {
	db, err := storm.Open("json_driver.db")

	if err != nil {
		panic(err)
	}
	logData := model.Log{
		ID:        util.CreateUID(),
		Content:   "Database INIT. Successfully.",
		CreatedAt: time.Now(),
	}
	err = db.Save(&logData)
	if err != nil {
		return
	}

	defer func(db *storm.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
}

func CreateUser() (data, bool) {

}
