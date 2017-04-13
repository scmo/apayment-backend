package models

type Request struct {
	Id             int64 `json:"id"`
	Contributions []*Contribution `orm:"rel(m2m)" json:"contributions"`
	User          *User `orm:"rel(fk)" json:"-"`
}