package dao

import "douban/model"

func InsertComment(comment model.Comment) error {
	_, err := dB.Exec("INSERT INTO comment(username, txt, commenttime, moviename,flag) "+"values(?, ?, ?, ?,?);", comment.Username, comment.Txt, comment.CommentTime, comment.MovieName,comment.Flag)
	return err
}

