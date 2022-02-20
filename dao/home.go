package dao

type MiniMovie struct {
	Mname string
	Mscore float32
}

func Homep()([]MiniMovie,error) {
	var Movies [] MiniMovie
	s:="%"
	str:="SELECT moviename,score FROM movie WHERE moviename LIKE ?"
	rows, err := dB.Query(str,s)
	if err != nil {
		return nil,err
	}
	defer rows.Close()

	for rows.Next() {
		var movie MiniMovie
		err := rows.Scan(&movie.Mname,&movie.Mscore)
		if err != nil {
			return nil,err
		}
		Movies = append(Movies, movie)
	}

	return Movies,nil
}