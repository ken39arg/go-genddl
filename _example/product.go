package example

import (
	"database/sql"
	"time"

	"github.com/mackee/go-genddl/index"
)

// NullUserID is optional UserID
type NullUserID = sql.NullInt32

// Product is product of user
//+table: product
type Product struct {
	ID             uint32        `db:"id,primarykey,autoincrement"`
	Name           string        `db:"name"`
	Type           uint32        `db:"type"`
	UserID         uint32        `db:"user_id"`
	ReceivedUserID NullUserID    `db:"received_user_id"`
	Description    string        `db:"description,text"`
	Size           sql.NullInt16 `db:"size"`
	CreatedAt      time.Time     `db:"created_at"`
	UpdatedAt      sql.NullTime  `db:"updated_at"`
}

func (s Product) _schemaIndex(methods index.Methods) []index.Definition {
	return []index.Definition{
		methods.Unique(s.UserID, s.Type),
		methods.Complex(s.UserID, s.CreatedAt),
		methods.ForeignKey(s.UserID, User{}.ID, index.ForeignKeyDeleteCascade, index.ForeignKeyUpdateCascade),
	}
}
