package delivery

import "github.com/ALTA-BE10-AthaDhaiffathin/Deployment-3.git/domain"

type InsertFormat struct {
	Judul    string `json:"judul"`
	Penerbit string `json:"penerbit"`
}

func (i *InsertFormat) ToDomain() domain.Book {
	return domain.Book{
		Judul:    i.Judul,
		Penerbit: i.Penerbit,
	}
}
