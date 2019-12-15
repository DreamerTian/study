package main

import (
	"github.com/boltdb/bolt"
	"fmt"
	"log"
)

func main() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.

	//创建或者打开数据库
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//创建表
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("BlockBucket"))
		if b != nil {
			data := b.Get([]byte("l"))
			fmt.Printf("%s",data)
		}
		// 返回nil,以便数据库处理相应操作
		return nil
	})
	//更新失败
	if err != nil {
		log.Panic(err)
	}
}
