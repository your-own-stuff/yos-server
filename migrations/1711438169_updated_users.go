package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		options := map[string]any{}
		if err := json.Unmarshal([]byte(`{
			"allowEmailAuth": true,
			"allowOAuth2Auth": true,
			"allowUsernameAuth": true,
			"exceptEmailDomains": null,
			"manageRule": "@request.auth.isAdmin = true",
			"minPasswordLength": 8,
			"onlyEmailDomains": null,
			"onlyVerified": false,
			"requireEmail": false
		}`), &options); err != nil {
			return err
		}
		collection.SetOptions(options)

		collection.ListRule = types.Pointer("id = @request.auth.id || @request.auth.isAdmin = true")
		collection.ViewRule = types.Pointer("id = @request.auth.id || @request.auth.isAdmin = true")
		collection.UpdateRule = types.Pointer("id = @request.auth.id || @request.auth.isAdmin = true")
		collection.DeleteRule = types.Pointer("id = @request.auth.id || @request.auth.isAdmin = true")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		options := map[string]any{}
		if err := json.Unmarshal([]byte(`{
			"allowEmailAuth": true,
			"allowOAuth2Auth": true,
			"allowUsernameAuth": true,
			"exceptEmailDomains": null,
			"manageRule": null,
			"minPasswordLength": 8,
			"onlyEmailDomains": null,
			"onlyVerified": false,
			"requireEmail": false
		}`), &options); err != nil {
			return err
		}
		collection.SetOptions(options)

		collection.ListRule = types.Pointer("id = @request.auth.id")
		collection.ViewRule = types.Pointer("id = @request.auth.id")
		collection.UpdateRule = types.Pointer("id = @request.auth.id")
		collection.DeleteRule = types.Pointer("id = @request.auth.id")

		return dao.SaveCollection(collection)
	})
}
