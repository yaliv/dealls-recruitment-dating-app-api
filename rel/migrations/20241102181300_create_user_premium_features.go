package migrations

import "github.com/go-rel/rel"

func MigrateCreateUserPremiumFeatures(schema *rel.Schema) {
	schema.CreateTable("user_premium_features", func(t *rel.Table) {
		t.BigID("id")
		t.DateTime("created_at", rel.Required(true), rel.Options("DEFAULT now()"))
		t.DateTime("updated_at", rel.Required(true), rel.Options("DEFAULT now()"))
		t.BigInt("user_id", rel.Required(true))
		t.BigInt("feature_id", rel.Required(true))

		t.ForeignKey("user_id", "users", "id")
		t.ForeignKey("feature_id", "premium_features", "id")
	})
}

func RollbackCreateUserPremiumFeatures(schema *rel.Schema) {
	schema.DropTable("user_premium_features")
}
