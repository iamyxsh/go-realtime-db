package dal

type PGNum uint64

type Page struct {
	Num  PGNum
	Data []byte
}
