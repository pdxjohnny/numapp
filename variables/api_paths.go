package variables

const (
	// APIPathAccountServer is the path to the
	APIPathAccountServer = "/account/:id"
	// APIPathAccount is the path to the
	APIPathAccount = "/api" + APIPathAccountServer
)

var (
	// BlankResponse is so that we can send something without an EOF
	BlankResponse = []byte("{}")
)
