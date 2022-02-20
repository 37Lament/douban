package dao

import (
	"douban/model"
)

func InsertMovie(movie model.Movie) error {
	str:="INSERT INTO movie(moviename, director, screenwriter, starring,style,area,language,releasetime,length,imdb) "+"values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	_, err := dB.Exec(str, movie.Name,movie.Director,movie.Screenwriter,movie.Starring,movie.Style,movie.Area,movie.Language,movie.ReleaseData,movie.Length,movie.IMDb)
	return err
}

func QueryMovie(movieid int)([]model.Movie,error) {
	var Movies []model.Movie
	rows, err := dB.Query("SELECT moviename, director, screenwriter, starring,style,area,language,releasetime,length,imdb,score FROM movie WHERE mid LIKE ?", movieid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var movie model.Movie
		err := rows.Scan(&movie.Name, &movie.Director, &movie.Screenwriter, &movie.Starring, &movie.Style, &movie.Area, &movie.Language, &movie.ReleaseData, &movie.Length, &movie.IMDb, &movie.Score)
		if err != nil {
			return nil, err
		}
		Movies = append(Movies, movie)
	}
	return Movies, nil
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



