package datatype

import (
	"database/sql"
	"encoding/json"

	"github.com/SofyanHadiA/linq/core/utils"
)

type JsonNullInt64 struct {
	sql.NullInt64
}

func (v JsonNullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	} else {
		return json.Marshal(nil)
	}
}

func (v *JsonNullInt64) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *int64
	err := json.Unmarshal(data, &x)

	if utils.HandleWarn(err) {
		return err
	}

	if x != nil {
		v.Valid = true
		v.Int64 = *x
	} else {
		v.Valid = false
	}
	return nil
}
