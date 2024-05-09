package model

import (
	"log"

	"github.com/TeamPentagon/grain-bee-backend/internal/database"
	"google.golang.org/protobuf/proto"
	//"time"
)

func (s *Admin) SaveAdmin() error {
	db, err := database.CreateLevelDBDatabase("Database/", "Common", "Shard_0.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	bytes, err := proto.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Put([]byte(s.AdminId), bytes, nil)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (g *Admin) GetAdmin() error {
	db, err := database.CreateLevelDBDatabase("Database/", "Common", "Shard_0.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	data, err := db.Get([]byte(g.AdminId), nil)
	if err != nil {
		log.Fatal(err)
	}
	err = proto.Unmarshal(data, g)
	if err != nil {
		log.Fatal(err)
	}
	return nil

}

func (u *Admin) UpdateAdmin() error {
	db, err := database.CreateLevelDBDatabase("Database/", "Common", "Shard_0.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	data, err := db.Get([]byte(u.AdminId), nil)
	if err != nil {
		log.Fatal(err)
	}
	err = proto.Unmarshal(data, u)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (d *Admin) DeleteAdmin() error {
	db, err := database.CreateLevelDBDatabase("Database/", "Common", "Shard_0.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	data, err := db.Get([]byte(d.AdminId), nil)
	if err != nil {
		log.Fatal(err)
	}
	err = proto.Unmarshal(data, d)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
