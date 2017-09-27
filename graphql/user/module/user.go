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

package module

import (
	"log"

	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2/bson"

	"github.com/fengyfei/gopher/graphql/user/mongo"
)

// user struct
type User struct {
	Login  string `json:"login"`
	Admin  string `json:"admin"`
	Active string `json:"active"`
}

// GetSingleInfo get single user information.
func GetSingleInfo(p graphql.ResolveParams) (interface{}, error) {
	var (
		u User
	)

	login := p.Args["login"].(string)

	err := mongo.MDB.C("users").Find(bson.M{"login": login}).One(&u)
	if err != nil {
		log.Printf("GetSingleInfo returned error: %v", err)
		return nil, err
	}

	return u, nil
}

// Create create single user.
func Create(p graphql.ResolveParams) (interface{}, error) {
	user := User{
		Login:  p.Args["login"].(string),
		Admin:  p.Args["admin"].(string),
		Active: p.Args["active"].(string),
	}

	err := mongo.MDB.C("users").Insert(&user)
	if err != nil {
		log.Printf("Create user returned error: %v", err)
	}

	return err == nil, err
}
