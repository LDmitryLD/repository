package models

type User struct {
	ID        int    `db:"id" db_type:"BIGSERIAL PRIMARY KEY" db_ops:""`
	FirstName string `db:"first_name" db_type:"VARCHAR(100)" db_ops:"create,update"`
	LastName  string `db:"last_name" db_type:"VARCHAR(100)" db_ops:"create,update"`
	Username  string `db:"username" db_type:"VARCHAR(100)" db_ops:"create,update"`
	Email     string `db:"email" db_type:"VARCHAR(100)" db_ops:"create,update"`
	Address   string `db:"address" db_type:"VARCHAR(100)" db_ops:"create,update"`
}

func (u User) TableName() string {
	return "users"
}

func (u User) GetID() int {
	return u.ID
}
