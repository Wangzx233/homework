package Model

type User struct {
	Uid string		`gorm:"primary_key"`
	Password string	`gorm:"type:varchar(256);"`
	Sign string		`gorm:"type:varchar(256);"`
}

type Records struct {
	Id int
	FromUid string
	ToUid string
	Money float64
	Remarks string
	Createtime string
}
