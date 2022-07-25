package dao

import (
	"fmt"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestVersion(t *testing.T) {
	dsn := "pig:pLlGJSpdg1RCp^80@tcp(10.20.152.211:3306)/secure_check?charset=utf8mb4&parseTime=True&loc=Local"
	db := GetConnection(dsn)
	var componentMavenVersionDB ComponentMavenVersionDB = ComponentMavenVersionDB{
		Id:            1000000000,
		ComponentCode: "ss",
		Version:       "ss",
		Timestamp:     time.Now(),
		Packaging:     "jar",
		IsBlack:       true,
		CreateTime:    time.Now()}
	componentMavenVersionDB.Save(db)
	var comVersion ComponentMavenVersionDB
	result := db.First(&comVersion)
	fmt.Printf("type:%T\n", result)
	fmt.Println(comVersion)
	var comVersions []ComponentMavenDB
	db.Find(&comVersions)
	fmt.Println(comVersions)
}
func versionInsert(db *gorm.DB) {
	var componentMavenVersionDB ComponentMavenVersionDB = ComponentMavenVersionDB{
		Id:            1000000000,
		ComponentCode: "ss",
		Version:       "ss",
		Timestamp:     time.Now(),
		Packaging:     "jar",
		IsBlack:       true,
		CreateTime:    time.Now()}
	componentMavenVersionDB.Save(db)
}
