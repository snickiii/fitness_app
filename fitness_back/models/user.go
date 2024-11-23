package models

type RoleType string

const (
	RoleTech     RoleType = "tech"
	RoleConsumer RoleType = "consumer"
)

type User struct {
	UserID   uint   `gorm:"primary_key;auto_increment"`
	Email    string `gorm:"unique;not null"`
	Name     string
	Username string   `gorm:"unique;not null"`
	Password string   `gorm:"not null"`
	Role     RoleType `gorm:"type:varchar(10);default:'consumer'"`
	Targets  []Target `gorm:"foreignKey:UserID"`
}

type ProfileResponse struct {
	UserID   uint
	Email    string
	Name     string
	Username string
	Role     RoleType
	Targets  []Target
}
