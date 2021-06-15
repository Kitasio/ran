package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
)

const (
	DRIVER_NAME      = "mysql"
	DATA_SOURCE_NAME = "kit:shitface1.0@/ran"
)

var db *sql.DB
var connectionError error
var store *sessions.CookieStore
var sessionName string = "goSession"

// TODO make the cookie secure
func init() {
	db, connectionError = sql.Open(DRIVER_NAME, DATA_SOURCE_NAME)
	if connectionError != nil {
		log.Fatal("error connecting to database : ", connectionError)
	}

	store = sessions.NewCookieStore([]byte("super-secret-key"))
}

type User struct {
	Uid      int    `json:"uid"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type clientQuery struct {
	Query string   `json:"query"`
	Data  []string `json:"data"`
}

var readQueries map[string]string = map[string]string{
	"getAllUsers":   "SELECT name FROM users",
	"getManagement": "SELECT * FROM management",
	"getCouncil":    "SELECT * FROM council",
	"allNews":       "SELECT * FROM news",
	"allAds":        "SELECT * FROM ads",
	"allUnits":      "SELECT * FROM units",
	"allPartners":   "SELECT * FROM partners",
	"singleNews":    "SELECT * FROM news WHERE id=?",
	"singleAd":      "SELECT * FROM ads WHERE id=?",
	"singleUnit":    "SELECT * FROM units WHERE id=?",
}

var writeQueries map[string]string = map[string]string{
	"createManagement": "INSERT management SET name=?, position=?, phone=?, email=?",
	"createCouncil":    "INSERT council SET name=?, position=?, phone=?, email=?",
	"createNews":       "INSERT news SET title=?, tag=?, date=?, body=?, img=?",
	"createAd":         "INSERT ads SET title=?, tag=?, date=?, time=?",
	"createUnit":       "INSERT units SET title=?, name=?, body=?, img=?",
	"createPartner":    "INSERT partners SET name=?, link=?, img=?",
}

var updateQueries map[string]string = map[string]string{
	"updateNews": "UPDATE news SET title=?, tag=?, date=?, body=?, img=? WHERE id=?",
	"updateAd":   "UPDATE ads SET title=?, tag=?, date=?, time=? WHERE id=?",
	"updateUnit": "UPDATE units SET title=?, name=?, body=?, img=? WHERE id=?",
}

var deleteQueries map[string]string = map[string]string{
	"deleteManagement": "DELETE from management WHERE id=?",
	"deleteCouncil":    "DELETE from council WHERE id=?",
	"deleteNews":       "DELETE from news WHERE id=?",
	"deleteAd":         "DELETE from ads WHERE id=?",
	"deleteUnit":       "DELETE from units WHERE id=?",
	"deletePartner":    "DELETE from partners WHERE id=?",
}

func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, sessionName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		var authenticated interface{} = session.Values["authenticated"]
		if authenticated != nil {
			isAuthenticated := session.Values["authenticated"].(bool)
			if isAuthenticated {
				endpoint(w, r)
			}
		}
	})
}

func Upload(w http.ResponseWriter, r *http.Request) {
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("static", "*"+handler.Filename)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	w.Write([]byte(tempFile.Name()))
}

func getJSON(sqlString string) ([]byte, error) {
	rows, err := db.Query(sqlString)
	if err != nil {
		return []byte(""), err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return []byte(""), err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return []byte(""), err
	}
	return jsonData, nil
}

func getJSONWithArg(sqlString, arg string) ([]byte, error) {
	rows, err := db.Query(sqlString, arg)
	if err != nil {
		return []byte(""), err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return []byte(""), err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return []byte(""), err
	}
	return jsonData, nil
}

func ReadRecords(w http.ResponseWriter, r *http.Request) {
	reqData := clientQuery{}
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		log.Print("Error decoding json")
	}
	if len(reqData.Data) < 1 {
		data, err := getJSON(readQueries[reqData.Query])
		if err != nil {
			log.Print("error while converting to json")
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	} else {
		data, err := getJSONWithArg(readQueries[reqData.Query], reqData.Data[0])
		log.Print("json with arg func", readQueries[reqData.Query], reqData.Data[0])
		if err != nil {
			log.Print("error while converting to json")
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func CreateRecord(w http.ResponseWriter, r *http.Request) {
	reqData := clientQuery{}
	err := json.NewDecoder(r.Body).Decode(&reqData)
	log.Printf("request data: %v", reqData)
	stmt, err := db.Prepare(writeQueries[reqData.Query])
	if err != nil {
		log.Print("error preparing query :: ", err)
		return
	}
	var queryParams []interface{}
	for i := 0; i < len(reqData.Data); i++ {
		queryParams = append(queryParams, string(reqData.Data[i]))
	}
	fmt.Println(queryParams, reqData.Data)
	result, err := stmt.Exec(queryParams...)
	if err != nil {
		log.Print("error executing query :: ", err)
		return
	}
	id, err := result.LastInsertId()
	fmt.Fprintf(w, "Last Inserted Record Id is  :: %s", strconv.FormatInt(id, 10))
}

func UpdateRecord(w http.ResponseWriter, r *http.Request) {
	reqData := clientQuery{}
	err := json.NewDecoder(r.Body).Decode(&reqData)
	log.Printf("request data: %v", reqData)
	stmt, err := db.Prepare(updateQueries[reqData.Query])
	if err != nil {
		log.Print("error preparing query :: ", err)
		return
	}
	var queryParams []interface{}
	for i := 0; i < len(reqData.Data); i++ {
		queryParams = append(queryParams, string(reqData.Data[i]))
	}
	fmt.Println(queryParams...)
	fmt.Println(updateQueries[reqData.Query])
	result, err := stmt.Exec(queryParams...)
	if err != nil {
		log.Print("error executing query :: ", err)
		return
	}
	id, err := result.LastInsertId()
	fmt.Fprintf(w, "Last Inserted Record Id is  :: %s", strconv.FormatInt(id, 10))
}

func DeleteRecord(w http.ResponseWriter, r *http.Request) {
	reqData := clientQuery{}
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		log.Print("Error decoding json data")
	}
	log.Printf("request data: %v", reqData)
	stmt, err := db.Prepare(deleteQueries[reqData.Query])
	if err != nil {
		log.Print("error occurred while preparing query :: ", err)
		return
	}

	var queryParams []interface{}
	for i := 0; i < len(reqData.Data); i++ {
		queryParams = append(queryParams, string(reqData.Data[i]))
	}
	fmt.Println(queryParams, reqData.Data)
	result, err := stmt.Exec(queryParams...)

	if err != nil {
		log.Print("error occurred while executing query :: ", err)
		return
	}
	rowsAffected, err := result.RowsAffected()
	fmt.Fprintf(w, "Number of rows deleted in database are :: %d", rowsAffected)
}

func Auth(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, sessionName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var authenticated interface{} = session.Values["authenticated"]
	if authenticated != nil {
		isAuthenticated := session.Values["authenticated"].(bool)
		if isAuthenticated {
			w.Write([]byte(session.Values["user"].(string)))
		}
	}
}

// TODO store passwords securely in the database
func Login(w http.ResponseWriter, r *http.Request) {
	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Print("error occured while reading data: ", err)
		return
	}

	row, err := db.Query("SELECT password FROM users WHERE name=?", user.Name)
	if err != nil {
		log.Print("error preparing the query")
		return
	}

	var password string
	for row.Next() {
		err = row.Scan(&password)
	}

	if password == user.Password {
		session, err := store.Get(r, sessionName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		session.Values["user"] = user.Name
		session.Values["authenticated"] = true
		session.Save(r, w)
		fmt.Fprintln(w, "authenticated:", session.Values["authenticated"])
	} else {
		session, err := store.Get(r, sessionName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		session.Save(r, w)
		w.WriteHeader(401)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, sessionName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	session.Values["authenticated"] = false
	session.Options.MaxAge = -1
	session.Save(r, w)
	fmt.Fprintln(w, "authenticated:", session.Values["authenticated"])
}
