package domain

type User struct {
	ID        uint   `gorm:"column:id;primaryKey" json:"id"`
	Email     string `gorm:"column:email" json:"email"`
	Password  string `gorm:"password" json:"-"`
	FirstName string `gorm:"first_name" json:"firstName"`
	LastName  string `gorm:"last_name" json:"lastName"`
}

func (User) TableName() string {
	return "TB_USER"
}
