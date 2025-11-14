package register

type request struct {
	Nickname string `form:"nickname" json:"nickname" binding:"required,max=40"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=8"`
}
