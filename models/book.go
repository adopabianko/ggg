package models

import (
	"github.com/adopabianko/ggg-boilerplate/config"

	_ "github.com/go-sql-driver/mysql"
)

// Fetch all book data
func GetAllBooks(book *[]Book) (err error) {
	if err = config.DB.Find(book).Error; err != nil {
		return err
	}

	return nil
}

// Insert new data
func CreateBook(book *Book) (err error) {
	if err = config.DB.Create(book).Error; err != nil {
		return err
	}

	return nil
}

// Fetch only one book by id
func GetBookByID(book *Book, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(book).Error; err != nil {
		return err
	}

	return nil
}

// Update book
func UpdateBook(book *Book) (err error) {
	if err = config.DB.Save(book).Error; err != nil {
		return err
	}

	return nil
}

// Delete book
func DeleteBook(book *Book, id string) (err error) {
	if err = config.DB.Where("id = ?", id).Delete(book).Error; err != nil {
		return err
	}

	return nil
}
