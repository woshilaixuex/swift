package model

type Salt struct {
	ID     uint64 `gorm:"primaryKey"`
	RoleID uint64 `gorm:"unique"`
	Salt   string
}
