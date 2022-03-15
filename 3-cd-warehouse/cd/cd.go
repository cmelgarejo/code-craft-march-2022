package cd

type CD struct {
	Title    string
	Artist   string
	Quantity int
	Ratings  []Rating
	// Rating   int
	// Comments []string
}

type Rating struct {
	Rating  int
	Comment string
}

const (
	NOT_ENOUGH_STOCK = "not enough stock"
	SUCCESS          = "success"
)

func (cd *CD) Rate(rating int, comment string) bool {
	cd.Ratings = append(cd.Ratings, Rating{Rating: rating, Comment: comment})
	return true
}

func (cd *CD) Buy(ccNo string, ccExpDate string) string {
	if cd.Quantity > 0 {
		cd.Quantity--
		return SUCCESS
	}
	return NOT_ENOUGH_STOCK
}
