package dao

import (
	"douban/model"
)

func InsertMovie(movie model.Movie) error {
	str:="INSERT INTO movie(moviename, director, screenwriter, starring,style,area,language,releasetime,length,imdb) "+"values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	_, err := dB.Exec(str, movie.Name,movie.Director,movie.Screenwriter,movie.Starring,movie.Style,movie.Area,movie.Language,movie.ReleaseData,movie.Length,movie.IMDb)
	return err
}

type Page struct {
	Movie []model.Movie
	Commentsid []int
}
func QueryMovie(movieid int)(Page,error) {
	var Movies []model.Movie
	var movie model.Movie
	str:="SELECT moviename, director, screenwriter, starring,style,area,language,releasetime,length,imdb,score FROM movie WHERE mid LIKE ?"
	rows, err := dB.Query(str, movieid)
	if err != nil {
		return Page{}, err
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&movie.Name, &movie.Director, &movie.Screenwriter, &movie.Starring, &movie.Style, &movie.Area, &movie.Language, &movie.ReleaseData, &movie.Length, &movie.IMDb, &movie.Score)
		if err != nil {
			return Page{}, err
		}
		Movies = append(Movies, movie)
	}
	var idx [] int
	var comment model.Comment
	str2:="SELECT id From comment WHERE moviename LIKE ?"
	row,err:=dB.Query(str2,movie.Name)
	if err != nil {
		return Page{}, err
	}

	defer row.Close()

	for row.Next(){
		err=row.Scan(&comment.Id)
		if err != nil {
			return Page{}, err
		}
		idx=append(idx,comment.Id)
	}
	PG:=Page{
		Movie:Movies,
		Commentsid: idx,
	}

	return PG, nil
}

func SelectMoviebyId(id int) (model.Movie, error) {
	var movie model.Movie
	rows:= dB.QueryRow("SELECT moviename  FROM movie WHERE mid = ? ", id)
	if rows.Err() != nil {
		return movie, rows.Err()
	}

	err := rows.Scan(&movie.Name)
	if err != nil {
		return movie, err
	}

	return movie, nil
}



