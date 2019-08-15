package database

import (
	"github.com/hizzely/osp-wsapi/models"
	"github.com/volatiletech/sqlboiler/boil"
)

// Dosen get single dosen account by their ID
func Dosen(id string) (*models.Dosen, error) {
	return models.FindDosen(Ctx, Db, id)
}

// DosenCreate insert new dosen account
func DosenCreate(id, namaLengkap, hashedPassword string) error {
	dosen := &models.Dosen{
		ID:          id,
		NamaLengkap: namaLengkap,
		Password:    hashedPassword,
	}
	return dosen.Insert(Ctx, Db, boil.Infer())
}

// DosenDelete by their ID
func DosenDelete(id string) (int64, error) {
	dosen, _ := models.FindDosen(Ctx, Db, id, "id")
	return dosen.Delete(Ctx, Db)
}