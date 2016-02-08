// Auto-generated - DO NOT EDIT

package pg_query

func (node A_Indices) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("A_Indices")

	if node.Lidx != nil {
		subCtx := FingerprintSubContext{}
		node.Lidx.Fingerprint(&subCtx, "Lidx")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("lidx")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.Uidx != nil {
		subCtx := FingerprintSubContext{}
		node.Uidx.Fingerprint(&subCtx, "Uidx")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("uidx")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}
