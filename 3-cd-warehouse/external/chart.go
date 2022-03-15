package external

type Album struct {
	Title       string
	Artist       string
	LowestPrice int
}

type AlbumChartSvc interface {
	Notify(title string, artist string, quantity int) bool
	GetAlbum(title string, artist string) Album
}
