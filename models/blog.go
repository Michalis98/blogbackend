package models

type Blog struct {
	Id     uint   `json:"id"`
	Title  string `json:title`
	Desc   string `json:desc`
	Image  string `json:image`
	UserID string `json:userid`
	//One blog belongs to a user (foreign key)
	User User `json:"user";gorm:foreignkey:UserID"`
}
