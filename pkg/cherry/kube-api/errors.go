package kubeErrors

import (
	bytes "bytes"
	template "text/template"

	cherry "git.containerum.net/ch/kube-client/pkg/cherry"
)

const ()

// ErrAdminRequired error
// User is not admin and has no permissions
func ErrAdminRequired(params ...func(*cherry.Err)) *cherry.Err {
	err := &cherry.Err{Message: "Admin access required", StatusHTTP: 403, ID: cherry.ErrID{SID: 0x2, Kind: 0x1}, Details: []string(nil)}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

// ErrRequiredHeadersNotProvided error
// Required headers not provided
func ErrRequiredHeadersNotProvided(params ...func(*cherry.Err)) *cherry.Err {
	err := &cherry.Err{Message: "Required headers not provided", StatusHTTP: 400, ID: cherry.ErrID{SID: 0x2, Kind: 0x2}, Details: []string(nil)}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

// ErrRequestValidationFailed error
// Validation error when parsing request
func ErrRequestValidationFailed(params ...func(*cherry.Err)) *cherry.Err {
	err := &cherry.Err{Message: "Request validation failed", StatusHTTP: 400, ID: cherry.ErrID{SID: 0x2, Kind: 0x3}, Details: []string(nil)}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

func ErrUnableGetResourcesList(params ...func(*cherry.Err)) *cherry.Err {
	err := &cherry.Err{Message: "Unable to get resources list", StatusHTTP: 500, ID: cherry.ErrID{SID: 0x2, Kind: 0x4}, Details: []string(nil)}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

func ErrUnableGetResource(params ...func(*cherry.Err)) *cherry.Err {
	err := &cherry.Err{Message: "Unable to get resource", StatusHTTP: 500, ID: cherry.ErrID{SID: 0x2, Kind: 0x5}, Details: []string(nil)}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

func ErrUnableCreateResource(params ...func(*cherry.Err)) *cherry.Err {
	err := &cherry.Err{Message: "Unable to create resource", StatusHTTP: 500, ID: cherry.ErrID{SID: 0x2, Kind: 0x6}, Details: []string(nil)}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

func ErrUnableUpdateResource(params ...func(*cherry.Err)) *cherry.Err {
	err := &cherry.Err{Message: "Unable to update resource", StatusHTTP: 500, ID: cherry.ErrID{SID: 0x2, Kind: 0x7}, Details: []string(nil)}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

func ErrUnableDeleteResource(params ...func(*cherry.Err)) *cherry.Err {
	err := &cherry.Err{Message: "Unable to delete resource", StatusHTTP: 500, ID: cherry.ErrID{SID: 0x2, Kind: 0x8}, Details: []string(nil)}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

func ErrResourceAlreadyExists(params ...func(*cherry.Err)) *cherry.Err {
	err := &cherry.Err{Message: "Resource with this name already exists", StatusHTTP: 409, ID: cherry.ErrID{SID: 0x2, Kind: 0x9}, Details: []string(nil)}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

func ErrResourceNotExist(params ...func(*cherry.Err)) *cherry.Err {
	err := &cherry.Err{Message: "Resource with this name doesn't exist", StatusHTTP: 404, ID: cherry.ErrID{SID: 0x2, Kind: 0xa}, Details: []string(nil)}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}
func renderTemplate(templText string) string {
	buf := &bytes.Buffer{}
	templ, err := template.New("").Parse(templText)
	if err != nil {
		return err.Error()
	}
	err = templ.Execute(buf, map[string]string{})
	if err != nil {
		return err.Error()
	}
	return buf.String()
}