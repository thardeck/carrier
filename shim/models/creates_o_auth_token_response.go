// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreatesOAuthTokenResponse creates o auth token response
//
// swagger:model createsOAuthTokenResponse
type CreatesOAuthTokenResponse struct {

	// An OAuth2 access token. When token_format=opaque is requested this value will be a random string that can only be validated using the UAA's /check_token or /introspect endpoints. When token_format=jwt is requested, this token will be a JSON Web Token suitable for offline validation by OAuth2 Resource Servers.
	AccessToken string `json:"access_token,omitempty"`

	// The number of seconds until the access token expires.
	ExpiresIn int64 `json:"expires_in,omitempty"`

	// An OpenID Connect ID token. This portion of the token response is only returned when clients are configured with the scope openid, the response_type includes id_token, and the user has granted approval to the client for the openid scope.
	IDToken string `json:"id_token,omitempty"`

	// A globally unique identifier for this access token. This identifier is used when revoking tokens.
	Jti string `json:"jti,omitempty"`

	// An OAuth2 refresh token. Clients typically use the refresh token to obtain a new access token without the need for the user to authenticate again. They do this by calling /oauth/token with grant_type=refresh_token. See here for more information. A refresh token will only be issued to clients that have refresh_token in their list of authorized_grant_types.
	RefreshToken string `json:"refresh_token,omitempty"`

	// A space-delimited list of scopes authorized by the user for this client. This list is the intersection of the scopes configured on the client, the group memberships of the user, and the user's approvals (when autoapprove: true is not configured on the client).
	Scope string `json:"scope,omitempty"`

	// The type of the access token issued. This field is mandated in RFC 6749. In the UAA, the only supported token_type is bearer.
	TokenType string `json:"token_type,omitempty"`
}

// Validate validates this creates o auth token response
func (m *CreatesOAuthTokenResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreatesOAuthTokenResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreatesOAuthTokenResponse) UnmarshalBinary(b []byte) error {
	var res CreatesOAuthTokenResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
