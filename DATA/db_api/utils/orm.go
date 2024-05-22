package utils

import "gorm.io/gorm"

func ORMInit() {
	//创建一个数据库的连接
	var err error
	db, err = gorm.Open("mysql", "root:12345@/demo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	//迁移the schema
	db.AutoMigrate(&todoModel{})
}
