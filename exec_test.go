package reviuw

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExec(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer(id,name) VALUES('P002','Ulla')"

	_, err := db.ExecContext(ctx, query)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert Data Customer")
}

func TestQuery(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	query := "SELECT id,name FROM customer "

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID :", id)
		fmt.Println("NAME :", name)
	}

	fmt.Println("Success Get All Data Customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id,name,email,balance,rating,created_at,birth_date,married from customer"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int
		var rating float64
		var created_at time.Time
		var birth_date sql.NullTime
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &created_at, &birth_date, &married)
		if err != nil {
			panic(err)
		}

		fmt.Println("============================================")
		fmt.Println("ID :", id)
		fmt.Println("NAME :", name)
		if email.Valid {
			fmt.Println("EMAIL :", email.String)
		}
		fmt.Println("BALANCE :", balance)
		fmt.Println("RATING :", rating)
		fmt.Println("CreatedAt :", created_at)
		if birth_date.Valid {
			fmt.Println("Birth_Date :", birth_date.Time)
		}
		fmt.Println("Married :", married)
	}
}

func TestSqlInjection(t *testing.T) {

	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "arif"
	password := "salah"

	query := "SELECT username from user where username = ? AND password = ? LIMIT 1"

	rows, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login")
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestExecSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "eko"
	password := "eko"

	query := "INSERT INTO user(username,password)VALUES(?,?)"

	ctx := context.Background()

	_, err := db.ExecContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Sukses Insert Data User")
}

type Product struct {
	id           int32
	name_product string
	price        int64
	quantity     int64
}

func TestExecAutoIncrement(t *testing.T) {

	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	product := Product{
		name_product: "Mie Ayam",
		price:        12000,
		quantity:     20,
	}

	query := "INSERT INTO product(product_name,price,quantity)VALUES(?,?,?)"

	result, err := db.ExecContext(ctx, query, product.name_product, product.price, product.quantity)

	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Sukses Insert New Product With ID :", id)
}

func TestQueryAja(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id,product_name,price,quantity from product"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		product := Product{}
		err := rows.Scan(&product.id, &product.name_product, &product.price, &product.quantity)
		if err != nil {
			panic(err)
		}
		fmt.Println("====================================")
		fmt.Println("ID :", product.id)
		fmt.Println("Name_Product :", product.name_product)
		fmt.Println("Price :", product.price)
		fmt.Println("Quantity :", product.quantity)
	}
}
