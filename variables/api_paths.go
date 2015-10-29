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
	// APIPathUserLoginServer is the path for a user to login
	APIPathUserLoginServer = "/user/login"
	// APIPathUserLogin is the for a user to login
	APIPathUserLogin = "/api" + APIPathUserLoginServer
	// APIPathUserRegisterServer is the path for a user to register
	APIPathUserRegisterServer = "/user/register"
	// APIPathUserRegister is the for a user to register
	APIPathUserRegister = "/api" + APIPathUserRegisterServer
)

var (
	// BlankResponse is so that we can send something without an EOF
	BlankResponse = []byte("{}")
)
