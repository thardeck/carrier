// Code generated by go-swagger; DO NOT EDIT.

package service_brokers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteServiceBrokerNoContentCode is the HTTP code returned for type DeleteServiceBrokerNoContent
const DeleteServiceBrokerNoContentCode int = 204

/*DeleteServiceBrokerNoContent successful response

swagger:response deleteServiceBrokerNoContent
*/
type DeleteServiceBrokerNoContent struct {
}

// NewDeleteServiceBrokerNoContent creates DeleteServiceBrokerNoContent with default headers values
func NewDeleteServiceBrokerNoContent() *DeleteServiceBrokerNoContent {

	return &DeleteServiceBrokerNoContent{}
}

// WriteResponse to the client
func (o *DeleteServiceBrokerNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}
