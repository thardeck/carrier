// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// RetrieveUserOKCode is the HTTP code returned for type RetrieveUserOK
const RetrieveUserOKCode int = 200

/*RetrieveUserOK successful response

swagger:response retrieveUserOK
*/
type RetrieveUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.RetrieveUserResponseResource `json:"body,omitempty"`
}

// NewRetrieveUserOK creates RetrieveUserOK with default headers values
func NewRetrieveUserOK() *RetrieveUserOK {

	return &RetrieveUserOK{}
}

// WithPayload adds the payload to the retrieve user o k response
func (o *RetrieveUserOK) WithPayload(payload *models.RetrieveUserResponseResource) *RetrieveUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve user o k response
func (o *RetrieveUserOK) SetPayload(payload *models.RetrieveUserResponseResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
