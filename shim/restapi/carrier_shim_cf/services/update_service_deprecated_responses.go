// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// UpdateServiceDeprecatedCreatedCode is the HTTP code returned for type UpdateServiceDeprecatedCreated
const UpdateServiceDeprecatedCreatedCode int = 201

/*UpdateServiceDeprecatedCreated successful response

swagger:response updateServiceDeprecatedCreated
*/
type UpdateServiceDeprecatedCreated struct {

	/*
	  In: Body
	*/
	Payload *models.UpdateServiceDeprecatedResponseResource `json:"body,omitempty"`
}

// NewUpdateServiceDeprecatedCreated creates UpdateServiceDeprecatedCreated with default headers values
func NewUpdateServiceDeprecatedCreated() *UpdateServiceDeprecatedCreated {

	return &UpdateServiceDeprecatedCreated{}
}

// WithPayload adds the payload to the update service deprecated created response
func (o *UpdateServiceDeprecatedCreated) WithPayload(payload *models.UpdateServiceDeprecatedResponseResource) *UpdateServiceDeprecatedCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update service deprecated created response
func (o *UpdateServiceDeprecatedCreated) SetPayload(payload *models.UpdateServiceDeprecatedResponseResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateServiceDeprecatedCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
