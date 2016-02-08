// Auto-generated - DO NOT EDIT

package pg_query

func (node AlterUserMappingStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterUserMappingStmt")

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

	if node.Servername != nil {
		ctx.WriteString("servername")
		ctx.WriteString(*node.Servername)
	}

	if node.Username != nil {
		ctx.WriteString("username")
		ctx.WriteString(*node.Username)
	}
}
