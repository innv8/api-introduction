package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/innv8/api-introduction/entities"
	"github.com/innv8/api-introduction/middleware"
	"github.com/innv8/api-introduction/models"
)

// http.ResponseWriter is used to write the response to the client
// http.Request pointer has all the data from the client
func (b *Base) MembersController(w http.ResponseWriter, r *http.Request) {

	log.Println("the connection that reached the controller ", r.Header.Get("Connection"))

	members, _ := models.FetchMembers(b.DB)
	middleware.JSONResponse(w, 200, members)
}

func (b *Base) MembersByPostionController(w http.ResponseWriter, r *http.Request) {

	log.Println("the connection that reached the controller ", r.Header.Get("Connection"))

	members, _ := models.FetchMembersByPosition(1, b.DB)
	middleware.JSONResponse(w, http.StatusOK, members)
}

func (b *Base) CreateMemberController(w http.ResponseWriter, r *http.Request) {
	// we will read the payload sent by the client into a struct Member
	var member entities.Member
	var err error

	// 1. we read the information from the client
	err = json.NewDecoder(r.Body).Decode(&member)
	if err != nil {
		log.Println("could not read the data from client because ", err)
		middleware.JSONResponse(w, http.StatusBadRequest, "no data was sent")
		return
	}
	// 2. pass the data to the model

	err = models.CreateMember(member, b.DB)
	if err != nil {
		log.Println("could not read the data from client because ", err)
		middleware.JSONResponse(w, http.StatusBadRequest, "could not create user")
		return
	}

	// 3. respond to the client

	middleware.JSONResponse(w, http.StatusOK, "ok user created")
}
