package service

// result codes
var (
	SUCCESS             = 0
	ENT_CHECK_ERROR     = 10001
	INVALID_CHANNEL     = 10002
	GET_USER_DATA_ERROR = 10003
	USERNAME_EXIST      = 10004
	INCORRECT_PASSWORD  = 10005
	ILLEGAL_BASE64_DATA = 10006
	ILLEGAL_CNT_TYPE    = 10007
	REGISTER_ERROR      = 10008
)

// error contents
var (
	ENT_CHECK_ERROR_MSG  = "ent request check error"
	INVALID_CHANNEL_MSG  = "no such request channel"
	ILLEGAL_CNT_TYPE_MSG = "content type isn't available in current channel"
)

// db
var (
	NO_USERNAME_MSG        = "user name doesn't exist"
	USERNAME_EXIST_MSG     = "user name already exists"
	INCORRECT_PASSWORD_MSG = "password is incorrect"
)
