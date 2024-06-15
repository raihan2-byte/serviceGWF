package entity

type BuyerAddress struct {
	ID          int
	UserID      int
	Province    string
	City        string
	SubDistrict string
	HomeAddress string
	User        User `gorm:"foreignKey:UserID"`
}
