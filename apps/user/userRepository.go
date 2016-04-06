package user

import(
    "linq/core/utils"
    . "linq/core/repository"
    . "linq/core/database"
)

type userRepository struct{
    countQuery string
}

func NewUserRepository() IRepository{
    return userRepository{
        countQuery : "SELECT COUNT(*) FROM users",
    }
}

func (repo userRepository) CountAll() int{
    var result int
    rows := DB.Resolve(repo.countQuery)
    
    for rows.Next() {  
        err := rows.Scan(&result)
        utils.HandleWarn(err)
    }
    
    utils.HandleWarn(rows.Err())
    return result
}

func (repo userRepository) GetAll() []IModel{
    var result = Users{}
    rows := DB.Resolve("select uid, username, password, email, last_login from users")
    
    for rows.Next() {  
        var user = User{}
        err := rows.Scan(&user.Uid, &user.Username, &user.Password, &user.Email, &user.LastLogin)
        utils.HandleWarn(err)
        result = append(result, user)
    }
    
    utils.HandleWarn(rows.Err())
    
    return result
}