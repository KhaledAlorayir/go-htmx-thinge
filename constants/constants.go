package constants

import "fmt"

var USER_PATH = "/user"
var CREATE_USER_ROUTE = fmt.Sprintf("%s/create-user", USER_PATH)

var AUTH_PATH = "/auth"
var LOGOUT_ROUTE = fmt.Sprintf("%s/logout", AUTH_PATH)

var AUTH_COOKIE_NAME = "auth_token"

var CONTEXT_AUTH_DATA = "AUTH_DATA"
