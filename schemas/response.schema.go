package schemas

import db "golang-crudsqlc-rest/db/sqlc"

// Response shapes used by the swagger docs. They mirror the gin.H payloads
// returned from the controllers.

type ContactResponse struct {
	Status  string     `json:"status" example:"successfully created contact"`
	Contact db.Contact `json:"contact"`
}

type ContactListResponse struct {
	Status   string       `json:"status" example:"Successfully retrieved all contacts"`
	Size     int          `json:"size" example:"1"`
	Contacts []db.Contact `json:"contacts"`
}

type ErrorResponse struct {
	Status  string `json:"status" example:"failed"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}
