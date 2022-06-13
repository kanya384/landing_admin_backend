package domain

type LandingContent struct {
	Advantages  []*Advantage
	Docs        []*Doc
	Editables   []*Editable
	Hod         []Year
	Plans       []*Plan
	Posters     []*Poster
	ProjectInfo []*ProjectInfo
	Video       []*Video
}
