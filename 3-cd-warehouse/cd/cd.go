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
