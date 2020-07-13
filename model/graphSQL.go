package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

var testType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Test",
	Description: "Test Model",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"number": &graphql.Field{
			Type: graphql.Int,
		},
		"alias": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var query = graphql.Field{
	Name: "QueryTestData",
	Type: testType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"title": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["id"].(int)
		title, _ := p.Args["title"].(string)
		fmt.Println(title)
		result := Test{}
		DB.Model(&Test{}).First(&result, id)
		fmt.Println(id, result)
		return result, nil
	},
}

var helloType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Test",
	Description: "Test Model",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var queryHello = graphql.Field{
	Name: "QueryTestData",
	Type: helloType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return gin.H{"id": 1, "name": "HeliantHuS"}, nil
	},
}

// 定义跟查询节点
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RootQuery",
	Description: "Root Query",
	Fields: graphql.Fields{
		"test": &query, // queryHello 参考schema/hello.go
		"hello": &queryHello,
	},
})

// 定义Schema用于http handler处理
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: nil, // 需要通过GraphQL更新数据，可以定义Mutation
})
