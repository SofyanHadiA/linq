package users

import(
    "errors"
    
	"linq/core/repository"
	"linq/core/utils"

	"github.com/satori/go.uuid"
)

type userService struct {
    repo repository.IRepository
}

func UserService(repo repository.IRepository) userService{
    return userService{
        repo: repo,
    }
}

func (service userService) CountAll() (int, error) {
	return service.repo.CountAll()
}

func (service userService) IsExist(id uuid.UUID) (bool, error) {
    return service.repo.IsExist(id)
}

func (service userService) GetAll(paging utils.Paging) (repository.IModels, error) {
	return service.repo.GetAll(paging)
}

func (service userService) Get(id uuid.UUID) (repository.IModel, error) {
    return service.repo.Get(id)
}

func (service userService) Insert(model repository.IModel) error {
	return service.repo.Insert(model)
}

func (service userService) Update(model repository.IModel) error {
    if exist, _ := service.repo.IsExist(model.GetId()); exist {
	    return service.repo.Update(model)
    }else{
        return userNotFoundError()
    }
}

func (service userService) UpdateUserPhoto(model repository.IModel) error {
    if exist, _ := service.repo.IsExist(model.GetId()); exist {
        userRepo := service.repo.(UserRepository)
        
	    err := userRepo.UpdateUserPhoto(model)
	    
	    return err
    }else{
        return userNotFoundError()
    }
}

func (service userService) Delete(model repository.IModel) error {
    if exist, _ := service.repo.IsExist(model.GetId()); exist {
	    err :=  service.repo.Delete(model)   
	    
	    return err
    }else{
        return userNotFoundError()
    }
}

func (service userService) DeleteBulk(userIds []uuid.UUID) error{
    for _, id := range userIds {
        if exist, _ := service.repo.IsExist(id); exist {
            return userNotFoundError()
        }
    }
    
	 err := service.repo.DeleteBulk(userIds)
	 return err
}

func (service userService) ValidatePassword(model repository.IModel) (bool, error) {
    userRepo := service.repo.(UserRepository)

	return userRepo.ValidatePassword(model)
}

func (service userService) ChangePassword(model repository.IModel) error {
    userRepo := service.repo.(UserRepository)
    
	return userRepo.ChangePassword(model)
}

func userNotFoundError() error {
    return errors.New("UserNotFound")
}