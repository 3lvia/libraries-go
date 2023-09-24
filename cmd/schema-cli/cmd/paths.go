package cmd

import (
	"sort"
	"strings"
)

type pathComputer interface {
	addSubject(subject string)
	sortedPaths() []string
}

func newPathComputer(path string) pathComputer {
	p := path
	if p == "" {
		p = "."
	}

	return &pathComputerImpl{
		path:  p,
		nodes: map[string]*pathComputerImpl{},
	}
}

type pathComputerImpl struct {
	path  string
	nodes map[string]*pathComputerImpl
}

func (p *pathComputerImpl) addSubject(subject string) {
	s := strings.Replace(subject, "-", "_", -1)
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)

	if len(s) > len("_value") && s[len(s)-len("_value"):] == "_value" {
		s = s[:len(s)-len("_value")]
	}

	i := strings.Index(s, ".")
	if i == -1 {
		p.nodes[s] = &pathComputerImpl{
			path:  s,
			nodes: map[string]*pathComputerImpl{},
		}
		return
	}

	k := s[:i]
	if k == "private" || k == "public" {
		s = s[i+1:]
		i = strings.Index(s, ".")
		k = s[:i] + "/" + k
	}
	if _, ok := p.nodes[k]; !ok {
		p.nodes[k] = &pathComputerImpl{
			path:  k,
			nodes: map[string]*pathComputerImpl{},
		}
	}

	p.nodes[k].addSubject(s[i+1:])
}

func (p *pathComputerImpl) sortedPaths() []string {
	result := p.addPaths("", nil)
	sort.Strings(result)
	return result
}

func (p *pathComputerImpl) addPaths(parent string, accumulator []string) []string {
	if p.path != "." {
		accumulator = append(accumulator, parent+"/"+p.path)
	}

	if len(p.nodes) == 0 {
		return accumulator
	}

	pp := parent + "/" + p.path
	if parent == "" {
		pp = p.path
	}

	for _, n := range p.nodes {
		accumulator = n.addPaths(pp, accumulator)
	}
	return accumulator
}
