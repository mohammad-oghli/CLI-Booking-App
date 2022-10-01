package models

import (
	"gorm.io/gorm"
	"oghli.pro/cli-booking-app/db_config"
)

var db *gorm.DB

type Conference struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"unique;not null"`
	TotalTickets uint   `gorm:"not null"`
	Users        []User `gorm:"foreignKey:Conf_ID"`
}

type User struct {
	gorm.Model
	ID            uint   `gorm:"primaryKey"`
	UserName      string `gorm:"unique;not null"`
	Password      string `gorm:"not null"`
	FirstName     string
	LastName      string
	Email         string
	BookedTickets uint
	Conf_ID       uint
}

func init() {
	db = db_config.Connect()
	db.AutoMigrate(&Conference{})
	db.AutoMigrate(&User{})
}

func (user *User) NewUser() {
	db.Omit("Conf_ID").Create(&user)
}

func GetUserByUserName(userName string) *User {
	var selected_user User
	db.Find(&selected_user, "User_Name = ?", userName).Limit(1)
	return &selected_user
}

func (user *User) UpdateUserBooking() {
	db.Omit("Conf_ID").Save(&user)
}

func (conf *Conference) NewConference() {
	db.Create(&conf)
}

func GetAllConferences() []Conference {
	var conf_list []Conference
	db.Preload("Users").Find(&conf_list)
	return conf_list
}

func GetConferenceByName(name string) *Conference {
	var selected_conf Conference
	db.Preload("Users").First(&selected_conf, "Name = ?", name)
	return &selected_conf
}

func GetConferenceById(Id uint) *Conference {
	var selected_conf Conference
	db.Preload("Users").First(&selected_conf, Id)
	return &selected_conf
}

func (conf *Conference) UpdateConferenceBooking() {
	db.Save(&conf)
}
