package models

import (
	"fmt"
	"net/http"
	"rest-api-sijamin/db"
)

type User struct {
	Username    string `json:"username"`
	IdPegawai   int    `json:"id_pegawai"`
	NamaPegawai string `json:"nama_pegawai"`
	IdGroup     int    `json:"id_group"`
	IdPtUnit    int    `json:"id_pt_unit"`
	TglLahir    string `json:"tgl_lahir"`
	NIP         string `json:"nip"`
	IdAuditor   int    `json:"id_auditor"`
	Aktif       string `json:"aktif"`
}

type ProdiAuditor struct {
	PKIdAuditPs int    `json:"pk_id_audit_ps"`
	IdAudit     int    `json:"id_audit"`
	IdPtUnit    int    `json:"id_pt_unit"`
	TglAudit    string `json:"tgl_audit"`
	Selesai     string `json:"selesai"`
	Kesimpulan  string `json:"kesimpulan"`
	KodePtUnit  string `json:"kode_pt_unit"`
	NamaPtUnit  string `json:"nama_pt_unit"`
	NamaAudit   string `json:"nama_audit"`
	PkIdPeriode int    `json:"pk_id_periode"`
	IdAuditor   int    `json:"id_auditor"`
}

func LoginAuditor(username, password string) (Response, error) {
	var res Response
	var user User

	con := db.CreateCon()

	sqlstatement := "SELECT username,id_pegawai,id_grup,id_pt_unit FROM `usr-pengguna` WHERE username = ? AND PASSWORD = MD5(?) AND id_grup = 41"

	con.QueryRow(sqlstatement, username, password).Scan(&user.Username, &user.IdPegawai, &user.IdGroup, &user.IdPtUnit)
	if user.Username == "" {
		res.Status = 0
		res.Message = "Username tidak ditemukan"
		return res, nil
	}

	sqlstatement2 := "SELECT nama_pegawai, tanggal_lahir,NIP,aktif FROM `str-pegawai` WHERE `pk_id_pegawai` = ?"

	con.QueryRow(sqlstatement2, user.IdPegawai).Scan(&user.NamaPegawai, &user.TglLahir, &user.NIP, &user.Aktif)

	sqlstatement3 := "SELECT pk_id_auditor FROM `ami-auditor` where `id_pegawai` = ?"

	con.QueryRow(sqlstatement3, user.IdPegawai).Scan(&user.IdAuditor)

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = user

	return res, nil

}

func ProdiYangAuditor(IdAuditor int) (Response, error) {
	var res Response
	var prodiAuditor ProdiAuditor
	var prodi []ProdiAuditor
	var nameAudit string
	var cek string
	con := db.CreateCon()

	sqlstatementCek := "SELECT count(*) as cek FROM `set-periode` WHERE aktif = 'Y'"

	con.QueryRow(sqlstatementCek).Scan(&cek)

	if cek == "0" {
		res.Status = http.StatusNoContent
		res.Message = "Tidak ada periode aktif"

		return res, nil
	}

	sqlstatement := "select `ami-audit_ps`.`pk_id_audit_ps`, `ami-audit_ps`.`id_audit`, `ami-audit_ps`.`id_pt_unit`,`ami-audit_ps`.`tgl_audit`, `ami-audit_ps`.`selesai`, `ami-audit_ps`.`kesimpulan`,`str-pt_unit`.`kode_pt_unit`,  `str-pt_unit`.`nama_pt_unit`, `ami-audit`.`nama_audit`,`set-periode`.`pk_id_periode`,`ami-auditor_audit_ps`.`id_auditor` from `ami-audit_ps` inner join `str-pt_unit` on `ami-audit_ps`.`id_pt_unit` = `str-pt_unit`.`pk_id_pt_unit` inner join `ami-audit` on `ami-audit_ps`.`id_audit` = `ami-audit`.`pk_id_audit` inner join `set-periode` on `ami-audit`.`id_periode` = `set-periode`.`pk_id_periode` inner join `ami-auditor_audit_ps` on `ami-audit_ps`.`pk_id_audit_ps` = `ami-auditor_audit_ps`.`id_audit_ps`where  `set-periode`.`aktif` = 'Y' AND `ami-auditor_audit_ps`.`id_auditor`=" + fmt.Sprintf("%d", IdAuditor)

	rows, err := con.Query(sqlstatement)
	if err != nil {
		return res, err
	}
	for rows.Next() {
		err := rows.Scan(&prodiAuditor.PKIdAuditPs, &prodiAuditor.IdAudit, &prodiAuditor.IdPtUnit, &prodiAuditor.TglAudit, &prodiAuditor.Selesai, &prodiAuditor.Kesimpulan, &prodiAuditor.KodePtUnit, &prodiAuditor.NamaPtUnit, &prodiAuditor.NamaAudit, &prodiAuditor.PkIdPeriode, &prodiAuditor.IdAuditor)
		if err != nil {
			return res, err
		}
		nameAudit = prodiAuditor.NamaAudit
		prodi = append(prodi, prodiAuditor)
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]interface{}{
		"nama_audit": nameAudit,
		"prodi":      prodi,
	}

	return res, nil
}
