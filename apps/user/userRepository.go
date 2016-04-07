package user

import(
    "linq/core/utils"
    . "linq/core/repository"
    . "linq/core/database"
)

type userRepository struct{
    db IDB
    countQuery string
    selectAllQuery string
    selectQuery string
}

func NewUserRepository(db IDB) IRepository{
    return userRepository{
        db : db,
        countQuery : "SELECT COUNT(*) FROM users",
        selectAllQuery : "SELECT uid, username, password, email, last_login FROM users",
        selectQuery : "SELECT uid, username, password, email, last_login FROM users WHERE uid = ?",
    }
}

func (repo userRepository) CountAll() int{
    var result int
    rows := repo.db.Resolve(repo.countQuery)
    
    for rows.Next() {  
        err := rows.Scan(&result)
        utils.HandleWarn(err)
    }
    
    utils.HandleWarn(rows.Err())
    return result
}

func (repo userRepository) GetAll() []IModel{
    var result = Users{}
    rows := repo.db.Resolve(repo.selectAllQuery)
    
    for rows.Next() {  
        var user = User{}
        err := rows.Scan(&user.Uid, &user.Username, &user.Password, &user.Email, &user.LastLogin)
        utils.HandleWarn(err)
        result = append(result, user)
    }
    
    utils.HandleWarn(rows.Err())
    
    if(len(result) > 0){
        return result
    }else{
        return nil
    }
}

func (repo userRepository) Get(id int) IModel{
    user := User{ Uid :-1 }
    rows := repo.db.Resolve(repo.selectQuery, id)
    
    for rows.Next() {  
        err := rows.Scan(&user.Uid, &user.Username, &user.Password, &user.Email, &user.LastLogin)
        utils.HandleWarn(err)
    }
    
    utils.HandleWarn(rows.Err())
    
    if(user.Uid > 0){
        return user
    }else{
        return nil
    }
}