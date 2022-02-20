package dao

import (
	"douban/model"
)


type Data1 struct {
	People []model.People
	Movie []model.Movie
}

func Searching(data string)(Data1,error) {
	var Movies []model.Movie
	var Peoples []model.People
	s:="%"+data+"%"
	str:="SELECT moviename, director, screenwriter,starring,style,area,language,length,score FROM movie WHERE moviename LIKE ?"
	row, err := dB.Query(str,s)
	if err != nil {
		return Data1{},err
	}

	defer row.Close()

	for row.Next() {
		var movie model.Movie
		err := row.Scan(&movie.Name, &movie.Director, &movie.Screenwriter, &movie.Starring, &movie.Style, &movie.Area, &movie.Language, &movie.Length, &movie.Score)
		if err != nil {
			return Data1{},err
		}
		Movies = append(Movies, movie)
	}


	str2:="SELECT name , birthdate, job FROM people where name LIKE ?"

	rows,err2 :=dB.Query(str2,s)
	if err2 != nil {
		return Data1{},err
	}
	defer rows.Close()
	for rows.Next() {
		var people1 model.People
		err = rows.Scan(&people1.Name,&people1.Birthdate,&people1.Job,)
		if err != nil {
			return Data1{},err
		}
		Peoples = append(Peoples,people1)
	}


	data2:= Data1{
		Movie: Movies,
		People: Peoples,
	}
	return data2, nil
}

