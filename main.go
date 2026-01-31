package main

import (
	"fmt"
	"sort"
)

// topological sort function using Kahn's Algorithm
func topologicalSort(graph map[string][]string) ([]string, bool) {
	indegree := make(map[string]int)

	// 1. initialize indegree of all nodes to 0
	for node := range graph {
		indegree[node] = 0
	}

	// 2. calculate indegree for each node
	for _, neighbors := range graph {
		for _, neighbor := range neighbors {
			indegree[neighbor]++
		}
	}

	// 3. add nodes with indegree 0 to queue
	queue := []string{}
	for node, deg := range indegree {
		if deg == 0 {
			queue = append(queue, node)
		}
	}
	// sort queue for deterministic order
	sort.Strings(queue)

	result := []string{}

	// 4. process queue
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		result = append(result, current)

		for _, neighbor := range graph[current] {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// 5. check for cycle
	if len(result) != len(indegree) {
		return nil, false
	}

	return result, true
}

func main() {
	// course prerequisite graph
	graph := map[string][]string{
		"Basic Mathematics": {"Data Structures"},
		"Data Structures":   {"Algorithms"},
		"Algorithms":        {"AI"},
		"Database":          {"AI"},
		"AI":                {},
	}

	order, valid := topologicalSort(graph)

	if !valid {
		fmt.Println("There is a prerequisite cycle! Invalid order.")
		return
	}

	fmt.Println("Valid course taking order:")
	for i, course := range order {
		fmt.Printf("%d. %s\n", i+1, course)
	}
}