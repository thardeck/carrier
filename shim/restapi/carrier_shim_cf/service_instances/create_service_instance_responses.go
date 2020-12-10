// Code generated by go-swagger; DO NOT EDIT.

package service_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// CreateServiceInstanceCreatedCode is the HTTP code returned for type CreateServiceInstanceCreated
const CreateServiceInstanceCreatedCode int = 201

/*CreateServiceInstanceCreated successful response

swagger:response createServiceInstanceCreated
*/
type CreateServiceInstanceCreated struct {

	/*
	  In: Body
	*/
	Payload *models.CreateServiceInstanceResponseResource `json:"body,omitempty"`
}

// NewCreateServiceInstanceCreated creates CreateServiceInstanceCreated with default headers values
func NewCreateServiceInstanceCreated() *CreateServiceInstanceCreated {

	return &CreateServiceInstanceCreated{}
}

// WithPayload adds the payload to the create service instance created response
func (o *CreateServiceInstanceCreated) WithPayload(payload *models.CreateServiceInstanceResponseResource) *CreateServiceInstanceCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create service instance created response
func (o *CreateServiceInstanceCreated) SetPayload(payload *models.CreateServiceInstanceResponseResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateServiceInstanceCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
