package controllers

import (
	"fmt"
	"net/http"
	"rest-api-sijamin/models"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate = validator.New()

func FetchAllAmiPenilaiaan(c echo.Context) error {
	result, err := models.FetchAllAmiPenilaiaan()
	if err != nil {
		return err
	}
	return c.JSON(200, result)
}

func ButirStandar(c echo.Context) error {
	IdAuditPs := c.Param("id_audit_ps")
	result, err := models.ButirStandar(IdAuditPs)
	if err != nil {
		return err
	}
	return c.JSON(200, result)
}

func InputAmiPenilaian(c echo.Context) error {
	var res models.Response
	input := &models.InputAmiPenilaianAuditor{}
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if err := validate.Struct(input); err != nil {
		fmt.Println(err)
		res.Status = http.StatusBadRequest
		res.Message = err.Error()
		return c.JSON(http.StatusInternalServerError, res)
	}

	result, err := models.InputAmiPenilaian(input)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())

	}
	return c.JSON(http.StatusOK, result)
}
func EditAmiPenilaian(c echo.Context) error {

	kondisi := c.FormValue("kondisi")
	temuan := c.FormValue("temuan")
	rekomendasi := c.FormValue("rekomendasi")
	tipeTemuan, err := strconv.Atoi(c.FormValue("tipe_temuan"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	status, err := strconv.Atoi(c.FormValue("status"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	PkIdPenilaianAuditPs, err := strconv.Atoi(c.FormValue("pk_id_penilaian_audit_ps"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	skor, err := strconv.ParseFloat(c.FormValue("skor"), 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	result, err := models.EditAmiPenilaian(kondisi, temuan, rekomendasi, tipeTemuan, status, PkIdPenilaianAuditPs, skor)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())

	}
	return c.JSON(http.StatusOK, result)
}

func AmiPenilaian(c echo.Context) error {
	IdAuditor := c.Param("id")
	idInt, _ := strconv.Atoi(IdAuditor)
	result, err := models.AmiPenilaian(idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())

	}
	return c.JSON(http.StatusOK, result)
}

func ClusteringTest(c echo.Context) error {
	id_pt_unit := c.Param("id_pt_unit")
	tipe := c.Param("tipe")
	result, err := models.ClusteringAMI(id_pt_unit, tipe)

	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func CekAmiPenilaian(c echo.Context) error {
	id_audit_ps := c.FormValue("id_audit_ps")
	id_pt_indikator_pernyataan_standar := c.FormValue("id_pt_indikator_pernyataan_standar")

	result, err := models.CekAmiPenilaian(id_audit_ps, id_pt_indikator_pernyataan_standar)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}

func CekDataStandar(c echo.Context) error {
	PkIdPtIndikatorPernyataanStandar := c.Param("pk_id_pt_indikator_pernyataan_standar")
	result, err := models.CekDataStandar(PkIdPtIndikatorPernyataanStandar)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}
