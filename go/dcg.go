package main

type Node struct {
	Id string `json:"id"`
}

type Edge struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

// For fast lookup children will point to the index of the node in the array
type NodeWithChildren struct {
	Node
	Children []int
}

type Visited map[int]bool
type RecStack map[int]bool

func IsCyclic(iNodes []Node, iEdges []Edge) bool {
	nodes := make([]NodeWithChildren, len(iNodes))

	// Complexity of this is O(n^3)..?
	for idx, iNode := range iNodes {
		node := NodeWithChildren{Node: iNode}

		for _, edge := range iEdges {
			if edge.Source == node.Id {
				node.Children = append(node.Children, getIndexForNodeWithId(edge.Target, iNodes))
			}
		}

		nodes[idx] = node
	}

	visited := make(Visited)
	recStack := make(RecStack)

	for i := range nodes {
		visited[i] = false
		recStack[i] = false
	}

	for i := range nodes {
		if isCyclicUtil(i, visited, recStack, nodes) {
			return true
		}
	}

	return false
}

func getIndexForNodeWithId(id string, nodes []Node) int {
	for idx, node := range nodes {
		if node.Id == id {
			return idx
		}
	}

	return -1
}

func isCyclicUtil(idx int, visited Visited, recStack RecStack, nodes []NodeWithChildren) bool {
	if recStack[idx] {
		return true
	}

	if visited[idx] {
		return false
	}

	visited[idx] = true
	recStack[idx] = true

	children := nodes[idx].Children

	for _, child := range children {
		if isCyclicUtil(child, visited, recStack, nodes) {
			return true
		}
	}

	recStack[idx] = false
	return false
}
