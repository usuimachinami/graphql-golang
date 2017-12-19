package gql

import (
	"app/db"
	"app/model"
	"fmt"
	"github.com/graphql-go/graphql"
	"log"
	"strconv"
)

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)

var titleType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Title",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"story": &graphql.Field{
				Type: graphql.NewList(storyType),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					title := p.Source.(model.Title)
					db := db.ConnectGORM()
					story := []model.Story{}
					db.Find(&story, "full_title_id=?", title.FullTitleId)
					log.Print(story)
					return story, nil
				},
			},
		},
	},
)

var storyType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Story",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"User": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, err := strconv.ParseInt(p.Args["id"].(string), 10, 64)
					if err == nil {
						db := db.ConnectGORM()
						user := model.User{}
						user.Id = idQuery
						db.First(&user, idQuery)
						return user, nil
					}
					return nil, err
				},
			},
			"Title": &graphql.Field{
				Type: graphql.NewList(titleType),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					//idQuery := p.Args["id"].(string)
					db := db.ConnectGORM()
					title := []model.Title{}
					//title.FullTitleId = idQuery
					db.Find(&title)
					//log.Print(idQuery)
					return title, nil
				},
			},
			"Story": &graphql.Field{
				Type: storyType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery := p.Args["id"].(string)
					db := db.ConnectGORM()
					story := model.Story{}
					story.FullStoryId = idQuery
					db.First(&story)
					log.Print(idQuery)
					return story, nil
				},
			},
		},
	},
)

func ExecuteQuery(query string) *graphql.Result {
	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
		},
	)

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}

	return result
}
