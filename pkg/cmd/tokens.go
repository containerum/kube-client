package cmd

import (
	"git.containerum.net/ch/json-types/auth"
	user "git.containerum.net/ch/json-types/user-manager"
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
func (client *Client) CheckToken(token, userFingerprint string) (auth.CheckTokenResponse, error) {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"access_token": token,
		}).
		SetResult(auth.CheckTokenResponse{}).
		SetHeaders(map[string]string{
			user.FingerprintHeader: userFingerprint,
			user.UserAgentHeader:   userAgent,
		}).Get(client.APIurl + getCheckToken)
	if err != nil {
		return auth.CheckTokenResponse{}, err
	}
	return *resp.Result().(*auth.CheckTokenResponse), nil
}

// ExtendToken -- consumes refresh JWT token and user fingerprint
// If they're correct returns new extended access and refresh token OR void tokens AND error.
// Old access and refresh token become inactive.
func (client *Client) ExtendToken(refreshToken, userFingerprint string) (model.Tokens, error) {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"refresh_token": refreshToken,
		}).
		SetResult(auth.Tokens{}).
		SetHeaders(map[string]string{
			user.FingerprintHeader: userFingerprint,
		}).Put(client.APIurl + getExtendToken)
	if err != nil {
		return model.Tokens{}, err
	}
	return *resp.Result().(*model.Tokens), nil
}
