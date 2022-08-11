// Package hclust contains methods for performing agglomerative hierarchical clustering.
package main

import (
	"fmt"
	"rest-api-sijamin/cluster"
	"rest-api-sijamin/dendrogram"
	"rest-api-sijamin/distance"
	"rest-api-sijamin/matrixop"
	"rest-api-sijamin/optimize"
	"rest-api-sijamin/sort"
	"rest-api-sijamin/tree"
	"rest-api-sijamin/typedef"
)

// Cluster references the main cluster method in the cluster subpackage.
// Cluster mereferensikan metode cluster utama dalam sub-paket cluster.
var Cluster = cluster.Cluster

// Dendrogram is an array of SubClusters.
// Dendrogram adalah larik dari SubCluster.
type Dendrogram []SubCluster

// Distance references the main distance method in the distance subpackage.
// Jarak mereferensikan metode jarak utama dalam subpaket jarak.
var Distance = distance.Distance

// GetNodeHeights gets the height for each dendrogram node by summing child branch lengths.
// GetNodeHeights mendapatkan tinggi untuk setiap simpul dendrogram dengan menjumlahkan panjang cabang anak.
var GetNodeHeight = dendrogram.GetNodeHeight

// Optimize references the main leaf optimization method in the optimize subpackage.
// Optimalkan mereferensikan metode optimasi daun utama dalam subpaket optimasi.
var Optimize = optimize.Optimize

// Sort references the main sort method in the sort subpackage
// Urutkan merujuk pada metode pengurutan utama dalam subpaket pengurutan
var Sort = sort.Sort

// SubCluster stores the node, distance and names of leafs for a subcluster.
// SubCluster menyimpan node, jarak dan nama daun untuk subcluster.
type SubCluster = typedef.SubCluster

// TreeLayout contains a tree in newick format and the leaf order.
// TreeLayout berisi pohon dalam format newick dan urutan daun.
type TreeLayout = tree.Tree

// Tree references the main method for generating the newick tree in the tree subpackage.
// Pohon mereferensikan metode utama untuk menghasilkan pohon baru di sub-paket pohon.
var Tree = tree.Create

func main() {
	matrix := [][]float64{
		{2, 4}, {8, 2}, {9, 3}, {1, 5}, {8.5, 1},
	}
	name := []string{"test1", "test2", "test3", "test4", "test5"}
	matrix = matrixop.Transpose(matrix)
	hasiljarak := Distance(matrix, "euclidean", true)

	cluster, _ := Cluster(hasiljarak, "complete")
	fmt.Println("cluster")
	fmt.Println(cluster)
	tree, _ := Tree(cluster, name)
	fmt.Println(tree)
	// fmt.Println(tree.Newick)
	// fmt.Println(tree.Order)
	dendrogram := Dendrogram(cluster)
	// for _, v := range cluster {
	// 	fmt.Println(v.Leafa)

	// 	fmt.Println(v.Leafb)

	// 	fmt.Println(v.Lengtha)
	// 	fmt.Println(v.Lengthb)
	// 	fmt.Println(v.Node)
	// 	fmt.Println(v)
	// 	fmt.Println("======================")

	// }

	// fmt.Println(dendrogram)

	// for _, v := range tree {

	// }

	// fmt.Println(tree)
	// X := [][]float64{
	// 	{0, 0},
	// 	{2, 2},
	// 	{1, 1},
	// 	{2, -1.2},
	// 	{3, 2.2},
	// 	{3.5, 0.5},
	// }

	// Ward's method

}
