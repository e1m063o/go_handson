package services

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

var DB *gorm.DB

//ConDB DBに接続する
func ConDB() {
	spec := revel.Config.StringDefault("db.spec", "")

	driver := revel.Config.StringDefault("db.driver", "")
	DB, _ = gorm.Open(driver, spec+"?parseTime=True&loc=Local")

	debug := revel.Config.StringDefault("db.debug", "")
	debugBool, err := strconv.ParseBool(debug)
	if err != nil {
		debugBool = false
	}
	DB.LogMode(debugBool)

}