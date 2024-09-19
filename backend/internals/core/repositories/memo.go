package repositories

import (
	"back-end/database"
	"back-end/internals/core/domain"
	"back-end/internals/core/helpers"
	"errors"
	"time"

	"gorm.io/gorm"
)

type MemoRepository struct {
	conn  *gorm.DB
	model domain.Memo
}

func NewMemoRepository() *MemoRepository {
	return &MemoRepository{
		conn:  database.DB.Conn,
		model: domain.MemoModel,
	}
}

func (r *MemoRepository) GetModel() domain.Memo {
	return r.model
}

func (r *MemoRepository) Create(createInput helpers.CreateMemoPayload, referOption ...helpers.Refer) error {
	var refer helpers.Refer

	if len(referOption) > 0 {
		refer = referOption[0]
	}

	memo := r.GetModel()
	memo.Content = createInput.Content
	memo.Title = createInput.Title
	memo.UserRefer = refer.UserId()

	if err := r.conn.Create(&memo).Error; err != nil {
		return err
	}
	return nil
}

func (r *MemoRepository) GetById(id int) (*domain.Memo, error) {

	if err := r.conn.Where("id = ?", id).First(&r.model).Error; err != nil {
		return nil, err
	}

	return &r.model, nil
}

func (r *MemoRepository) GetAll(referOption ...helpers.Refer) ([]*domain.Memo, error) {
	refer := helpers.NewRefer()

	if len(referOption) > 0 {
		*refer = referOption[0]
	} else {
		return nil, errors.New("could not get user ID")
	}

	var memos []*domain.Memo
	if err := r.conn.Where("user_refer = ?", refer.UserId()).Find(&memos).Error; err != nil {
		return nil, err
	}

	// if err := r.conn.Find(&memos).Error; err != nil {
	// 	return nil, err
	// }
	// return memos, nil
	return memos, nil
}

func (r *MemoRepository) Update(id int, updateUser helpers.UpdateMemoPayload) (*domain.Memo, error) {

	memo, err := r.GetById(id)
	if err != nil {
		return nil, err
	}

	if updateUser.Title != "" {
		memo.Title = updateUser.Title
		memo.UpdatedAt = time.Now()
	}

	if updateUser.Content != "" {
		memo.Content = updateUser.Content
		memo.UpdatedAt = time.Now()
	}

	err = r.conn.Save(&memo).Error
	if err != nil {
		return nil, err
	}

	return memo, nil
}

func (r *MemoRepository) Delete(id int) error {

	memo, err := r.GetById(id)
	if err != nil {
		return err
	}

	err = r.conn.Delete(&memo).Error
	if err != nil {
		return err
	}
	return nil
}
