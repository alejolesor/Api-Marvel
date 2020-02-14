package models

//ResponseGeneral ...
type ResponseGeneral struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

//Data ...
type Data struct {
	ID      int             `json:"_id"`
	Total   int             `json:"total"`
	Results []ResultsComics `json:"results"`
}

//ResultsComics ...
type ResultsComics struct {
	IDCommic    int    `json:"id"`
	Tittle      string `json:"title"`
	Description string `json:"description"`
	Modified    string `json:"modified"`
}
