package model

import (
	"log"

	"github.com/TeamPentagon/grain-bee-backend/internal/database"
	"google.golang.org/protobuf/proto"
	//"time"
)

func (s *Product) SaveProduct() error {
	db, err := database.CreateLevelDBDatabase("Database/", "Common", "Shard_0.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	bytes, err := proto.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Put([]byte(s.ProductId), bytes, nil)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (g *Product) GetProduct() error {
	db, err := database.CreateLevelDBDatabase("Database/", "Common", "Shard_0.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	data, err := db.Get([]byte(userName), nil)
	if err != nil {
		log.Fatal(err)
	}
	err = proto.Unmarshal(data, g)
	if err != nil {
		log.Fatal(err)
	}
	return nil

}


func (u *Product) UpdateProduct() error {
	db, err := database.CreateLevelDBDatabase("Database/", "Common", "Shard_0.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	data, err := db.Get([]byte(UserName), nil)
	if err != nil {
		log.Fatal(err)
	}
	err = proto.Unmarshal(data, u)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (d *Product) DeleteProduct() error {
	db, err := database.CreateLevelDBDatabase("Database/", "Common", "Shard_0.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	data, err := db.Get([]byte(UserName), nil)
	if err != nil {
		log.Fatal(err)
	}
	err = proto.Unmarshal(data, d)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
