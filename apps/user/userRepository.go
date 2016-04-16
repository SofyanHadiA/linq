package user

import (
	. "linq/core/database"
	. "linq/core/repository"
	"linq/core/utils"
)

type userRepository struct {
	db             	IDB
	countQuery    	string
	selectAllQuery 	string
	selectQuery    	string
	insertQuery	   	string
	updateQuery		string
}

func UserRepository(db IDB) IRepository {
	return userRepository{
		db:             db,
		countQuery:     "SELECT COUNT(*) FROM users",
		selectAllQuery: "SELECT uid, username, password, email, last_login FROM users",
		selectQuery:    "SELECT uid, username, password, email, last_login FROM users WHERE uid = ?",
		insertQuery:	"INSERT INTO users (username, password, email) VALUES(?, ?, ?)",
		updateQuery:	"UPDATE users set username=?, password=?, email=? WHERE uid=?",
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

func (repo userRepository) GetAll() []IModel {
	var result = Users{}
	rows := repo.db.Resolve(repo.selectAllQuery)

	for rows.Next() {
		var user = User{}
		err := rows.Scan(&user.Uid, &user.Username, &user.Password, &user.Email, &user.LastLogin)
		utils.HandleWarn(err)
		result = append(result, user)
	}

	utils.HandleWarn(rows.Err())

	if len(result) > 0 {
		return result
	} else {
		return nil
	}
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

