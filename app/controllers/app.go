package controllers

import (
	"github.com/revel/revel"
	"net/http"
	"log"
	"io/ioutil"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Search() revel.Result {
	var url string
	var apiKey string

	apiKey, _ = revel.Config.String("app.nutrition_api_key")
	url = "http://api.nal.usda.gov/ndb/search/?format=json&q="+
				"butter"+
				"&sort=n&max=25&offset=0"+
				"&api_key="+apiKey

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
			defer res.Body.Close()
			log.Println(ioutil.ReadAll(res.Body))
			log.Println(res.Body)
		}

	return c.Render()
}
