package Model

type User struct {
	Uid string		`gorm:"primary_key"`
	Name string 	`gorm:"type:varchar(256);"`
	Password string	`gorm:"type:varchar(256);"`
	Sign string		`gorm:"type:varchar(256);"`
}

type Reply struct {
	Id int			`gorm:"primary_key;auto_increment;"`
	//Update_time time.Time
	Reply_id int	`form:"reply_id"`
	Content string	`form:"content"`
	From_uid string	`form:"from_uid"`
	To_uid string	`form:"to_uid"`
	Power string	`form:"power"`
}
