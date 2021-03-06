package query

import (
	"errors"

	"github.com/charliekenney23/go-graphql-todo/app/graphql/types"
	"github.com/charliekenney23/go-graphql-todo/app/model"
	"github.com/charliekenney23/go-graphql-todo/app/shared"
	"github.com/graphql-go/graphql"
)

var getUser = &graphql.Field{
	Type:        types.User,
	Description: "Get user by username",
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		var user model.User

		username := params.Args["username"].(string)

		if err := shared.SharedApp.DB.Preload("Tasks").Find(&user, "username = ?", username).Error; err != nil {
			return nil, errors.New("Could not find user")
		}

		return user, nil
	},
}

var getAllUsers = &graphql.Field{
	Type:        graphql.NewList(types.User),
	Description: "Get all users",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		var users []model.User

		if err := shared.SharedApp.DB.Find(&users).Error; err != nil {
			return nil, errors.New("Could not resolve users")
		}

		return users, nil
	},
}
