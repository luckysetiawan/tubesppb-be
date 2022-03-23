package data

type User struct {
	UID             int    `form:"uid" json:"uid"`
	Username        string `form:"username" json:"username"`
	Profile_picture string `form:"profile_picture" json:"profile_picture"`
	Friend_mode     bool   `form:"friend_mode" json:"friend_mode"`
}
