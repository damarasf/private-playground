package repository

import (
	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/db"
)

type UserRepository struct {
	db db.DB
}

func NewUserRepository(db db.DB) UserRepository {
	return UserRepository{db}
}

func (u *UserRepository) LoadOrCreate() ([]User, error) {
	// records, err := u.db.Load("users")
	// if err != nil {
	// 	records = [][]string{
	// 		{"username", "password", "status"},
	// 	}
	// 	if err := u.db.Save("users", records); err != nil {
	// 		return nil, err
	// 	}
	// }

	// result := make([]User, 0)
	// for i := 1; i < len(records); i++ {
	// 	user := User{
	// 		Username: records[i][0],
	// 		Password: records[i][1],
	// 		Loggedin: records[i][2] == "true",
	// 	}
	// 	result = append(result, user)
	// }

	return nil, nil
	// return result, nil
}

func (u *UserRepository) SelectAll() ([]User, error) {
	users, err := u.LoadOrCreate()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u UserRepository) Login(username string, password string) (*string, error) {
	users, err := u.LoadOrCreate()
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(users); i++ {
		if users[i].Username == username && users[i].Password == password {
			users[i].Loggedin = true
			if err := u.Save(users); err != nil {
				return nil, err
			}
			return &users[i].Username, nil
		}
	}

	return nil, nil
}

func (u *UserRepository) FindLoggedinUser() (*string, error) {
	users, err := u.LoadOrCreate()
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(users); i++ {
		if users[i].Loggedin {
			return &users[i].Username, nil
		}
	}

	return nil, nil
}

func (u *UserRepository) Logout(username string) error {
	return nil // TODO: replace this
}

func (u *UserRepository) Save(users []User) error {
	return nil // TODO: replace this
}

func (u *UserRepository) changeStatus(username string, status bool) error {
	users, err := u.LoadOrCreate()
	if err != nil {
		return err
	}

	for i := 0; i < len(users); i++ {
		if users[i].Username == username {
			users[i].Loggedin = status
			if err := u.Save(users); err != nil {
				return err
			}
			return nil
		}
	}

	return nil
}

func (u *UserRepository) LogoutAll() error {
	return nil // TODO: replace this
}
