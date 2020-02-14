package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Api-Marvel/internal/environments"
	"github.com/Api-Marvel/internal/models"
)

//GetIndexString ...
func GetIndexString() string {
	return "Consumo de index correcto"
}

//ConsumingAPIMarvell , verbHTTP(GET, POST, PUT) , methodComics(v1/public/commics) ...
func ConsumingAPIMarvell(verbHTTP string, methodComics string) *http.Response {
	urlConcat := environments.GetURLMarvell(methodComics)
	client := &http.Client{}
	req, err := http.NewRequest(verbHTTP, urlConcat, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", `*/*`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

//GetAllComics map del json de comics...
func GetAllComics() models.ResponseGeneral {
	var methodComics = "public/comics?"
	var responsgeneral models.ResponseGeneral
	response := ConsumingAPIMarvell("GET", methodComics)
	defer response.Body.Close()
	databyte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	errmars := json.Unmarshal(databyte, &responsgeneral)
	if err != nil {
		log.Fatal(errmars)
	}
	return responsgeneral
}
