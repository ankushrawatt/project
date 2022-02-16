package helper

import (
	"assignment/database"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserInfo struct {
	ID     string `db:"id" json:"id"`
	Name   string `db:"name" json:"name"`
	Email  string `db:"email" json:"email"`
	Mobile string `db:"mobile_no" json:"mobile_no"`
	Userid string `db:"userid" json:"userid"`
}

type UserLog struct {
	ID string `db:"id" json:"id"`
	//	Name      string    `db:"name" json:"name"`
	//	Password  string    `db:"password" json:"password"`
	Userid    string    `db:"userid" json:"userid"`
	Createdat time.Time `db:"createdat" json:"createdat"`
}

type LoginInfo struct {
	Userid   string `db:"userid" json:"userid"`
	Password string `db:"password" json:"password"`
}

func HashPassword(password string) string {
	pass, _ := bcrypt.GenerateFromPassword([]byte(password), 0)
	return string(pass)
}

func CheckHashPassword(password, pass string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(pass), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

func Newuser(id, email, password, name, mobile, userid string) (*UserInfo, error) {
	SQL := `INSERT INTO users(id,name,email,password,mobile_no,userid) VALUES($1,$2,$3,$4,$5,$6) returning id,name,email,mobile_no,userid`
	var user UserInfo
	pass := HashPassword(password)
	err := database.Assignment.Get(&user, SQL, id, name, email, pass, mobile, userid)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(mobile, email string) error {
	SQL := `UPDATE users SET mobile_no=$1 WHERE email=$2`
	//var user UserInfo
	_, err := database.Assignment.Exec(SQL, mobile, email)
	if err != nil {
		return err
	}
	return nil
}

func Delete(email, password string) error {
	SQL := `DELETE FROM users WHERE email=$1 AND password=$2`
	_, err := database.Assignment.Exec(SQL, email, password)
	if err != nil {
		return err
	}
	return nil
}

func Login(id, password, email string) (string, error) {
	SQL := `SELECT userid,password FROM users where email=$1 `
	//var user LoginInfo
	var pass, user string
	err := database.Assignment.QueryRowx(SQL, email).Scan(&user, &pass)
	//fmt.Println(user)
	if err != nil {
		return "", err
	}
	hash, passwordErr := CheckHashPassword(password, pass)

	if hash {
		SQL = `INSERT INTO session(id,userid) VALUES($1,$2)returning id`
		var loggedUser UserLog
		NewErr := database.Assignment.Get(&loggedUser, SQL, id, user)
		if NewErr != nil {
			return "", NewErr
		}
		return "", nil
	}
	return "WRONG PASSWORD OR USERNAME", passwordErr
}

func GetSession(sessionToken string) (*UserInfo, error) {
	SQL := `SELECT  u.userid,u.email,u.mobile_no,u.name FROM users u JOIN session s on u.userid = s.userid WHERE s.id = $1`
	var user UserInfo
	err := database.Assignment.Get(&user, SQL, sessionToken)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, nil
}

//func LoggedUser(Token, userID string) (*UserLog, error) {
//	SQL := ``
//	var user UserLog
//	err := database.Assignment.Get(&user, SQL)
//	if err != nil {
//		return nil, err
//	}
//	return &user, nil
//}

//func GetSession(SessionTokken string) (*UserInfo, error) {
//	SQL := `SELECT u.userid,u.name,u.email FROM users u JOIN session s ON u.id=s.id where s.id=$1`
//	fmt.Println(SessionTokken)
//	var user UserInfo
//	err := database.Assignment.Get(&user, SQL, SessionTokken)
//	if err != nil {
//		return nil, err
//	}
//	return &user, nil
//}

func LogoutUser(id string) error {
	SQL := `DELETE FROM session WHERE id=$1`
	_, err := database.Assignment.Exec(SQL, id)
	if err != nil {
		return err
	}
	return nil
}
