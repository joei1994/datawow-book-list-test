package list

import (
	mysql "github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type AddTagColumn struct{}

func (m *AddTagColumn) GetName() string {
	return "AddTagColumnToBooksTable"
}

func (m *AddTagColumn) Up(con *sqlx.DB) {
	table := mysql.ChangeTable("books", con)
	table.String("tag", 50)

	table.MustExec()
}

func (m *AddTagColumn) Down(con *sqlx.DB) {
	table := mysql.ChangeTable("books", con)
	table.Column("tags").Drop()

	table.MustExec()
}
