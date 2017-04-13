package models

type Request struct {
	Id            uint32
	Contributions []*Contribution `orm:"rel(m2m)" json:"contributions"`
	User          *User `orm:"rel(fk)" json:"-"`
}