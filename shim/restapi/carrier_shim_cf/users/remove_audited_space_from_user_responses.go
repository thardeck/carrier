// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// RemoveAuditedSpaceFromUserCreatedCode is the HTTP code returned for type RemoveAuditedSpaceFromUserCreated
const RemoveAuditedSpaceFromUserCreatedCode int = 201

/*RemoveAuditedSpaceFromUserCreated successful response

swagger:response removeAuditedSpaceFromUserCreated
*/
type RemoveAuditedSpaceFromUserCreated struct {

	/*
	  In: Body
	*/
	Payload *models.RemoveAuditedSpaceFromUserResponseResource `json:"body,omitempty"`
}

// NewRemoveAuditedSpaceFromUserCreated creates RemoveAuditedSpaceFromUserCreated with default headers values
func NewRemoveAuditedSpaceFromUserCreated() *RemoveAuditedSpaceFromUserCreated {

	return &RemoveAuditedSpaceFromUserCreated{}
}

// WithPayload adds the payload to the remove audited space from user created response
func (o *RemoveAuditedSpaceFromUserCreated) WithPayload(payload *models.RemoveAuditedSpaceFromUserResponseResource) *RemoveAuditedSpaceFromUserCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove audited space from user created response
func (o *RemoveAuditedSpaceFromUserCreated) SetPayload(payload *models.RemoveAuditedSpaceFromUserResponseResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveAuditedSpaceFromUserCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
