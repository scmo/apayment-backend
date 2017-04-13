package services

import (
	"github.com/scmo/foodchain-backend/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

func CreateUser(u *models.User) error {
	o := orm.NewOrm()
	err := CreatePlant(u.Plant)
	_, err = o.Insert(u)
	m2m := o.QueryM2M(u, "Roles")
	for _, rolePtr := range u.Roles {
		if _, id, err := o.ReadOrCreate(rolePtr, "Name"); err == nil {
			rolePtr.Id = id
		} else {
			beego.Error("ReadOrCreate ", err.Error())
			return err
		}
		_, err := m2m.Add(rolePtr)
		if err != nil {
			beego.Error("Many2Many Add ", err.Error())
			return err
		}
	}
	return err
}

func CheckLogin(_username string, _password string) (models.User, error) {
	o := orm.NewOrm()
	user := models.User{Username: _username, Password: _password}
	err := o.Read(&user, "Username", "Password")
	beego.Debug(user)
	return user, err;
}

func GetAllUsers() ([]*models.User) {
	o := orm.NewOrm()
	var users []*models.User
	o.QueryTable(new(models.User)).All(&users)
	for _, user := range users {
		o.LoadRelated(user, "Roles")
	}
	return users
}

func GetUserById(_id int64) (*models.User, error) {
	o := orm.NewOrm()
	user := models.User{Id: _id}
	err := o.Read(&user)
	if err == orm.ErrNoRows {
		beego.Error("No result found.")
		return nil, err
	} else if err == orm.ErrMissPK {
		beego.Error("No primary key found.")
		return nil, err
	}
	o.LoadRelated(&user, "Roles")
	return &user, nil
}

func GetUserByUsername(_username string) (*models.User, error) {
	o := orm.NewOrm()
	user := models.User{Username: _username}
	err := o.Read(&user, "Username")
	if err == orm.ErrNoRows {
		beego.Error("No result found.")
		return nil, err
	} else if err == orm.ErrMissPK {
		beego.Error("No primary key found.")
		return nil, err
	}
	o.LoadRelated(&user, "Roles")
	return &user, nil
}