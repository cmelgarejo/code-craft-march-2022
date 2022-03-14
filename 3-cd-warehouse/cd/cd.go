package cd

type CD struct {
	Title    string
	Artist   string
	Rating   int
	Quantity int
	Comments  []string
}

type Rating struct {
	Rating  int
	Comment string
}
