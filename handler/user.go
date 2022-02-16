package handler

import (
	"assignment/database/helper"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type UserInfo struct {
	Name     string `json:"Name"`
	UserID   string `json:"UserID"`
	Email    string `json:"Email"`
	Mobile   string `json:"mobile"`
	Password string `json:"Password"`
}

type token struct {
	ID string `json:"ID"`
}

//var (
//	jwtkey = []byte("secret-key")
//	store  = sessions.NewCookieStore(jwtkey)
//)

//type Claims struct {
//	Userid string `json:"userid"`
//	jwt.StandardClaims
//}

type LoginInfo struct {
	UserId   string `json:"userId"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup(writer http.ResponseWriter, request *http.Request) {
	var info UserInfo

	err := json.NewDecoder(request.Body).Decode(&info)
	if err != nil {
		writer.WriteHeader(http.StatusBadGateway)
		return
	}
	fmt.Println(info.Mobile)
	id := uuid.New()
	userId, newErr := helper.Newuser(id.String(), info.Email, info.Password, info.Name, info.Mobile, info.UserID)
	if newErr != nil {
		writer.WriteHeader(http.StatusBadGateway)
		return
	}
	jsonData, jsonErr := json.Marshal(userId)
	if jsonErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Write([]byte(fmt.Sprintf("Thank you for signing %s", userId.Name)))
	writer.Write(jsonData)
}

func Update(writer http.ResponseWriter, request *http.Request) {
	var info UserInfo
	err := json.NewDecoder(request.Body).Decode(&info)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	newErr := helper.UpdateUser(info.Mobile, info.Email)
	if newErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Write([]byte(fmt.Sprintf("Record updated successfully...")))
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	var info UserInfo
	err := json.NewDecoder(request.Body).Decode(&info)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	newErr := helper.Delete(info.Email, info.Password)
	if newErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Write([]byte(fmt.Sprintf("Record deleted successfully...")))

}

func Login(writer http.ResponseWriter, request *http.Request) {
	var info LoginInfo
	err := json.NewDecoder(request.Body).Decode(&info)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	//	request.Header.Get("x-api-key")
	id := uuid.New()
	loginUser, newErr := helper.Login(id.String(), info.Password, info.Email)

	if newErr != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	//	user, userErr := helper.Login

	//session, _ := store.Get(request, "Cookie")
	//session.Values["authenticated"] = true
	//session.Save(request, writer)

	//expirationTime := time.Now().Add(time.Minute * 5)
	//claims := &Claims{
	//	Userid: info.UserId,
	//	StandardClaims: jwt.StandardClaims{
	//		ExpiresAt: expirationTime.Unix(),
	//	},
	//}
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//tokenString, err := token.SignedString(jwtkey)
	//if err != nil {
	//	writer.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//http.SetCookie(writer, &http.Cookie{
	//	Name:    "token",
	//	Value:   tokenString,
	//	Expires: expirationTime,
	//})

	fmt.Println(loginUser)

	//userLogin, loggedErr := helper.LoggedUser(id.String(), loginUser)
	//if loggedErr != nil {
	//	writer.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//request.Header.Get(id.String())
	jsonData, jsonErr := json.Marshal(loginUser)
	if jsonErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Write([]byte(fmt.Sprintf("Tokken: %s", id.String())))
	writer.Write(jsonData)
}

func Home(writer http.ResponseWriter, request *http.Request) {
	//session, _ := store.Get(request, "Cookie")
	//if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
	//	http.Error(writer, "Forbidden", http.StatusForbidden)
	//	return
	//}
	writer.Write([]byte(fmt.Sprintf("Hello user....")))
}

func Logout(writer http.ResponseWriter, request *http.Request) {
	var token token
	json.NewDecoder(request.Body).Decode(&token)
	err := helper.LogoutUser(token.ID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	writer.Write([]byte(fmt.Sprintf("logged out")))
}
