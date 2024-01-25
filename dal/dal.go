package dal

import (
	"fmt"
	"os"
)

type Dal struct {
	File     *os.File
	PageSize int

	*Freelist
}

func NewDal(path string, pageSize int) (*Dal, error) {
	// RDWR = ReadWrite, CREATE = Create if not exist
	// 0666 = Permission to be written by owner only
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	dal := &Dal{
		file,
		pageSize,
		NewFreelist(0),
	}
	return dal, nil
}

func (d *Dal) Close() error {
	if d.File != nil {
		err := d.File.Close()
		if err != nil {
			return fmt.Errorf("could not close file: %s", err)
		}
		d.File = nil
	}
	return nil
}

func (d *Dal) AllocateEmptyPage() *Page {
	return &Page{
		Data: make([]byte, d.PageSize),
	}
}

func (d *Dal) ReadPage(pageNum PGNum) (*Page, error) {
	p := d.AllocateEmptyPage()

	offset := int(pageNum) * d.PageSize

	_, err := d.File.ReadAt(p.Data, int64(offset))
	if err != nil {
		return nil, err
	}
	return p, err
}

func (d *Dal) WritePage(p *Page) error {
	offset := int64(p.Num) * int64(d.PageSize)
	_, err := d.File.WriteAt(p.Data, offset)
	return err
}
