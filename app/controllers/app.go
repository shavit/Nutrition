package controllers

import (
	"github.com/revel/revel"
	"net/http"
	"log"
	"encoding/json"
)

type App struct {
	*revel.Controller
}

type NutrientItem struct {
	Id string `json:"ndbno"`
	Name string `json:"name"`
	Group string `json:"group"`
	Offset int `json:"offset"`
}
type NutrientList struct {
	SearchTerm string `json:"q"`
	Sr string `json:"sr"`
	Start int `json:"start"`
	End int `json:"end"`
	Total int `json:"total"`
	Group string `json:"group"`
	Sort string `json:"sort"`
	Items []NutrientItem `json:"item"`
}
type NutrientReport struct {
	List NutrientList
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Search() revel.Result {
	var url string
	var apiKey string
	entries := new(NutrientReport)

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
			json.NewDecoder(res.Body).Decode(entries)
		}

	return c.Render(entries)
}
