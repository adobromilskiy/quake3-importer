package importer

type (
	XMLMatch struct {
		Map      string      `xml:"map,attr"`
		Type     string      `xml:"type,attr"`
		Duration uint        `xml:"duration,attr"`
		Datetime string      `xml:"datetime,attr"`
		Players  []XMLPlayer `xml:"player"`
	}

	XMLPlayer struct {
		Name     string   `xml:"name,attr"`
		Stats    []Stat   `xml:"stat"`
		Weapons  []Weapon `xml:"weapons>weapon"`
		Items    []Item   `xml:"items>item"`
		Powerups []Item   `xml:"powerups>item"`
	}

	Stat struct {
		Name  string `xml:"name,attr"`
		Value int    `xml:"value,attr"`
	}

	Match struct {
		Map      string   `json:"map" bson:"map"`
		Type     string   `json:"type" bson:"type"`
		Duration uint     `json:"duration" bson:"duration"`
		Datetime string   `json:"datetime" bson:"datetime"`
		Players  []Player `json:"players" bson:"players"`
	}

	Player struct {
		Name        string   `json:"name" bson:"name"`
		Score       int      `json:"score" bson:"score"`
		Kills       uint     `json:"kills" bson:"kills"`
		Deaths      uint     `json:"deaths" bson:"deaths"`
		Suicides    uint     `json:"suicides" bson:"suicides"`
		DamageGiven uint     `json:"damage_given" bson:"damage_given"`
		DamageTaken uint     `json:"damage_taken" bson:"damage_taken"`
		HealthTotal uint     `json:"health_total" bson:"health_total"`
		ArmorTotal  uint     `json:"armor_total" bson:"armor_total"`
		Weapons     []Weapon `json:"weapons" bson:"weapons"`
		Items       []Item   `json:"items" bson:"items"`
		Powerups    []Item   `json:"powerups,omitempty" bson:"powerups,omitempty"`
	}

	Item struct {
		Name    string `xml:"name,attr" json:"name" bson:"name"`
		Pickups uint   `xml:"pickups,attr" json:"pickups" bson:"pickups"`
		Time    uint   `xml:"time,attr,omitempty" json:"time,omitempty" bson:"time,omitempty"`
	}

	Weapon struct {
		Name  string `xml:"name,attr" json:"name" bson:"name"`
		Hits  uint   `xml:"hits,attr" json:"hits" bson:"hits"`
		Shots uint   `xml:"shots,attr" json:"shots" bson:"shots"`
		Kills uint   `xml:"kills,attr" json:"kills" bson:"kills"`
	}
)
