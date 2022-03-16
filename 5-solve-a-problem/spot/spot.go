package spot

type Spot struct {
	Name        string
	Description string
	Lat, Long   float64
	Available   bool     // Is the spot available?
	Price       float32  // in dollars
	Size        float32  // In sq. meters
	ImageURLs   []string // URLs of images from the spot
}
