package models

import (
	db "go_api/database"
)

type Members struct {
	ID int64 `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	RememberToken string `json:"remember_token" form:"remember_token"`
}

var Users []Members

//添加
func (user Members) Insert() (id int64, err error) {

	//添加数据
	result := db.PostgresDb.Create(&user)
	id =user.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//列表
func (user *Members) Users() (users []Members, err error) {
	if err = db.PostgresDb.Find(&users).Error; err != nil {
		return
	}
	return
}

//修改
func (user *Members) Update(id int64) (updateUser Members, err error) {

	if err = db.PostgresDb.Select([]string{"id", "username"}).First(&updateUser, id).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = db.PostgresDb.Model(&updateUser).Updates(&user).Error; err != nil {
		return
	}
	return
}

//删除数据
func (user *Members) Destroy(id int64) (Result Members, err error) {

	if err = db.PostgresDb.Select([]string{"id"}).First(&user, id).Error; err != nil {
		return
	}

	if err = db.PostgresDb.Delete(&user).Error; err != nil {
		return
	}
	Result = *user
	return
}