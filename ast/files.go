package ast

import "github.com/pingcap/parser/format"

// LoadFileStmt is the statement node for loading statistic.
type LoadFileStmt struct {
	stmtNode

	DsName string
	Path   string
}

// Restore implements Node interface.
func (n *LoadFileStmt) Restore(ctx *format.RestoreCtx) error {
	ctx.WriteKeyWord("LOAD LOCALFILE ")
	ctx.WriteString(n.Path)
	ctx.WriteString(n.DsName)
	return nil
}

// Accept implements Node Accept interface.
func (n *LoadFileStmt) Accept(v Visitor) (Node, bool) {
	newNode, skipChildren := v.Enter(n)
	if skipChildren {
		return v.Leave(newNode)
	}
	n = newNode.(*LoadFileStmt)
	return v.Leave(n)
}
