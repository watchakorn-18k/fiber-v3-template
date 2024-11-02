package entities

type CarbonResponse struct {
	Url         string      `json:"url"`
	Green       interface{} `json:"green"`
	Bytes       int         `json:"bytes"`
	CleanerThan float64     `json:"cleanerThan"`
	Rating      string      `json:"rating"`
}
