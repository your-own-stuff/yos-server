package types

type Systemstatus struct {
	Id	  	string	`db:"id" json:"id"`
	Name  	string	`db:"name" json:"name"`
	Value 	string	`db:"value" json:"value"`
}