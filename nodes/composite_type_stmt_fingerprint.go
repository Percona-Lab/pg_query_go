// Auto-generated - DO NOT EDIT

package pg_query

func (node CompositeTypeStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CompositeTypeStmt")

	if len(node.Coldeflist.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Coldeflist.Fingerprint(&subCtx, "Coldeflist")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("coldeflist")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.Typevar != nil {
		subCtx := FingerprintSubContext{}
		node.Typevar.Fingerprint(&subCtx, "Typevar")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("typevar")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}
