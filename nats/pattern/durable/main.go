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
 *     Initial: 2017/10/29        Feng Yifei
 */

package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"sync"
	"time"

	stan "github.com/nats-io/go-nats-streaming"
)

func logCloser(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Printf("close error: %s", err)
	}
}

func startSubscriber(conn stan.Conn, wg *sync.WaitGroup, n int, d chan<- struct{}) {
	var (
		i   int
		err error
		sub stan.Subscription
	)

	sub, err = conn.Subscribe("counter", func(msg *stan.Msg) {
		// Print the value and whether it was redelivered.
		fmt.Printf("seq = %d [redelivered = %v]\n", msg.Sequence, msg.Redelivered)

		// Add jitter..
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		// Mark it is done.
		wg.Done()

		i++

		msg.Ack()

		if i == n {
			sub.Close()
			d <- struct{}{}
		}
	}, stan.DurableName("i-will-remember"), stan.MaxInflight(1), stan.SetManualAckMode())
	if err != nil {
		log.Print(err)
	}
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	conn, err := stan.Connect(
		"test-cluster",
		"test-client",
		stan.NatsURL("nats://localhost:4222"),
	)
	if err != nil {
		return err
	}
	defer logCloser(conn)

	wg := &sync.WaitGroup{}

	go func() {
		done := make(chan struct{})
		startSubscriber(conn, wg, 5, done)
		<-done
		log.Print("subscriber disconnected..")
		log.Print("reconnecting..")
		startSubscriber(conn, wg, 5, done)
	}()

	// Publish up to 10.
	for i := 0; i < 10; i++ {
		wg.Add(1)

		err := conn.Publish("counter", nil)
		if err != nil {
			return err
		}
	}

	// Wait until all messages have been processed.
	wg.Wait()

	return nil
}