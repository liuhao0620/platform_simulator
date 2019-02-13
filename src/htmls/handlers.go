package htmls

import (
	"data"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"
)

var template_file_name = "htmls/index.html"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Index Get")
	t, err := template.ParseFiles(template_file_name)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "Unknown Error")
		return
	}
	t.Execute(w, map[string]interface{}{
		"display_index":                 "block",
		"display_registe_result":        "none",
		"display_changepassword_result": "none",
		"display_registe":               "none",
		"display_changepassword":        "none"})
}

func RegisteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Registe Get")
	t, err := template.ParseFiles(template_file_name)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "Unknown Error")
		return
	}
	params, err := url.ParseQuery(r.RequestURI[strings.Index(r.RequestURI, "?")+1:])
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "Unknown Error")
		return
	}
	if params["username"] == nil {
		t.Execute(w, map[string]interface{}{
			"display_index":                 "none",
			"display_registe_result":        "none",
			"display_changepassword_result": "none",
			"display_registe":               "block",
			"display_changepassword":        "none"})
		return
	}

	username := params["username"][0]
	registe_result := username + " 注册成功"
	password, err := data.RedisGetPassword(username)
	if err != nil {
		registe_result = username + " 注册失败：" + err.Error()
	} else {
		if password == "" {
			password = params["password"][0]
			err = data.RedisSetPassword(username, password)
			if err != nil {
				registe_result = username + " 注册失败：" + err.Error()
			}
		} else {
			registe_result = username + " 注册失败： 用户名已存在"
		}
	}
	t.Execute(w, map[string]interface{}{
		"display_index":                 "block",
		"display_registe_result":        "block",
		"registe_result":                registe_result,
		"display_changepassword_result": "none",
		"display_registe":               "none",
		"display_changepassword":        "none"})
}

func ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ChangePassword Get")
	t, err := template.ParseFiles(template_file_name)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "Unknown Error")
		return
	}
	params, err := url.ParseQuery(r.RequestURI[strings.Index(r.RequestURI, "?")+1:])
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "Unknown Error")
		return
	}
	if params["username"] == nil {
		t.Execute(w, map[string]interface{}{
			"display_index":                 "none",
			"display_registe_result":        "none",
			"display_changepassword_result": "none",
			"display_registe":               "none",
			"display_changepassword":        "block"})
		return
	}
	username := params["username"][0]
	changepassword_result := username + " 密码修改成功"
	password, err := data.RedisGetPassword(username)
	if err != nil {
		changepassword_result = username + " 密码修改失败：" + err.Error()
	} else {
		if password != "" {
			old_password := params["password"][0]
			new_password := params["new_password"][0]
			if password == old_password {
				err = data.RedisSetPassword(username, new_password)
				if err != nil {
					changepassword_result = username + " 密码修改失败：" + err.Error()
				}
			} else {
				changepassword_result = username + " 密码修改失败：旧密码不正确"
			}
		} else {
			changepassword_result = username + " 密码修改失败： 用户名不存在，请先注册"
		}
	}
	t.Execute(w, map[string]interface{}{
		"display_index":                 "block",
		"display_registe_result":        "none",
		"display_changepassword_result": "block",
		"changepassword_result":         changepassword_result,
		"display_registe":               "none",
		"display_changepassword":        "none"})
}

func CheckUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CheckUser Get")
	params, err := url.ParseQuery(r.RequestURI[strings.Index(r.RequestURI, "?")+1:])
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "Unknown Error")
		return
	}
	username := params["username"][0]
	check_password := params["password"][0]
	password, err := data.RedisGetPassword(username)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, "Unknown Error")
		return
	}
	if password == check_password {
		fmt.Fprintln(w, "true")
	} else {
		fmt.Fprintln(w, "false")
	}
}
