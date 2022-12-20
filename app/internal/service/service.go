package service

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return Service{db: db}
}

//Реализация методов обращения в базу данных

type Orderreservations struct {
	OrderId uuid.UUID `json:"orderId" gorm:"type:uuid; not null"`
	GoodId  uuid.UUID `json:"goodId" gorm:"type:uuid; not null"`
}

type Good struct {
	Id    uuid.UUID `json:"id" gorm:"type:uuid; unique; primary_key;"`
	Name  string    `json:"name" gorm:"type:string;"`
	Price int       `json:"price"`
}

func (s *Service) Reserve(orderId uuid.UUID, goods []uuid.UUID) error {
	for _, goodId := range goods {
		err := s.db.Create(Orderreservations{
			OrderId: orderId,
			GoodId:  goodId,
		}).Error

		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) CancelReservation(orderId uuid.UUID) error {
	return s.db.Delete(&Orderreservations{}, orderId).Error
}

func (s *Service) Get(orderId uuid.UUID) (reservation *Orderreservations, err error) {
	err = s.db.Model(reservation).Where(orderId).First(&reservation).Error
	return
}
