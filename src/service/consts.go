package service

// result codes
var (
	SUCCESS             = 0
	ENT_CHECK_ERROR     = 10001
	INVALID_CHANNEL     = 10002
	GET_USER_DATA_ERROR = 10003
	INCORRECT_PASSWORD  = 10004
)

// error contents
var (
	ENT_CHECK_ERROR_MSG = "ent request check error"
	INVALID_CHANNEL_MSG = "no such request channel"
)

// db
var (
	NO_USER_NAME_MSG       = "user name doesn't exist"
	INCORRECT_PASSWORD_MSG = "password is incorrect"
)
