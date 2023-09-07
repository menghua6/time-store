package db

import (
	"fmt"
	"log"
	"time"

	"github.com/menghua6/time-store/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlDb *gorm.DB

func InitMysql() error {
	url := "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local"
	constr := fmt.Sprintf(url, MysqlUserName, MysqlPassword, MysqlHost, MysqlPort, MysqlDbName)
	//连接数据库
	db, err := gorm.Open(mysql.Open(constr), &gorm.Config{})
	if err != nil {
		return err
	}
	log.Println("MysqlDb初始化成功")

	db.AutoMigrate(entity.Job{})

	mysqlDb = db

	return nil
}

func GetJobsByTime(st time.Time, et time.Time) ([]entity.Job, error) {
	jobs := make([]entity.Job, 0)
	dbRes := mysqlDb.Where("start-time <= ?", et).Where("end-time >= ?", st).Find(&jobs)
	if dbRes.Error != nil {
		return nil, dbRes.Error
	}
	return jobs, nil
}
