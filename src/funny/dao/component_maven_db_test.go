package dao

import (
	"fmt"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestComponentMaven(t *testing.T) {
	dsn := "pig:pLlGJSpdg1RCp^80@tcp(10.20.152.211:3306)/secure_check?charset=utf8mb4&parseTime=True&loc=Local"
	db := GetConnection(dsn)
	var com ComponentMavenDB
	result := db.First(&com)
	fmt.Printf("type:%T\n", result)
	fmt.Println(com)
	var coms []ComponentMavenDB
	db.Find(&coms)
	fmt.Println(com)
}
func testComponentMavenInsert(db *gorm.DB) {
	var componentMavenDb ComponentMavenDB = ComponentMavenDB{
		Code:           "sss",
		Group:          "ss",
		Artifact:       "ss",
		Merge:          "ss",
		LastUpdateTime: time.Now(),
		CreateTime:     time.Now()}
	componentMavenDb.Save(db)
}
