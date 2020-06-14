package hello

import "github.com/rianekacahya/graphql/internal/v1/entity"

func (DI *hello) GetHello() (*entity.Hello, error) {
	return &entity.Hello{
		ID: 1,
		Nama: "Rian Eka Cahya",
		Alamat: "Bekasi",
	}, nil
}