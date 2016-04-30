package users

import (
	"fmt"

	. "linq/core/database"
	. "linq/core/repository"
	"linq/core/utils"

	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
)

type UserRepository struct {
	db IDB
}

func NewUserRepository(db IDB) IRepository {
	return UserRepository{
		db: db,
	}
}

func (repo UserRepository) CountAll() (int, error) {
	countQuery := "SELECT COUNT(*) FROM users WHERE deleted = 0"

	var result int
	row, err := repo.db.ResolveSingle(countQuery)
	row.Scan(&result)
	utils.HandleWarn(err)
	return result, err
}

func (repo UserRepository) IsExist(id uuid.UUID) (bool, error) {
	isExistQuery := "SELECT EXISTS(SELECT * FROM users WHERE uid=? AND deleted = 0)"

	var result bool
	row, err := repo.db.ResolveSingle(isExistQuery, id)
	row.Scan(&result)
	return result, err
}

func (repo UserRepository) GetAll(paging utils.Paging) (IModels, error) {
	query := "SELECT * FROM users WHERE deleted=0 "

	if paging.Keyword != "" {
		query += ` AND (username LIKE '%?%' OR email LIKE '%?%' OR first_name LIKE '%?%' OR last_name LIKE '%?%') `
	}

	if paging.Order > 0 {
		var columnMap string

		switch paging.Order {
		case 0:
			columnMap = "uid"
		case 1:
			columnMap = "username"
		case 2:
			columnMap = "email"
		case 3:
			columnMap = "first_name"
		default:
			columnMap = "username"
		}

		query += fmt.Sprintf(" ORDER BY %s %s ", columnMap, paging.OrderDir)
	}

	if paging.Length > 0 {
		query += fmt.Sprintf(" LIMIT %d ", paging.Length)
	} else {
		query += " LIMIT 25 "
	}

	rows := &sqlx.Rows{}
	var err error

	if paging.Keyword != "" {
		rows, err = repo.db.Resolve(query, paging.Keyword)
	} else {
		rows, err = repo.db.Resolve(query)
	}
	utils.HandleWarn(err)

	result := Users{}

	for rows.Next() {
		var user = &User{}
		err := rows.StructScan(&user)
		utils.HandleWarn(err)
		result = append(result, (*user))
	}

	return &result, err
}

func (repo UserRepository) Get(id uuid.UUID) (IModel, error) {
	selectQuery := "SELECT * FROM users WHERE uid = ? AND deleted=0 "

	user := &User{}
	rows, err := repo.db.ResolveSingle(selectQuery, id)
	utils.HandleWarn(err)
	rows.StructScan(user)

	return user, err
}

func (repo UserRepository) Insert(model IModel) error {
	insertQuery := `INSERT INTO users 
		(uid, username, email, first_name, last_name, phone_number, address, country, city, state, zip ) 
		VALUES(:uid, :username, :email, :first_name, :last_name, :phone_number, :address, :country, :city, :state, :zip)`

	user, _ := model.(*User)
	user.Uid = uuid.NewV4()

	_, err := repo.db.Execute(insertQuery, user)

	return err
}

func (repo UserRepository) Update(model IModel) error {
	updateQuery := `UPDATE users SET username=:username, email=:email, first_name=:first_name, last_name=:last_name, phone_number=:phone_number,
		address=:address, country=:country, city=:city, state=:state, zip=:zip WHERE uid=:uid`

	user, _ := model.(*User)

	_, err := repo.db.Execute(updateQuery, user)

	return err
}

func (repo UserRepository) UpdateUserPhoto(model IModel) error {
	updateQuery := `UPDATE users SET avatar=:avatar WHERE uid=:uid`

	user, _ := model.(*User)

	_, err := repo.db.Execute(updateQuery, user)

	return err
}

func (repo UserRepository) Delete(model IModel) error {
	deleteQuery := "UPDATE users SET deleted=1 WHERE uid=:uid"

	user, _ := model.(*User)
	_, err := repo.db.Execute(deleteQuery, user)

	return err
}

func (repo UserRepository) DeleteBulk(users []uuid.UUID) error {
	deleteQuery := "UPDATE users SET deleted=1 WHERE uid IN(?)"
	_, err := repo.db.ExecuteBulk(deleteQuery, users)

	return err
}

func (repo UserRepository) ValidatePassword(uid uuid.UUID, password string) (bool, error) {
	isValidPasswordQuery := "SELECT EXISTS(SELECT * FROM users WHERE uid=? AND password=?)"

	var result bool
	row, err := repo.db.ResolveSingle(isValidPasswordQuery, uid.String(), password)
	row.Scan(&result)

	return result, err
}

func (repo UserRepository) ChangePassword(uid uuid.UUID, password string) error {
	updatePasswordQuery := "UPDATE users SET password=? WHERE uid=?"

	_, err := repo.db.ExecuteArgs(updatePasswordQuery, password, uid.String())

	return err
}
