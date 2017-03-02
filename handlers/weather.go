package handlers

import (
	"github.com/yujie0121/canvas/libhttp"
	"html/template"
	"net/http"
	"github.com/yujie0121/canvas/remote"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	city := r.FormValue("city")
	cityInfo :=remote.GetAirData(city)
	tmpl, err := template.ParseFiles("templates/dashboard.html.tmpl", "templates/weather.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	tmpl.Execute(w, cityInfo)
}
