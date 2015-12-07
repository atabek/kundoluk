package logbook

import (
        "html/template"
        "net/http"
        "time"

        "appengine"
        "appengine/datastore"
        "appengine/user"
	"strconv"
)

// [START greeting_struct]
type Entry struct {
	Date    time.Time        
	Author  string

	Kreading int
	Rreading int
	Creading int
}
// [END greeting_struct]

func init() {
        http.HandleFunc("/", root)
        http.HandleFunc("/sign", sign)
}

// guestbookKey returns the key used for all guestbook entries.
func logbookKey(c appengine.Context) *datastore.Key {
        // The string "default_logbook" here could be varied to have multiple guestbooks.
        return datastore.NewKey(c, "Logbook", "default_logbook", 0, nil)
}

// [START func_root]
func root(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
	u := user.Current(c)
	author := u.String()
        // Ancestor queries, as shown here, are strongly consistent with the High
        // Replication Datastore. Queries that span entity groups are eventually
        // consistent. If we omitted the .Ancestor from this query there would be
        // a slight chance that Greeting that had just been written would not
        // show up in a query.
        // [START query]
        q := datastore.NewQuery("Entry").Ancestor(logbookKey(c)).Filter("Author =", author).Order("-Date").Limit(10)
        // [END query]
        // [START getall]
	//fmt.Println(user.Current(c))
        entries := make([]Entry, 0, 10)
        if _, err := q.GetAll(c, &entries); err != nil {
                http.Error(w, err.Error() + " Get all queries", http.StatusInternalServerError)
                return
        }
        // [END getall]
        if err := logbookTemplate.Execute(w, entries); err != nil {
                http.Error(w, err.Error() + " template", http.StatusInternalServerError)
        }
}
// [END func_root]

var logbookTemplate = template.Must(template.ParseFiles("templates/home.gtpl"))

// [START func_sign]
func sign(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
	u := user.Current(c)
	author := u.String()
	

	kreading := r.FormValue("Kreading")
	rreading := r.FormValue("Rreading")
	creading := r.FormValue("Creading")

	// number of kreading
	nkreading, err := strconv.Atoi(kreading)
	if err != nil {
		http.Error(w, err.Error() + " creading", http.StatusInternalServerError)
                return
	}

	// number of rreading
	nrreading, err := strconv.Atoi(rreading)
	if err != nil {
		http.Error(w, err.Error() + " creading", http.StatusInternalServerError)
                return
	}

	// number of creading
	ncreading, err := strconv.Atoi(creading)
	if err != nil {
		http.Error(w, err.Error() + " creading", http.StatusInternalServerError)
                return
	}

        e := Entry{
		Author: author,
                Kreading: nkreading,
		Rreading: nrreading,
		Creading: ncreading,
                Date:    time.Now(),
        }
        // We set the same parent key on every Greeting entity to ensure each Greeting
        // is in the same entity group. Queries across the single entity group
        // will be consistent. However, the write rate to a single entity group
        // should be limited to ~1/second.
        key := datastore.NewIncompleteKey(c, "Entry", logbookKey(c))
        _, err = datastore.Put(c, key, &e)
        if err != nil {
                http.Error(w, err.Error() + " putting datastore", http.StatusInternalServerError)
                return
        }
        http.Redirect(w, r, "/", http.StatusFound)
}
// [END func_sign]
