package cd

type CD struct {
	Title    string
	Artist   string
	Price    int
	Quantity int
	Ratings  []Rating
}

type Rating struct {
	Rating  int
	Comment string
}

const (
	NOT_ENOUGH_STOCK = "not enough stock"
	SUCCESS          = "success"
)

func NewCD(title, artist string, quantity, price int, ratings []Rating) CD {
	return CD{
		Title:    title,
		Artist:   artist,
		Price:    price,
		Quantity: quantity,
		Ratings:  ratings,
	}
}

func (cd *CD) Rate(rating int, comment string) bool {
	cd.Ratings = append(cd.Ratings, Rating{Rating: rating, Comment: comment})
	return true
}

func (cd *CD) GetPrice() int {
	return cd.Price
}
