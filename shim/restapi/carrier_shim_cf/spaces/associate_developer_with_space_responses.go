// Code generated by go-swagger; DO NOT EDIT.

package spaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// AssociateDeveloperWithSpaceCreatedCode is the HTTP code returned for type AssociateDeveloperWithSpaceCreated
const AssociateDeveloperWithSpaceCreatedCode int = 201

/*AssociateDeveloperWithSpaceCreated successful response

swagger:response associateDeveloperWithSpaceCreated
*/
type AssociateDeveloperWithSpaceCreated struct {

	/*
	  In: Body
	*/
	Payload *models.AssociateDeveloperWithSpaceResponseResource `json:"body,omitempty"`
}

// NewAssociateDeveloperWithSpaceCreated creates AssociateDeveloperWithSpaceCreated with default headers values
func NewAssociateDeveloperWithSpaceCreated() *AssociateDeveloperWithSpaceCreated {

	return &AssociateDeveloperWithSpaceCreated{}
}

// WithPayload adds the payload to the associate developer with space created response
func (o *AssociateDeveloperWithSpaceCreated) WithPayload(payload *models.AssociateDeveloperWithSpaceResponseResource) *AssociateDeveloperWithSpaceCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the associate developer with space created response
func (o *AssociateDeveloperWithSpaceCreated) SetPayload(payload *models.AssociateDeveloperWithSpaceResponseResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AssociateDeveloperWithSpaceCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
