package dao

import "douban/model"

func InsertP(people model.People) error {
	str:="INSERT INTO people(name,birthdate, birthplace, job, imdb) "+"values(?, ?, ?, ?, ?);"
	_, err := dB.Exec(str, people.Name,people.Birthdate,people.Birthplace,people.Job,people.IMDb)
	return err
}

func QueryP(peopleid int)([]model.People,error) {
	var Peoples []model.People
	rows, err := dB.Query("SELECT  name ,birthdate, birthplace, job, imdb FROM people WHERE id LIKE ?", peopleid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var people model.People
		err := rows.Scan(&people.Name,&people.Birthplace,&people.Birthdate,&people.Job,&people.IMDb)
		if err != nil {
			return nil, err
		}
		Peoples = append(Peoples, people)
	}
	return Peoples, nil
}
