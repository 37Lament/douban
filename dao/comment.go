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
	sqlStr := "delete from comment where id = ?"
	ret, err := dB.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return err
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return err
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
	return nil
}
