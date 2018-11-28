package service

import "errors"

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

func GetUserDataFromSql(userName string) (userData *UserData, err error) {
	pwd, ok := sim_db[userName]
	if !ok {
		return nil, errors.New(NO_USER_NAME_MSG)
	}
	userData = &UserData{
		UserName: userName,
		Password: pwd,
	}
	return userData, nil
}

func ContainUserName(userName string) (ok bool, err error) {
   pwd, ok := sim_db[userName]
   if ok {
      return true, nil
   }
   return false, nil
   
}