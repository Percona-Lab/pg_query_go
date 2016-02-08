// Auto-generated - DO NOT EDIT

package pg_query

func (node ExplainStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("ExplainStmt")

	if len(node.Options.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Options.Fingerprint(&subCtx, "Options")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("options")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.Query != nil {
		subCtx := FingerprintSubContext{}
		node.Query.Fingerprint(&subCtx, "Query")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("query")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}
