package remote

import (
	"net/http"
	"log"
	"encoding/json"
	"io/ioutil"
	"github.com/canvas/model"
	"github.com/canvas/persistence/redis"
	"github.com/canvas/util"
	"strings"
	"time"
)


func GetAirData(cityName string) model.City {

	var c model.City

	if strings.Trim(cityName, "") == "" {
		cityName = "上海"
	}


	/*data , err := redis.Get(util.KEY_PREFIX+cityName)
	if data != nil {
		err = json.Unmarshal(data, &c)
		if err != nil{
			log.Fatal("Unmarshal data error!")
		}
	}*/

	jsonStr := redis.GetValue(util.KEY_PREFIX+cityName)
	if  strings.Trim(jsonStr, " ") != ""{
		err := json.Unmarshal([]byte(jsonStr), &c)
		if err != nil{
			log.Fatal("Unmarshal data error!")
		}
	}else {
		resp, err := http.Get(util.WEATHER_URL + "?city=" +cityName)

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil{
			log.Fatal("fetch weather data error!")
		}


		err = json.Unmarshal(body, &c)
		if err != nil{
			log.Fatal("Unmarshal data error!")
		}
		if(body != nil){
			jsonValue, err := json.Marshal(c)
			if err != nil {
				log.Fatal("transfer json error!")
			}

			//放入缓存 30秒过期
			go redis.SetValue(util.KEY_PREFIX+cityName, string(jsonValue), time.Second*30)

			log.Println("city["+cityName+"] info put in redis successfully!")
		}

	}

	return c
}
