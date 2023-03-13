package models

import (
	
	"gorm.io/gorm"
)

type Fammembers struct {
	gorm.Model
	Firstname  string
	Happiness int
	Urlstr string
};

//create a user
func CreateUser(db *gorm.DB, User *Fammembers) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	};
	return nil
}

//get users
func GetUsers(db *gorm.DB, User *[]Fammembers) (err error) {
	err = db.Find(User).Error;
	if err != nil {
		return err;
	};	return nil}

//get user by id
func GetUser(db *gorm.DB, User *Fammembers, id int) (err error) {
	err = db.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	};
	return nil
}

//update user
func UpdateUser(db *gorm.DB, User *Fammembers) (err error) {
	db.Save(User)
	return nil
}

//delete user
func DeleteUser(db *gorm.DB, User *Fammembers, id int) (err error) {
	db.Where("id = ?", id).Delete(User)
	return nil
}