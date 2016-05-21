package datatype

import (
	"database/sql"
	"encoding/json"

	"github.com/SofyanHadiA/linq/core/utils"
)

type JsonNullString struct {
	sql.NullString
}

func (v JsonNullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	} else {
		return json.Marshal(nil)
	}
}

func (v *JsonNullString) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *string
	err := json.Unmarshal(data, &x)

	if utils.HandleWarn(err) {
		return err
	}

	if x != nil {
		v.Valid = true
		v.String = *x
	} else {
		v.Valid = false
	}
	return nil
}
