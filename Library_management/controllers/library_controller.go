package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
)

var library = services.NewLibrary()

func RunLibraryConsole() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Exit")
		fmt.Print("Enter choice: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter book ID: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Enter book title: ")
			scanner.Scan()
			title := scanner.Text()
			fmt.Print("Enter author name: ")
			scanner.Scan()
			author := scanner.Text()
			book := models.Book{ID: id, Title: title, Author: author, Status: "Available"}
			library.AddBook(book)
		case "2":
			fmt.Print("Enter book ID to remove: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			library.RemoveBook(id)
		case "3":
			fmt.Print("Enter member ID: ")
			scanner.Scan()
			memberID, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Enter book ID to borrow: ")
			scanner.Scan()
			bookID, _ := strconv.Atoi(scanner.Text())
			err := library.BorrowBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "4":
			fmt.Print("Enter member ID: ")
			scanner.Scan()
			memberID, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Enter book ID to return: ")
			scanner.Scan()
			bookID, _ := strconv.Atoi(scanner.Text())
			err := library.ReturnBook(bookID, memberID)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "5":
			books := library.ListAvailableBooks()
			for _, book := range books {
				fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
			}
		case "6":
			fmt.Print("Enter member ID: ")
			scanner.Scan()
			memberID, _ := strconv.Atoi(scanner.Text())
			books := library.ListBorrowedBooks(memberID)
			for _, book := range books {
				fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
			}
		case "7":
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
