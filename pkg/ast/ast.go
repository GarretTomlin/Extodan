package ast

import "strconv"

type NodeType string

const (
	NodeProgram             NodeType = "Program"
	NodeFunctionDeclaration NodeType = "FunctionDeclaration"
	NodeParameter           NodeType = "Parameter"
	NodeReturnStatement     NodeType = "ReturnStatement"
	NodeExpression          NodeType = "Expression"
	NodeTerm                NodeType = "Term"
	NodeOperator            NodeType = "Operator"
	NodeIdentifier          NodeType = "Identifier"
	NodeInteger             NodeType = "Integer"
	NodeMissingExpression NodeType = "MissingExpression"

)

type Node struct {
	Type     NodeType
	Value    string
	Children []Node
}

func NewNode(nodeType NodeType, value string) Node {
	return Node{
		Type:  nodeType,
		Value: value,
	}
}

func AddChild(parent Node, child Node) Node {
	parent.Children = append(parent.Children, child)
	return parent
}

// Evaluate evaluates the AST node and returns the result.
func (n Node) Evaluate() int {
	switch n.Type {
	case NodeInteger:
		value, _ := strconv.Atoi(n.Value)
		return value
	case NodeExpression:
		left := n.Children[0].Evaluate()
		right := n.Children[1].Evaluate()
		switch n.Value {
		case "+":
			return left + right
		// Add other operators here
		}
	default:
		// Handle other node types
		return 0
	}
	return 0
}
