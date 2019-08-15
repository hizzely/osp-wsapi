package database

import (
	"github.com/hizzely/osp-wsapi/models"
	"github.com/volatiletech/sqlboiler/boil"
)

// ClassroomAll get all classroom record
func ClassroomAll() (models.KelasMHSlice, error) {
	return models.KelasMHS().All(Ctx, Db)
}

// ClassroomCreate create new classroom
func ClassroomCreate(namaKelas string) error {
	data := &models.KelasMH{NamaKelas: namaKelas}
	return data.Insert(Ctx, Db, boil.Infer())
}
