package shoppingpal

type Item struct {
	Name         string  `json:"Name"`
	Price        float64 `json:"price"`
	Rating       string  `json:"rating"`
	Summary      string  `json:"summary"`
	Manufacturer string  `json:"manufacturer"`
	Sources      Source  `json:"source"`
}

type Source struct {
	Name          string `json:"Name"`
	Author        string `json:"author"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	URL           string `json:"url"`
	DatePublished string `json:"date"`
}

type ItemList []Item
