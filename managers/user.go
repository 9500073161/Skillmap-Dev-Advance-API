package managers

import "github.com/9500073161/skill-map-app/models"

type UserManager struct {
	//dbClient
}

func NewUserManager()*UserManager{
	return &UserManager{}
}
func (userMgr *UserManager) Create(user *models.User)(*models.User,error){
return nil,nil
}