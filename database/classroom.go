package database

import (
	"github.com/hizzely/osp-wsapi/models"
	"github.com/volatiletech/sqlboiler/boil"
)

// Classroom get by ID
func Classroom(id int) (*models.KelasMH, error) {
	return models.FindKelasMH(Ctx, Db, id)
}

// ClassroomAll get all classroom record
func ClassroomAll() (models.KelasMHSlice, error) {
	return models.KelasMHS().All(Ctx, Db)
}

// ClassroomCreate create new classroom
func ClassroomCreate(namaKelas string) error {
	data := &models.KelasMH{NamaKelas: namaKelas}
	return data.Insert(Ctx, Db, boil.Infer())
}
