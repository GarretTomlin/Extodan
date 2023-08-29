package ast

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
