package external

type Album struct {
	Title       string
	Artist      string
	LowestPrice int
}

type AlbumChartSvc interface {
	Notify(title string, artist string, quantity int) bool
	GetAlbum(title string, artist string) Album
}

// func NewAlbumChartSvc() *AlbumChartSvc {
type AlbumChartSvcImpl struct{}

func (s AlbumChartSvcImpl) Notify(title string, artist string, quantity int) bool {
	return true
}

func (s AlbumChartSvcImpl) GetAlbum(title string, artist string) Album {
	return Album{
		Title:       title,
		Artist:      artist,
		LowestPrice: 100,
	}
}

func NewAlbumChartSvc() AlbumChartSvc {
	return AlbumChartSvcImpl{}
}
