package main

import (
	"bufio"
	"fmt"
	"io"
	"maps"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Maximum number of passengers: %d", solve("./input.txt"))
}

func solve(inputPath string) int {
	f, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()

	hub, goals, nodes, quotas := getNodes(getLines(f))

	fmt.Printf("hub: %s, goals: %v\nnodes: %v\n", hub, goals, nodes)
	for from, children := range quotas {
		for to, q := range children {
			fmt.Printf("%4s |- %2.d -> %s\n", from, q, to)
		}
	}

	g := make(Graph)
	for from, children := range quotas {
		for to, q := range children {
			g.AddEdge(from, to, q)
		}
	}

	// add final node super and connect all goals to it
	for _, n := range goals {
		g.AddEdge(n, "SUPER", math.MaxInt)
	}

	return EdmondsKarp(g, hub, "SUPER")
}

type Graph map[string][]*Edge

func (g Graph) AddEdge(from, to string, cap int) {
	forward := &Edge{to, cap, nil}
	back := &Edge{from, 0, nil}
	forward.Rev = back
	back.Rev = forward

	g[from] = append(g[from], forward)
	g[to] = append(g[to], back)
}

func EdmondsKarp(g Graph, hub, super string) int {
	res := 0
	for {
		prev := make(map[string]*Edge)
		prev[hub] = nil // avoid cycle

		q := []string{hub}
		for len(q) > 0 {
			node := q[0]
			q = q[1:]
			if node == super {
				break
			}
			for _, e := range g[node] {
				if _, ok := prev[e.To]; !ok && e.Cap > 0 {
					prev[e.To] = e
					q = append(q, e.To)
				}
			}
		}
		if _, ok := prev[super]; !ok {
			break
		}

		// find capacity of the path
		minCap := math.MaxInt
		for e := prev[super]; e != nil; e = prev[e.Rev.To] {
			minCap = min(minCap, e.Cap)
		}
		for e := prev[super]; e != nil; e = prev[e.Rev.To] {
			e.Cap -= minCap
			e.Rev.Cap += minCap
		}

		res += minCap
	}

	return res
}

type Edge struct {
	To  string
	Cap int
	Rev *Edge
}

func getLines(f io.ReadCloser) []string {
	scan := bufio.NewScanner(f)
	lines := make([]string, 0, 64)
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}
	return lines
}

func getNodes(raw []string) (
	hub string,
	goals []string,
	nodes []string,
	quotas map[string]map[string]int,
) {
	nodesMap := make(map[string]struct{})
	goalsMap := make(map[string]struct{})
	quotas = make(map[string]map[string]int)

	for _, line := range raw {
		// nodes
		if strings.Contains(line, "TRANSMISSION") {
			lr := strings.Split(line, " RELAYS ")
			l := strings.Split(lr[0], " ")
			r := strings.Split(lr[1], " ")
			lNode, rNode := l[len(l)-1], r[0]
			nodesMap[lNode] = struct{}{}
			nodesMap[rNode] = struct{}{}
			q, err := strconv.Atoi(r[len(r)-1])
			if err != nil {
				panic(err)
			}
			if _, ok := quotas[lNode]; !ok {
				quotas[lNode] = make(map[string]int)
			}
			quotas[lNode][rNode] = q

		}
		// start point | hub
		if strings.Contains(line, "ALERT") {
			ws := strings.Split(line, " ")
			hub = ws[len(ws)-1]
		}
		// goal points
		if strings.Contains(line, "CRITICAL") {
			ws := strings.Split(
				strings.Split(
					line,
					" CRITICAL: FINAL ARRIVAL POINTS ARE ")[1],
				", ",
			)
			for _, w := range ws {
				goalsMap[w] = struct{}{}
			}
		}
	}

	goals = slices.Collect(maps.Keys(goalsMap))
	nodes = slices.Collect(maps.Keys(nodesMap))

	return hub, goals, nodes, quotas
}
