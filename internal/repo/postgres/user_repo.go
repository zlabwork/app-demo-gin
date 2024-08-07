package postgres

import (
	"app/internal/entity"
	"app/internal/help"
	"context"
	"gorm.io/gorm"
	"time"
)

func NewUserRepo() (*UserRepo, error) {

	conn := help.Db
	return &UserRepo{
		Conn: conn,
	}, nil
}

type UserRepo struct {
	Conn *gorm.DB
}

func (ur *UserRepo) GetOne(ctx context.Context, uid int64) (*entity.User, error) {

	data := &entity.User{}
	err := ur.Conn.Where("uid = ?", uid).First(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (ur *UserRepo) GetMany(ctx context.Context, uid []int64) ([]entity.User, error) {

	var data []entity.User
	err := ur.Conn.Where("uid IN (?)", uid).Find(&data).Error
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

	// 更新非零值的字段
	user.UpdatedAt = time.Now().Unix()
	return ur.Conn.Model(&user).Where("uid = ?", user.Uid).Updates(user).Limit(1).Error
}

func (ur *UserRepo) Delete(ctx context.Context, uid int64) error {

	return ur.Conn.Where("uid = ?", uid).Delete(&entity.User{}).Error
}
