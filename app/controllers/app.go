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
type Nutrient struct {
	id int `json:"nutrient_id"`
	Name string `json:"name"`
	Group string `json:"group"`
	Unit string `json:"unit"`
	Value float64 `json:"value"`
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
type FoodReport struct {
	Sr string `json:"sr"`
	Type string `json:"type"`
	Food FoodReportItem `json:"food"`
}
type FoodReportItem struct {
	id string `json:"ndbno"`
	Name string `json:"name"`
	Sd string `json:"sd"`
	Fg string `json:"fg"`
	Nutrients []Nutrient `json:"nutrients"`
}


type NutrientReport struct {
	List NutrientList
	Report FoodReport `json:"report"`
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Search() revel.Result {
	var url string
	var apiKey string
	var term = c.Params.Get("q")
	entries := new(NutrientReport)

	apiKey, _ = revel.Config.String("app.nutrition_api_key")
	url = "http://api.nal.usda.gov/ndb/search/?format=json"+
				"&sort=n&max=25&offset=0"+
				"&q="+term+
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

func (c App) Nutrient(id string) revel.Result {
	var url string
	var apiKey string
	nutrient := new(NutrientReport)

	apiKey, _ = revel.Config.String("app.nutrition_api_key")
	url = "http://api.nal.usda.gov/ndb/reports/?format=json"+
				"&type=f"+
				"&ndbno="+string(id)+
				"&api_key="+apiKey

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
			defer res.Body.Close()
			json.NewDecoder(res.Body).Decode(nutrient)
		}

	return c.Render(nutrient)
}
