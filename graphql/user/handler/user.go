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

package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/labstack/echo"

	"github.com/fengyfei/gopher/graphql/user/module"
)

func UserHandler(c echo.Context) error {
	params := graphql.Params{
		Schema:        module.UserSchema,
		RequestString: c.Request().URL.Query().Get("query"),
	}

	result := graphql.Do(params)
	if result.HasErrors() {
		log.Printf("Failed due to errors: %v\n", result.Errors)
		err := errors.New(fmt.Sprintf("%v", result.Errors))
		return err
	}

	return c.JSON(http.StatusOK, result.Data)
}
