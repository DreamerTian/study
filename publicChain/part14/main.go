package main

import (
	"fmt"
	"study/publicChain/part14/BLC"
	"github.com/boltdb/bolt"
	"log"
)

func main() {

	//block := BLC.NewBlock("Test",1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
	//
	//fmt.Printf("%d\n",block.Nonce)
	//fmt.Printf("%x\n",block.Hash)

	//创建或者打开数据库
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//err = db.Update(func(tx *bolt.Tx) error {
	//	//去找blocks的表
	//	b := tx.Bucket([]byte("blocks"))
	//	//如果没有就创建
	//	if b == nil{
	//		b,err := tx.CreateBucket([]byte("blocks"))
	//		if err != nil {
	//			log.Panic("Blocks table create error......")
	//		}
	//		err = b.Put([]byte("l"),block.Serialize())
	//		if err != nil {
	//			log.Panic("存储失败")
	//		}
	//	}
	//	// 返回nil,以便数据库处理相应操作
	//	return nil
	//})
	////更新失败
	//if err != nil {
	//	log.Panic(err)
	//}

	err = db.View(func(tx *bolt.Tx) error {
		//去找blocks的表
		b := tx.Bucket([]byte("blocks"))

		blockData := b.Get([]byte("l"))

		fmt.Printf("%v",BLC.DeserializeBlock(blockData))

		// 返回nil,以便数据库处理相应操作
		return nil
	})

	//更新失败
	if err != nil {
		log.Panic(err)
	}
}
