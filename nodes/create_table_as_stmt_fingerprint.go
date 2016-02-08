// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node CreateTableAsStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("CreateTableAsStmt")

	if node.Into != nil {
		subCtx := FingerprintSubContext{}
		node.Into.Fingerprint(&subCtx, "Into")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("into")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.IsSelectInto {
		ctx.WriteString("is_select_into")
		ctx.WriteString(strconv.FormatBool(node.IsSelectInto))
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

	if int(node.Relkind) != 0 {
		ctx.WriteString("relkind")
		ctx.WriteString(strconv.Itoa(int(node.Relkind)))
	}
}
