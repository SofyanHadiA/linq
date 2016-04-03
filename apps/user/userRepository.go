package user

import(
    "linq/core/utils"
    . "linq/core/database"
)

func getAllUser() Users{
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