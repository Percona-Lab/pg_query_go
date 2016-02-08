// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RowCompareExpr) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RowCompareExpr")

	if len(node.Inputcollids.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Inputcollids.Fingerprint(&subCtx, "Inputcollids")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("inputcollids")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.Largs.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Largs.Fingerprint(&subCtx, "Largs")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("largs")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.Opfamilies.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Opfamilies.Fingerprint(&subCtx, "Opfamilies")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("opfamilies")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.Opnos.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Opnos.Fingerprint(&subCtx, "Opnos")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("opnos")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.Rargs.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Rargs.Fingerprint(&subCtx, "Rargs")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("rargs")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if int(node.Rctype) != 0 {
		ctx.WriteString("rctype")
		ctx.WriteString(strconv.Itoa(int(node.Rctype)))
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
