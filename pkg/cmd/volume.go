package cmd

import (
	"net/http"

	"git.containerum.net/ch/kube-client/pkg/model"
)

const (
	resourceVolumeRootPath   = "/volume"
	resourceVolumePath       = "/volume/{volume}"
	resourceVolumeNamePath   = "/volume/{volume}/name"
	resourceVolumeAccessPath = "/volume/{volume}/access"
)

// DeleteVolume -- deletes Volume with provided volume name
func (client *Client) DeleteVolume(volumeName string) error {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"volume": volumeName,
		}).SetError(model.ResourceError{}).
		Delete(client.ResourceAddr + resourceVolumePath)
	return catchErr(err, resp, http.StatusAccepted, http.StatusOK)
}

// GetVolume -- return User Volume by name,
// consumes optional userID param
func (client *Client) GetVolume(volumeName string, userID *string) (model.ResourceVolume, error) {
	req := client.Request.
		SetPathParams(map[string]string{
			"volume": volumeName,
		}).
		SetResult(model.ResourceVolume{}).
		SetError(model.ResourceError{})
	if userID != nil {
		req.SetQueryParam("user-id", *userID)
	}
	resp, err := req.Get(client.ResourceAddr + resourceVolumePath)
	if err := catchErr(err, resp, http.StatusOK); err != nil {
		return model.ResourceVolume{}, err
	}
	return *resp.Result().(*model.ResourceVolume), nil
}

// GetVolumeList -- get list of volumes,
// consumes optional user ID and filter parameters.
// Returns new_access_level as access if user role = user.
// Should have filters: not deleted, limited, not limited, owner, not owner.
func (client *Client) GetVolumeList(userID, filter *string) ([]model.ResourceVolume, error) {
	req := client.Request.
		SetResult([]model.ResourceVolume{}).
		SetError(model.ResourceError{})
	if userID != nil {
		req.SetQueryParam("user-id", *userID)
	}
	if filter != nil {
		req.SetQueryParam("user-id", *filter)
	}
	resp, err := req.Get(client.ResourceAddr + resourceVolumeRootPath)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, resp.Error().(*model.ResourceError)
	}
	return *resp.Result().(*[]model.ResourceVolume), nil
}

//RenameVolume -- change volume name
func (client *Client) RenameVolume(volumeName, newName string) error {
	_, err := client.Request.
		SetPathParams(map[string]string{
			"volume": volumeName,
		}).
		SetBody(model.ResourceUpdateName{Label: newName}).
		Put(client.ResourceAddr + resourceVolumeNamePath)
	return err
}

// SetVolumeAccess -- sets User Volume access
func (client *Client) SetVolumeAccess(volumeName, username, access string) error {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"volume": volumeName,
		}).
		SetBody(model.ResourceUpdateUserAccess{
			Username: username,
			Access:   access,
		}).
		SetError(model.ResourceError{}).
		Put(client.ResourceAddr + resourceVolumeAccessPath)
	return catchErr(err, resp, http.StatusOK, http.StatusAccepted)
}

// DeleteVolumeAccess -- deletes user Volume access
func (client *Client) DeleteVolumeAccess(volumeName, username string) error {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"volume": volumeName,
		}).
		SetBody(model.ResourceUpdateUserAccess{
			Username: username,
		}).
		SetError(model.ResourceError{}).
		Delete(client.ResourceAddr + resourceVolumeAccessPath)
	return catchErr(err, resp, http.StatusOK, http.StatusAccepted)
}
