package shoppingpal

type Item struct {
	Id           string  `json:id`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Rating       string  `json:"rating"`
	Summary      string  `json:"summary"`
	Manufacturer string  `json:"manufacturer"`
	Sources      Source  `json:"source"`
}

type Source struct {
	Name          string `json:"name"`
	Author        string `json:"author"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	URL           string `json:"url"`
	DatePublished string `json:"date"`
}

type Items []Item

func (I Items) append(itm Item) {
	I = append(I, itm)
}

// Return Items
func NewItems() Items {
	source := Source{
		Name:          "N1",
		Author:        "A1",
		Title:         "T1",
		Content:       "C1",
		URL:           "https://test",
		DatePublished: "09-08-2019",
	}

	item1 := Item{
		Id:      "10001",
		Name:    "Book",
		Price:   40,
		Summary: "My Book",
		Sources: source,
	}

	item2 := Item{
		Id:      "10002",
		Name:    "Book2",
		Price:   20,
		Summary: "My Book 2",
		Sources: source,
	}

	ItemList := Items{
		item1,
		item2,
	}

	return ItemList
}
