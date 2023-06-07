package SMT

import (   
    "fmt"
    "encoding/json"
    
    "src/topology"
    "src/stream"
)

func SteninerTree(topology *topology.Topology, flows *stream.Flows, cost float64) *Trees {
    trees := &Trees{}
    
    for _, flow := range flows.TSNFlows {
        tree := GetTree(topology, flow, cost)
        trees.TSNTrees = append(trees.TSNTrees, tree)
    }
    fmt.Println("Finish TSN all SteninerTree")

    for _, flow := range flows.AVBFlows {
        tree := GetTree(topology, flow, cost)
        trees.AVBTrees = append(trees.AVBTrees, tree)
    }
    fmt.Println("Finish AVB all SteninerTree")

    return trees
}

func GetTree(topology *topology.Topology, flow *stream.Flow, cost float64) *Tree {
    t := DeepCopy(topology) //Duplicate of Topology
    t.AddT2S(flow.Source, cost)
    t.AddS2L(flow.Destinations, cost)
    //t.Show_Topology()

    tree := &Tree{}
    for _, terminal := range flow.Destinations {  
        graphs := AddGarph(t, terminal+2000)
        graphs = Dijkstra(graphs, flow.Source+1000, terminal+2000)
        //fmt.Printf("%d --> %d Shortest Path:\n", flow.Source, terminal)
        //graphs.Show_Graph()

        tree.AddTree(graphs, flow.Source+1000, terminal+2000)
    }

    return tree
}

func (tree *Tree) AddTree(graphs *Graphs, strat int, terminal int) {
    path := terminal
    n := &Node{ID: terminal}
    tree.Nodes = append(tree.Nodes, n)

    for path != strat {
        graph := graphs.findGraph(path)
        node, b := tree.InTree(graph.Path)
              
        connection := &Connection{
            FromNodeID: graph.Path,
            ToNodeID: graph.ID,
            Cost: 750000.,
        }

        if !(node.InNode(connection)) {
            node.Connections = append(node.Connections, connection)
        }
        
        if !(b){
            tree.Nodes = append(tree.Nodes, node)
        }
        path = graph.Path
    }
}

func (tree *Tree) InTree(id int) (*Node, bool) {
    for _, node := range tree.Nodes{
        if node.ID == id  {
            return node, true
        }
    }
    return &Node{ID: id}, false
}

func (node *Node) InNode(connection *Connection) bool {
    for _, c := range node.Connections {
        if c.ToNodeID == connection.ToNodeID {
            return true
        }
    }
    return false
}

func AddGarph(topology *topology.Topology, terminal int) *Graphs {
    graphs := &Graphs{}
    //Talker
    for _, t := range topology.Talker {
        gt := &Graph{}
        gt.ID = t.ID
        gt.AddEdge(t.Connections)
        graphs.Graphs = append(graphs.Graphs, gt)
    }

    //Switch
    for _, s := range topology.Switch {
        gs := &Graph{}
        gs.ID = s.ID
        gs.AddEdge(s.Connections)
        graphs.Graphs = append(graphs.Graphs, gs)
    }

    //Listener 
    gl := &Graph{}
    gl.ID = terminal
    graphs.Graphs = append(graphs.Graphs, gl)

    return graphs
}

func (graph *Graph) AddEdge(connections []*topology.Connection){
    for _, c := range connections {
        edge := &Edge{
            Strat: c.FromNodeID,
            End: c.ToNodeID,
            Cost: 1,
        }
        graph.Edges = append(graph.Edges, edge)  
    }
}

func (graphs *Graphs) findGraph(id int) *Graph {
    for _, graph := range graphs.Graphs {
        if graph.ID == id {
            return graph
        }
    }
    return &Graph{}
}

func DeepCopy(t1 *topology.Topology) (*topology.Topology) {
    if buf, err := json.Marshal(t1); err != nil {
        return nil
    } else {
        t2 := &topology.Topology{}
        if err = json.Unmarshal(buf, t2); err != nil {
            return nil
        }
        return t2
    }
}
