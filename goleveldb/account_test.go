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
 *     Initial: 2017/10/09        Jia Chenhui
 */

package account_test

import (
	"log"
	"math/rand"
	"testing"

	"github.com/fengyfei/gopher/goleveldb"
)

func BenchmarkAccountServiceProvider_Create(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := account.AccountService.Create("test", 100)
		if err != nil {
			log.Printf("CreateOne testing error: %v\n", err)
		}
	}
}

func BenchmarkAccountServiceProvider_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := account.AccountService.Get(i)
		if err != nil {
			log.Printf("GetOne sequential testing error: %v\n", err)
		}
	}
}

func BenchmarkAccountServiceProvider_GetRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		id := rand.Intn(b.N)
		_, err := account.AccountService.Get(id)
		if err != nil {
			log.Printf("GetOne random testing returned error: %v\n", err)
		}
	}
}
