package controllers

import (
	"encoding/json"
	"net/http"
)

// this is a sample function that is accessed through the '/' endpoint
// we can call it a handler, or in some cases a controller (we will talk about controllers in another lesson)
// w contains a resource that is used to write a response to the client
// r is a pointer to the data that has been sent to our API by the client.

// To test this, open
// localhost:5000 in a browser

// another way of testing this in a terminal
/*
curl localhost:5000
*/

// another way (I recommend this)
// You can use Postman/ Insomnia
func (b *Base) HomeController(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"hello": "world",
	}
	responseBytes, _ := json.Marshal(response)
	/*
		by default, the type of data we send to the client is plain text (text/plain)
		other common types of data are:
			- JSON (application/json)
			- XML (application/xml)
		If we want to specify the content type, we add a content-type header
		a header is just data sent along with the payload
		headers are case insensitive. So content-type and Content-Type are the same
		Also, header values can only be strings
		most common headers are
			- content-type
			- content-length : the length of data in bytes
			- Date			: the date and time that the data was sent


	*/
	w.Header().Set("content-type", "application/json")
	w.Header().Set("expires", "300")

	/*
		We can also specify the http code.
		HTTP codes are numeric codes that tell the client whether everything was ok or something happened.
		samples (commonly used)

		200 : Ok
		201 : Created (e.g if you are creating a new record, you can return this code to show the record was created
		202 : Accepted

		400 : Bad request e.g the client has sent bad data
		401	: unauthorized
		404 : not found
		405: method not allowed

		500: Server error

		For more information: https://developer.mozilla.org/en-US/docs/Web/HTTP/Status


	*/

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(responseBytes)
}
