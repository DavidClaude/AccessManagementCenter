package service

import (
	"errors"
)

type UserData struct {
	UserName string
	Password string
}

var sim_db map[string]string

func init() {
	sim_db = make(map[string]string)
	sim_db["davidclaude"] = "success0325"
	sim_db["user0"] = "123456"
}

func RequestSqlForUserData(userName string) (userData *UserData, err error) {
	pwd, ok := sim_db[userName]
	if !ok {
		return nil, errors.New(NO_USERNAME_MSG)
	}
	userData = &UserData{
		UserName: userName,
		Password: pwd,
	}
	return userData, nil
}

func RequestSqlForRegister(userName, password string) (err error) {
	if _, ok := sim_db[userName]; ok {
		return errors.New(USERNAME_EXIST_MSG)
	}
	sim_db[userName] = password
	return nil
}
