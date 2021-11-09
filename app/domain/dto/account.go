package dto

import "yatter-backend-go/app/domain/object"

type Account struct {
	Username    string          `json:"username"`
	DisplayName *string         `json:"display_name"`
	CreateAt    object.DateTime `json:"create_at"`
	Avatar      *string         `json:"avatar"`
	Header      *string         `json:"header"`
	Note        *string         `json:"note"`
}
