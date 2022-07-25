package dao

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type ComponentMavenDB struct {
	Code, Group, Artifact, Merge string
	LastUpdateTime               time.Time
	CreateTime                   time.Time
}

func (c *ComponentMavenDB) Save(db *gorm.DB) {
	result := db.Create(c)
	fmt.Println(result.Error)
	fmt.Println(c.Code)
	fmt.Println(result.RowsAffected)
}

/*func (ComponentMavenDB) TableName() string {
	return "component_maven_db"
}*/
