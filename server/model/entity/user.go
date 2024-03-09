package entity

type User struct {
	Id        int    `gorm:"primary_key" json:"id"`
	Username  string `json:"username" form:"username" gorm:"unique"`
	Password  string `json:"password" form:"password"`
	Nickname  string `json:"nickname" form:"nickname"`
	Email     string `json:"email" form:"email"`
	Phone     string `json:"phone" form:"phone"`
	GithubUID int    `json:"github_uid" form:"github_uid"`
}
