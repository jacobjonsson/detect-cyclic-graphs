export type Node = { id: string };
export type Edge = { source: string; target: string };
export type NodeOrEdge = Node | Edge;

type Visited = Record<string, boolean>;
type RecStack = Record<string, boolean>;

export function isCyclic(nodes: Node[], edges: Edge[]): boolean {
  let visited: Visited = nodes.reduce(
    (acc, node) => ({ ...acc, [node.id]: false }),
    {} as Visited
  );
  let recStack: RecStack = nodes.reduce(
    (acc, node) => ({ ...acc, [node.id]: false }),
    {} as RecStack
  );

  for (const node of nodes) {
    if (isCyclicUtil(node, visited, recStack, nodes, edges)) {
      return true;
    }
  }

  return false;
}

function isCyclicUtil(
  node: Node,
  visited: Visited,
  recStack: RecStack,
  nodes: Node[],
  edges: Edge[]
): boolean {
  if (recStack[node.id]) {
    return true;
  }

  if (visited[node.id]) {
    return false;
  }

  visited[node.id] = true;
  recStack[node.id] = true;

  let nodeEdges = edges
    .filter((edge) => edge.source === node.id)
    .map((edge) => edge.target);
  let children = nodes.filter((n) => nodeEdges.includes(n.id));

  for (const node of children) {
    if (isCyclicUtil(node, visited, recStack, nodes, edges)) {
      return true;
    }
  }

  recStack[node.id] = false;

  return false;
}
