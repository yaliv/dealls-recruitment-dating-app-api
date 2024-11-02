package migrations

import "github.com/go-rel/rel"

func MigrateCreatePremiumFeatures(schema *rel.Schema) {
	schema.CreateTable("premium_features", func(t *rel.Table) {
		t.BigID("id")
		t.DateTime("created_at", rel.Required(true), rel.Options("DEFAULT now()"))
		t.DateTime("updated_at", rel.Required(true), rel.Options("DEFAULT now()"))
		t.Text("description", rel.Required(true))
		t.Int("price", rel.Required(true))
	})
}

func RollbackCreatePremiumFeatures(schema *rel.Schema) {
	schema.DropTable("premium_features")
}
