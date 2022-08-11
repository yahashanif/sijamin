package tree

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"rest-api-sijamin/typedef"
)

// Tree references a tree in newick format and the leaf order.
type Tree struct {
	Newick string
	Order  []string
}

var n = 0
var c string

// Create generates a newick tree in string format and returns the order
// of the clustering.
func Create(dendrogram []typedef.SubCluster, names []string) (tree Tree, L []Level, err error) {
	// Return if names length does not match matix length.
	if len(names) != len(dendrogram)+1 {
		err = errors.New("The names vector must have the same dimension as the leaf number")
		return
	}

	// Dendrogram clusters/leaf number.
	n := len(dendrogram)

	// Create map of nodes to dendrogram indicies.
	nodeMap := make(map[int]int, n)
	fmt.Println(n)
	for i, cluster := range dendrogram {
		nodeMap[cluster.Node] = i
	}
	fmt.Println("node")
	fmt.Println(dendrogram)

	// Begin with top node, iterate through left and right branches and add to
	// ordering.
	level, _ := Descend(n, 2*n, nodeMap, dendrogram, names)
	tree.Newick = level.Newick
	tree.Order = level.Order

	fmt.Println("level Tree")
	fmt.Println(level.NewickArr)
	return
}
func LevelCluster(dendrogram []typedef.SubCluster, names []string) (level Level, cluster []ClusterData, err error) {
	var clus ClusterData
	// Return if names length does not match matix length.
	if len(names) != len(dendrogram)+1 {
		err = errors.New("The names vector must have the same dimension as the leaf number")
		return
	}

	// Dendrogram clusters/leaf number.
	n := len(dendrogram)

	// Create map of nodes to dendrogram indicies.
	nodeMap := make(map[int]int, n)
	fmt.Println(n)
	for i, cluster := range dendrogram {
		nodeMap[cluster.Node] = i
	}

	// Begin with top node, iterate through left and right branches and add to
	// ordering.

	level, clusterString := Descend(n, 2*n, nodeMap, dendrogram, names)
	// tree.Newick = level.Newick
	// tree.Order = level.Order

	fmt.Println("level Tree")
	fmt.Println(cluster)
	testsplit := strings.Split(clusterString, ",")
	fmt.Println("testsplit")
	fmt.Println(testsplit)
	fmt.Println(len(testsplit))
	for i, s := range testsplit {
		if i == len(testsplit)-1 {

		} else {
			clus.NameCluster = "Cluster " + strconv.Itoa(i+1)
			clus.Cluster = []string{s}
			cluster = append(cluster, clus)

		}
	}
	c = ""

	return
}
