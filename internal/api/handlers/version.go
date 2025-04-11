package api_handlers

import (
	"amsys/internal/api/handlers/pr_parser"
	orm_models "amsys/internal/api/models"
	api_services "amsys/internal/api/services"
	"log"

	"encoding/json"
	"net/http"
)

func GetProjectHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	projectId := query.Get("projectId")

	if projectId == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "projectId query parameter is required"})
		return
	}

	results, err := api_services.GetVersionsByRepo(projectId)
	if err {
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string][]orm_models.ControlVersion{
		"results": results,
	})
}

func CreateVersionHandler(w http.ResponseWriter, r *http.Request) {
	body := make(map[string]interface{})
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	// Safely convert id to uint
	var id uint
	if idFloat, ok := body["issue_id"].(float64); ok {
		id = uint(idFloat)
	}

	version := orm_models.ControlVersion{
		ID:           id,
		IssueId:      body["issue_id"].(string),
		RepoOwner:    body["repo_owner"].(string),
		RepoName:     body["repo_name"].(string),
		TargetBranch: body["target_branch"].(string),
		PrID:         body["pr_id"].(string),
		Title:        body["title"].(string),
		Description:  body["description"].(string),
		Version:      body["version"].(string),
		AppLayer:     body["app_layer"].(string),
		IssueType:    body["issue_type"].(string),
		RawMessage:   body["raw_message"].(string),
		MergedAt:     body["merged_at"].(string),
	}

	if err := api_services.CreateVersion(version); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "failed to create version"})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "version created successfully"})
}

func CreateVersionFromGithubHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	body := make(map[string]any)
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid request body"})
		return
	}

	log.Printf("Request Body: %+v", body)
	var prTitleData map[string]interface{}
	if prTitleInterface, ok := body["pr_title"]; ok {
		prTitleData = prTitleInterface.(map[string]interface{})
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "pr_title is missing"})
		return
	}

	log.Printf("PR Fields: %+v", prTitleData)
	prTitle, err := pr_parser.ParsePRTitle(prTitleData["title"].(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid pr_title format"})
		return
	}
	log.Printf("PR Title Fields: %+v", prTitle)

	var prData map[string]interface{}
	if prInterface, ok := body["pr_title"]; ok {
		prData = prInterface.(map[string]interface{})
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "pr is missing"})
		return
	}

	version := orm_models.ControlVersion{
		IssueId:      prTitle.IssueId,
		IssueType:    prTitle.IssueType,
		Title:        prTitle.Title,
		Version:      prTitle.Version,
		AppLayer:     prTitle.AppLayer,
		RepoOwner:    body["repo_owner"].(string),
		RepoName:     body["repo_name"].(string),
		TargetBranch: body["target_branch"].(string),
		PrID:         prData["id"].(string),
		Description:  prData["body"].(string),
		MergedAt:     body["timestamp"].(string),
		RawMessage:   prData["body"].(string),
	}

	if err := api_services.CreateVersion(version); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "failed to create version"})
		return
	}

	w.WriteHeader(http.StatusOK)

}
