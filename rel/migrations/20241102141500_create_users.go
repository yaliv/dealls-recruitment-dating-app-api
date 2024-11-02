package migrations

import "github.com/go-rel/rel"

func MigrateCreateUsers(schema *rel.Schema) {
	schema.CreateTable("users", func(t *rel.Table) {
		t.BigID("id")
		t.DateTime("created_at", rel.Required(true), rel.Options("DEFAULT now()"))
		t.DateTime("updated_at", rel.Required(true), rel.Options("DEFAULT now()"))
		t.DateTime("deactivated_at")
		t.Text("email", rel.Required(true), rel.Unique(true))
		t.Text("secret", rel.Required(true))
	})
}

func RollbackCreateUsers(schema *rel.Schema) {
	schema.DropTable("users")
}
