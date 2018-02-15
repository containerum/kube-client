package cmd

import (
	"net/http"

	"git.containerum.net/ch/kube-client/pkg/model"
)

const (
	userInfoPath = "/user/info"
)

// GetProfileInfo -- returns user info
func (client *Client) GetProfileInfo() (model.User, error) {
	resp, err := client.Request.
		SetResult(model.User{}).
		SetError(model.ResourceError{}).
		Get(client.userManagerURL + userInfoPath)
	if err := catchErr(err, resp, http.StatusOK); err != nil {
		return model.User{}, err
	}
	return *resp.Result().(*model.User), nil
}
