package cmd

import (
	"net"

	"git.containerum.net/ch/json-types/user-manager"
)

const (
	getCheckToken = "/token/{access_token}"
	userAgent     = "kube-client"
)

func (client *Client) CheckToken(token, userFingerprint string, clientIP net.IP) (interface{}, error) {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"access_token": token,
		}).SetHeaders(map[string]string{
		"X-User-Fingerprint": userFingerprint,
		user.ClientIPHeader:  clientIP.String(),
		user.UserAgentHeader: userAgent,
	}).Get(client.serverURL + getCheckToken)
	resp.Result()
	if err != nil {
		return nil, err
	}
	return nil, nil
}
