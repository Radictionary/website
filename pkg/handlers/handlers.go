package handlers

import (
	"fmt"
	"net"
	"net/http"

	"github.com/Radictionary/website/pkg/config"
	"github.com/Radictionary/website/pkg/database"
	"github.com/Radictionary/website/pkg/models"
	"github.com/Radictionary/website/pkg/render"
	"github.com/dgraph-io/badger/v4"
)

var greeting string

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Delete(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	remoteIP, _, err := net.SplitHostPort(r.RemoteAddr)
	config.Handle(err, "Spliting the RemoteAddr into Host and Port")
	db := database.CallDatabase()
	defer db.Close()
	_ = db.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(string(remoteIP)))
		if err != nil {
			config.Handle(err, "Error Deleting IP Address record")

		}
		return nil
	  })

}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	var newUser bool
	remoteIP := r.RemoteAddr
	remoteIP, _, err := net.SplitHostPort(r.RemoteAddr)
	config.Handle(err, "Spliting the RemoteAddr into Host and Port")

	db := database.CallDatabase()
	defer db.Close()

	_ = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(string(remoteIP)))
		if err != nil {
			fmt.Println("NOT FOUND:", item, remoteIP)
			newUser = true
			return err

		} else {
			newUser = false
		}

		return err
	})
	database.ViewDatabase(db)
	if newUser {
		fmt.Println("THIS DEVICE IS NEW!", remoteIP)
		err := db.Update(func(txn *badger.Txn) error {
			err := txn.Set([]byte(string(remoteIP)), []byte(remoteIP))
			return err
		})
		config.Handle(err, "Could not register new IP address into DB")
		greeting = "Welcome"
	} else {
		fmt.Println("THIS DEVICE HAS BEEN HERE BEFORE!", remoteIP)
		greeting = "Weclome back!"
	}

	render.RenderTemplate(w, "home.html", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// send data to the template
	render.RenderTemplate(w, "about.html", &models.TemplateData{
		String: greeting,
	})
}
