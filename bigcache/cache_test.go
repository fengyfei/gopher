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
 *     Initial: 2017/09/16        Jia Chenhui
 */

package cache_test

import (
	"log"
	"testing"

	"github.com/fengyfei/gopher/bigcache"
)

func BenchmarkCacheServiceProvider_SetOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		key := "key" + cache.IntToStr(i)
		cache.CacheServer.SetOne(key, "test")
	}
}

func BenchmarkCacheServiceProvider_SetMany(b *testing.B) {
	for i := 0; i < b.N; i++ {
		key := "key" + cache.IntToStr(i)
		cache.CacheServer.SetMany(key, "test")
	}
}

func BenchmarkCacheServiceProvider_GetOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := cache.CacheServer.GetOne("key1")
		if err != nil {
			log.Printf("GetOne testing: %s", err.Error())
		}
	}
}

func BenchmarkCacheServiceProvider_GetAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := cache.CacheServer.GetAll()
		if err != nil {
			log.Printf("GetAll testing: %s", err.Error())
		}
	}
}
