package commands

// ConfigOptions is used to set viper defaults
var ConfigOptions = map[string]interface{}{
	"get": map[string]interface{}{
		"id": map[string]interface{}{
			"value": "",
			"help":  "The id's doc to return",
		},
	},
	"put": map[string]interface{}{},
	"http": map[string]interface{}{
		"addr": map[string]interface{}{
			"value": "0.0.0.0",
			"help":  "Address to bind to",
		},
		"port": map[string]interface{}{
			"value": 8080,
			"help":  "Port to bind to",
		},
		"cert": map[string]interface{}{
			"value": "keys/http/cert.pem",
			"help":  "Certificate for https server",
		},
		"key": map[string]interface{}{
			"value": "keys/http/key.pem",
			"help":  "Key for https server",
		},
		"static": map[string]interface{}{
			"value": "static",
			"help":  "Directory which holds static content",
		},
	},
}
