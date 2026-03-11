package database

import (
	"MackaWebsite/internal/models"
	"database/sql"
	"fmt"

	_ "github.com/glebarez/go-sqlite"
)

var db *sql.DB

func ConnectToDb() error {
	var err error
	if db, err = sql.Open("sqlite", "./app.db"); err != nil {
		fmt.Println("Error when connecting to the db: ", err)
		return err
	}

	return err
}

func InitializeDatatable() error {
	sql := `
		CREATE TABLE IF NOT EXISTS admins(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL UNIQUE
		);
	`

	if _, err := db.Exec(sql); err != nil {
		fmt.Println("Error when initializing DTs: ", err)
		return err
	}

	sql = `
		CREATE TABLE IF NOT EXISTS landing_page(
			title TEXT NOT NULL PRIMARY KEY,
			subtitle TEXT NOT NULL,
			text TEXT NOT NULL,
			image_link TEXT NOT NULL
		);
	`

	if _, err := db.Exec(sql); err != nil {
		fmt.Println("Error when initializing DTs: ", err)
		return err
	}

	sql = `
		CREATE TABLE IF NOT EXISTS blogs(
			title TEXT NOT NULL PRIMARY KEY,
			subtitle TEXT NOT NULL,
			text TEXT NOT NULL,
			image_link TEXT NOT NULL
		);
	`

	if _, err := db.Exec(sql); err != nil {
		fmt.Println("Error when initializing DTs: ", err)
		return err
	}

	return nil
}

func CreatTheFirstUser() error {
	sql := "INSERT INTO admins (username, password) VALUES (?, ?)"
	admin := "admin"
	password := "password"
	_, err := db.Exec(sql, admin, password)
	if err != nil {
		fmt.Println("Error when inserting the first admin: ", err)
		return err
	}

	return nil
}

func SelectUserPasswordByUsername(username string) (string, string, error) {
	query := "SELECT id, password FROM admins WHERE username = ?"
	row := db.QueryRow(query, username)
	var id, password string
	if err := row.Scan(&id, &password); err != nil {
		fmt.Println("Error when selecting user password from db by username: ", err)
		return "", "", err
	}

	return id, password, nil
}

func SelectLandingPageFromDb() (*models.LandingPage, error) {
	query := "SELECT * FROM landing_page"
	row := db.QueryRow(query)
	var landingPage models.LandingPage
	if err := row.Scan(&landingPage.Title, &landingPage.SubTitle, &landingPage.Text); err != nil {
		fmt.Println("Error when selecting landingPage from db: ", err)
		return nil, err
	}

	return &landingPage, nil
}

func SelectBlogsFromDb() ([]models.Blog, error) {
	query := "SELECT * FROM blogs"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error when selecting blogs from db: ", err)
		return nil, err
	}
	defer rows.Close()

	var blogs []models.Blog
	for rows.Next() {
		var c models.Blog
		if err := rows.Scan(&c.Title, &c.SubTitle, &c.Text); err != nil {
			fmt.Println("Error when selecting blogs from db: ", err)
			return nil, err
		}
		blogs = append(blogs, c)
	}

	return blogs, nil
}

func SelectBlogFromDbById(title string) (*models.Blog, error) {
	query := "SELECT * FROM blogs WHERE title = ?"
	row := db.QueryRow(query, title)
	var blog models.Blog
	if err := row.Scan(&blog.Title, &blog.SubTitle, &blog.Text); err != nil {
		fmt.Println("Error when selecting blog from db: ", err)
		return nil, err
	}

	return &blog, nil
}

func InsertBlogIntoDb(blog *models.Blog) error {
	sql := "INSERT INTO blogs (title, subtitle, text, image_link) VALUES (?,?,?,?)"
	if _, err := db.Exec(sql, blog.Title, blog.SubTitle, blog.Text, blog.ImageLink); err != nil {
		fmt.Println("Error when selecting blog from db: ", err)
		return err
	}

	return nil
}

func UpdateBlogImageLinkByTitleInDb(imageLink string, title string) error {
	sql := "UPDATE blogs SET image_link = ? WHERE title = ?"
	if _, err := db.Exec(sql, imageLink, title); err != nil {
		fmt.Println("Error when updating blog image link by title: ", err)
		return err
	}

	return nil
}

func UpdateBlogByTitleInDb(blog *models.Blog) error {
	query := "UPDATE blogs SET title = ?, subtitle = ?, text = ?, image_link = ? WHERE title = ?"
	if _, err := db.Exec(query, blog.Title, blog.SubTitle, blog.Text, blog.ImageLink, blog.Title); err != nil {
		fmt.Println("Error when updating blog in db: ", err)
		return err
	}

	return nil
}
