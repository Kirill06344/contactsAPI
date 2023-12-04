package controllers

import (
	"contactsBook/models"
	u "contactsBook/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error!"))
		return
	}

	contact.UserId = user
	resp := contact.CreateContact()
	u.Respond(w, resp)
}

var GetContacts = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetContacts(id)
	resp := u.Message(true, "All available contacts in your contact book.")
	resp["data"] = data
	u.Respond(w, resp)
}

var DeleteContact = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	contact := &models.Contact{}
	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error!"))
		return
	}
	contact.UserId = id

	data := models.DeleteContact(contact)
	if data == nil {
		u.Respond(w, u.Message(false, "Contact doesn't exist!"))
		return
	}
	resp := u.Message(true, "Contact successfully deleted!")
	u.Respond(w, resp)
}

var UpdateContact = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	contact := &models.Contact{}
	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error!"))
		return
	}
	if contact.ID == 0 {
		u.Respond(w, u.Message(false, "Write ID of record to change.You can get it with GET /me/contacts"))
		return
	}
	contact.UserId = id

	data := models.UpdateContact(contact)
	if data == nil {
		u.Respond(w, u.Message(false, fmt.Sprintf("You need contact with ID %d", contact.ID)))
		return
	}
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
