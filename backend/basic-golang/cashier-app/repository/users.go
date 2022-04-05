package repository

import (
	"errors"

	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/db"
)

type UserRepository struct {
	db db.DB
}

func NewUserRepository(db db.DB) UserRepository {
	return UserRepository{db}
}

func (u *UserRepository) LoadOrCreate() ([]User, error) {
	records, err := u.db.Load("users")
	if err != nil {
		records = [][]string{
			{"username", "password", "loggedin"},
		}
		if err := u.db.Save("users", records); err != nil {
			return nil, err
		}
	}

	result := make([]User, 0)
	for i := 1; i < len(records); i++ {
		user := User{
			Username: records[i][0],
			Password: records[i][1],
			Loggedin: records[i][2] == "true",
		}
		result = append(result, user)
	}

	return result, nil
}

func (u *UserRepository) SelectAll() ([]User, error) {
	return u.LoadOrCreate()
}

func (u UserRepository) Login(username string, password string) (*string, error) {
	if err := u.LogoutAll(); err != nil {
		return nil, err
	}

	users, err := u.LoadOrCreate()
	if err != nil {
		return nil, err
	}

	for _, users := range users {
		if users.Username == username && users.Password == password {
			if err := u.changeStatus(username, true); err != nil {
				return nil, err
			}
			return &users.Username, nil

		}
	}

	return nil, errors.New("Login Failed")
}

func (u *UserRepository) FindLoggedinUser() (*string, error) {
	users, err := u.SelectAll()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Loggedin {
			return &user.Username, nil
		}
	}

	return nil, errors.New("no user is logged in")
}

func (u *UserRepository) Logout(username string) error {
	users, err := u.FindLoggedinUser()
	if err != nil {
		return err
	}

	if *users == username {
		if err := u.changeStatus(username, false); err != nil {
			return err
		}
	}

	return nil
}

func (u *UserRepository) Save(users []User) error {
	record := [][]string{
		{"username", "password", "loggedin"},
	}

	for _, user := range users {
		newRow := []string{
			user.Username, user.Password,
		}

		if user.Loggedin == true {
			newRow = append(newRow, "true")
		} else {
			newRow = append(newRow, "false")
		}

		record = append(record, newRow)
	}

	return u.db.Save("users", record)
}

func (u *UserRepository) changeStatus(username string, status bool) error {
	users, err := u.LoadOrCreate()
	if err != nil {
		return err
	}

	for i := 0; i < len(users); i++ {
		if users[i].Username == username {
			users[i].Loggedin = true
		}
	}

	if err := u.Save(users); err != nil {
		return err
	}

	return u.Save(users)
}

func (u *UserRepository) LogoutAll() error {
	users, err := u.LoadOrCreate()
	if err != nil {
		return err
	}

	for i := 0; i < len(users); i++ {
		users[i].Loggedin = false
	}

	if err := u.Save(users); err != nil {
		return err
	}

	return u.Save(users)
}
