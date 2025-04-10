package services

import (
	"errors"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	books   map[int]models.Book
	members map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		books:   make(map[int]models.Book),
		members: make(map[int]models.Member),
	}
}

func (l *Library) AddBook(book models.Book) {
	l.books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) {
	delete(l.books, bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, exists := l.books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Status == "Borrowed" {
		return errors.New("book already borrowed")
	}

	book.Status = "Borrowed"
	l.books[bookID] = book

	member := l.members[memberID]
	member.ID = memberID // in case it's a new member
	member.BorrowedBooks = append(member.BorrowedBooks, bookID)
	l.members[memberID] = member
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, exists := l.books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	book.Status = "Available"
	l.books[bookID] = book

	member, exists := l.members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	updatedBooks := []int{}
	for _, id := range member.BorrowedBooks {
		if id != bookID {
			updatedBooks = append(updatedBooks, id)
		}
	}
	member.BorrowedBooks = updatedBooks
	l.members[memberID] = member
	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	var available []models.Book
	for _, book := range l.books {
		if book.Status == "Available" {
			available = append(available, book)
		}
	}
	return available
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	var borrowed []models.Book
	member, exists := l.members[memberID]
	if !exists {
		return borrowed
	}
	for _, id := range member.BorrowedBooks {
		if book, ok := l.books[id]; ok {
			borrowed = append(borrowed, book)
		}
	}
	return borrowed
}
