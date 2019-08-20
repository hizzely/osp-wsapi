package database

import "github.com/hizzely/osp-wsapi/models"

// Matkul get matkul by ID
func Matkul(id int) (*models.MataKuliah, error) {
	return models.FindMataKuliah(Ctx, Db, id)
}

// MatkulAll get all matkul
func MatkulAll() (models.MataKuliahSlice, error) {
	return models.MataKuliahs().All(Ctx, Db)
}
