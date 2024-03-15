package dto

import "time"

type AdminInfoResponse struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	LoginTime    time.Time `json:"login_time"`
	Avatar       string    `json:"avatar"`
	Introduction string    `json:"introduction"`
	Roles        []string  `json:"roles"`
}

type AdminChangeRequest struct {
	OldPassword string `json:"old_password" label:"旧密码" rule:"notnull"`
	NewPassword string `json:"new_password" label:"新密码" rule:"notnull"`
}
