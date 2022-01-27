package domain

type User struct {
	ID        int    `gorm:"column:id;primaryKey;autoIncrement"`
	FirstName string `gorm:"column:first_name"`
	Email     string `gorm:"column:email;unique"`
	Password  string `gorm:"column:password"`
}
