package api

import (
	"net/http"
	"log"
	"github.com/canvas/remote"
	"fmt"
)

func weather(w http.ResponseWriter, r *http.Request)  {
	city := r.FormValue("city")
	cityInfo :=remote.GetAirData(city)
	fmt.Fprint(w, cityInfo)
}

func Serve()  {
	http.HandleFunc("/weather", weather)
	err := http.ListenAndServe(":9000",nil)
	if err != nil {
		log.Fatal("serve error!")
	}
}
