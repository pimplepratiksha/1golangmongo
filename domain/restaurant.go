package domain

//Restaurant is a domain object
type Restaurant struct {
	DBID         ID      `json:"_id" bson:"_id"`
	Name         string  `json:"name" bson:"name"`
	Address      string  `json:"address" bson:"address"`
	AddressLine2 string  `json:"address line 2" bson:"address line 2"`
	URL          string  `json:"URL" bson:"URL"`
	Outcode      string  `json:"outcode" bson:"outcode"`
	Postcode     string  `json:"postcode" bson:"postcode"`
	Rating       float32 `json:"rating" bson:"rating"`
	TypeOfFood   string  `json:"type_of_food" bson:"type_of_food"`
}
