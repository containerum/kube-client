package cmd

import (
	"net/http"

	"git.containerum.net/ch/json-types/auth"
	"git.containerum.net/ch/kube-client/pkg/model"
)

const (
	getCheckToken  = "/token/{access_token}"
	getExtendToken = "/token/{refresh_token}"
	userAgent      = "kube-client"
)

// CheckToken -- consumes JWT token, user fingerprint
// If they're correct returns user access data:
// list of namespaces and list of volumes OR uninitialized structure AND error
func (client *Client) CheckToken(token, userFingerprint string) (model.Tokens, error) {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"access_token": token,
		}).
		SetResult(auth.CheckTokenResponse{}).
		Get(client.APIurl + getCheckToken)
	if err != nil {
		return model.Tokens{}, err
	}
	return *resp.Result().(*model.Tokens), nil
}

// ExtendToken -- consumes refresh JWT token and user fingerprint
// If they're correct returns new extended access and refresh token OR void tokens AND error.
// Old access and refresh token become inactive.
func (client *Client) ExtendToken(refreshToken, userFingerprint string) (model.Tokens, error) {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"refresh_token": refreshToken,
		}).
		SetResult(model.Tokens{}).
		Put(client.APIurl + getExtendToken)
	if err := catchErr(err, resp, http.StatusOK, http.StatusAccepted); err != nil {
		return model.Tokens{}, err
	}
	return *resp.Result().(*model.Tokens), nil
}
