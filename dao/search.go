package dao

import "douban/model"

func Searching(data string)([]model.Movie,error) {
	var Movies []model.Movie
	s:="%"+data+"%"
	rows, err := dB.Query("SELECT moviename, director, screenwriter,starring,style,area,language,releasetime,length,imdb,score FROM movie WHERE (moviename LIKE ?) or (starring LIKE ?) or(director LIKE ?)", s,s,s)
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

