package models

import (
	"log"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&User{},
		// &RoleGroup{}, // thêm model khác ở đây nếu có
	)
	if err != nil {
		log.Fatal("❌ Lỗi migrate database: ", err)
	}
	log.Println("✅ Auto migrate database thành công!")
}
