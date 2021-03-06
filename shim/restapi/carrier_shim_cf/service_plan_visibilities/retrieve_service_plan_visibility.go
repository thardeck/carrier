// Code generated by go-swagger; DO NOT EDIT.

package service_plan_visibilities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RetrieveServicePlanVisibilityHandlerFunc turns a function with the right signature into a retrieve service plan visibility handler
type RetrieveServicePlanVisibilityHandlerFunc func(RetrieveServicePlanVisibilityParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RetrieveServicePlanVisibilityHandlerFunc) Handle(params RetrieveServicePlanVisibilityParams) middleware.Responder {
	return fn(params)
}

// RetrieveServicePlanVisibilityHandler interface for that can handle valid retrieve service plan visibility params
type RetrieveServicePlanVisibilityHandler interface {
	Handle(RetrieveServicePlanVisibilityParams) middleware.Responder
}

// NewRetrieveServicePlanVisibility creates a new http.Handler for the retrieve service plan visibility operation
func NewRetrieveServicePlanVisibility(ctx *middleware.Context, handler RetrieveServicePlanVisibilityHandler) *RetrieveServicePlanVisibility {
	return &RetrieveServicePlanVisibility{Context: ctx, Handler: handler}
}

/*RetrieveServicePlanVisibility swagger:route GET /service_plan_visibilities/{guid} servicePlanVisibilities retrieveServicePlanVisibility

Retrieve a Particular Service Plan Visibility

curl --insecure -i %s/v2/service_plan_visibilities/{guid} -X GET -H 'Authorization: %s'

*/
type RetrieveServicePlanVisibility struct {
	Context *middleware.Context
	Handler RetrieveServicePlanVisibilityHandler
}

func (o *RetrieveServicePlanVisibility) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRetrieveServicePlanVisibilityParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
