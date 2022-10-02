package main

type Node struct {
	token Token
	node  []*Node
}

type TreeBase struct {
	root *Node
	line int
}

var trees []*TreeBase

func (n *Node) add(node *Node) {
	n.node = append(n.node, node)
}

func (n *Node) addToken(token Token) {
	n.token = token
}

func treeWalk() {
	for _,token := range tokens {
		if token.tokenType == VAR {
			trees = append(trees, &Node{token: token})
		}
	}

}