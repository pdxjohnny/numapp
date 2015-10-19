package put

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

// Run is the function to be run for the cli
func Run() {
	dec := json.NewDecoder(os.Stdin)
	for {
		var doc map[string]interface{}
		err := dec.Decode(&doc)
		if err == io.EOF {
			return
		} else if err != nil {
			log.Println(err)
			return
		}
		err = Put(&doc)
		if err != nil {
			log.Println(err)
		}
	}
}
