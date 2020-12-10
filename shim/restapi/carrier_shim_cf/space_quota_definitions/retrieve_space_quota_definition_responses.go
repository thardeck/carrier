// Code generated by go-swagger; DO NOT EDIT.

package space_quota_definitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// RetrieveSpaceQuotaDefinitionOKCode is the HTTP code returned for type RetrieveSpaceQuotaDefinitionOK
const RetrieveSpaceQuotaDefinitionOKCode int = 200

/*RetrieveSpaceQuotaDefinitionOK successful response

swagger:response retrieveSpaceQuotaDefinitionOK
*/
type RetrieveSpaceQuotaDefinitionOK struct {

	/*
	  In: Body
	*/
	Payload *models.RetrieveSpaceQuotaDefinitionResponseResource `json:"body,omitempty"`
}

// NewRetrieveSpaceQuotaDefinitionOK creates RetrieveSpaceQuotaDefinitionOK with default headers values
func NewRetrieveSpaceQuotaDefinitionOK() *RetrieveSpaceQuotaDefinitionOK {

	return &RetrieveSpaceQuotaDefinitionOK{}
}

// WithPayload adds the payload to the retrieve space quota definition o k response
func (o *RetrieveSpaceQuotaDefinitionOK) WithPayload(payload *models.RetrieveSpaceQuotaDefinitionResponseResource) *RetrieveSpaceQuotaDefinitionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve space quota definition o k response
func (o *RetrieveSpaceQuotaDefinitionOK) SetPayload(payload *models.RetrieveSpaceQuotaDefinitionResponseResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveSpaceQuotaDefinitionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
