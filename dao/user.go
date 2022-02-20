package dao

import "douban/model"

func UpdatePassword(username, newPassword string) error {
	_, err := dB.Exec("UPDATE user SET password = ? WHERE username = ?", newPassword, username)
	return err
}

func SelectUserByUsername(username string) (model.User, error) {
	user := model.User{}

	row := dB.QueryRow("SELECT id, password  FROM user WHERE username = ? ", username)
	if row.Err() != nil {
		return user, row.Err()
	}

	err := row.Scan(&user.Id, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}

func InsertUser(user model.User) error {
	str:="INSERT INTO user(username, password,time) "+"values(?, ?,?);"
	_, err := dB.Exec(str, user.Username, user.Password,user.Time)
	return err
}

type Miniuser struct {
	Id int
	Username string
	Time string
	Txt string
}

func QueryU(userid int)([]Miniuser,error) {
	var Users []Miniuser
	rows, err := dB.Query("SELECT  id ,username, time, txt FROM user WHERE id LIKE ?", userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user Miniuser
		err := rows.Scan(&user.Id,&user.Username,&user.Time,&user.Txt)
		if err != nil {
			return nil, err
		}
		Users = append(Users, user)
	}
	return Users, nil
}

func UpdateTxt(txt string,username string) error {
	_, err := dB.Exec("UPDATE user SET txt = ? WHERE username = ?", txt,username)
	return err
}
