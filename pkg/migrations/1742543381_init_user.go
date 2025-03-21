package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		superusers, err := app.FindCollectionByNameOrId(core.CollectionNameSuperusers)
		if err != nil {
			return err
		}

		record := core.NewRecord(superusers)
		record.Set("email", "admin@eki.local")
		record.Set("password", "P@ssw0rd")

		return app.Save(record)
	}, func(app core.App) error {
		// add down queries...

		return nil
	})
}
