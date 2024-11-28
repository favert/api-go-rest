package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/favert/api-go-rest/database"
	"github.com/favert/api-go-rest/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	if database.UseDB() == 1 {
		fmt.Println("Retrieve all data from DB")
		var p []models.Personalidade
		database.DB.Find(&p)
		json.NewEncoder(w).Encode(p)
	} else {
		fmt.Println("Retrieve all data from mock")
		json.NewEncoder(w).Encode(models.Personalidades)
	}
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if database.UseDB() == 1 {
		fmt.Println("Retrieve unique data from DB")
		var personalidade models.Personalidade
		database.DB.First(&personalidade, id)
		json.NewEncoder(w).Encode(personalidade)
	} else {
		fmt.Println("Retrieve unique data from mock")
		for _, personalidade := range models.Personalidades {
			if strconv.Itoa(personalidade.Id) == id {
				fmt.Println("Found id:", id, " on mock")
				json.NewEncoder(w).Encode(personalidade)
			}
		}
	}
}

func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&novaPersonalidade)
	if database.UseDB() == 1 {
		database.DB.Create(&novaPersonalidade)
	} else {
		models.Personalidades = append(models.Personalidades, novaPersonalidade)
	}
	json.NewEncoder(w).Encode(novaPersonalidade)
}

func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if database.UseDB() == 1 {
		fmt.Println("Delete unique data from DB")
		var personalidade models.Personalidade
		database.DB.Delete(&personalidade, id)
		json.NewEncoder(w).Encode(personalidade)
	} else {
		fmt.Println("Delete unique data from mock")
		var newModels []models.Personalidade
		for _, personalidade := range models.Personalidades {
			if strconv.Itoa(personalidade.Id) == id {
				fmt.Println("(Delete) Found id: ", id, " on mock")
				json.NewEncoder(w).Encode(personalidade)
			} else {
				newModels = append(newModels, personalidade)
			}
		}
		models.Personalidades = newModels
	}
}

func EditaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if database.UseDB() == 1 {
		fmt.Println("Edit unique data from DB")
		var personalidade models.Personalidade
		database.DB.First(&personalidade, id)
		json.NewDecoder(r.Body).Decode(&personalidade)
		database.DB.Save(&personalidade)
		json.NewEncoder(w).Encode(personalidade)
	} else {
		fmt.Println("Edit unique data from mock")
		var newModels []models.Personalidade
		for _, personalidade := range models.Personalidades {
			if strconv.Itoa(personalidade.Id) == id {
				var personalidadeNova models.Personalidade
				json.NewDecoder(r.Body).Decode(&personalidadeNova)
				fmt.Println("(Edit) Found id: ", id, " on mock")
				newModels = append(newModels, personalidadeNova)
			} else {
				newModels = append(newModels, personalidade)
			}
		}
		models.Personalidades = newModels
	}
}
