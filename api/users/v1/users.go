package v1

import "github.com/gogf/gf/v2/frame/g"

type GetUsersReq struct {
	g.Meta `path:"/users" tags:"Users" method:"get" summary:"Get all users"`
	Page   int `json:"page" d:"1" description:"Page number"`
	Limit  int `json:"limit" d:"10" description:"Number of items per page"`
}

type GetUsersRes struct {
	Users []User `json:"users"`
	Total int    `json:"total"`
}

type GetUserReq struct {
	g.Meta `path:"/users/:id" tags:"Users" method:"get" summary:"Get user by ID"`
	ID     int `json:"id" v:"required|min:1" description:"User ID"`
}

type GetUserRes struct {
	User *User `json:"user"`
}

type CreateUserReq struct {
	g.Meta   `path:"/users" tags:"Users" method:"post" summary:"Create new user"`
	Name     string `json:"name" v:"required" description:"User name"`
	Email    string `json:"email" v:"required|email" description:"User email"`
	Password string `json:"password" v:"required|min:6" description:"User password"`
}

type CreateUserRes struct {
	User *User `json:"user"`
}

type UpdateUserReq struct {
	g.Meta   `path:"/users/:id" tags:"Users" method:"put" summary:"Update user"`
	ID       int    `json:"id" v:"required|min:1" description:"User ID"`
	Name     string `json:"name" description:"User name"`
	Email    string `json:"email" v:"email" description:"User email"`
	Password string `json:"password" v:"min:6" description:"User password"`
}

type UpdateUserRes struct {
	User *User `json:"user"`
}

type DeleteUserReq struct {
	g.Meta `path:"/users/:id" tags:"Users" method:"delete" summary:"Delete user"`
	ID     int `json:"id" v:"required|min:1" description:"User ID"`
}

type DeleteUserRes struct {
	Success bool `json:"success"`
}

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
