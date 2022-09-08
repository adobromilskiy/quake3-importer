package importer

type (
	Match struct {
		Map      string   `xml:"map,attr"`
		Type     string   `xml:"type,attr"`
		Duration int      `xml:"duration,attr"`
		Datetime string   `xml:"datetime,attr"`
		Players  []Player `xml:"player"`
	}

	Player struct {
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

	Item struct {
		Name    string `xml:"name,attr"`
		Pickups int    `xml:"pickups,attr"`
		Time    int    `xml:"time,attr,omitempty"`
	}

	Weapon struct {
		Name  string `xml:"name,attr"`
		Hits  int    `xml:"hits,attr"`
		Shots int    `xml:"shots,attr"`
		Kills int    `xml:"kills,attr"`
	}
)
