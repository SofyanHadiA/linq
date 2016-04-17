package user

import (
	"fmt"
	
	. "linq/core/database"
	. "linq/core/repository"
	"linq/core/utils"
)

type userRepository struct {
	db             	IDB
	countQuery    	string
	isExistQuery 	string
	selectAllQuery 	string
	selectQuery    	string
	insertQuery	   	string
	updateQuery		string
	deleteQuery		string
}

func UserRepository(db IDB) IRepository {
	return userRepository{
		db:             db,
		countQuery:     "SELECT COUNT(*) FROM users",
		isExistQuery:	"SELECT EXISTS(SELECT * FROM users WHERE uid=?)",
		selectAllQuery: "SELECT uid, username, password, email, last_login FROM users",
		selectQuery:    "SELECT uid, username, password, email, last_login FROM users WHERE uid = ?",
		insertQuery:	"INSERT INTO users (username, password, email) VALUES(?, ?, ?)",
		updateQuery:	"UPDATE users set username=?, password=?, email=? WHERE uid=?",
		deleteQuery:	"DELETE FROM users WHERE uid=?",
	}
}

func (repo userRepository) CountAll() int {
	var result int
	rows := repo.db.Resolve(repo.countQuery)

	for rows.Next() {
		err := rows.Scan(&result)
		utils.HandleWarn(err)
	}

	utils.HandleWarn(rows.Err())
	return result
}

func (repo userRepository) IsExist(id int) bool {
	var result bool
	err := repo.db.ResolveSingle(repo.isExistQuery, id).Scan(&result)
	utils.HandleWarn(err)
	return result
}

func (repo userRepository) GetAll(keyword string, order string, orderDir string) []IModel {
	var result = Users{}
	
	query:= repo.selectAllQuery
	
	if(keyword!=""){
		query = query + fmt.Sprintf(" WHERE username like '%%%s%%' ", keyword)
	}
	
	if(order!=""){
		query=query + fmt.Sprintf(" ORDER BY %s %s ", order, orderDir)
	}

	rows := repo.db.Resolve(query)

	for rows.Next() {
		var user = User{}
		err := rows.Scan(&user.Uid, &user.Username, &user.Password, &user.Email, &user.LastLogin)
		utils.HandleWarn(err)
		result = append(result, user)
	}
	utils.HandleWarn(rows.Err())
	
	return result
}

func (repo userRepository) Get(id int) IModel {
	user := User{Uid: -1}
	rows := repo.db.Resolve(repo.selectQuery, id)

	for rows.Next() {
		err := rows.Scan(&user.Uid, &user.Username, &user.Password, &user.Email, &user.LastLogin)
		utils.HandleWarn(err)
	}

	utils.HandleWarn(rows.Err())

	if user.Uid > 0 {
		return user
	} else {
		return nil
	}
}

func (repo userRepository) Insert(model IModel) IModel {
	repo.db.Execute(repo.insertQuery, model.InsertVal()...)
	return model
}

func (repo userRepository) Update(model IModel) IModel {
	repo.db.Execute(repo.updateQuery, model.UpdateVal()...)
	return model
}

func (repo userRepository) Delete(model IModel) IModel {
	repo.db.Execute(repo.deleteQuery, model.GetId())
	return model
}
