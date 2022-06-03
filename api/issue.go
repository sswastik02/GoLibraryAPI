package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sswastik02/GoLibraryAPI/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IssueReq struct{
	Username string `json:"username"`
	BookID uint `json:"book_id"`
}


func(r* Repository) getBookFromId(bookID uint) (*models.Book,int,error) {

	book := models.Book{}
	response := r.DB.Where("id = ?",bookID).First(&book)
	err := response.Error

	if err != nil {
		if err == gorm.ErrRecordNotFound{
		return nil,http.StatusNotFound,fmt.Errorf("book not found")
		}
		return nil,http.StatusInternalServerError,fmt.Errorf("error while fetching book")
	}

	return &book,http.StatusOK,nil
}

func getIssueFromRequest(context *fiber.Ctx) (*IssueReq,int,error) {
	issueReq := IssueReq{}
	err := context.BodyParser(&issueReq)
	
	if err != nil {
		return nil,http.StatusBadRequest,fmt.Errorf("request should contain username and book id")
	}
	return &issueReq,http.StatusOK,nil
}

func(r* Repository) createIssueRecordForUserWithBook(book *models.Book,username string) (int,error) {
	issue := models.Issue{
		Username: username,
		Book: *book,
	}
	issueRecord := models.IssuesRecord{
		Username: username,
		Issues: []models.Issue{},
	}

	response := r.DB.FirstOrCreate(&issueRecord,"username = ?",username)
	err := response.Error
	
	if err != nil {
		return http.StatusInternalServerError,fmt.Errorf("error while issueing record")
	}
	
	issueRecord.Issues = append(issueRecord.Issues, issue)
	response = r.DB.Updates(&issueRecord)
	err = response.Error

	if err != nil {
		if(strings.Contains(err.Error(),"duplicate key value violates unique constraint")) {
			return http.StatusConflict ,fmt.Errorf("cannot issue book, already issued")
		}
		return http.StatusInternalServerError,fmt.Errorf("error adding new issue record")
	}

	return http.StatusOK,nil
}



func(r* Repository) entryIssue(context *fiber.Ctx) error {
	
	issueReq,httpcode,err := getIssueFromRequest(context)

	if err != nil {
		context.Status(httpcode).JSON(
			&fiber.Map{
				"message":err.Error(),
			},
		)
		return nil
	}


	book,httpcode,err := r.getBookFromId(issueReq.BookID)

	if err != nil {
		context.Status(httpcode).JSON(
			&fiber.Map{
				"message":err.Error(),
			},
		)
		return nil
	}

	httpcode, err = r.createIssueRecordForUserWithBook(book,issueReq.Username)
	if err != nil {
		context.Status(httpcode).JSON(
			&fiber.Map{
				"message":err.Error(),
			},
		)
		return nil
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message":"Book Issued Successfully",
		},
	)

	return nil

	
}

func(r* Repository) getIssues(context *fiber.Ctx) error {
	user,err := getUserFromContext(context)
	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message":err.Error(),
			},
		)
		return nil
	}
	var response *gorm.DB
	issueRecord := models.IssuesRecord{}
	issueRecords := []models.IssuesRecord{}
	if user.Admin {
		response = r.DB.Preload("Issues.Book").Preload(clause.Associations).Find(&issueRecords)
	} else {
		response = r.DB.Preload("Issues.Book").Preload(clause.Associations).Find(&issueRecord,"username = ?",user.Username)
	}
	err = response.Error

	if err != nil {
		context.Status(http.StatusNotFound).JSON(
			&fiber.Map{
				"message":"Could not find issue record for user",
			},
		)
		return nil
	}

	if(user.Admin) {
		context.Status(http.StatusOK).JSON(
			&fiber.Map{
			"message":"Issue Records Found",
			"data":issueRecords,
			},
		)
	} else {

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message":"Issue Record Found",
			"data": issueRecord,
		},
	)

	}
	return nil
}

func(r* Repository) deleteIssue(context *fiber.Ctx) error {
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"message":"must contain issue id",
			},
		)
		return nil
	}

	uid , err := strconv.Atoi(id)

	if err != nil {
		log.Fatal("Error while converting ID to uint type")
	}

	issue := models.Issue{IssueID: uint(uid)}
	response := r.DB.Delete(&issue)
	err = response.Error
	if err != nil || response.RowsAffected<1 {
		if response.RowsAffected < 1 {
			context.Status(http.StatusNotFound).JSON(
				&fiber.Map{
					"message":"Issue Not found",
				},
			)
			return nil
		}

		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message":"Error while Deleting issue",
			},
		)

		return nil
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message":"Deleted Issue Successfuly",
		},
	)

	return nil

}