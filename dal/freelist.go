package dal

type Freelist struct {
	MaxPage       PGNum
	ReleasedPages []PGNum
}

func NewFreelist(initialPage PGNum) *Freelist {
	return &Freelist{
		MaxPage:       initialPage,
		ReleasedPages: []PGNum{},
	}
}

func (fr *Freelist) GetNextPage() PGNum {
	if len(fr.ReleasedPages) != 0 {
		pageID := fr.ReleasedPages[len(fr.ReleasedPages)-1]
		fr.ReleasedPages = fr.ReleasedPages[:len(fr.ReleasedPages)-1]
		return pageID
	}
	fr.MaxPage += 1
	return fr.MaxPage
}

func (fr *Freelist) ReleasePage(page PGNum) {
	fr.ReleasedPages = append(fr.ReleasedPages, page)
}
