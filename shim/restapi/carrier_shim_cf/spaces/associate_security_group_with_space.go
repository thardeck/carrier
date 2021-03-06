// Code generated by go-swagger; DO NOT EDIT.

package spaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AssociateSecurityGroupWithSpaceHandlerFunc turns a function with the right signature into a associate security group with space handler
type AssociateSecurityGroupWithSpaceHandlerFunc func(AssociateSecurityGroupWithSpaceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AssociateSecurityGroupWithSpaceHandlerFunc) Handle(params AssociateSecurityGroupWithSpaceParams) middleware.Responder {
	return fn(params)
}

// AssociateSecurityGroupWithSpaceHandler interface for that can handle valid associate security group with space params
type AssociateSecurityGroupWithSpaceHandler interface {
	Handle(AssociateSecurityGroupWithSpaceParams) middleware.Responder
}

// NewAssociateSecurityGroupWithSpace creates a new http.Handler for the associate security group with space operation
func NewAssociateSecurityGroupWithSpace(ctx *middleware.Context, handler AssociateSecurityGroupWithSpaceHandler) *AssociateSecurityGroupWithSpace {
	return &AssociateSecurityGroupWithSpace{Context: ctx, Handler: handler}
}

/*AssociateSecurityGroupWithSpace swagger:route PUT /spaces/{guid}/security_groups/{security_group_guid} spaces associateSecurityGroupWithSpace

Associate Security Group with the Space

curl --insecure -i %s/v2/spaces/{guid}/security_groups/{security_group_guid} -X PUT -H 'Authorization: %s'

*/
type AssociateSecurityGroupWithSpace struct {
	Context *middleware.Context
	Handler AssociateSecurityGroupWithSpaceHandler
}

func (o *AssociateSecurityGroupWithSpace) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAssociateSecurityGroupWithSpaceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
