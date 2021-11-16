package main

type Node struct {
	Id string `json:"id"`
}

type Edge struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type Visited map[string]bool
type RecStack map[string]bool

func IsCyclic(nodes []Node, edges []Edge) bool {
	visited := make(Visited)
	recStack := make(RecStack)

	for _, node := range nodes {
		visited[node.Id] = false
		recStack[node.Id] = false
	}

	for _, node := range nodes {
		if isCyclicUtil(node, visited, recStack, nodes, edges) {
			return true
		}
	}

	return false
}

func isCyclicUtil(node Node, visited Visited, recStack RecStack, nodes []Node, edges []Edge) bool {
	if recStack[node.Id] {
		return true
	}

	if visited[node.Id] {
		return false
	}

	visited[node.Id] = true
	recStack[node.Id] = true

	children := getChildren(node, nodes, edges)

	for _, child := range children {
		if isCyclicUtil(child, visited, recStack, nodes, edges) {
			return true
		}
	}

	recStack[node.Id] = false
	return false
}

func getChildren(node Node, nodes []Node, edges []Edge) []Node {
	children := make([]Node, 0)
	for _, edge := range edges {
		if edge.Source == node.Id {
			child := getNode(edge.Target, nodes)
			if child != nil {
				children = append(children, *child)
			}
		}
	}
	return children
}

func getNode(id string, nodes []Node) *Node {
	for _, node := range nodes {
		if node.Id == id {
			return &node
		}
	}

	return nil
}
