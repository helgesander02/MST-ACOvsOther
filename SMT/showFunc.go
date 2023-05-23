package SMT

import (
	"fmt"
)

func (graphs *Graphs) Show_Graph() {
    for _, g := range graphs.Graphs {
        fmt.Printf("%d\t", g.ID)
        fmt.Printf("%v\t", g.Visited)
        fmt.Printf("%d\t", g.Cost)
        fmt.Printf("%d\n", g.Path)
		//for _, e := range g.Edges {
			//fmt.Printf("%d-->%d cost:%d\n", e.Strat, e.End, e.Cost)
		//}
    } 
}

func (trees *Trees) Show_Trees() {
    tsn := 1
    for _,tree := range trees.TSNTrees{  
        fmt.Printf("TSN Tree %d \n", tsn)
        tree.Show_Tree()
        tsn++
    }
    avb := 1
    for _,tree := range trees.AVBTrees{
        fmt.Printf("AVB Tree %d \n", avb)
        tree.Show_Tree()
        avb++
    }
}

func (tree *Tree) Show_Tree() {
    for _, node := range tree.Nodes{
        for _, c := range node.Connections {
            fmt.Printf("%d --> %d \n", c.FromNodeID, c.ToNodeID)
        }      
    }
}