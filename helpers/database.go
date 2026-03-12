package database 

import (
	"database/sql"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func OpenDB() *sql.DB{
	godotenv.Load();
	dns := os.Getenv("DB_DNS")
	db, err := sql.Open("mysql", dns);
	
	if err != nil{
		panic(err);
	}
	return db;
}


func CloseDB(db *sql.DB){
	defer db.Close();
}

func Trydb(){
	fmt.Println("opening the database");
	db :=OpenDB()
	fmt.Println(db);
	fmt.Println("database opened")
	CloseDB(db)
	fmt.Println("db closed");
}