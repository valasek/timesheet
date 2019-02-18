// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package api

import (
	"encoding/json"
	"net/http"
)

// ProjectsGetAll returns list of all projects
func (api *API) ProjectsGetAll(w http.ResponseWriter, req *http.Request) {
	projects := api.projects.ProjectsGetAll()
	json.NewEncoder(w).Encode(projects)
}
