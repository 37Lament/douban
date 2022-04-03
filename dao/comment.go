package dao

import (
	"douban/model"
	"fmt"
)

func InsertComment(comment model.Comment) error {
	str:="INSERT INTO comment(username, txt, commenttime, moviename,flag,score) "+"values(?, ?, ?, ?,?,?);"
	_, err := dB.Exec(str, comment.Username, comment.Txt, comment.CommentTime, comment.MovieName,comment.Flag,comment.Score)
	return err
}
func InsertlevelComment(comment model.Comment) error {
	str:="INSERT INTO comment(username, txt, commenttime,postid) "+"values(?, ?, ?, ?);"
	_, err := dB.Exec(str, comment.Username, comment.Txt, comment.CommentTime, comment.Postid)
	return err
}
type commentf struct {
	Username string
	Moviename string
	Txt string
	CommentTime string
	Flag bool
}

type commentss struct {
	Username string
	Txt string
	Commenttime string
}

type Bigcomment struct {
	Fcomment commentf
	Coment2s  []commentss
}

func ShowComment(id int)(Bigcomment, error)  {
	var comment commentf
	var Comments []commentss
	str:="SELECT username,commenttime,txt,flag,moviename FROM comment WHERE id = ? "
	str2:="SELECT username,commenttime,txt FROM comment WHERE postid = ? "

	row:= dB.QueryRow(str, id)
	if row.Err() != nil {
		return Bigcomment{}, row.Err()
	}

	err := row.Scan(&comment.Username,&comment.CommentTime,&comment.Txt,&comment.Flag,&comment.Moviename)
	if err != nil {
		return Bigcomment{}, err
	}

	rows, err := dB.Query(str2,id)
	if err != nil {
		return Bigcomment{},err
	}

	defer rows.Close()

	for rows.Next() {
		var comment2 commentss
		err := rows.Scan(&comment2.Username,&comment2.Commenttime,&comment2.Txt)
		if err != nil {
			return Bigcomment{},err
		}
		Comments=append(Comments,comment2)
	}

	Wholec:= Bigcomment{
		Fcomment: comment,
		Coment2s: Comments,
	}

	return Wholec,nil
}

func DeletC(id int)error  {
	comment:=model.Comment{,1,,,,,,}
	dB.Delete(&comment)
}
