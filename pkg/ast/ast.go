package ast

// NodeType represents the type of an AST node.
type NodeType string

const (
	NodeProgram              NodeType = "Program"              
	NodeFunctionDeclaration NodeType = "FunctionDeclaration" 
	NodeParameter           NodeType = "Parameter"           
	NodeReturnStatement     NodeType = "ReturnStatement"     
	NodeExpression          NodeType = "Expression"          
	NodeTerm                NodeType = "Term"               
	NodeOperator            NodeType = "Operator"            
	NodeIdentifier          NodeType = "Identifier"         
	NodeInteger             NodeType = "Integer"            

// Node represents a node in the Abstract Syntax Tree (AST).
type Node struct {
	Type     NodeType // The type of the node.
	Value    string   // The value associated with the node (e.g., token value or identifier name).
	Children []Node   // The child nodes of the current node.
}

// NewNode creates a new Node instance with the specified node type and value.
func NewNode(nodeType NodeType, value string) Node {
	return Node{
		Type:  nodeType,
		Value: value,
	}
}

// AddChild adds a child node to the parent node's Children slice and returns the modified parent node.
func AddChild(parent Node, child Node) Node {
	parent.Children = append(parent.Children, child)
	return parent
}
