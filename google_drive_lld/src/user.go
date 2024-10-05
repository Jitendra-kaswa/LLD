package src

type User struct {
	id    int
	name  string
	email string
}

func NewUser(id int, name string, email string) *User {
	return &User{
		id:    id,
		name:  name,
		email: email,
	}
}

func (u User) GetId() int {
	return u.id
}
func (u User) GetName() string {
	return u.name
}
func (u User) GetEmail() string {
	return u.email
}

type IUserRepository interface {
	GetUserById(id int) *User
	CreateUser(name, email string) *User
}

type UserRepositoryMap struct {
	userMap       map[int]*User
	currentUserId int
}

func NewUserRepositoryMap() IUserRepository {
	return &UserRepositoryMap{
		userMap:       make(map[int]*User),
		currentUserId: 1,
	}
}

func (urm UserRepositoryMap) GetUserById(id int) *User {
	return urm.userMap[id]
}

func (urm *UserRepositoryMap) CreateUser(name, email string) *User {
	newUser := NewUser(urm.currentUserId, name, email)
	urm.userMap[newUser.GetId()] = newUser
	urm.currentUserId++
	return newUser
}
