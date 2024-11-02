package migrations

import "github.com/go-rel/rel"

func MigrateCreateUserProfiles(schema *rel.Schema) {
	schema.CreateTable("user_profiles", func(t *rel.Table) {
		t.BigID("id")
		t.DateTime("created_at", rel.Required(true), rel.Options("DEFAULT now()"))
		t.DateTime("updated_at", rel.Required(true), rel.Options("DEFAULT now()"))
		t.BigInt("user_id", rel.Required(true), rel.Unique(true))
		t.Bool("verified", rel.Required(true), rel.Default(false))
		t.Text("name")
		t.SmallInt("age")
		t.Text("bio")
		t.Text("pic_url")

		t.ForeignKey("user_id", "users", "id")
	})
}

func RollbackCreateUserProfiles(schema *rel.Schema) {
	schema.DropTable("user_profiles")
}
