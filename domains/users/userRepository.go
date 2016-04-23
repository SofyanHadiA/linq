package users

import (
	"fmt"
	
	. "linq/core/database"
	. "linq/core/repository"
	"linq/core/utils"

	"github.com/jmoiron/sqlx"
    "github.com/satori/go.uuid"
)

type userRepository struct {
	db             	IDB
	countQuery    	string
	isExistQuery 	string
	selectAllQuery 	string
	selectQuery    	string
	deleteQuery		string
	deleteBulkQuery		string
}

func UserRepository(db IDB) IRepository {
	return userRepository{
		db:             db,
		countQuery:     "SELECT COUNT(*) FROM users",
		isExistQuery:	"SELECT EXISTS(SELECT * FROM users WHERE uid=?)",
	}
}

func (repo userRepository) CountAll() int {
	var result int
	err := repo.db.ResolveSingle(repo.countQuery).Scan(&result)
	utils.HandleWarn(err)
	return result
}

func (repo userRepository) IsExist(id uuid.UUID) bool {
	var result bool
	err := repo.db.ResolveSingle(repo.isExistQuery, id).Scan(&result)
	utils.HandleWarn(err)
	return result
}

func (repo userRepository) GetAll(keyword string, length int, order int, orderDir string) IModels {
	query:= "SELECT * FROM users WHERE deleted=0 "
	
	if(keyword!=""){
		query += ` AND (username LIKE '%?%' OR email LIKE '%?%' OR first_name LIKE '%?%' OR last_name LIKE '%?%') `
	}
	
	if(order > 0){
		var columnMap string
		
		switch order{
			case 0 : columnMap = "uid"
			case 1 : columnMap = "username"
			case 2 : columnMap = "email"
			case 3 : columnMap = "first_name"
			default : columnMap = "username"
		}
		
		query += fmt.Sprintf(" ORDER BY %s %s ", columnMap, orderDir)
	}
	
	if(length>0){
		query += fmt.Sprintf(" LIMIT %d ", length)
	}else{
		query += " LIMIT 25 "
	}
	
	rows := &sqlx.Rows{}
	
	if(keyword!=""){
		rows = repo.db.Resolve(query, keyword)
	}else{
		rows = repo.db.Resolve(query)
	}

	result := Users{}
	
	for rows.Next() {
		var user = &User{}
		err := rows.StructScan(&user)
		utils.HandleWarn(err)
		result = append(result, (*user))
	}
	utils.HandleWarn(rows.Err())

	return &result
}

func (repo userRepository) Get(id uuid.UUID) IModel {
	selectQuery := "SELECT * FROM users WHERE uid = ? AND deleted=0 "

	user := &User{}
	err := repo.db.ResolveSingle(selectQuery, id).StructScan(user)
	
	utils.HandleWarn(err)

	return user
}

func (repo userRepository) Insert(model IModel) IModel {
	insertQuery := `INSERT INTO users 
		(uid, username, email, first_name, last_name, password, phone_number, address, country, city, state, zip ) 
		VALUES(:uid, :username, :email, :first_name, :last_name, :password, :phone_number, :address, :country, :city, :state, :zip)`
	
	user, _ := model.(*User)
	user.Uid = uuid.NewV4()
	
	result := repo.db.Execute(insertQuery, user)
	
	utils.Log.Info("-----", result)
	
	return user
}

func (repo userRepository) Update(model IModel) IModel {
	updateQuery := `UPDATE users SET username=:username, email=:email, first_name=:first_name, last_name=:last_name, password=:password, phone_number=:phone_number,
		address=:address, country=:country, city=:city, state=:state, zip=:zip WHERE uid=:uid`
	
	user, _ := model.(*User)
	
	repo.db.Execute(updateQuery, user)
		
	return user
}

func (repo userRepository) Delete(model IModel) IModel {
	deleteQuery := "UPDATE users SET deleted=1 WHERE uid=:uid"
	
	user, _ := model.(*User)
	
	repo.db.Execute(deleteQuery, user)
	return model
}

func (repo userRepository) DeleteBulk(users []uuid.UUID) error {
	deleteQuery := "UPDATE users SET deleted=1 WHERE uid IN"
	
	inClause := "("
	i:= 0
	for _, user := range users {
         if(repo.IsExist(user)){
         	if(i>0){
         		inClause += ", "
         	}
         	inClause += fmt.Sprintf("'%s'", user)
         	i++
         }
    }
	inClause += ")"
	
	repo.db.Execute(deleteQuery + inClause, nil)
	return nil
}
