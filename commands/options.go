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
}
