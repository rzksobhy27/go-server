package tables

type User struct {
	Name     string `gorm:"primaryKey,unique,not null"`
	Password string `gorm:"not null"`
}
