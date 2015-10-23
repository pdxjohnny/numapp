package user

import (
	"log"
	"net/http"

	customhttp "github.com/pdxjohnny/numapp/http"
	"github.com/pdxjohnny/numapp/user/api"
	"github.com/spf13/viper"
)

// Run starts the http(s) server for the cli
func Run() {
	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", *api.MakeHandler()))
	err := customhttp.ServeMux(
		mux,
		viper.GetString("addr"),
		viper.GetString("port"),
		viper.GetString("cert"),
		viper.GetString("key"),
	)
	if err != nil {
		log.Println(err)
		return
	}
}
