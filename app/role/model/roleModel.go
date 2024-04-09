package model

type User struct {
	ID          uint64 `gorm:"primaryKey"`
	CreatedTime int64  `gorm:"autoCreateTime"`
	UpdatedTime int64  `gorm:"autoUpdateTime"`
	DeletedTime int64  `gorm:"autoDeleteTime"`
	DelState    int8
	Role        int8
	Mobile      string `gorm:"unique"`
	PassWord    string
	Nickname    string
}
