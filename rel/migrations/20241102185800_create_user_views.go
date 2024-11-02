package migrations

import "github.com/go-rel/rel"

func MigrateCreateUserViews(schema *rel.Schema) {
	schema.CreateTable("user_views", func(t *rel.Table) {
		t.BigID("id")
		t.DateTime("created_at", rel.Required(true), rel.Options("DEFAULT now()"))
		t.DateTime("updated_at", rel.Required(true), rel.Options("DEFAULT now()"))
		t.BigInt("viewer_id", rel.Required(true))
		t.BigInt("target_id", rel.Required(true))
		t.Date("view_date", rel.Required(true), rel.Options("DEFAULT now()"))
		t.Bool("like", rel.Required(true))

		t.ForeignKey("viewer_id", "users", "id")
		t.ForeignKey("target_id", "users", "id")
		t.Unique([]string{"viewer_id", "target_id", "view_date"})
	})
}

func RollbackCreateUserViews(schema *rel.Schema) {
	schema.DropTable("user_views")
}
