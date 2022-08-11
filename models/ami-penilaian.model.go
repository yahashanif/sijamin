package models

import (
	"fmt"
	"net/http"
	"rest-api-sijamin/db"
	"strconv"
)

type Standar struct {
	PkIdStandar       int    `json:"pk_id_standar"`
	NamaStandar       string `json:"nama_standar"`
	IdKategoriStandar int    `json:"id_kategori_standar"`
	Aktif             string `json:"aktif"`
	IdTipeStandar     int    `json:"id_tipe_standar"`
	KodeStandar       string `json:"kode_standar"`
}

func FetchAllAmiPenilaiaan() (Response, error) {
	var res Response
	var standar Standar
	var arrStandar []Standar

	con := db.CreateCon()

	sqlstatement := "SELECT * FROM `spmi-standar`"

	rows, err := con.Query(sqlstatement)
	if err != nil {
		return res, err
	}
	for rows.Next() {
		err = rows.Scan(&standar.PkIdStandar, &standar.NamaStandar, &standar.IdKategoriStandar, &standar.Aktif, &standar.IdTipeStandar, &standar.KodeStandar)
		if err != nil {
			fmt.Println(err)
			return res, err
		}
		arrStandar = append(arrStandar, standar)
	}
	res.Status = 1
	res.Message = "Success"
	res.Data = arrStandar
	return res, nil
	// SELECT * FROM `spmi-pernyataan_standar`
	// inner join `spmi-standar` on
	// `spmi-pernyataan_standar`.`id_standar` = `spmi-standar`.pk_id_standar
	// inner join `pt-indikator_pernyataan_standar`
	//  on `pt-indikator_pernyataan_standar`.`id_pernyataan_standar` = `spmi-pernyataan_standar`.`pk_id_pernyataan_standar`
	// inner join `ie-indikator` on `ie-indikator`.`pk_id_indikator` = `pt-indikator_pernyataan_standar`.`id_indikator`
	// where  `pt-indikator_pernyataan_standar`.`id_pt_unit` = 1
}

type SPMI_pernyataan_standar struct {
	PkIdPernyataanStandar        int                          `json:"pk_id_pernyataan_standar"`
	UraianPernyataanStandar      string                       `json:"uraian_pernyataan_standar"`
	IdKategoriPernyataanStandar  int                          `json:"id_kategori_pernyataan_standar"`
	IdStandar                    int                          `json:"id_standar"`
	KodeButir                    string                       `json:"kode_butir"`
	PtIndikatorPernyataanStandar PtIndikatorPernyataanStandar `json:"pt_indikator_pernyataan_standar"`
	AmiPenilaianAuditor          AmiPenilaianAuditor          `json:"ami_penilaian_auditor"`
}

type StandarSpmi struct {
	PkIdStandar       int    `json:"pk_id_standar"`
	NamaStandar       string `json:"nama_standar"`
	IdKategoriStandar int    `json:"id_kategori_standar"`
	KodeStandar       string `json:"kode_standar"`
}

type PtIndikatorPernyataanStandar struct {
	PkIdPtIndikatorPernyataanStandar int       `json:"pk_id_pt_indikator_pernyataan_standar"`
	IdIndikator                      int       `json:"id_indikator"`
	IdPernyataanStandar              int       `json:"id_pernyataan_standar"`
	IdPtUnit                         int       `json:"id_pt_unit"`
	Indikator                        Indikator `json:"indikator"`
}

type Indikator struct {
	PkIdIndikator int    `json:"pk_id_indikator"`
	NamaIndikator string `json:"nama_indikator"`
	Bobot         string `json:"bobot"`
}

type ResponseButirStandar struct {
	PkIdStandar int    `json:"pk_id_standar"`
	NamaStandar string `json:"nama_standar"`
	KodeStandar string `json:"kode_standar"`
	IdAuditPs   int    `json:"id_audit_ps"`

	SpmiPernyataanStandar []SPMI_pernyataan_standar `json:"spmi_pernyataan_standar"`
}

// INSERT INTO `ami-penilaian` (`pk_id_penilaian_audit_ps`, `id_audit_ps`, `id_pt_indikator_pernyataan_standar`, `kondisi`, `skor`, `capaian`, `temuan`, `tipe_temuan`, `rekomendasi`, `status`) VALUES (NULL, '1', '2', 'njknjknjkn', '0', '0', 'Dokumen StrukturOrganisasi sudah diarsipkan dalam bentuk digital, tapi belum berupa asrsip fisik di jurusan', '1', 'Sebaiknya juga ada dalam bentuk arsip fisik', '4');

type InputAmiPenilaianAuditor struct {
	IdAuditPs                      int     `json:"id_audit_ps" validate:"required"`
	IdPtIndikatorPernyataanStandar int     `json:"id_pt_indikator_pernyataan_standar" validate:"required"`
	Kondisi                        string  `json:"kondisi" validate:"required"`
	Skor                           float64 `json:"skor" `
	Temuan                         string  `json:"temuan" validate:"required"`
	TipeTemuan                     int     `json:"tipe_temuan" validate:"required"`
	Rekomendasi                    string  `json:"rekomendasi" validate:"required"`
	Status                         int     `json:"status" validate:"required"`
}
type AmiPenilaianAuditor struct {
	PkIdPenilaianAuditPs           int          `json:"pk_id_penilaian_audit_ps"`
	IdAuditPs                      int          `json:"id_audit_ps"`
	IdPtIndikatorPernyataanStandar int          `json:"id_pt_indikator_pernyataan_standar"`
	Kondisi                        string       `json:"kondisi"`
	Skor                           float64      `json:"skor"`
	Capaian                        string       `json:"capaian"`
	Temuan                         string       `json:"temuan"`
	TipeTemuan                     int          `json:"tipe_temuan"`
	Rekomendasi                    string       `json:"rekomendasi"`
	Status                         int          `json:"status"`
	Indikator                      Indikator    `json:"indikator"`
	ProdiAuditor                   ProdiAuditor `json:"prodi_auditor"`
}

type ResponseAmiPenilaianAuditor struct {
	AmiPenilaianAuditor []AmiPenilaianAuditor `json:"ami_penilaian_auditor"`
}

func ButirStandar(IdAuditPs string) (Response, error) {
	var res Response
	var spmi_pernyataan_standar SPMI_pernyataan_standar
	var resbutir ResponseButirStandar
	var arrResButir []ResponseButirStandar
	IdAPs, err := strconv.Atoi(IdAuditPs)
	if err != nil {
		return res, err
	}
	resbutir.IdAuditPs = IdAPs

	con := db.CreateCon()

	sqlstatement := "SELECT `pk_id_standar`,`nama_standar`,`kode_standar` FROM `spmi-standar`"

	rows, err := con.Query(sqlstatement)
	if err != nil {
		return res, err
	}
	for rows.Next() {
		var arrSpmiPernyataanStandar []SPMI_pernyataan_standar
		err = rows.Scan(&resbutir.PkIdStandar, &resbutir.NamaStandar, &resbutir.KodeStandar)

		if err != nil {
			fmt.Println(err)
			return res, err
		}

		sqlstatement2 := "SELECT `pk_id_pernyataan_standar`,`uraian_pernyataan_standar`,`id_kategori_pernyataan_standar`,`id_standar`,kode_butir FROM `spmi-pernyataan_standar` where `id_standar` = " + fmt.Sprintf("%d", resbutir.PkIdStandar)
		fmt.Println(sqlstatement2)
		rows2, err := con.Query(sqlstatement2)
		if err != nil {
			return res, err
		}
		for rows2.Next() {

			err = rows2.Scan(&spmi_pernyataan_standar.PkIdPernyataanStandar, &spmi_pernyataan_standar.UraianPernyataanStandar, &spmi_pernyataan_standar.IdKategoriPernyataanStandar, &spmi_pernyataan_standar.IdStandar, &spmi_pernyataan_standar.KodeButir)
			if err != nil {
				fmt.Println(err)
				return res, err
			}

			sqlstatement3 := "SELECT `pk_id_pt_indikator_pernyataan_standar`,`id_indikator`,`id_pernyataan_standar`,`id_pt_unit` FROM `pt-indikator_pernyataan_standar` where `id_pt_unit`=1 AND `id_pernyataan_standar` = " + fmt.Sprintf("%d", spmi_pernyataan_standar.PkIdPernyataanStandar)
			rows3, err := con.Query(sqlstatement3)
			if err != nil {
				return res, err
			}

			for rows3.Next() {
				var ami_penilaian_auditor AmiPenilaianAuditor
				err = rows3.Scan(&spmi_pernyataan_standar.PtIndikatorPernyataanStandar.PkIdPtIndikatorPernyataanStandar, &spmi_pernyataan_standar.PtIndikatorPernyataanStandar.IdIndikator, &spmi_pernyataan_standar.PtIndikatorPernyataanStandar.IdPernyataanStandar, &spmi_pernyataan_standar.PtIndikatorPernyataanStandar.IdPtUnit)
				if err != nil {
					fmt.Println(err)
					return res, err
				}
				sqlStatementAmi := "SELECT * FROM `ami-penilaian` WHERE `id_audit_ps` = " + IdAuditPs + " AND `id_pt_indikator_pernyataan_standar` =" + strconv.Itoa(spmi_pernyataan_standar.PtIndikatorPernyataanStandar.PkIdPtIndikatorPernyataanStandar)
				fmt.Println(sqlStatementAmi)
				con.QueryRow(sqlStatementAmi).Scan(&ami_penilaian_auditor.PkIdPenilaianAuditPs, &ami_penilaian_auditor.IdAuditPs, &ami_penilaian_auditor.IdPtIndikatorPernyataanStandar, &ami_penilaian_auditor.Kondisi, &ami_penilaian_auditor.Skor, &ami_penilaian_auditor.Capaian, &ami_penilaian_auditor.Temuan, &ami_penilaian_auditor.TipeTemuan, &ami_penilaian_auditor.Rekomendasi, &ami_penilaian_auditor.Status)

				spmi_pernyataan_standar.AmiPenilaianAuditor = ami_penilaian_auditor
				sqlstatement4 := "SELECT `pk_id_indikator`,`nama_indikator`,`bobot` FROM `ie-indikator` where `lam` = '1' and `pk_id_indikator` = " + fmt.Sprintf("%d", spmi_pernyataan_standar.PtIndikatorPernyataanStandar.IdIndikator)
				rows4, err := con.Query(sqlstatement4)
				if err != nil {
					fmt.Println(err)

					return res, err
				}
				for rows4.Next() {
					err = rows4.Scan(&spmi_pernyataan_standar.PtIndikatorPernyataanStandar.Indikator.PkIdIndikator, &spmi_pernyataan_standar.PtIndikatorPernyataanStandar.Indikator.NamaIndikator, &spmi_pernyataan_standar.PtIndikatorPernyataanStandar.Indikator.Bobot)
					if err != nil {
						fmt.Println(err)
						return res, err
					}
				}
			}
			arrSpmiPernyataanStandar = append(arrSpmiPernyataanStandar, spmi_pernyataan_standar)
			fmt.Println(spmi_pernyataan_standar.PtIndikatorPernyataanStandar)

		}
		resbutir.SpmiPernyataanStandar = arrSpmiPernyataanStandar

		arrResButir = append(arrResButir, resbutir)
	}
	res.Status = http.StatusOK
	res.Data = arrResButir
	res.Message = "SUKSES"
	return res, nil
}

/*
* Input Data Penilaian
 */

func InputAmiPenilaian(input *InputAmiPenilaianAuditor) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlstatement := "INSERT INTO `ami-penilaian` (`pk_id_penilaian_audit_ps`, `id_audit_ps`, `id_pt_indikator_pernyataan_standar`, `kondisi`, `skor`, `capaian`, `temuan`, `tipe_temuan`, `rekomendasi`, `status`) VALUES (NULL, ?, ?, ?, ?, '0', ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlstatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(input.IdAuditPs, input.IdPtIndikatorPernyataanStandar, input.Kondisi, input.Skor, input.Temuan, input.TipeTemuan, input.Rekomendasi, input.Status)

	if err != nil {
		return res, err
	}
	strconv.Atoi("2")

	LastInsertId, err := result.LastInsertId()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Berhasil Input Penilaian"
	res.Data = LastInsertId

	return res, nil

}
func EditAmiPenilaian(kondisi, temuan, rekomendasi string, tipe_temuan, status, PkIdPenilaianAuditPs int, skor float64) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlstatement := "UPDATE `ami-penilaian` SET `kondisi` = ?, `skor` = ?, `temuan` = ?, `tipe_temuan` = ?, `rekomendasi` = ?, `status` = ? WHERE `ami-penilaian`.`pk_id_penilaian_audit_ps` = ?"

	stmt, err := con.Prepare(sqlstatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(kondisi, skor, temuan, tipe_temuan, rekomendasi, status, PkIdPenilaianAuditPs)

	if err != nil {
		return res, err
	}
	strconv.Atoi("2")

	LastInsertId, err := result.LastInsertId()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Berhasil Edit Penilaian"
	res.Data = LastInsertId

	return res, nil

}

func AmiPenilaian(IdAuditor int) (Response, error) {
	var res Response
	var ami_penilaian_auditor AmiPenilaianAuditor
	var arrAmi []AmiPenilaianAuditor
	var indikator Indikator
	var prodiAuditor ProdiAuditor

	con := db.CreateCon()
	sqlstatement := "SELECT * FROM `ami-penilaian`"
	rows, err := con.Query(sqlstatement)
	if err != nil {
		return res, err
	}
	for rows.Next() {
		err = rows.Scan(&ami_penilaian_auditor.PkIdPenilaianAuditPs, &ami_penilaian_auditor.IdAuditPs, &ami_penilaian_auditor.IdPtIndikatorPernyataanStandar, &ami_penilaian_auditor.Kondisi, &ami_penilaian_auditor.Skor, &ami_penilaian_auditor.Capaian, &ami_penilaian_auditor.Temuan, &ami_penilaian_auditor.TipeTemuan, &ami_penilaian_auditor.Rekomendasi, &ami_penilaian_auditor.Status)
		if err != nil {
			fmt.Println(err)
			return res, err
		}
		sqlstatement2 := "SELECT `pk_id_indikator`,`nama_indikator`,`bobot` FROM `ie-indikator` where `pk_id_indikator` = " + fmt.Sprintf("%d", ami_penilaian_auditor.IdPtIndikatorPernyataanStandar)

		con.QueryRow(sqlstatement2).Scan(&indikator.PkIdIndikator, &indikator.NamaIndikator, &indikator.Bobot)

		ami_penilaian_auditor.Indikator = indikator

		sqlstatement3 := "select `ami-audit_ps`.`pk_id_audit_ps`, `ami-audit_ps`.`id_audit`, `ami-audit_ps`.`id_pt_unit`,`ami-audit_ps`.`tgl_audit`, `ami-audit_ps`.`selesai`, `ami-audit_ps`.`kesimpulan`,`str-pt_unit`.`kode_pt_unit`,  `str-pt_unit`.`nama_pt_unit`, `ami-audit`.`nama_audit`,`set-periode`.`pk_id_periode`,`ami-auditor_audit_ps`.`id_auditor` from `ami-audit_ps` inner join `str-pt_unit` on `ami-audit_ps`.`id_pt_unit` = `str-pt_unit`.`pk_id_pt_unit` inner join `ami-audit` on `ami-audit_ps`.`id_audit` = `ami-audit`.`pk_id_audit` inner join `set-periode` on `ami-audit`.`id_periode` = `set-periode`.`pk_id_periode` inner join `ami-auditor_audit_ps` on `ami-audit_ps`.`pk_id_audit_ps` = `ami-auditor_audit_ps`.`id_audit_ps`where  `ami-audit_ps`.`pk_id_audit_ps` = " + fmt.Sprintf("%d", ami_penilaian_auditor.IdAuditPs) + " AND `ami-auditor_audit_ps`.`id_auditor` = " + fmt.Sprintf("%d", IdAuditor)

		con.QueryRow(sqlstatement3).Scan(&prodiAuditor.PKIdAuditPs, &prodiAuditor.IdAudit, &prodiAuditor.IdPtUnit, &prodiAuditor.TglAudit, &prodiAuditor.Selesai, &prodiAuditor.Kesimpulan, &prodiAuditor.KodePtUnit, &prodiAuditor.NamaPtUnit, &prodiAuditor.NamaAudit, &prodiAuditor.PkIdPeriode, &prodiAuditor.IdAuditor)

		ami_penilaian_auditor.ProdiAuditor = prodiAuditor

		arrAmi = append(arrAmi, ami_penilaian_auditor)

	}
	res.Status = http.StatusOK
	res.Message = "OK"
	res.Data = arrAmi

	return res, nil

}

func CekAmiPenilaian(id_audit_ps, id_pt_indikator_pernyataan_standar string) (Response, error) {
	var res Response
	var ami_penilaian_auditor AmiPenilaianAuditor

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM `ami-penilaian` WHERE `id_audit_ps` = " + id_audit_ps + " AND `id_pt_indikator_pernyataan_standar` =" + id_pt_indikator_pernyataan_standar

	err := con.QueryRow(sqlStatement).Scan(&ami_penilaian_auditor.PkIdPenilaianAuditPs, &ami_penilaian_auditor.IdAuditPs, &ami_penilaian_auditor.IdPtIndikatorPernyataanStandar, &ami_penilaian_auditor.Kondisi, &ami_penilaian_auditor.Skor, &ami_penilaian_auditor.Capaian, &ami_penilaian_auditor.Temuan, &ami_penilaian_auditor.TipeTemuan, &ami_penilaian_auditor.Rekomendasi, &ami_penilaian_auditor.Status)

	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "Belum Ada Penilaian"
	} else {
		res.Status = http.StatusOK
		res.Message = "Ada Data"
		res.Data = ami_penilaian_auditor
	}
	return res, nil
}
