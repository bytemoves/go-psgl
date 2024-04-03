import (

    "database/sql"

    _ "github.com/lib/pq"

)


db, err := sql.Open("postgres", "postgres://your_username:your_username@host:port/your_database_name?sslmode=disable")

if err != nil {

    log.Fatal(err)

}

defer db.Close()


err = db.Ping()

if err != nil {

    log.Fatal(err)

}

fmt.Println("Successfully connected to PostgreSQL!")


rows, err := db.Query("SELECT * FROM your_table")

if err != nil {

    log.Fatal(err)

}

defer rows.Close()



// Process query results

for rows.Next() {

    // Process each row

}
