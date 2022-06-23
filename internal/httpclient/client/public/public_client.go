// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new public API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for public API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DisconnectUser(params *DisconnectUserParams, opts ...ClientOption) error

	DiscoverOpenIDConfiguration(params *DiscoverOpenIDConfigurationParams, opts ...ClientOption) (*DiscoverOpenIDConfigurationOK, error)

	DynamicClientRegistrationCreateOAuth2Client(params *DynamicClientRegistrationCreateOAuth2ClientParams, opts ...ClientOption) (*DynamicClientRegistrationCreateOAuth2ClientCreated, error)

	DynamicClientRegistrationDeleteOAuth2Client(params *DynamicClientRegistrationDeleteOAuth2ClientParams, opts ...ClientOption) (*DynamicClientRegistrationDeleteOAuth2ClientNoContent, error)

	DynamicClientRegistrationGetOAuth2Client(params *DynamicClientRegistrationGetOAuth2ClientParams, opts ...ClientOption) (*DynamicClientRegistrationGetOAuth2ClientOK, error)

	DynamicClientRegistrationUpdateOAuth2Client(params *DynamicClientRegistrationUpdateOAuth2ClientParams, opts ...ClientOption) (*DynamicClientRegistrationUpdateOAuth2ClientOK, error)

	IsInstanceReady(params *IsInstanceReadyParams, opts ...ClientOption) (*IsInstanceReadyOK, error)

	Oauth2Token(params *Oauth2TokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Oauth2TokenOK, error)

	OauthAuth(params *OauthAuthParams, opts ...ClientOption) error

	RevokeOAuth2Token(params *RevokeOAuth2TokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RevokeOAuth2TokenOK, error)

	Userinfo(params *UserinfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserinfoOK, error)

	WellKnown(params *WellKnownParams, opts ...ClientOption) (*WellKnownOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DisconnectUser opens ID connect front backchannel enabled logout

  This endpoint initiates and completes user logout at Ory Hydra and initiates OpenID Connect Front-/Back-channel logout:

https://openid.net/specs/openid-connect-frontchannel-1_0.html
https://openid.net/specs/openid-connect-backchannel-1_0.html

Back-channel logout is performed asynchronously and does not affect logout flow.
*/
func (a *Client) DisconnectUser(params *DisconnectUserParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisconnectUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "disconnectUser",
		Method:             "GET",
		PathPattern:        "/oauth2/sessions/logout",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DisconnectUserReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
  DiscoverOpenIDConfiguration opens ID connect discovery

  The well known endpoint an be used to retrieve information for OpenID Connect clients. We encourage you to not roll
your own OpenID Connect client but to use an OpenID Connect client library instead. You can learn more on this
flow at https://openid.net/specs/openid-connect-discovery-1_0.html .

Popular libraries for OpenID Connect clients include oidc-client-js (JavaScript), go-oidc (Golang), and others.
For a full list of clients go here: https://openid.net/developers/certified/
*/
func (a *Client) DiscoverOpenIDConfiguration(params *DiscoverOpenIDConfigurationParams, opts ...ClientOption) (*DiscoverOpenIDConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDiscoverOpenIDConfigurationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "discoverOpenIDConfiguration",
		Method:             "GET",
		PathPattern:        "/.well-known/openid-configuration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DiscoverOpenIDConfigurationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DiscoverOpenIDConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for discoverOpenIDConfiguration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DynamicClientRegistrationCreateOAuth2Client registers an o auth 2 0 client using the open ID o auth2 dynamic client registration management protocol

  This endpoint behaves like the administrative counterpart (`createOAuth2Client`) but is capable of facing the
public internet directly and can be used in self-service. It implements the OpenID Connect
Dynamic Client Registration Protocol. This feature needs to be enabled in the configuration. This endpoint
is disabled by default. It can be enabled by an administrator.

Please note that using this endpoint you are not able to choose the `client_secret` nor the `client_id` as those
values will be server generated when specifying `token_endpoint_auth_method` as `client_secret_basic` or
`client_secret_post`.

The `client_secret` will be returned in the response and you will not be able to retrieve it later on.
Write the secret down and keep it somewhere safe.
*/
func (a *Client) DynamicClientRegistrationCreateOAuth2Client(params *DynamicClientRegistrationCreateOAuth2ClientParams, opts ...ClientOption) (*DynamicClientRegistrationCreateOAuth2ClientCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDynamicClientRegistrationCreateOAuth2ClientParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dynamicClientRegistrationCreateOAuth2Client",
		Method:             "POST",
		PathPattern:        "/oauth2/register",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DynamicClientRegistrationCreateOAuth2ClientReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DynamicClientRegistrationCreateOAuth2ClientCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DynamicClientRegistrationCreateOAuth2ClientDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DynamicClientRegistrationDeleteOAuth2Client deletes an o auth 2 0 client using the open ID o auth2 dynamic client registration management protocol

  This endpoint behaves like the administrative counterpart (`deleteOAuth2Client`) but is capable of facing the
public internet directly and can be used in self-service. It implements the OpenID Connect
Dynamic Client Registration Protocol. This feature needs to be enabled in the configuration. This endpoint
is disabled by default. It can be enabled by an administrator.

To use this endpoint, you will need to present the client's authentication credentials. If the OAuth2 Client
uses the Token Endpoint Authentication Method `client_secret_post`, you need to present the client secret in the URL query.
If it uses `client_secret_basic`, present the Client ID and the Client Secret in the Authorization header.

OAuth 2.0 clients are used to perform OAuth 2.0 and OpenID Connect flows. Usually, OAuth 2.0 clients are
generated for applications which want to consume your OAuth 2.0 or OpenID Connect capabilities.
*/
func (a *Client) DynamicClientRegistrationDeleteOAuth2Client(params *DynamicClientRegistrationDeleteOAuth2ClientParams, opts ...ClientOption) (*DynamicClientRegistrationDeleteOAuth2ClientNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDynamicClientRegistrationDeleteOAuth2ClientParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dynamicClientRegistrationDeleteOAuth2Client",
		Method:             "DELETE",
		PathPattern:        "/oauth2/register/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DynamicClientRegistrationDeleteOAuth2ClientReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DynamicClientRegistrationDeleteOAuth2ClientNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DynamicClientRegistrationDeleteOAuth2ClientDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DynamicClientRegistrationGetOAuth2Client gets an o auth 2 0 client using the open ID o auth2 dynamic client registration management protocol

  This endpoint behaves like the administrative counterpart (`getOAuth2Client`) but is capable of facing the
public internet directly and can be used in self-service. It implements the OpenID Connect
Dynamic Client Registration Protocol. This feature needs to be enabled in the configuration. This endpoint
is disabled by default. It can be enabled by an administrator.

To use this endpoint, you will need to present the client's authentication credentials. If the OAuth2 Client
uses the Token Endpoint Authentication Method `client_secret_post`, you need to present the client secret in the URL query.
If it uses `client_secret_basic`, present the Client ID and the Client Secret in the Authorization header.

OAuth 2.0 clients are used to perform OAuth 2.0 and OpenID Connect flows. Usually, OAuth 2.0 clients are
generated for applications which want to consume your OAuth 2.0 or OpenID Connect capabilities.
*/
func (a *Client) DynamicClientRegistrationGetOAuth2Client(params *DynamicClientRegistrationGetOAuth2ClientParams, opts ...ClientOption) (*DynamicClientRegistrationGetOAuth2ClientOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDynamicClientRegistrationGetOAuth2ClientParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dynamicClientRegistrationGetOAuth2Client",
		Method:             "GET",
		PathPattern:        "/oauth2/register/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DynamicClientRegistrationGetOAuth2ClientReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DynamicClientRegistrationGetOAuth2ClientOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DynamicClientRegistrationGetOAuth2ClientDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DynamicClientRegistrationUpdateOAuth2Client updates an o auth 2 0 client using the open ID o auth2 dynamic client registration management protocol

  This endpoint behaves like the administrative counterpart (`updateOAuth2Client`) but is capable of facing the
public internet directly and can be used in self-service. It implements the OpenID Connect
Dynamic Client Registration Protocol. This feature needs to be enabled in the configuration. This endpoint
is disabled by default. It can be enabled by an administrator.

If you pass `client_secret` the secret will be updated and returned via the API.
This is the only time you will be able to retrieve the client secret, so write it down and keep it safe.

To use this endpoint, you will need to present the client's authentication credentials. If the OAuth2 Client
uses the Token Endpoint Authentication Method `client_secret_post`, you need to present the client secret in the URL query.
If it uses `client_secret_basic`, present the Client ID and the Client Secret in the Authorization header.

OAuth 2.0 clients are used to perform OAuth 2.0 and OpenID Connect flows. Usually, OAuth 2.0 clients are
generated for applications which want to consume your OAuth 2.0 or OpenID Connect capabilities.
*/
func (a *Client) DynamicClientRegistrationUpdateOAuth2Client(params *DynamicClientRegistrationUpdateOAuth2ClientParams, opts ...ClientOption) (*DynamicClientRegistrationUpdateOAuth2ClientOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDynamicClientRegistrationUpdateOAuth2ClientParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "dynamicClientRegistrationUpdateOAuth2Client",
		Method:             "PUT",
		PathPattern:        "/oauth2/register/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DynamicClientRegistrationUpdateOAuth2ClientReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DynamicClientRegistrationUpdateOAuth2ClientOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DynamicClientRegistrationUpdateOAuth2ClientDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  IsInstanceReady checks readiness status

  This endpoint returns a 200 status code when the HTTP server is up running and the environment dependencies (e.g.
the database) are responsive as well.

If the service supports TLS Edge Termination, this endpoint does not require the
`X-Forwarded-Proto` header to be set.

Be aware that if you are running multiple nodes of this service, the health status will never
refer to the cluster state, only to a single instance.
*/
func (a *Client) IsInstanceReady(params *IsInstanceReadyParams, opts ...ClientOption) (*IsInstanceReadyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIsInstanceReadyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "isInstanceReady",
		Method:             "GET",
		PathPattern:        "/health/ready",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &IsInstanceReadyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IsInstanceReadyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for isInstanceReady: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Oauth2Token thes o auth 2 0 token endpoint

  The client makes a request to the token endpoint by sending the
following parameters using the "application/x-www-form-urlencoded" HTTP
request entity-body.

> Do not implement a client for this endpoint yourself. Use a library. There are many libraries
> available for any programming language. You can find a list of libraries here: https://oauth.net/code/
>
> Do note that Hydra SDK does not implement this endpoint properly. Use one of the libraries listed above!
*/
func (a *Client) Oauth2Token(params *Oauth2TokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*Oauth2TokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOauth2TokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "oauth2Token",
		Method:             "POST",
		PathPattern:        "/oauth2/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &Oauth2TokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Oauth2TokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for oauth2Token: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  OauthAuth thes o auth 2 0 authorize endpoint

  This endpoint is not documented here because you should never use your own implementation to perform OAuth2 flows.
OAuth2 is a very popular protocol and a library for your programming language will exists.

To learn more about this flow please refer to the specification: https://tools.ietf.org/html/rfc6749
*/
func (a *Client) OauthAuth(params *OauthAuthParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOauthAuthParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "oauthAuth",
		Method:             "GET",
		PathPattern:        "/oauth2/auth",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &OauthAuthReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
  RevokeOAuth2Token revokes o auth2 tokens

  Revoking a token (both access and refresh) means that the tokens will be invalid. A revoked access token can no
longer be used to make access requests, and a revoked refresh token can no longer be used to refresh an access token.
Revoking a refresh token also invalidates the access token that was created with it. A token may only be revoked by
the client the token was generated for.
*/
func (a *Client) RevokeOAuth2Token(params *RevokeOAuth2TokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RevokeOAuth2TokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRevokeOAuth2TokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "revokeOAuth2Token",
		Method:             "POST",
		PathPattern:        "/oauth2/revoke",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RevokeOAuth2TokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RevokeOAuth2TokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for revokeOAuth2Token: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Userinfo opens ID connect userinfo

  This endpoint returns the payload of the ID Token, including the idTokenExtra values, of
the provided OAuth 2.0 Access Token.

For more information please [refer to the spec](http://openid.net/specs/openid-connect-core-1_0.html#UserInfo).

In the case of authentication error, a WWW-Authenticate header might be set in the response
with more information about the error. See [the spec](https://datatracker.ietf.org/doc/html/rfc6750#section-3)
for more details about header format.
*/
func (a *Client) Userinfo(params *UserinfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UserinfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserinfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "userinfo",
		Method:             "GET",
		PathPattern:        "/userinfo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserinfoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UserinfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for userinfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  WellKnown JSONs web keys discovery

  This endpoint returns JSON Web Keys to be used as public keys for verifying OpenID Connect ID Tokens and,
if enabled, OAuth 2.0 JWT Access Tokens. This endpoint can be used with client libraries like
[node-jwks-rsa](https://github.com/auth0/node-jwks-rsa) among others.
*/
func (a *Client) WellKnown(params *WellKnownParams, opts ...ClientOption) (*WellKnownOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWellKnownParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "wellKnown",
		Method:             "GET",
		PathPattern:        "/.well-known/jwks.json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &WellKnownReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WellKnownOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for wellKnown: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
