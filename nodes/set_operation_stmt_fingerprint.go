// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node SetOperationStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("SetOperationStmt")

	if node.All {
		ctx.WriteString("all")
		ctx.WriteString(strconv.FormatBool(node.All))
	}

	if len(node.ColCollations.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.ColCollations.Fingerprint(&subCtx, "ColCollations")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("colCollations")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.ColTypes.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.ColTypes.Fingerprint(&subCtx, "ColTypes")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("colTypes")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.ColTypmods.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.ColTypmods.Fingerprint(&subCtx, "ColTypmods")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("colTypmods")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.GroupClauses.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.GroupClauses.Fingerprint(&subCtx, "GroupClauses")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("groupClauses")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.Larg != nil {
		subCtx := FingerprintSubContext{}
		node.Larg.Fingerprint(&subCtx, "Larg")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("larg")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if int(node.Op) != 0 {
		ctx.WriteString("op")
		ctx.WriteString(strconv.Itoa(int(node.Op)))
	}

	if node.Rarg != nil {
		subCtx := FingerprintSubContext{}
		node.Rarg.Fingerprint(&subCtx, "Rarg")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("rarg")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}
