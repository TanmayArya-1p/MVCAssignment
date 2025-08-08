package models

import (
	"database/sql"
	"errors"
	types "inorder/pkg/types"
	utils "inorder/pkg/utils"
	"time"
)

func CreateUser(user *types.User) (types.UserID, error) {
	if user.Username == "" || user.Password == "" {
		return 0, errors.New("username and password cannot be empty")
	}

	if user.Role == "" {
		user.Role = types.UserRole
	}

	var hashedPassword string
	var err error
	hashedPassword, err = utils.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	user.HashedPassword = hashedPassword

	res, err := db.Exec("INSERT INTO users (username,password,role) VALUES (?,?,?)", user.Username, hashedPassword, user.Role)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	user.ID = types.UserID(id)
	return user.ID, nil
}

func GetUserByID(id types.UserID) (*types.User, error) {
	var user types.User
	var row *sql.Row

	row = db.QueryRow("SELECT id,username,password,role,created_at FROM users WHERE id = ?", id)
	var temp []uint8
	err := row.Scan(&user.ID, &user.Username, &user.HashedPassword, &user.Role, &temp)
	if err != nil {
		return nil, utils.ErrUserNotFound
	}
	user.CreatedAt, err = time.Parse(time.DateTime, string(temp))
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetAllUsers(page types.Page) ([]*types.User, error) {
	var rows *sql.Rows
	var err error
	rows, err = db.Query("SELECT id,username,password,role,created_at FROM users LIMIT ? OFFSET ?", page.Limit, page.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var otpt []*types.User = make([]*types.User, 0)
	var ifAny bool = rows.Next()
	if !ifAny {
		return otpt, nil
	}
	for {
		var user types.User
		var temp []uint8
		if err := rows.Scan(&user.ID, &user.Username, &user.HashedPassword, &user.Role, &temp); err != nil {
			if err == sql.ErrNoRows {
				break
			}
			return nil, err
		}

		var err error
		user.CreatedAt, err = time.Parse(time.DateTime, string(temp))
		if err != nil {
			return nil, err
		}
		otpt = append(otpt, &user)
		var isNext = rows.Next()
		if !isNext {
			break
		}
	}
	return otpt, nil
}

func GetUserByUsername(username string) (*types.User, error) {
	var user types.User
	var row *sql.Row

	row = db.QueryRow("SELECT id,username,password,role,created_at FROM users WHERE username = ?", username)

	var temp []uint8
	err := row.Scan(&user.ID, &user.Username, &user.HashedPassword, &user.Role, &temp)
	if err != nil {
		return nil, utils.ErrUserNotFound
	}
	user.CreatedAt, err = time.Parse(time.DateTime, string(temp))
	if err != nil {
		return nil, err
	}
	return &user, nil
}

type UserUpdateInstruction struct {
	User              *types.User
	Username          string
	PlaintextPassword string
	Role              types.Role
}

func UpdateUser(upd *UserUpdateInstruction) (*types.User, error) {
	if upd.User == nil {
		return nil, utils.ErrUserNotFound
	}
	if upd.Username != "" {
		upd.User.Username = upd.Username
	}
	if upd.PlaintextPassword != "" {
		upd.User.Password = upd.PlaintextPassword
	}
	if upd.Role != types.Role("") {
		upd.User.Role = upd.Role
	}

	if upd.PlaintextPassword != "" {
		var hashedPass string
		var err error

		hashedPass, err = utils.HashPassword(upd.PlaintextPassword)
		if err != nil {
			return nil, err
		}
		upd.User.HashedPassword = hashedPass
	}

	var err error
	if upd.PlaintextPassword != "" {
		_, err = db.Exec("UPDATE users SET username = ?, password = ?, role = ? WHERE id = ?", upd.User.Username, upd.User.HashedPassword, string(upd.User.Role), upd.User.ID)
	} else {
		_, err = db.Exec("UPDATE users SET username = ?, role = ? WHERE id = ?", upd.User.Username, string(upd.User.Role), upd.User.ID)
	}
	if err != nil {
		return nil, err
	}

	return upd.User, nil
}

func DeleteUserById(id types.UserID) error {
	var err error
	_, err = db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
