// Auto-generated from postgres/src/include/nodes/primnodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type ConvertRowtypeExpr struct {
	Xpr        Expr  `json:"xpr"`
	Arg        *Expr `json:"arg"`        /* input expression */
	Resulttype Oid   `json:"resulttype"` /* output type (always a composite type) */

	/* Like RowExpr, we deliberately omit a typmod and collation here */
	Convertformat CoercionForm `json:"convertformat"` /* how to display this node */
	Location      int          `json:"location"`      /* token location, or -1 if unknown */
}

func (node ConvertRowtypeExpr) MarshalJSON() ([]byte, error) {
	type ConvertRowtypeExprMarshalAlias ConvertRowtypeExpr
	return json.Marshal(map[string]interface{}{
		"CONVERTROWTYPEEXPR": (*ConvertRowtypeExprMarshalAlias)(&node),
	})
}

func (node *ConvertRowtypeExpr) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["xpr"] != nil {
		err = json.Unmarshal(fields["xpr"], &node.Xpr)
		if err != nil {
			return
		}
	}

	if fields["arg"] != nil {
		var nodePtr *Node
		nodePtr, err = UnmarshalNodePtrJSON(fields["arg"])
		if err != nil {
			return
		}
		if nodePtr != nil && *nodePtr != nil {
			val := (*nodePtr).(Expr)
			node.Arg = &val
		}
	}

	if fields["resulttype"] != nil {
		err = json.Unmarshal(fields["resulttype"], &node.Resulttype)
		if err != nil {
			return
		}
	}

	if fields["convertformat"] != nil {
		err = json.Unmarshal(fields["convertformat"], &node.Convertformat)
		if err != nil {
			return
		}
	}

	if fields["location"] != nil {
		err = json.Unmarshal(fields["location"], &node.Location)
		if err != nil {
			return
		}
	}

	return
}