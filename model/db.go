package model

import (
	"fmt"
	"ginblog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var db *gorm.DB
var err error

func InitDb() {


	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
<<<<<<< HEAD
		DisableForeignKeyConstraintWhenMigrating: true, // 外键约束
		SkipDefaultTransaction: true,  // 禁用默认事务（提高运行速度）
=======
		//Logger:logger.Default.LogMode(logger.Silent), // gorm日志模式：silent
		DisableForeignKeyConstraintWhenMigrating: true,// 外键约束
		SkipDefaultTransaction: true, // 禁用默认事务（提高运行速度）
>>>>>>> 3399ea91f3e8c684fc63dc5c06fe24dd46f9b27a
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,  // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
		},
	})

	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
	}

	_ = db.AutoMigrate(&User{}, &Article{}, &Category{})

	sqlDB, _ := db.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	//db.Close()
}
