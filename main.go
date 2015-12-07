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
	Mreading int
	Preading int
	Technight int
	Orozo int
	Video int
	Taspi int
	Kaza int
	Ezb int
	Kop int
	Sabak int
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
	mreading := r.FormValue("Mreading")
	preading := r.FormValue("Preading")
	technight := r.FormValue("Technight")
	orozo     := r.FormValue("Orozo")
	video     := r.FormValue("Video")
	taspi     := r.FormValue("Taspi")
	kaza      := r.FormValue("Kaza")
	ezb       := r.FormValue("Ezb")
	kop       := r.FormValue("Kop")
	sabak     := r.FormValue("Sabak")


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

	// number of mreading
	nmreading, err := strconv.Atoi(mreading)
	if err != nil {
		http.Error(w, err.Error() + " mreading", http.StatusInternalServerError)
                return
	}

	// number of preading
	npreading, err := strconv.Atoi(preading)
	if err != nil {
		http.Error(w, err.Error() + " preading", http.StatusInternalServerError)
                return
	}

	// number of technight
	ntechnight, err := strconv.Atoi(technight)
	if err != nil {
		http.Error(w, err.Error() + " technight", http.StatusInternalServerError)
                return
	}

	// number of orozo
	norozo, err := strconv.Atoi(orozo)
	if err != nil {
		http.Error(w, err.Error() + " orozo", http.StatusInternalServerError)
                return
	}

	// number of video
	nvideo, err := strconv.Atoi(video)
	if err != nil {
		http.Error(w, err.Error() + " video", http.StatusInternalServerError)
                return
	}

	// number of video
	ntaspi, err := strconv.Atoi(taspi)
	if err != nil {
		http.Error(w, err.Error() + " taspi", http.StatusInternalServerError)
                return
	}

	// number of kaza
	nkaza, err := strconv.Atoi(kaza)
	if err != nil {
		http.Error(w, err.Error() + " kaza", http.StatusInternalServerError)
                return
	}

	// number of ezb
	nezb, err := strconv.Atoi(ezb)
	if err != nil {
		http.Error(w, err.Error() + " ezb", http.StatusInternalServerError)
                return
	}

	// number of kop
	nkop, err := strconv.Atoi(kop)
	if err != nil {
		http.Error(w, err.Error() + " kop", http.StatusInternalServerError)
                return
	}

	// number of sabak
	nsabak, err := strconv.Atoi(sabak)
	if err != nil {
		http.Error(w, err.Error() + " sabak", http.StatusInternalServerError)
                return
	}

        e := Entry{
		Author: author,
		Date:    time.Now(),                
		Kreading: nkreading,
		Rreading: nrreading,
		Creading: ncreading,
                Mreading: nmreading,
		Preading: npreading,
		Technight: ntechnight,
		Orozo:    norozo,
		Video:    nvideo,
		Taspi:    ntaspi,
		Kaza:     nkaza,
		Ezb:      nezb,
		Kop:      nkop,
		Sabak:    nsabak,
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
