// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterSystemStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterSystemStmt")

	if node.Setstmt != nil {
		subCtx := FingerprintSubContext{}
		node.Setstmt.Fingerprint(&subCtx, "Setstmt")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("setstmt")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}
