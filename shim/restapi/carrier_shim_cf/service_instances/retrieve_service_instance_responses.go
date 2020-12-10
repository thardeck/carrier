// Code generated by go-swagger; DO NOT EDIT.

package service_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// RetrieveServiceInstanceOKCode is the HTTP code returned for type RetrieveServiceInstanceOK
const RetrieveServiceInstanceOKCode int = 200

/*RetrieveServiceInstanceOK successful response

swagger:response retrieveServiceInstanceOK
*/
type RetrieveServiceInstanceOK struct {

	/*
	  In: Body
	*/
	Payload *models.RetrieveServiceInstanceResponseResource `json:"body,omitempty"`
}

// NewRetrieveServiceInstanceOK creates RetrieveServiceInstanceOK with default headers values
func NewRetrieveServiceInstanceOK() *RetrieveServiceInstanceOK {

	return &RetrieveServiceInstanceOK{}
}

// WithPayload adds the payload to the retrieve service instance o k response
func (o *RetrieveServiceInstanceOK) WithPayload(payload *models.RetrieveServiceInstanceResponseResource) *RetrieveServiceInstanceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve service instance o k response
func (o *RetrieveServiceInstanceOK) SetPayload(payload *models.RetrieveServiceInstanceResponseResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveServiceInstanceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
