package database

import (
	"log"
	"strings"

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

// DosenPresenceSessionCreateClassroom create new classroom session
func DosenPresenceSessionCreateClassroom(matkulID, kelasID int, dosenID, judul, deskripsi, sessionCode string) error {
	session := models.SesiPresensi{
		Kode:         sessionCode,
		DosenID:      dosenID,
		MatkulID:     matkulID,
		KelasID:      kelasID,
		KetJudul:     judul,
		KetDeskripsi: deskripsi,
		Status:       "aktif",
	}
	return session.Insert(Ctx, Db, boil.Infer())
}

// DosenPresenceSessionDetail get detail from ID
func DosenPresenceSessionDetail(id int) map[string]interface{} {
	details, _ := models.FindSesiPresensi(Ctx, Db, id)
	dosen, _ := details.Dosen().One(Ctx, Db)
	matkul, _ := details.Matkul().One(Ctx, Db)
	kelas, _ := details.Kelas().One(Ctx, Db)

	return map[string]interface{} {
		"details": details,
		"dosen": dosen,
		"matkul": matkul,
		"kelas": kelas,
	}
}

// DosenPresenceSessionUpdate update session data
func DosenPresenceSessionUpdate(id int, judul, deskripsi, status string) (int64, error) {
	var query strings.Builder
	query.WriteString(`UPDATE sesi_presensi
	SET ket_judul = ?, ket_deskripsi = ?, status = ?, updated_at = CURRENT_TIMESTAMP`)
	if status == "selesai" {
		query.WriteString(`, ended_at = CURRENT_TIMESTAMP`)
	}
	query.WriteString(` WHERE id = ?`)

	result, err := Db.Exec(query.String(), judul, deskripsi, status, id)
	if err != nil {
		log.Println(err)
		return 0, err	
	}

	return result.RowsAffected()
}

// DosenPresenceSessionDelete deletes one record by ID
func DosenPresenceSessionDelete(id int) (int64, error) {
	session, _ := models.FindSesiPresensi(Ctx, Db, id)
	return session.Delete(Ctx, Db)
}

// DosenPresenceSessionRefreshCode refreshes the code
func DosenPresenceSessionRefreshCode(id int, newCode string) (int64, error) {
	session, _ := models.FindSesiPresensi(Ctx, Db, id)
	session.Kode = newCode
	return session.Update(Ctx, Db, boil.Infer())
}