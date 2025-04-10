package web

import (
	orm_models "amsys/internal/api/models"
	api_services "amsys/internal/api/services"

	"html/template"
	"net/http"
	"strconv"
	"time"
)

// HomeHandler serves the home page with the sidebar and user list
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"internal/web/templates/base.html",
		"internal/web/templates/navbar.html",
		"internal/web/templates/home.html",
	))
	tmpl.Execute(w, nil)
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"internal/web/templates/base.html",
		"internal/web/templates/navbar.html",
		"internal/web/templates/users.html",
	))
	tmpl.Execute(w, nil)
}

type PullRequest struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	RepoName    string    `json:"repo_name"`
	PrId        string    `json:"pr_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Tags        []string  `json:"tags"`
	Version     string    `json:"version"`
	MergedAt    time.Time `json:"merged_at"`
}

const pageSize = 10 // Number of items per page

func TableHandler(w http.ResponseWriter, r *http.Request) {
	versions, err := api_services.GetVersionsByRepo("sinhantb")
	if err {
		http.Error(w, "Failed to fetch versions", http.StatusInternalServerError)
		return
	}

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}

	start := (page - 1) * pageSize
	end := start + pageSize
	if start >= len(versions) {
		start = len(versions)
	}
	if end > len(versions) {
		end = len(versions)
	}

	data := struct {
		Versions []orm_models.ControlVersion
		PrevPage int
		NextPage int
	}{
		Versions: versions[start:end],
		PrevPage: page - 1,
		NextPage: page + 1,
	}

	tmpl, _ := template.ParseFiles(
		"internal/web/templates/changelog/cards_section.html",
		"internal/web/templates/changelog/icon_selector.html",
	)
	tmpl.Execute(w, data)
}
