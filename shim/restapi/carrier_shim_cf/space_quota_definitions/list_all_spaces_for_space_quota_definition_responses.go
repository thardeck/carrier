// Code generated by go-swagger; DO NOT EDIT.

package space_quota_definitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/suse/carrier/shim/models"
)

// ListAllSpacesForSpaceQuotaDefinitionOKCode is the HTTP code returned for type ListAllSpacesForSpaceQuotaDefinitionOK
const ListAllSpacesForSpaceQuotaDefinitionOKCode int = 200

/*ListAllSpacesForSpaceQuotaDefinitionOK successful response

swagger:response listAllSpacesForSpaceQuotaDefinitionOK
*/
type ListAllSpacesForSpaceQuotaDefinitionOK struct {

	/*
	  In: Body
	*/
	Payload *models.ListAllSpacesForSpaceQuotaDefinitionResponsePaged `json:"body,omitempty"`
}

// NewListAllSpacesForSpaceQuotaDefinitionOK creates ListAllSpacesForSpaceQuotaDefinitionOK with default headers values
func NewListAllSpacesForSpaceQuotaDefinitionOK() *ListAllSpacesForSpaceQuotaDefinitionOK {

	return &ListAllSpacesForSpaceQuotaDefinitionOK{}
}

// WithPayload adds the payload to the list all spaces for space quota definition o k response
func (o *ListAllSpacesForSpaceQuotaDefinitionOK) WithPayload(payload *models.ListAllSpacesForSpaceQuotaDefinitionResponsePaged) *ListAllSpacesForSpaceQuotaDefinitionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list all spaces for space quota definition o k response
func (o *ListAllSpacesForSpaceQuotaDefinitionOK) SetPayload(payload *models.ListAllSpacesForSpaceQuotaDefinitionResponsePaged) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListAllSpacesForSpaceQuotaDefinitionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
