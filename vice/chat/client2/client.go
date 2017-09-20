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
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/matryer/vice/queues/nsq"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	defer func() {
		signal.Stop(c)
		cancel()
	}()

	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()

	transport := nsq.New()
	s := transport.Send("receiveFrom2")
	r := transport.Receive("sendTo2")

	go func() {
		for msg := range r {
			fmt.Println("Receive message from client1:", string(msg))
		}
	}()

	go func() {
		log.Println("Say something...")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			b := scanner.Bytes()
			if len(b) == 0 {
				continue
			}
			s <- b
		}
		if err := scanner.Err(); err != nil {
			log.Printf("Scan returned error: %s", err.Error())
		}
	}()

	<-ctx.Done()
	transport.Stop()
	<-transport.Done()

	log.Println("Client1 send message successed.")
}
