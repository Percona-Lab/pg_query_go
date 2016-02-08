// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RangeTblEntry) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RangeTblEntry")

	if node.Alias != nil {
		subCtx := FingerprintSubContext{}
		node.Alias.Fingerprint(&subCtx, "Alias")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("alias")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.CheckAsUser != 0 {
		ctx.WriteString("checkAsUser")
		ctx.WriteString(strconv.Itoa(int(node.CheckAsUser)))
	}

	if len(node.Ctecolcollations.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Ctecolcollations.Fingerprint(&subCtx, "Ctecolcollations")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("ctecolcollations")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.Ctecoltypes.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Ctecoltypes.Fingerprint(&subCtx, "Ctecoltypes")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("ctecoltypes")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.Ctecoltypmods.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Ctecoltypmods.Fingerprint(&subCtx, "Ctecoltypmods")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("ctecoltypmods")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.Ctelevelsup != 0 {
		ctx.WriteString("ctelevelsup")
		ctx.WriteString(strconv.Itoa(int(node.Ctelevelsup)))
	}

	if node.Ctename != nil {
		ctx.WriteString("ctename")
		ctx.WriteString(*node.Ctename)
	}

	if node.Eref != nil {
		subCtx := FingerprintSubContext{}
		node.Eref.Fingerprint(&subCtx, "Eref")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("eref")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.Funcordinality {
		ctx.WriteString("funcordinality")
		ctx.WriteString(strconv.FormatBool(node.Funcordinality))
	}

	if len(node.Functions.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Functions.Fingerprint(&subCtx, "Functions")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("functions")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.InFromCl {
		ctx.WriteString("inFromCl")
		ctx.WriteString(strconv.FormatBool(node.InFromCl))
	}

	if node.Inh {
		ctx.WriteString("inh")
		ctx.WriteString(strconv.FormatBool(node.Inh))
	}

	if len(node.Joinaliasvars.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.Joinaliasvars.Fingerprint(&subCtx, "Joinaliasvars")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("joinaliasvars")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if int(node.Jointype) != 0 {
		ctx.WriteString("jointype")
		ctx.WriteString(strconv.Itoa(int(node.Jointype)))
	}

	if node.Lateral {
		ctx.WriteString("lateral")
		ctx.WriteString(strconv.FormatBool(node.Lateral))
	}

	ctx.WriteString("modifiedCols")
	for _, val := range node.ModifiedCols {
		ctx.WriteString(strconv.Itoa(int(val)))
	}

	if node.Relid != 0 {
		ctx.WriteString("relid")
		ctx.WriteString(strconv.Itoa(int(node.Relid)))
	}

	if node.Relkind != 0 {
		ctx.WriteString("relkind")
		ctx.WriteString(string(node.Relkind))

	}

	if node.RequiredPerms != 0 {
		ctx.WriteString("requiredPerms")
		ctx.WriteString(strconv.Itoa(int(node.RequiredPerms)))
	}

	if int(node.Rtekind) != 0 {
		ctx.WriteString("rtekind")
		ctx.WriteString(strconv.Itoa(int(node.Rtekind)))
	}

	if len(node.SecurityQuals.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.SecurityQuals.Fingerprint(&subCtx, "SecurityQuals")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("securityQuals")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.SecurityBarrier {
		ctx.WriteString("security_barrier")
		ctx.WriteString(strconv.FormatBool(node.SecurityBarrier))
	}

	ctx.WriteString("selectedCols")
	for _, val := range node.SelectedCols {
		ctx.WriteString(strconv.Itoa(int(val)))
	}

	if node.SelfReference {
		ctx.WriteString("self_reference")
		ctx.WriteString(strconv.FormatBool(node.SelfReference))
	}

	if node.Subquery != nil {
		subCtx := FingerprintSubContext{}
		node.Subquery.Fingerprint(&subCtx, "Subquery")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("subquery")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.ValuesCollations.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.ValuesCollations.Fingerprint(&subCtx, "ValuesCollations")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("values_collations")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.ValuesLists.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.ValuesLists.Fingerprint(&subCtx, "ValuesLists")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("values_lists")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}
