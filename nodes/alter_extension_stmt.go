// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type AlterExtensionStmt struct {
	Extname *string `json:"extname"`
	Options []Node  `json:"options"` /* List of DefElem nodes */
}

func (node AlterExtensionStmt) MarshalJSON() ([]byte, error) {
	type AlterExtensionStmtMarshalAlias AlterExtensionStmt
	return json.Marshal(map[string]interface{}{
		"ALTEREXTENSIONSTMT": (*AlterExtensionStmtMarshalAlias)(&node),
	})
}

func (node *AlterExtensionStmt) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["extname"] != nil {
		err = json.Unmarshal(fields["extname"], &node.Extname)
		if err != nil {
			return
		}
	}

	if fields["options"] != nil {
		node.Options, err = UnmarshalNodeArrayJSON(fields["options"])
		if err != nil {
			return
		}
	}

	return
}