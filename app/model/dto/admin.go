package dto

type AdminInfoResponse struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	LoginTime string   `json:"loginTime"`
	Avatar    string   `json:"avatar"`
	Remark    string   `json:"remark"`
	Roles     []string `json:"roles"`
	Nickname  string   `json:"nickname"`
	Email     string   `json:"email"`
	Phone     string   `json:"phone"`
}

type AdminChangeRequest struct {
	OldPassword string `json:"old_password" label:"旧密码" rule:"notnull"`
	NewPassword string `json:"new_password" label:"新密码" rule:"notnull"`
}
