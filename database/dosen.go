package database

import (
	"log"

	"github.com/hizzely/osp-wsapi/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
)

// DosenKelasMatkul tables join
type DosenKelasMatkul struct {
	KelasID         int    `boil:"kelas_id" json:"id_kelas"`
	NamaKelas       string `boil:"nama_kelas" json:"nama_kelas"`
	MatkulID        int    `boil:"matkul_id" json:"id_matkul"`
	NamaMatkul      string `boil:"nama_matkul" json:"nama_matkul"`
	DeskripsiMatkul string `boil:"deskripsi_matkul" json:"deskripsi_matkul"`
}

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

// DosenLectureSubject get subjects by dosen id
func DosenLectureSubject(id string) ([]DosenKelasMatkul, error) {
	var model []DosenKelasMatkul
	err := queries.Raw(`
		SELECT * FROM dosen_jadwal
		LEFT JOIN dosen ON dosen.id = dosen_jadwal.dosen_id
		LEFT JOIN kelas_mhs ON kelas_mhs.id = dosen_jadwal.kelas_id
		LEFT JOIN mata_kuliah ON mata_kuliah.id = dosen_jadwal.matkul_id
		WHERE dosen_jadwal.dosen_id = ?`, id,
	).Bind(Ctx, Db, &model)
	return model, err
}

// DosenLectureSubjectCreate new subject
func DosenLectureSubjectCreate(dosenID string, matkulID, kelasID int) error {
	tx, err := Db.BeginTx(Ctx, nil)

	if err != nil {
		log.Println("Tx Error", err)
		return err
	}

	jadwal := models.DosenJadwal{
		DosenID:  dosenID,
		MatkulID: matkulID,
		KelasID:  kelasID,
	}

	dosen, _ := models.FindDosen(Ctx, Db, dosenID, "id")
	jadwal.SetDosen(Ctx, Db, false, dosen)

	matkul, _ := models.FindMataKuliah(Ctx, Db, matkulID, "id")
	jadwal.SetMatkul(Ctx, Db, false, matkul)

	kelas, _ := models.FindKelasMH(Ctx, Db, kelasID, "id")
	jadwal.SetKelas(Ctx, Db, false, kelas)

	jadwal.Insert(Ctx, Db, boil.Infer())
	commitErr := tx.Commit()

	if commitErr != nil {
		log.Println("Commit error!", commitErr)
		tx.Rollback()
	}

	return commitErr
}
