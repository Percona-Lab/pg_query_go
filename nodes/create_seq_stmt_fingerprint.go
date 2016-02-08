// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateSeqStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateSeqStmt")

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

	if node.OwnerId != 0 {
		ctx.WriteString("ownerId")
		ctx.WriteString(strconv.Itoa(int(node.OwnerId)))
	}

	if node.Sequence != nil {
		subCtx := FingerprintSubContext{}
		node.Sequence.Fingerprint(&subCtx, "Sequence")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("sequence")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}
