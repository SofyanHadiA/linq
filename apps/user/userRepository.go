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
	insertQuery		string
	updateQuery		string
	deleteQuery		string
	deleteBulkQuery		string
}

func UserRepository(db IDB) IRepository {
	return userRepository{
		db:             db,
		countQuery:     "SELECT COUNT(*) FROM users",
		isExistQuery:	"SELECT EXISTS(SELECT * FROM users WHERE uid=?)",
		selectAllQuery: "SELECT uid, username, email, first_name, last_name, last_login FROM users",
		selectQuery:    "SELECT uid, username, email, first_name, last_name, last_login FROM users WHERE uid = ?",
		insertQuery:    "INSERT INTO users (uid, username, email, first_name, last_name, password ) VALUES(?, ?, ?, ?, ?, ?)",
		updateQuery:	"UPDATE users SET username=?, email=?, first_name=?, last_name=?, password=? WHERE uid=?",
		deleteQuery:	"DELETE FROM users WHERE uid=?",
		deleteBulkQuery:"DELETE FROM users WHERE uid in ",
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

func (repo userRepository) IsExist(id string) bool {
	var result bool
	err := repo.db.ResolveSingle(repo.isExistQuery, id).Scan(&result)
	utils.HandleWarn(err)
	return result
}

func (repo userRepository) GetAll(keyword string, length int, order int, orderDir string) []IModel {
	query:= repo.selectAllQuery
	
	if(keyword!=""){
		query += fmt.Sprintf(" WHERE username LIKE '%%%s%%' OR email LIKE '%%%s%%'", keyword)
	}
	
	if(order > 0){
		var columnMap = map[int]string{
			0 : "email",
			1 : "email",
			2 : "first_name",
		}
		
		query += fmt.Sprintf(" ORDER BY %s %s ", columnMap[order], orderDir)
	}
	
	if(length>0){
		query += fmt.Sprintf(" LIMIT %d ", length)
	}else{
		query += " LIMIT 25 "
	}
	
	rows := repo.db.Resolve(query)

	result := Users{}
	
	for rows.Next() {
		var user = &User{}
		err := rows.Scan(&user.Uid, &user.Username, &user.Email, &user.FirstName, &user.LastName, &user.LastLogin)
		utils.HandleWarn(err)
		result = append(result, user)
	}
	utils.HandleWarn(rows.Err())
	
	return result
}

func (repo userRepository) Get(id string) IModel {
	user := &User{}
	rows := repo.db.Resolve(repo.selectQuery, id)

	for rows.Next() {
		err := rows.Scan(&user.Uid, &user.Username, &user.Email, &user.FirstName, &user.LastName, &user.LastLogin)
		utils.HandleWarn(err)
	}

	utils.HandleWarn(rows.Err())

	return user
}

func (repo userRepository) Insert(user IModel) IModel {
	repo.db.Execute(repo.insertQuery, user.InsertVal()...)
	return user
}

func (repo userRepository) Update(model IModel) IModel {
	repo.db.Execute(repo.updateQuery, model.UpdateVal()...)
	return model
}

func (repo userRepository) Delete(model IModel) IModel {
	repo.db.Execute(repo.deleteQuery, model.GetId())
	return model
}

func (repo userRepository) DeleteBulk(users []string) error {
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
	
	println(repo.deleteBulkQuery + inClause)
	
	repo.db.Execute(repo.deleteBulkQuery + inClause)
	return nil
}
