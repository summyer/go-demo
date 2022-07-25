package dao

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type ComponentMavenVersionDB struct {
	Id            int64
	ComponentCode string
	Version       string
	Timestamp     time.Time
	Packaging     string
	IsBlack       bool
	CreateTime    time.Time
}

func (c *ComponentMavenVersionDB) Save(db *gorm.DB) {
	result := db.Create(c)
	fmt.Println(result.Error)
	fmt.Println(c.Id)
	fmt.Println(result.RowsAffected)
}
