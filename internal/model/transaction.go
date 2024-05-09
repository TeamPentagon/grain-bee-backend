package model

import (
	"log"

	"github.com/TeamPentagon/grain-bee-backend/internal/database"
	"google.golang.org/protobuf/proto"
	//"time"
)

func (s *Transaction) SaveTransaction() error {
	db, err := database.CreateLevelDBDatabase("Database/", "Common", "Shard_0.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	bytes, err := proto.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Put([]byte(s.TransactionId), bytes, nil)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (g *Transaction) GetTransaction() error {
	db, err := database.CreateLevelDBDatabase("Database/", "Common", "Shard_0.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	data, err := db.Get([]byte(g.TransactionId), nil)
	if err != nil {
		log.Fatal(err)
	}
	err = proto.Unmarshal(data, g)
	if err != nil {
		log.Fatal(err)
	}
	return nil

}

func (u *Transaction) UpdateTransaction() error {
	db, err := database.CreateLevelDBDatabase("Database/", "Common", "Shard_0.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	data, err := db.Get([]byte(u.TransactionId), nil)
	if err != nil {
		log.Fatal(err)
	}
	err = proto.Unmarshal(data, u)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (d *Transaction) DeleteTransaction() error {
	db, err := database.CreateLevelDBDatabase("Database/", "Common", "Shard_0.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	data, err := db.Get([]byte(d.TransactionId), nil)
	if err != nil {
		log.Fatal(err)
	}
	err = proto.Unmarshal(data, d)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
