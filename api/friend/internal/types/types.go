// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.3

package types

type AddFriendRequest struct {
	AdderEmail  string `json:"adder_email"`
	TargetEmail string `json:"target_email"`
	Content     string `json:"content"`
}

type AddFriendResponse struct {
	AddMessage string `json:"add_message"`
}

type DeleteFriendRequest struct {
	DeleteEmail string `json:"delete_email"`
	TargetEmail string `json:"target_email"`
}

type DeleteFriendResponse struct {
	DeleteMessage string `json:"delete_message"`
}

type FriendInfor struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

type GetFriendsRequest struct {
	UserEmail string `json:"user_email"`
}

type GetFriendsResponse struct {
	Friends []FriendInfor `json:"friend_infor"`
}
