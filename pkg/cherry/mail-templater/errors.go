package mailerr

import (
	"net/http"

	"git.containerum.net/ch/kube-client/pkg/cherry"
)

var mailErr = cherry.BuildErr(7)

//Errors
func ErrAdminRequired() *cherry.Err {
	return mailErr("Admin access required", http.StatusForbidden, 1)
}
func ErrPermissionsError() *cherry.Err {
	return mailErr("Unable to verify permissions", http.StatusForbidden, 1)
}
func ErrRequiredHeadersNotProvided() *cherry.Err {
	return mailErr("Required headers not provided", http.StatusForbidden, 2)
}
func ErrRequestValidationFailed() *cherry.Err {
	return mailErr("Request validation failed", http.StatusBadRequest, 3)
}

func ErrUnableGetTemplatesList() *cherry.Err {
	return mailErr("Unable to get templates list", http.StatusInternalServerError, 4)
}
func ErrUnableGetTemplate() *cherry.Err {
	return mailErr("Unable to get template", http.StatusInternalServerError, 5)
}
func ErrUnableSaveTemplate() *cherry.Err {
	return mailErr("Unable to save template", http.StatusInternalServerError, 6)
}
func ErrUnableUpdateTemplate() *cherry.Err {
	return mailErr("Unable to update template", http.StatusInternalServerError, 7)
}
func ErrUnableDeleteTemplate() *cherry.Err {
	return mailErr("Unable to delete template", http.StatusInternalServerError, 8)
}
func ErrTemplateAlreadyExists() *cherry.Err {
	return mailErr("Template with this name already exists", http.StatusConflict, 9)
}
func ErrTemplateNotExist() *cherry.Err {
	return mailErr("Template with this name doesn't exist", http.StatusNotFound, 10)
}
func ErrTemplateVersionNotExist() *cherry.Err {
	return mailErr("Template with this name and version doesn't exist", http.StatusNotFound, 11)
}

func ErrUnableGetMessagesList() *cherry.Err {
	return mailErr("Unable to get templates list", http.StatusInternalServerError, 12)
}
func ErrUnableGetMessage() *cherry.Err {
	return mailErr("Unable to get template", http.StatusInternalServerError, 13)
}
func ErrUnableSaveMessage() *cherry.Err {
	return mailErr("Unable to save template", http.StatusInternalServerError, 14)
}
func ErrMessageNotExist() *cherry.Err {
	return mailErr("Message with this id doesn't exist", http.StatusNotFound, 15)
}

func ErrMailSendFailed() *cherry.Err {
	return mailErr("Mail send failed", http.StatusInternalServerError, 16)
}
