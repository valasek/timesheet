package api

import (
	"net/http"
	"encoding/json"
)

// ProjectsGetAll returns list of all projects
func (api *API) ProjectsGetAll(w http.ResponseWriter, req *http.Request) {
	projects := api.projects.ProjectsGetAll()
	json.NewEncoder(w).Encode(projects)
}
