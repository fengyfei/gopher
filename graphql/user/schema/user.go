/*
 * MIT License
 *
 * Copyright (c) 2017 SmartestEE Co., Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2017/09/27        Jia Chenhui
 */

package schema

import (
	"fmt"
	"log"

	"github.com/graphql-go/graphql"

	"github.com/fengyfei/gopher/graphql/user/module"
)

var (
	UserSchema graphql.Schema
)

func init() {
	s, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	UserSchema = s
}

var (
	rootQuery    = graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	rootMutation = graphql.ObjectConfig{Name: "RootMutation", Fields: mutations}

	schemaConfig = graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: graphql.NewObject(rootMutation),
	}
)

var (
	// user data structure
	userType = graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"login": &graphql.Field{
				Type: graphql.String,
			},
			"admin": &graphql.Field{
				Type: graphql.String,
			},
			"active": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
)

var (
	// query data
	// get: curl -g 'http://localhost:8989/graphql?query={user(login:"jch"){login,admin,active}}'
	// An example GraphQL query might look like:
	/*
		{
			user(login: "leon") {
				login, admin, active
			  }
			}
	*/
	fields = graphql.Fields{
		"hello": &graphql.Field{
			Type:        graphql.String,
			Description: "Hello world",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				fmt.Println(p.Args)
				return "world", nil
			},
		},
		"user": &graphql.Field{
			Type:        userType,
			Description: "Get single user info",
			Args: graphql.FieldConfigArgument{
				"login": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: module.GetSingleInfo,
		},
	}

	// mutation data
	// create: curl -g 'http://localhost:8989/graphql?query=mutation+_{addNewUser(login:"jch",admin:"yes",active:"yes"){login,admin,active}}'
	// An example GraphQL mutation might look like:
	/*
		mutation {
		  addNewUser(login: "leon", admin: "true", active: "true") {
		    login, admin, active
		  }
		}
	*/
	mutations = graphql.Fields{
		"addNewUser": &graphql.Field{
			Type:        userType,
			Description: "Add new user",
			Args: graphql.FieldConfigArgument{
				"login": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"admin": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"active": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: module.Create,
		},
	}
)
