package models

import (
	"fmt"
	"net/http"
	"rest-api-sijamin/cluster"
	"rest-api-sijamin/db"
	"rest-api-sijamin/distance"
	"rest-api-sijamin/matrixop"
	"rest-api-sijamin/tree"
	"rest-api-sijamin/typedef"
)

type DataCek struct {
	NamaIndikator           string `json:"nama_indikator"`
	UraianPernyataanStandar string `json:"uraian_pernyataan_standar"`
	KodeStandar             string `json:"kode_standar"`
	KodeButir               string `json:"kode_butir"`
}

var Distance = distance.Distance
var Cluster = cluster.Cluster
var Tree = tree.LevelCluster

func ClusteringAMI(id_pt_unit, tipe string) (Response, error) {
	type SubCluster = typedef.SubCluster
	type Dendrogram []SubCluster
	var res Response
	var idPeriode string
	var mat [][]float64
	con := db.CreateCon()
	var sqlStatement string
	if tipe == "prodi" {
		sqlStatement = "SELECT pk_id_audit_ps FROM `ami-audit_ps` inner join `ami-audit` on `ami-audit_ps`.`id_audit` = `ami-audit`.`pk_id_audit`  WHERE `ami-audit_ps`.id_pt_unit = " + string(id_pt_unit) + " group by `ami-audit`.`id_periode`"
	} else {
		var idbantu string

		sqlbantu := "SELECT pk_id_audit FROM `ami-audit`ORDER BY pk_id_audit DESC limit 1"
		err := con.QueryRow(sqlbantu).Scan(&idbantu)
		if err != nil {
			return res, err
		}
		sqlStatement = "SELECT pk_id_audit_ps FROM `ami-audit_ps` where id_audit = " + idbantu
	}
	fmt.Println(sqlStatement)
	row, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}
	for row.Next() {
		err = row.Scan(&idPeriode)
		if err != nil {
			return res, err
		}
		sqlStatementSkor := "Select skor from `ami-penilaian` where id_audit_ps = '" + idPeriode + "' order by id_pt_indikator_pernyataan_standar "
		skor, err := con.Query(sqlStatementSkor)
		if err != nil {
			return res, err
		}
		var skorList []float64
		for skor.Next() {
			var skordata float64
			err = skor.Scan(&skordata)
			if err != nil {
				return res, err
			}
			skorList = append(skorList, skordata)
		}
		mat = append(mat, skorList)
		fmt.Print(skorList)
		fmt.Print(idPeriode)

	}

	sqlStatementPtIDIndikator := "SELECT id_pt_indikator_pernyataan_standar FROM `ami-penilaian` group by id_pt_indikator_pernyataan_standar order by id_pt_indikator_pernyataan_standar "
	ptIDIndikator, err := con.Query(sqlStatementPtIDIndikator)
	if err != nil {
		return res, err
	}
	var ptIDIndikatorList []string
	for ptIDIndikator.Next() {
		var ptIDIndikatorData string
		err = ptIDIndikator.Scan(&ptIDIndikatorData)
		if err != nil {
			return res, err
		}
		ptIDIndikatorList = append(ptIDIndikatorList, ptIDIndikatorData)
	}

	matrix := [][]float64{
		// {2, 4}, {8, 2}, {9, 3}, {1, 5}, {8.5, 1}, {2, 4}, {8, 2}, {9, 3}, {1, 5}, {8.5, 1}, {2, 4}, {8, 2}, {9, 3}, {1, 5}, {8.5, 1},
		{2, 4}, {8, 2}, {9, 3}, {1, 5}, {8.5, 1},
	}

	// matrik := matrix
	// name := []string{"test1", "test2", "test3", "test4", "test5", "test6", "test7", "test8", "test10", "test11", "test12", "test113", "test14", "test15", "test16"}
	// name := []string{"test1", "test2", "test3", "test4", "test5"}
	matrix = matrixop.Transpose(matrix)
	fmt.Println(matrix)
	hasiljarak := Distance(mat, "euclidean", true)

	cluster, _ := Cluster(hasiljarak, "complete")
	fmt.Println("cluster")
	fmt.Println(cluster)
	tree, c, _ := Tree(cluster, ptIDIndikatorList)
	fmt.Println(tree)
	dendrogram := Dendrogram(cluster)
	fmt.Println("dendrogram")
	fmt.Println(dendrogram)
	fmt.Println(c)
	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = map[string]interface{}{
		"cluster":     c,
		"matrik":      mat,
		"Hasil Jarak": hasiljarak,
		"tree":        tree,
	}
	for _, v := range cluster {
		fmt.Println(v)
	}

	return res, nil
}

func CekDataStandar(PkIdPtIndikatorPernyataanStandar string) (Response, error) {
	var res Response
	var data DataCek
	con := db.CreateCon()
	sqlStatement := "select `ie-indikator`.nama_indikator, `spmi-pernyataan_standar`.uraian_pernyataan_standar, `spmi-standar`.`kode_standar`,`spmi-pernyataan_standar`.`kode_butir` from `pt-indikator_pernyataan_standar` inner join `ie-indikator` on `pt-indikator_pernyataan_standar`.`id_indikator` = `ie-indikator`.`pk_id_indikator` inner join `spmi-pernyataan_standar` on `pt-indikator_pernyataan_standar`.id_pernyataan_standar = `spmi-pernyataan_standar`.pk_id_pernyataan_standar INNER JOIN `spmi-standar` on `spmi-pernyataan_standar`.`id_standar` = `spmi-standar`.`pk_id_standar` where `pt-indikator_pernyataan_standar`.`pk_id_pt_indikator_pernyataan_standar` = " + PkIdPtIndikatorPernyataanStandar
	con.QueryRow(sqlStatement).Scan(&data.NamaIndikator, &data.UraianPernyataanStandar, &data.KodeStandar, &data.KodeButir)
	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = data
	return res, nil

}
