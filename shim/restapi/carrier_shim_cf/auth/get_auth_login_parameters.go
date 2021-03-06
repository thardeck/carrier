// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetAuthLoginParams creates a new GetAuthLoginParams object
// no default values defined in spec.
func NewGetAuthLoginParams() GetAuthLoginParams {

	return GetAuthLoginParams{}
}

// GetAuthLoginParams contains all the bound params for the get auth login operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetAuthLogin
type GetAuthLoginParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Use the configured prompts of the OpenID Connect Provider with the given origin key in the response. Fallback to zone values if no prompts are configured or origin is invalid.
	  In: query
	*/
	Origin *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAuthLoginParams() beforehand.
func (o *GetAuthLoginParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qOrigin, qhkOrigin, _ := qs.GetOK("origin")
	if err := o.bindOrigin(qOrigin, qhkOrigin, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindOrigin binds and validates parameter Origin from query.
func (o *GetAuthLoginParams) bindOrigin(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Origin = &raw

	return nil
}
