package routes

import (
	"rest-api-sijamin/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	// e.GET("/butir-nilai", controllers.FetchAllAmiPenilaiaan)
	e.GET("/butir-nilai/:id_audit_ps", controllers.ButirStandar)
	e.POST("/Login", controllers.LoginAuditor)
	e.POST("/penilaian", controllers.InputAmiPenilaian)
	e.POST("/Editpenilaian", controllers.EditAmiPenilaian)
	e.POST("/ProdiYangAuditor", controllers.ProdiYangAuditor)
	e.GET("/ami-penilaian/:id", controllers.AmiPenilaian)
	e.GET("/testclustering/:id_pt_unit/:tipe", controllers.ClusteringTest)
	e.POST("/cekAmiPenilaian", controllers.CekAmiPenilaian)
	e.GET("/cekDataStandar/:pk_id_pt_indikator_pernyataan_standar", controllers.CekDataStandar)
	return e
}
