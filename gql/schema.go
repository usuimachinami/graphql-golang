package gql

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"graphql-golang/db"
	"graphql-golang/model"
	"log"
)

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "user",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
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

var favoriteType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Favorite",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"full_title_id": &graphql.Field{
				Type: graphql.String,
			},
			"user_id": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var titleType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "title",
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
			"stories": &graphql.Field{
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
		Name: "story",
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
		Name: "query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)
					db := db.ConnectGORM()
					user := model.User{}
					db.Where("id = ?", id).First(&user)
					return user, nil
				},
			},
			"users": &graphql.Field{
				Type: graphql.NewList(userType),
				Args: graphql.FieldConfigArgument{
					"order": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					order := p.Args["order"].(string)
					if order != "DESC" {
						order = "ASC"
					}
					db := db.ConnectGORM()
					users := []model.User{}
					db.Order("id "+order).Find(&users)
					return users, nil
				},
			},
			"title": &graphql.Field{
				Type: titleType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					db := db.ConnectGORM()
					id := p.Args["id"].(string)
					title := model.Title{}
					db.Where("full_title_id = ?", id).Find(&title)
					return title, nil
				},
			},
			"titles": &graphql.Field{
				Type: graphql.NewList(titleType),
				Args: graphql.FieldConfigArgument{
					"order": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					order := p.Args["order"].(string)
					if order != "DESC" {
						order = "ASC"
					}
					db := db.ConnectGORM()
					titles := []model.Title{}
					db.Order("full_title_id "+order).Find(&titles)
					return titles, nil
				},
			},
			"story": &graphql.Field{
				Type: storyType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)
					db := db.ConnectGORM()
					story := model.Story{}
					story.FullStoryId = id
					db.Where("full_story_id = ?", id).First(&story)
					return story, nil
				},
			},
			"stories": &graphql.Field{
				Type: graphql.NewList(storyType),
				Args: graphql.FieldConfigArgument{
					"order": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					order := p.Args["order"].(string)
					if order != "DESC" {
						order = "ASC"
					}
					db := db.ConnectGORM()
					stories := []model.Story{}
					db.Order("full_story_id "+order).Find(&stories)
					return stories, nil
				},
			},
			"favorites": &graphql.Field{
				Type: graphql.NewList(favoriteType),
				Args: graphql.FieldConfigArgument{
					"user_id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"order": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					userId := p.Args["user_id"].(string)
					order := p.Args["order"].(string)
					if order != "DESC" {
						order = "ASC"
					}
					db := db.ConnectGORM()
					favorites := []model.Favorite{}
					db.Where("user_id = ?", userId).Order("id "+order).Find(&favorites)
					return favorites, nil
				},
			},
		},
	},
)

var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "mutation",
		Fields: graphql.Fields{
			"createFavorite": &graphql.Field{
				Type: favoriteType,
				Args: graphql.FieldConfigArgument{
					"user_id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"full_title_id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					userId, _ := params.Args["user_id"].(int)
					fullTitleId, _ := params.Args["full_title_id"].(string)

					newFavorite := &model.Favorite{
						UserId: userId,
						FullTitleId: fullTitleId,
					}
					db := db.ConnectGORM()
					db.Create(&newFavorite)

					return newFavorite,nil
				},
			},
		},
	},
)


func ExecuteQuery(query string) *graphql.Result {
	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
			Mutation: mutationType,
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
