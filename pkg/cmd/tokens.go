package cmd

import (
	"net"

	"git.containerum.net/ch/json-types/auth"
	user "git.containerum.net/ch/json-types/user-manager"
)

const (
	getCheckToken = "/token/{access_token}"
	userAgent     = "kube-client"
)

func (client *Client) CheckToken(token, userFingerprint string, clientIP net.IP) (auth.CheckTokenResponse, error) {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"access_token": token,
		}).
		SetResult(auth.CheckTokenResponse{}).
		SetHeaders(map[string]string{
			"X-User-Fingerprint": userFingerprint,
			user.ClientIPHeader:  clientIP.String(),
			user.UserAgentHeader: userAgent,
		}).Get(client.serverURL + getCheckToken)
	if err != nil {
		return auth.CheckTokenResponse{}, err
	}
	return *resp.Result().(*auth.CheckTokenResponse), nil
}
