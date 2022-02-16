package model

import "time"

type User struct {
	Id       string `db:"id" json:"id"`     //primary
	Name     string `db:"name" json:"name"` //foreign
	Password string `db:"password" json:"password"`
	Email    string `db:"email" json:"email"`         //unique
	Mobile   string `db:"mobile_no" json:"mobile_No"` //unique
	Userid   string `db:"userid" json:"userid"`
}
type LoggedUser struct {
	ID        string    `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Password  string    `db:"password" json:"password"`
	Userid    string    `db:"userid" json:"userid"`
	Createdat time.Time `db:"createdat" json:"createdat"`
}
