package cmd

import "git.containerum.net/ch/kube-client/pkg/model"

const (
	resourceVolumeRootPath = "/volume"
	resourceVolumePath     = resourceVolumeRootPath + "/{volume}"
)

// CreateVolume -- creates Volume for User by Tariff
func (client *Client) CreateVolume(volume model.ResourceCreateVolume) error {
	_, err := client.Request.
		SetBody(volume).
		Post(client.resourceServiceAddr + resourceVolumeRootPath)
	return err
}

// DeleteVolume -- deletes Volume with provided volume name
func (client *Client) DeleteVolume(volumeName string) error {
	_, err := client.Request.
		SetPathParams(map[string]string{
			"volume": volumeName,
		}).
		Delete(client.resourceServiceAddr + resourceVolumePath)
	return err
}

// GetVolume -- return User Volume by name,
// consumes optional userID param
func (client *Client) GetVolume(volumeName string, userID *string) (model.ResourceVolume, error) {
	req := client.Request.
		SetPathParams(map[string]string{
			"volume": volumeName,
		}).
		SetResult(&model.ResourceVolume{})
	if userID != nil {
		req.SetQueryParam("user-id", *userID)
	}
	resp, err := req.Get(client.resourceServiceAddr + resourceVolumePath)
	if err != nil {
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
		SetResult([]model.ResourceVolume{})
	if userID != nil {
		req.SetQueryParam("user-id", *userID)
	}
	if filter != nil {
		req.SetQueryParam("user-id", *filter)
	}
	resp, err := req.Get(client.resourceServiceAddr + resourceVolumeRootPath)
	if err != nil {
		return nil, err
	}
	return *resp.Result().(*[]model.ResourceVolume), nil
}
