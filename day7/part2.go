package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	dir      bool
	name     string
	children []*Node
	parent   *Node
	size     int // either the size of the file for dir == false or the size of all underlying files if dir == true
}

func (node Node) hasChildren() bool {
	if node.children != nil {
		return true
	}
	return false
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	root := &Node{dir: true, name: "/", children: []*Node{}}
	curNode := root
	// build tree
	for scanner.Scan() {
		elements := strings.Split(scanner.Text(), " ")
		if len(elements) == 3 && elements[2] == "/" { // root dir
			continue
		}
		if elements[0] == "$" { // run a command
			if elements[1] == "ls" { // list contents of current directory
				continue
			} else { // change to a directory
				if elements[2] == ".." {
					curNode = curNode.parent
					continue
				}
				if curNode.hasChildren() {
					for _, child := range curNode.children {
						if child.dir && child.name == elements[2] {
							curNode = child
						}
					}
				} else {
					child := Node{dir: true, name: elements[2], parent: curNode}
					curNode.children = append(curNode.children, &child)
					curNode = &child
				}
			}
		} else {
			if elements[0] == "dir" { // add or noop a directory as a child
				add := true
				if curNode.hasChildren() {
					for _, child := range curNode.children {
						if child.dir && child.name == elements[1] {
							add = false
						}
					}
					if add {
						curNode.children = append(curNode.children, &Node{dir: true, name: elements[1], parent: curNode})
					}
				} else {
					curNode.children = append(curNode.children, &Node{dir: true, name: elements[1], parent: curNode})
				}
			} else { // file
				add := true
				if curNode.hasChildren() {
					for _, child := range curNode.children {
						if !child.dir && child.name == elements[1] {
							add = false
						}
					}
					if add {
						size := 0
						size, _ = strconv.Atoi(elements[0])
						curNode.children = append(curNode.children,
							&Node{dir: false, name: elements[1], parent: curNode, size: size})
					}
				} else {
					size := 0
					size, _ = strconv.Atoi(elements[0])
					curNode.children = append(curNode.children,
						&Node{dir: false, name: elements[1], parent: curNode, size: size})
				}
			}
		}
	}

	// calc sizes
	queue := make([]*Node, 0)
	queue = append(queue, root)
	candidates := make([]*Node, 0)
	seenRoot := false
	unused := 0
	required := 30000000
	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]
		next.calcSize()
		if !seenRoot {
			unused = 70000000 - next.size
			seenRoot = true
		} else {
			if next.size + unused >= required && next.dir {
				candidates = append(candidates, next)
			}
		}
		if next.hasChildren() {
			for _, child := range next.children {
				queue = append(queue, child)
			}	
		}
	}
	
	smallest := 0
	name := ""
	for i, node := range candidates {
		if(i == 0) {
			smallest = node.size
			name = node.name
		} else {
			if node.size < smallest {
				smallest = node.size
				name = node.name
			}
		}
	}
	fmt.Println(name)
	fmt.Println(smallest)
}

func (node *Node) calcSize() {
	sum := 0
	if node == nil {
		return
	}
	if node.hasChildren() {
		for _, child := range node.children {
			if child.dir {
				child.calcSize()
			} 
			sum += child.size
		}
		node.size = sum
	} 
}
