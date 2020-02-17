package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Api-Marvel/api/server"
	"github.com/Api-Marvel/internal/environments"
	"github.com/Api-Marvel/internal/models"
	"gopkg.in/mgo.v2/bson"
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

//CreateComic ...
func CreateComic(comic *models.ResultsComics) int {
	collection := server.ConnectedDB()
	results, err := collection.InsertOne(context.TODO(), comic)
	if err != nil {
		server.GetError(err)
		return 2
	}
	fmt.Println(results)
	return 1
}

//GetComicdb ...
func GetComicdb() []models.ResultsComics {
	var comics []models.ResultsComics
	collection := server.ConnectedDB()
	result, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		server.GetError(err)
	}
	defer result.Close(context.TODO())
	for result.Next(context.TODO()) {
		var comic models.ResultsComics
		err := result.Decode(&comic)
		if err != nil {
			log.Fatal(err)
		}
		comics = append(comics, comic)
	}

	return comics

}
