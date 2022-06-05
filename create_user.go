package RestApiServer

import (
	"RestApiServer/mlog"
	"database/sql"
	"net/http"
)

func GetNewDataBaseConnect() (*sql.DB, error) {
	sqlDB, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		mlog.Error("Database not connected. Error: %v", err)
		return nil, err
	}
	mlog.Info("Database connected")
	return sqlDB, nil
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	c := Users{}
	db, err := GetNewDataBaseConnect()
	if err != nil {
		mlog.Error("Database connect error: %v", err)
	}
	statement, _ := db.Prepare("INSERT INTO user (name,sirname,patronymic,number,email,birthdate) VALUES (?,?,?,?,?,?)")
	statement.Exec(statement, c.Id, c.NSP.Name, c.NSP.Sirname, c.NSP.Patronymic)
	if err != nil {
		return
	}
	checkErr(err)
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
