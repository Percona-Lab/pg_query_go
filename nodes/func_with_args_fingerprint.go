// Auto-generated - DO NOT EDIT

package pg_query

func (node FuncWithArgs) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("FuncWithArgs")

	if len(node.Funcargs.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Funcargs.Fingerprint(&subCtx, "Funcargs")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("funcargs")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.Funcname.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Funcname.Fingerprint(&subCtx, "Funcname")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("funcname")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}
