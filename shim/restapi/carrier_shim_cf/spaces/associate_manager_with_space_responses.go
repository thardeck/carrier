// Code generated by go-swagger; DO NOT EDIT.

package spaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// AssociateManagerWithSpaceCreatedCode is the HTTP code returned for type AssociateManagerWithSpaceCreated
const AssociateManagerWithSpaceCreatedCode int = 201

/*AssociateManagerWithSpaceCreated successful response

swagger:response associateManagerWithSpaceCreated
*/
type AssociateManagerWithSpaceCreated struct {

	/*
	  In: Body
	*/
	Payload *models.AssociateManagerWithSpaceResponseResource `json:"body,omitempty"`
}

// NewAssociateManagerWithSpaceCreated creates AssociateManagerWithSpaceCreated with default headers values
func NewAssociateManagerWithSpaceCreated() *AssociateManagerWithSpaceCreated {

	return &AssociateManagerWithSpaceCreated{}
}

// WithPayload adds the payload to the associate manager with space created response
func (o *AssociateManagerWithSpaceCreated) WithPayload(payload *models.AssociateManagerWithSpaceResponseResource) *AssociateManagerWithSpaceCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the associate manager with space created response
func (o *AssociateManagerWithSpaceCreated) SetPayload(payload *models.AssociateManagerWithSpaceResponseResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AssociateManagerWithSpaceCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
