package dao

import (
	"douban/model"
)

func InsertMovie(movie model.Movie) error {
	_, err := dB.Exec("INSERT INTO movie(moviename, director, screenwriter, starring,style,area,language,releasetime,length,imdb,score) "+"values(?, ?, ?, ?,?,?,?,?,?,?,?);", movie.Name,movie.Director,movie.Screenwriter,movie.Starring,movie.Style,movie.Area,movie.Language,movie.ReleaseData,movie.Length,movie.IMDb,movie.Score)
	return err
}

func QueryMovie(movieName string)([]model.Movie,error) {
	var Movies []model.Movie
	rows, err := dB.Query("SELECT moviename, director, screenwriter, starring,style,area,language,releasetime,length,imdb,score FROM movie WHERE moviename LIKE ?", movieName)
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


