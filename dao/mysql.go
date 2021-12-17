package dao

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"runtime"
)

var (
	DB *gorm.DB
)

//func CreateDatabase(dsn string) error {
//	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `buuble` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;")
//	db, err := sql.Open("mysql", dsn)
//	if err != nil {
//		return err
//	}
//	defer func(db *sql.DB) {
//		err = db.Close()
//		if err != nil {
//			fmt.Println(err)
//		}
//	}(db)
//	if err = db.Ping(); err != nil {
//		return err
//	}
//	_, err = db.Exec(createSql)
//	return err
//}
func InitMySQL() (err error) {
	//if err = CreateDatabase("root:161702@tcp(127.0.0.1:3306)/"); err != nil {
	//	fmt.Printf("%s", err)
	//	return err
	//}
	sysType := runtime.GOOS
	var pass string
	if sysType == "linux" {
		pass = "123456"
	}

	if sysType == "windows" {
		pass = "123456"
	}
	dsn := "root:" + pass + "@tcp(127.0.0.1:3306)/buuble?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//db.AutoMigrate(models.Todo{})
	return
}
