// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterFunctionStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterFunctionStmt")

	if len(node.Actions.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Actions.Fingerprint(&subCtx, "Actions")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("actions")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.Func != nil {
		subCtx := FingerprintSubContext{}
		node.Func.Fingerprint(&subCtx, "Func")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("func")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}
