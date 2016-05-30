package users

import (
	"crypto/sha1"
	"errors"
	"fmt"

	"github.com/SofyanHadiA/linq/core"
	"github.com/SofyanHadiA/linq/core/utils"

	"github.com/satori/go.uuid"
)

type UserService struct {
	repo core.IRepository
}

func NewUserService(repo core.IRepository) UserService {
	return UserService{
		repo: repo,
	}
}

func (service UserService) CountAll() (int, error) {
	return service.repo.CountAll()
}

func (service UserService) IsExist(id uuid.UUID) (bool, error) {
	return service.repo.IsExist(id)
}

func (service UserService) GetAll(paging utils.Paging) (core.IModels, error) {
	return service.repo.GetAll(paging)
}

func (service UserService) Get(id uuid.UUID) (core.IModel, error) {
	return service.repo.Get(id)
}

func (service UserService) Create(model core.IModel) error {
	return service.repo.Insert(model)
}

func (service UserService) Modify(model core.IModel) error {
	if exist, _ := service.repo.IsExist(model.GetId()); exist {
		return service.repo.Update(model)
	} else {
		return userNotFoundError()
	}
}

func (service UserService) UpdateUserPhoto(model core.IModel) error {
	if exist, _ := service.repo.IsExist(model.GetId()); exist {
		userRepo := service.repo.(userRepository)
		err := userRepo.UpdateUserPhoto(model)
		return err
	} else {
		return userNotFoundError()
	}
}

func (service UserService) Remove(model core.IModel) error {
	if exist, _ := service.repo.IsExist(model.GetId()); exist {
		err := service.repo.Delete(model)

		return err
	} else {
		return userNotFoundError()
	}
}

func (service UserService) RemoveBulk(userIds []uuid.UUID) error {
	for _, id := range userIds {
		if exist, _ := service.repo.IsExist(id); !exist {
			return userNotFoundError()
		}
	}
	err := service.repo.DeleteBulk(userIds)
	return err
}

func (service UserService) ValidatePassword(uid uuid.UUID, password string) (bool, error) {
	userRepo := service.repo.(userRepository)
	return userRepo.ValidatePassword(uid, password)
}

func (service UserService) ChangePassword(userCred *UserCredential) error {
	userRepo := service.repo.(userRepository)

	if userCred.PasswordNew != userCred.PasswordConfirm {
		return errors.New("PasswordConfirmNotMatch")
	} else if len(userCred.PasswordNew) < 6 {
		return errors.New("PasswordTooShort")
	} else if valid, err := service.ValidatePassword(userCred.Uid, sha1ToString(sha1.Sum([]byte(userCred.Password)))); !valid || err != nil {
		return errors.New("WrongPassword")
	} else {
		return userRepo.ChangePassword(userCred.Uid, sha1ToString(sha1.Sum([]byte(userCred.PasswordNew))))
	}
}

func userNotFoundError() error {
	return errors.New("UserNotFound")
}

func sha1ToString(c [20]byte) string {
	return string(fmt.Sprintf("%x", c))
}
