package model

type City struct {
	Desc string `json:"desc"`
	Status int32 `json:"status"`
	Data struct{
		Wendu string `json:"wendu"`
		Ganmao string `json:"ganmao"`
		Forecast []Weather `json:"forecast"`
		Yesterday struct{
			Fl string `json:"fl"`
			Fx string `json:"fx"`
			Hi string `json:"hi"`
			Status string `json:"type"` 
			Low string `json:"low"`
			Date string `json:"date"`
			  } `json:"yesterday"`
		Aqi string `json:"aqi"`
		City string `json:"city"`
	     } `json:"data"`


}

type Weather struct {
	Fengxiang string `json:"fengxiang"`
	Fengli string `json:"fengli"`
	High string `json:"high"`
	Status string `json:"type"`
	Low string `json:"low"`
	Date string `json:"date"`

}

