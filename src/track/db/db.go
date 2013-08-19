package db

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
//    "io/ioutil"
//    "os"
    "time"
)

func getConnection() (con *sql.DB) {
    con, err := sql.Open ("mysql", "tracktron:tracktron@/track")
    
    if err != nil {
        panic (err)
    }
    
    return
}
func FindVendors (searchString string) (vendors [] string, err error) {
    vendors = []string {}
    
    con := getConnection()
    defer con.Close()

    rows, err := con.Query ("SELECT DISTINCT(vendor) from purchases where vendor LIKE '%" +  searchString + "%'")
    
    if err != nil {
        panic (err)
    }
    
    for rows.Next() {
        var vendor string
        rows.Scan (&vendor)

        vendors = append (vendors, vendor)
    }
    
    return
}

func FindItems (searchString string) (items [] string, err error) {
    items = []string {}
    
    con := getConnection()
    defer con.Close()

    rows, err := con.Query ("SELECT DISTINCT(item) from purchases where item LIKE '%" +  searchString + "%'")
    
    if err != nil {
        panic (err)
    }
    
    for rows.Next() {
        var item string
        rows.Scan (&item)

        items = append (items, item)
    }
    
    return
}

func GetItemsToBuy () (items []string, err error) {
    items = []string {}
    con := getConnection()
    defer con.Close()

    rows, err := con.Query ("SELECT DISTINCT(item) from list")
    
    if err != nil {
        panic (err)
    }
    
    for rows.Next() {
        var item string
        rows.Scan (&item)

        items = append (items, item)
    }
    
    return
}

func RemoveFromItemsToBuy (item string) (err error) {
    con := getConnection()
    defer con.Close()

    stmt, err := con.Prepare ("DELETE FROM list WHERE item=?")
    
    if err != nil {
        panic (err)
    }
    
    _, err = stmt.Exec (item)
    return
}

func AddItemToBuy (item string) (err error) {
    con := getConnection()
    defer con.Close()

    stmt, err := con.Prepare ("INSERT INTO list (item) VALUES (?)")
    
    if err != nil {
        return 
    }

    _, err = stmt.Exec (item)
    return
}

func Save (item string, quantity float32, price float32, vendor string, date time.Time) (err error) {
    con := getConnection()
    defer con.Close()

    stmt, err := con.Prepare ("INSERT INTO purchases (item, price, quantity, vendor, date) VALUES (?,?,?,?,?)")
    
    if err != nil {
        return 
    }

    _, err = stmt.Exec (item, price, quantity, vendor, date)
    return
}
