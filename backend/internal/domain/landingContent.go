package domain

type LandingContent struct {
	Advantages      []*Advantage
	AdvantagePhotos []AdvantagePhoto
	Docs            []*Doc
	Editables       []*Editable
	Years           []Year
	Months          []Month
	Photos          []HodPhoto
	Plans           []*Plan
	Posters         []*Poster
	ProjectInfo     []*ProjectInfo
	Video           []*Video
}
