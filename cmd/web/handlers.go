package main

import (
	"database/sql"
	"html/template"
	"net/http"
	"net/smtp"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Неверный метод запроса", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		app.notFound(w) // Использование помощника notFound()
		return
	}
	files := []string{
		"./ui/html/home_page.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Не получен доступ к файлам страницы", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}
}

func (app *application) send_to_mail(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		app.notFound(w)
		return
	}

	name := r.FormValue("Name")
	phone := r.FormValue("phone")
	email := r.FormValue("email")
	goods := r.FormValue("goods")
	from := r.FormValue("from")
	to := r.FormValue("to")

	sendData := "Subject: Заявка на грузоперевозку \n" + "Оставлена заявка: \n" + "Имя:" + name + "\n Телефон: " + phone + "\n Почта: " + email + "\n Груз: " +
		goods + "\n Откуда: " + from + "\n Куда: " + to

	err := sendMail(sendData)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func sendMail(text string) error {
	sender := "markrivelli48410@gmail.com"
	senderPass := "rrmo nwzf eeep nenr"
	address := "WGSeries2@yandex.ru"
	host := "smtp.gmail.com"
	port := "587"
	auth := smtp.PlainAuth(
		"",
		sender,
		senderPass,
		host,
	)

	err := smtp.SendMail(
		host+":"+port,
		auth,
		sender,
		[]string{address},
		[]byte(text),
	)
	if err != nil {
		return err
	}
	return nil
}

func dbConnection() *sql.DB {
	db, errsql := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lab3")
	if errsql != nil {
		panic(errsql)
	}
	return db
}
