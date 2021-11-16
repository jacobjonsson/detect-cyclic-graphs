import type { Node, Edge } from "./index";
import { isCyclic } from "./index";
import testData from "../test-data.json";

test.each(testData)("isCyclic(%p)", (data) => {
  const nodes: Node[] = data.nodes;
  const edges: Edge[] = data.edges;

  expect(isCyclic(nodes, edges)).toBe(data.cyclic);
});
