package postgres

import (
	"app/internal/entity"
	"context"
	"gorm.io/gorm"
	"time"
)

func NewUserRepo() (*UserRepo, error) {

	conn, err := getHandle()
	if err != nil {
		return nil, err
	}
	return &UserRepo{
		Conn: conn,
	}, nil
}

type UserRepo struct {
	Conn *gorm.DB
}

func (ur *UserRepo) GetOne(ctx context.Context, uid int64) (*entity.User, error) {

	data := &entity.User{}
	err := ur.Conn.First(data, uid).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (ur *UserRepo) GetMany(ctx context.Context, id []int64) ([]entity.User, error) {

	var data []entity.User
	err := ur.Conn.Where("id IN (?)", id).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (ur *UserRepo) Create(ctx context.Context, user *entity.User) error {

	user.CreatedAt = time.Now().Unix()
	return ur.Conn.Create(&user).Error
}

func (ur *UserRepo) Update(ctx context.Context, user *entity.User) error {

	_, err := ur.GetOne(ctx, user.Uid)
	if err != nil {
		return err
	}

	return ur.Conn.Save(user).Error
}

func (ur *UserRepo) Delete(ctx context.Context, id int64) error {

	return ur.Conn.Delete(&entity.User{}, id).Error
}
