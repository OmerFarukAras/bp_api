package json_driver

import (
	"bp_api/model"
	"bp_api/util"
	"github.com/asdine/storm/v3"
)

type Database struct {
	cl *storm.DB
}

func Init() (Database, bool) {
	db, err := storm.Open("json_driver.db")
	if err != nil {
		return Database{}, false
	}
	dbs := Database{cl: db}

	dbs.cl = db

	/*logData := model.Log{
		ID:        "util.CreateUaID()",
		Content:   "Database INIT. Successfully.",
		CreatedAt: time.Now(),
	}*/

	//err = db.Save(&logData)
	/*if err != nil {
		return dbs, false
	}*/

	defer func(db *storm.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)
	return dbs, true
}

func (dbs *Database) CreateUser(user model.User) (model.User, bool) {
	ok := util.CheckString(user.ID)
	if !ok {
		user.CreateID()
	}

	user.SetCreationTime()
	user.SetUpdateTime()

	err := dbs.cl.Save(&user)
	if err != nil {
		return user, false
	}
	return user, true
}

func (dbs *Database) Test() {
	println("Db tested.")
}
