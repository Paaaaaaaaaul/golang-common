package model

import (
	"os"
)

func G_model(projectName string) error {
	if err := os.Mkdir(projectName+"/model", 755); err != nil {
		return err
	}

	file, err := os.OpenFile(projectName+"/model/example.go", os.O_CREATE|os.O_RDWR, 755)
	if err != nil {
		return err
	}

	if _, err := file.WriteString(model_temple); err != nil {
		return err
	}

	return nil
}

var model_temple = `package model

type User struct {
	Id     int    ` + "`gorm:\"column:id;primary_key;AUTO_INCREMENT\"`" + `
	UserId int64  ` + "`gorm:\"column:userId\"`" + `
	Name   string  ` + "`gorm:\"column:string\"`" + `
}

func (u *User) TableName() string {
	return "user"
}
`
