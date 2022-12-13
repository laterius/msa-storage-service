package mixtures

import (
	"github.com/ezn-go/mixture"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/google/uuid"
	"github.com/laterius/service_architecture_hw3/app/internal/service"
)

func (g Good) TableName() string {
	return "couriers"
}

type Good struct {
	Id    uuid.UUID `json:"id" gorm:"type:uuid; unique; primary_key;"`
	Name  string    `json:"name" gorm:"type:string;"`
	Price int       `json:"price"`
}

func init() {

	mx := &gormigrate.Migration{
		ID:       "0002",
		Migrate:  mixture.CreateTableM(&service.Good{}),
		Rollback: mixture.DropTableR(&service.Good{}),
	}

	goods := []Good{
		{Id: uuid.MustParse("497a545c-24df-4f9c-aab3-5a779c76a4f0"), Name: "good 1", Price: 100},
		{Id: uuid.MustParse("6b1641ae-3f10-4f4d-86a1-789731f24710"), Name: "good 2", Price: 200},
		{Id: uuid.MustParse("2077de75-a757-4f69-b5f3-2a52132afb52"), Name: "good 3", Price: 300},
	}

	mx = &gormigrate.Migration{
		ID:       "0002",
		Migrate:  mixture.CreateBatchM(goods),
		Rollback: mixture.DeleteBatchR(goods),
	}

	mixture.Add(mixture.ForAnyEnv, mx)
}
