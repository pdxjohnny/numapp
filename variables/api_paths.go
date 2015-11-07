package variables

const (
	// APIPathAccountServer is the path to the
	APIPathAccountServer = "/account/:id"
	// APIPathAccount is the path to the
	APIPathAccount = "/api" + APIPathAccountServer
	// APIPathUserServer is the path to a users settings
	APIPathUserServer = "/user/:id"
	// APIPathUser is the path to a users settings
	APIPathUser = "/api" + APIPathUserServer
	// APIPathLoginUserServer is the path for a user to login
	APIPathLoginUserServer = "/user/login"
	// APIPathLoginUser is the for a user to login
	APIPathLoginUser = "/api" + APIPathLoginUserServer
	// APIPathRefreshUserServer is the path for a user to refresh JWT
	APIPathRefreshUserServer = "/user/refresh"
	// APIPathRefreshUser is the for a user to refresh JWT
	APIPathRefreshUser = "/api" + APIPathRefreshUserServer
	// APIPathRegisterUserServer is the path for a user to register
	APIPathRegisterUserServer = "/user/register"
	// APIPathRegisterUser is the for a user to register
	APIPathRegisterUser = "/api" + APIPathRegisterUserServer
)

var (
	// BlankResponse is so that we can send something without an EOF
	BlankResponse = []byte("{}")
)
