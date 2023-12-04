package models

import (
	u "contactsBook/utils"
	"fmt"
	"github.com/dongri/phonenumber"
	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserId uint   `json:"user_id"`
}

func (contact *Contact) ValidateContact() (map[string]interface{}, bool) {

	if contact.Name == "" {
		return u.Message(false, "Name cannot be empty!"), false
	}

	if !validatePhoneNumber(contact.Phone) {
		return u.Message(false, "Invalid phone number!"), false
	}

	if contact.UserId <= 0 {
		return u.Message(false, "User not found!"), false
	}

	return u.Message(true, "success"), true
}

func validatePhoneNumber(phone string) bool {
	phone = phonenumber.Parse(phone, "RU")
	return phone != ""
}

func (contact *Contact) CreateContact() map[string]interface{} {

	if response, ok := contact.ValidateContact(); !ok {
		return response
	}

	GetDB().Create(contact)

	resp := u.Message(true, "success")
	resp["contact"] = contact
	return resp
}

func GetContact(id uint) *Contact {

	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

func GetContacts(user uint) []*Contact {

	contactsSlice := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&contactsSlice).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contactsSlice
}

func DeleteContact(contact *Contact) *Contact {
	dbResp := GetDB().Unscoped().Delete(contact)
	if dbResp.Error != nil {
		return nil
	} else if dbResp.RowsAffected < 1 {
		fmt.Errorf("row with id=%d cannot be deleted because it doesn't exist", contact.ID)
		return nil
	}

	return contact
}

func UpdateContact(contact *Contact) *Contact {
	db := GetDB()
	existedContact := &Contact{}
	db.Where("id = ?", contact.ID).First(existedContact)
	if existedContact.ID == 0 {
		return nil
	}
	db.Save(&contact)
	return contact
}
