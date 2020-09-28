package handlers

import (
	"github.com/asaskevich/govalidator"
	"html/template"
	"math/rand"
	"net/http"
	"shortlinkapp/pkg/models"

	"shortlinkapp/pkg/database"

	"github.com/julienschmidt/httprouter"
)

type LinksHandler struct {
	Repo *database.Repository
	Tmpl *template.Template
}

func (handler *LinksHandler) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := handler.Tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}

func (handler *LinksHandler) GetLink(w http.ResponseWriter,
								r *http.Request,
								ps httprouter.Params) {

	link := ps.ByName("link")
	shortLink, err := handler.Repo.Get(link)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, shortLink.Link, http.StatusMovedPermanently)
}

func (handler *LinksHandler) CreateShortLink(w http.ResponseWriter,
										r *http.Request,
										_ httprouter.Params) {
	link := r.FormValue("link")
	name := r.FormValue("name")

	// Валидация URL
	urlCond := govalidator.IsURL(link)
	if !urlCond {
		http.Error(w, "incorrect URL", http.StatusBadRequest)
		return
	}
	// Проверка на доступность
	_, err := http.Get(link)
	if err != nil {
		http.Error(w, "unavailable link: " + err.Error(), http.StatusBadRequest)
		return
	}

	if name == "" {
		// Генерация "имени" ссылки
		linkLength := 6
		name = GenerateLinkName(linkLength)
	} else {
		// Проверка корректности пользовательского "имени" ссылки
		cond := CheckLinkName(name)
		if !cond {
			http.Error(w, "incorrect link name", http.StatusBadRequest)
			return
		}
	}

	_, err = handler.Repo.Set(&models.ShortLink{ID: name, Link: link})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = handler.Tmpl.ExecuteTemplate(w, "result.html",
		struct{ Link string } {
			Link: "http://localhost:9000/" + name,
	})
	if err != nil {
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
	//log.Fatal(fmt.Fprintln(w, "short link: localhost:9000/" + name))
}

// Проверка пользовательского имени ссылки на соответствие "A..Z, a..z, -"
func CheckLinkName(name string) bool {
	arr := []rune(name)
	var num int
	for _, v := range arr {
		num = int(v)
		if !(num >= 65 && num <= 90 ||
			num >= 97 && num <= 122 ||
			num == 45) {
			return false
		}
	}
	return true
}

func GenerateLinkName(length int) string {
	runes := make([]rune, length)
	for i := range runes {
		// Генерация ссылки по номерам букв в ASCII
		num := rand.Intn(57) + 65
		for {
			if num > 90 && num < 97 {
				num = rand.Intn(57) + 65
			} else {
				break
			}
		}
		runes[i] = rune(num)
	}
	return string(runes)
}