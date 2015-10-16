package get

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Run is the function to be run for the cli
func Run() {
	data, err := Get(viper.GetString("num"))
	if err != nil {
		log.Println("ERROR geting number", err)
		return
	}
	dump, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Println("ERROR dumping number", err)
		return
	}
	fmt.Println(dump)
}
