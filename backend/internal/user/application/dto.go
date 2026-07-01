package application

import domainUser "hobby-blog/internal/user/domain"

type AuthResult struct {
    User  *domainUser.User
    Token string
}
