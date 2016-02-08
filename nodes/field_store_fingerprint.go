// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node FieldStore) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("FieldStore")

	if node.Arg != nil {
		subCtx := FingerprintSubContext{}
		node.Arg.Fingerprint(&subCtx, "Arg")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("arg")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.Fieldnums.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Fieldnums.Fingerprint(&subCtx, "Fieldnums")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("fieldnums")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.Newvals.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Newvals.Fingerprint(&subCtx, "Newvals")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("newvals")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.Resulttype != 0 {
		ctx.WriteString("resulttype")
		ctx.WriteString(strconv.Itoa(int(node.Resulttype)))
	}

	if node.Xpr != nil {
		subCtx := FingerprintSubContext{}
		node.Xpr.Fingerprint(&subCtx, "Xpr")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("xpr")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}
