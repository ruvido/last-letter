package server

import(
	"log"
	"os"
	_ "github.com/ruvido/letter/migrations"
	"github.com/spf13/viper"


    "github.com/pocketbase/pocketbase"
    "github.com/pocketbase/pocketbase/apis"
    "github.com/pocketbase/pocketbase/core"


	"github.com/pocketbase/pocketbase/forms"
    "github.com/pocketbase/pocketbase/models"
    "github.com/pocketbase/pocketbase/models/schema"
    "github.com/pocketbase/pocketbase/tools/types"
)


func Start (){

	log.Println("Starting megaserver!")
	log.Println("Config file used:", viper.ConfigFileUsed())

	app := pocketbase.New()

	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
	
		// this makes the magic!
		settingUpLetters( app )
		// publishScheduledLetters ( app )

		// serving static files, e.g. your newsletter archives 
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./public"), false))
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

func settingUpLetters( app *pocketbase.PocketBase ) {
	// collection, err := app.Dao().FindCollectionByNameOrId("letters")
	collection, err := app.Dao().FindCollectionByNameOrId("letters")
	if err != nil {
		// log.Fatal(err)
		log.Println(err)
		if collection == nil {
			log.Println("collection is missing... creating")
			createCollection(app, "letters")

		}
	}
			// remove users collection if (empty) -> not needed
			deleteCollection(app, "users")

}

func deleteCollection( app *pocketbase.PocketBase, collectionName string ) {
	// record, err := app.Dao().FindFirstRecordByFilter(
	_, err := app.Dao().FindFirstRecordByFilter(
		collectionName, "id != ''",
		// collectionName, "status = 'public' && category = {:category}",
		// dbx.Params{ "category": "news" },
	)
	if err != nil {
		log.Println(err)
		collection, err := app.Dao().FindCollectionByNameOrId(collectionName)
		if err  != nil {
			log.Println(err)
		}
		if err := app.Dao().DeleteCollection(collection); err != nil {
			log.Println(err)
		}
	} 

}

func createCollection( app *pocketbase.PocketBase, collectionName string ) {
	collection := &models.Collection{}

	form := forms.NewCollectionUpsert(app, collection)
	form.Name = collectionName
	form.Type = models.CollectionTypeBase
	form.ListRule = nil
	form.ViewRule = types.Pointer("@request.auth.id != ''")
	form.CreateRule = types.Pointer("")
	form.UpdateRule = types.Pointer("@request.auth.id != ''")
	form.DeleteRule = nil
	form.Schema.AddField(&schema.SchemaField{
		Name:     "title",
		Type:     schema.FieldTypeText,
		Required: true,
		Options: &schema.TextOptions{
			Max: types.Pointer(10),
		},
	})
	form.Schema.AddField(&schema.SchemaField{
		Name:     "date",
		Type:     schema.FieldTypeDate,
		Required: true,
	})
	form.Schema.AddField(&schema.SchemaField{
		Name:     "data",
		Type:     schema.FieldTypeJson,
		Required: true,
		Options: &schema.JsonOptions{
			MaxSize: 2000000,
		},
	})

	// form.Schema.AddField(&schema.SchemaField{
	// 	Name:     "user",
	// 	Type:     schema.FieldTypeRelation,
	// 	Required: true,
	// 	Options: &schema.RelationOptions{
	// 		MaxSelect:     types.Pointer(1),
	// 		CollectionId:  "ae40239d2bc4477",
	// 		CascadeDelete: true,
	// 	},
	// })

	// validate and submit (internally it calls app.Dao().SaveCollection(collection) in a transaction)
	if err := form.Submit(); err != nil {
		log.Fatal(err)
	}
}

func publishScheduledLetters () {
	log.Println("publish scheduled letters")
}

