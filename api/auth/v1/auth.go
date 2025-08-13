package v1

import "github.com/gogf/gf/v2/frame/g"

type RegisterReq struct {
	g.Meta   `path:"/auth/register" tags:"Authentication" method:"post" summary:"Register a new user"`
	Email    string `json:"email" v:"required|email" description:"User email"`
	Username string `json:"username" v:"required|length:3,30" description:"Username"`
	Password string `json:"password" v:"required|min:6" description:"Password"`
}

type RegisterRes struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    *UserData   `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type LoginReq struct {
	g.Meta   `path:"/auth/login" tags:"Authentication" method:"post" summary:"User login"`
	Email    string `json:"email" v:"required|email" description:"User email"`
	Password string `json:"password" v:"required" description:"Password"`
}

type LoginRes struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    *LoginData `json:"data,omitempty"`
	Error   string    `json:"error,omitempty"`
}

type CheckUsernameReq struct {
	g.Meta   `path:"/auth/check-username" tags:"Authentication" method:"post" summary:"Check username availability"`
	Username string `json:"username" v:"required" description:"Username to check"`
}

type CheckUsernameRes struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Data    *UsernameAvailability `json:"data,omitempty"`
	Error   string               `json:"error,omitempty"`
}

type UserData struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type LoginData struct {
	AccessToken string    `json:"access_token"`
	User        *UserData `json:"user"`
}

type UsernameAvailability struct {
	Username  string `json:"username"`
	Available bool   `json:"available"`
}