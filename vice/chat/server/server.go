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
 *     Initial: 2017/09/20        Jia Chenhui
 */

package main

import (
	"context"
	"log"

	"github.com/matryer/vice/queues/nsq"
)

func Transponder(ctx context.Context, r1, r2 <-chan []byte, s2, s1 chan<- []byte, errs <-chan error) {
	for {
		select {
		case <-ctx.Done():
			log.Println("Chat finished.")
		case err := <-errs:
			log.Printf("Transponder retruned error: %s", err.Error())
		case msg := <-r1:
			log.Println("Receive message from r1, start transpond.")
			s2 <- msg
		case msg := <-r2:
			log.Println("Receive message from r2, start transpond.")
			s1 <- msg
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	transport := nsq.New()
	defer func() {
		transport.Stop()
		<-transport.Done()
	}()

	r1 := transport.Receive("receiveFrom1")
	s2 := transport.Send("sendTo2")
	r2 := transport.Receive("receiveFrom2")
	s1 := transport.Send("sendTo1")

	Transponder(ctx, r1, r2, s2, s1, transport.ErrChan())
}
