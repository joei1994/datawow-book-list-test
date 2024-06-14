package list

import (
	mysql "github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateBooksTable struct{}

func (m *CreateBooksTable) GetName() string {
	return "CreateBooksTable"
}

func (m *CreateBooksTable) Up(con *sqlx.DB) {
	table := mysql.NewTable("books", con)
	table.Column("id").Type("int unsigned").Autoincrement()
	table.PrimaryKey("id")
	table.String("title", 500).NotNull()
	table.String("author", 100).NotNull()
	table.String("genre", 100).NotNull()
	table.String("year", 10).NotNull()
	table.String("created_by", 100).NotNull()
	table.Column("deleted_at").Type("datetime").Nullable()
	table.WithTimestamps()

	table.MustExec()
}

func (m *CreateBooksTable) Down(con *sqlx.DB) {
	mysql.DropTable("books", con).MustExec()
}
