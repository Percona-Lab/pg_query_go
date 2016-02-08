// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterOwnerStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterOwnerStmt")

	if node.Newowner != nil {
		ctx.WriteString("newowner")
		ctx.WriteString(*node.Newowner)
	}

	if len(node.Objarg.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Objarg.Fingerprint(&subCtx, "Objarg")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("objarg")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.Object.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Object.Fingerprint(&subCtx, "Object")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("object")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if int(node.ObjectType) != 0 {
		ctx.WriteString("objectType")
		ctx.WriteString(strconv.Itoa(int(node.ObjectType)))
	}

	if node.Relation != nil {
		subCtx := FingerprintSubContext{}
		node.Relation.Fingerprint(&subCtx, "Relation")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("relation")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}
