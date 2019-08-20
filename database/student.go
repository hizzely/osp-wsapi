package database

import (
	"log"
	"strings"

	"github.com/hizzely/osp-wsapi/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// StudentPresenceHistoryModel for binding joined table
type StudentPresenceHistoryModel struct {
	models.PresensiMH `boil:",bind"`
	models.SesiPresensi `boil:",bind"`
	models.Dosen `boil:",bind"`
	models.MataKuliah `boil:",bind"`
	models.KelasMH `boil:",bind"`
}

// Student get single student by id
func Student(id string) (*models.Mahasiswa, error) {
	return models.FindMahasiswa(Ctx, Db, id)
}

// StudentCreate a new student account
func StudentCreate(id, namaLengkap, password, status string, kelasID int) error {
	student := &models.Mahasiswa {
		ID: id,
		NamaLengkap: namaLengkap,
		Password: password,
		KelasID: kelasID,
		Status: status,
		TahunMasuk: 2019,
	}
	return student.Insert(Ctx, Db, boil.Infer())
}

// StudentDelete single student account
func StudentDelete(id string) (int64, error) {
	student, _ := models.FindMahasiswa(Ctx, Db, id)
	return student.Delete(Ctx, Db)
}

// StudentPresenceHistory get all student presence history
func StudentPresenceHistory(id string, matkulID int) ([]StudentPresenceHistoryModel, error) {
	var model []StudentPresenceHistoryModel

	var query strings.Builder
	query.WriteString(`SELECT * FROM presensi_mhs
	LEFT JOIN sesi_presensi ON sesi_presensi.id = presensi_mhs.session_id
	LEFT JOIN dosen ON dosen.id = sesi_presensi.dosen_id
	LEFT JOIN mata_kuliah ON mata_kuliah.id = sesi_presensi.matkul_id
	LEFT JOIN kelas_mhs ON kelas_mhs.id = sesi_presensi.kelas_id
	WHERE presensi_mhs.mahasiswa_id = ?`)
	
	var err error
	if matkulID != 0 {
		query.WriteString(` AND sesi_presensi.matkul_id = ?`)
		query.WriteString(` ORDER BY presensi_mhs.created_at DESC`)
		err = queries.Raw(query.String(), id, matkulID).Bind(Ctx, Db, &model)
	}
	query.WriteString(` ORDER BY presensi_mhs.created_at DESC`)
	err = queries.Raw(query.String(), id).Bind(Ctx, Db, &model)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return model, nil
}

// StudentPresenceByCode using code for presence
func StudentPresenceByCode(npm, sessionCode string) int {
	session, err := models.SesiPresensis(qm.Where("kode = ?", sessionCode)).One(Ctx, Db)
	if err != nil || session == nil{
		log.Println(err)
		return 0
	}

	// Check for existing presence data for this session
	presenceExist, err := models.PresensiMHS(qm.Where("mahasiswa_id = ? AND session_id = ?", npm, session.ID)).Exists(Ctx, Db)
	if presenceExist {
		return 2
	}

	// If not, create new presence data
	presence := &models.PresensiMH{
		MahasiswaID: npm,
		SessionID: session.ID,
	}
	insertErr := presence.Insert(Ctx, Db, boil.Infer())
	if insertErr != nil {
		log.Println(insertErr)
		return 0
	}
	return 1
}

// StudentPresenceByRfid using rfid code for presence
func StudentPresenceByRfid(rfidCode, sessionCode string) int {
	session, err := models.SesiPresensis(qm.Where("kode = ?", sessionCode)).One(Ctx, Db)
	if err != nil || session == nil{
		log.Println(err)
		return 0
	}

	student, err := models.Mahasiswas(qm.Where("kode_rfid = ?", rfidCode)).One(Ctx, Db)
	if err != nil || student == nil{
		log.Println(err)
		return 0
	}
	
	// Check for existing presence data for this session
	presenceExist, err := models.PresensiMHS(qm.Where("mahasiswa_id = ? AND session_id = ?", student.ID, session.ID)).Exists(Ctx, Db)
	if presenceExist {
		return 2
	}

	// If not, create new presence data
	presence := &models.PresensiMH{
		MahasiswaID: student.ID,
		SessionID: session.ID,
	}
	insertErr := presence.Insert(Ctx, Db, boil.Infer())
	if insertErr != nil {
		log.Println(insertErr)
		return 0
	}
	return 1
}