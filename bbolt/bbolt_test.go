/*
 * MIT License
 *
 * Copyright (c) 2017 SmartestEE Co., Ltd..
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
 *     Initial: 2017/09/14        Yang Chenglong
 */

package bbolt

import (
	"log"
	"strconv"
	"testing"
)

func BenchmarkUserServiceProvider_Create(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := UserService.Create("test", 100)
		if err != nil {
			log.Printf("create testing: %v", err)
		}
	}
}

func BenchmarkUserServiceProvider_Get(b *testing.B) {
	loop := make([]string, 1000000)

	for i := 0; i < b.N; i++ {
		for t := range loop {
			ts := strconv.Itoa(t)
			name := "test" + ts
			_, err := UserService.Get(name)
			if err != nil {
				log.Printf("get testing: %v", err)
			}
		}
	}
}

func BenchmarkUserServiceProvider_CreateOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := UserService.CreateOne("test", 100)
		if err != nil {
			log.Printf("create testing: %v", err)
		}
	}
}

func BenchmarkUserServiceProvider_GetOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := UserService.GetOne(10)
		if err != nil {
			log.Printf("get testing: %v", err)
		}
	}
}
