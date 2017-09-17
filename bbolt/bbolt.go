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
	"encoding/binary"
	"encoding/json"
	"log"

	"github.com/coreos/bbolt"
	"fmt"
)

type UserServiceProvider struct{}

var (
	UserService *UserServiceProvider = &UserServiceProvider{}
	UserDB      *bolt.DB
)

func init() {
	Open()
}

func Open() error {
	db, err := bolt.Open("user.db", 0666, nil)
	if err != nil {
		log.Printf("[open] init db error: %v", err)
		return err
	}
	UserDB = db

	tx, err := UserDB.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err = tx.CreateBucketIfNotExists([]byte("user")); err != nil {
		return err
	}

	return tx.Commit()
}

type User struct {
	Id      uint64
	Name    string
	Payload int
}

func (usp *UserServiceProvider) CreateOne(name string, payload int) (uint64, error) {
	tx, err := UserDB.Begin(true)
	if err != nil {
		log.Printf("[create] begin txn error: %v", err)
		return 0, err
	}
	defer tx.Rollback()

	bucket := tx.Bucket([]byte("user"))

	id, err := bucket.NextSequence()
	if err != nil {
		return 0, err
	}

	user := User{
		Id:      id,
		Name:    name,
		Payload: payload,
	}

	if data, err := json.Marshal(&user); err != nil {
		log.Printf("marshal error: %v", err)
		return 0, err
	} else if err := bucket.Put(intToByte(int(id)), data); err != nil {
		log.Printf("put error: %v", err)
		return 0, err
	}

	return id, tx.Commit()
}

func (usp *UserServiceProvider) Create(name string, payload int) error {
	var (
		payloadByte []byte
		err         error
	)

	tx, err := UserDB.Begin(true)
	if err != nil {
		log.Printf("[create] begin txn error: %v", err)
		return err
	}
	defer tx.Rollback()

	b := tx.Bucket([]byte("user"))

	if payloadByte, err = json.Marshal(&payload); err != nil {
		log.Printf("[create] marshal error: %v", err)
		return err
	}

	nameByte, err := json.Marshal(&name)
	if err != nil {
		log.Printf("[create] marshal error: %v", err)
		return err
	}

	err = b.Put(nameByte, payloadByte)
	if err != nil {
		log.Printf("[create] put error: %v", err)
		return err
	}

	return tx.Commit()
}

func (usp *UserServiceProvider) GetOne(id uint64) (*User, error) {
	tx, err := UserDB.Begin(false)
	if err != nil {
		log.Printf("[get] begin txn error: %v", err)
		return nil, err
	}
	defer tx.Rollback()

	var a User

	if v := tx.Bucket([]byte("user")).Get(intToByte(int(id))); v == nil {
		log.Print("get no record")
		return nil, nil
	} else if err := json.Unmarshal(v, &a); err != nil {
		log.Printf("unmarshal error: %v", err)
		return nil, err
	}

	return &a, nil
}

func (usp *UserServiceProvider) Get(name string) (int, error) {
	var (
		payload int
	)

	tx, err := UserDB.Begin(false)
	if err != nil {
		log.Printf("[get] begin txn error: %v", err)
		return 0, err
	}
	defer tx.Rollback()

	nameByte, err := json.Marshal(&name)
	if err != nil {
		log.Printf("[get] marshal error: %v", err)
		return 0, err
	}

	if v := tx.Bucket([]byte("user")).Get(nameByte); v == nil {
		log.Print("[get] return nil value")
		return 0, nil
	} else if err := json.Unmarshal(v, &payload); err != nil {
		log.Printf("[get] unmarshal error: %v", err)
		return 0, err
	}

	return payload, nil
}

func intToByte(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
