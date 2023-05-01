package handlers

import (
	dto "dumbflix-api/dto/result"
	transactiondto "dumbflix-api/dto/transaction"
	"dumbflix-api/models"
	repository "dumbflix-api/repositories"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gopkg.in/gomail.v2"
)

type handlerTransaction struct {
	TransactionRepository repository.TransactionRepository
}

func HandlerTransaction(TransactionRepository repository.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

func (h *handlerTransaction) FindTransactions(c echo.Context) error {
	transactions, err := h.TransactionRepository.FindTransactions()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: transactions})
}

func (h *handlerTransaction) GetTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := h.TransactionRepository.GetTransaction(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: transaction})
}

func (h *handlerTransaction) AddTransaction(c echo.Context) error {
	request := new(transactiondto.TransactionAddRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	Startdate := time.Now()
	Duedate := time.Now().Add(time.Hour * 24 * time.Duration(request.Days))

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	var transactionIsMatch = false
	var transactionId int
	for !transactionIsMatch {
		transactionId = int(time.Now().Unix())
		transactionData, _ := h.TransactionRepository.GetTransaction(transactionId)
		if transactionData.ID == 0 {
			transactionIsMatch = true
		}
	}

	transaction := models.Transaction{
		ID:        transactionId,
		StartDate: Startdate,
		DueDate:   Duedate,
		UserID:    int(userId),
		Price:     request.Price,
		Status:    "pending",
	}

	dataTransactions, err := h.TransactionRepository.AddTransaction(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	// 1. Initiate Snap client
	var s = snap.Client{}
	s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)
	// Use to midtrans.Production if you want Production Environment (accept real transaction).

	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(dataTransactions.ID),
			GrossAmt: int64(dataTransactions.Price),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: dataTransactions.User.Name,
			Email: dataTransactions.User.Email,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ := s.CreateTransaction(req)
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: snapResp})
}

func (h *handlerTransaction) DeleteTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.TransactionRepository.DeleteTransaction(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaction(data)})
}
func (h *handlerTransaction) Notification(c echo.Context) error {
	var notificationPayload map[string]interface{}

	if err := c.Bind(&notificationPayload); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)

	order_id, _ := strconv.Atoi(orderId)

	fmt.Print("This the payload:", notificationPayload)

	transaction, _ := h.TransactionRepository.GetTransaction(order_id)
	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {
			h.TransactionRepository.EditTransaction("pending", order_id)
		} else if fraudStatus == "accept" {
			SendMail("success", transaction)
			_, err := h.TransactionRepository.EditTransaction("success", order_id)
			if err != nil {
				fmt.Println(err)
			}
		}
	} else if transactionStatus == "settlement" {
		SendMail("success", transaction)
		_, err := h.TransactionRepository.EditTransaction("success", order_id)
		if err != nil {
			fmt.Println(err)
		}
	} else if transactionStatus == "deny" {
		h.TransactionRepository.EditTransaction("failed", order_id)
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		h.TransactionRepository.EditTransaction("failed", order_id)
	} else if transactionStatus == "pending" {
		h.TransactionRepository.EditTransaction("pending", order_id)
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: notificationPayload})
}

func SendMail(status string, transaction models.Transaction) {

	if status != transaction.Status && (status == "success") {
		var CONFIG_SMTP_HOST = "smtp.gmail.com"
		var CONFIG_SMTP_PORT = 587
		var CONFIG_SENDER_NAME = "Dumbflix <dumbflix@gmail.com>"
		var CONFIG_AUTH_EMAIL = os.Getenv("EMAIL_SYSTEM")
		var CONFIG_AUTH_PASSWORD = os.Getenv("PASSWORD_SYSTEM")

		var totalPrice = strconv.Itoa(transaction.Price)

		mailer := gomail.NewMessage()
		mailer.SetHeader("From", CONFIG_SENDER_NAME)
		mailer.SetHeader("To", transaction.User.Email)
		mailer.SetHeader("Subject", "Transaction Status")
		mailer.SetBody("text/html", fmt.Sprintf(`<!DOCTYPE html>
		<html lang="en">
		  <head>
		  <meta charset="UTF-8" />
		  <meta http-equiv="X-UA-Compatible" content="IE=edge" />
		  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
		  <title>Document</title>
		  <style>
			h1 {
			color: brown;
			}
		  </style>
		  </head>
		  <body>
		  <h2>Product payment :</h2>
		  <ul style="list-style-type:none;">
			<li>Total payment: Rp.%s</li>
			<li>Status : <b>%s</b></li>
		  </ul>
		  </body>
		</html>`, totalPrice, status))

		dialer := gomail.NewDialer(
			CONFIG_SMTP_HOST,
			CONFIG_SMTP_PORT,
			CONFIG_AUTH_EMAIL,
			CONFIG_AUTH_PASSWORD,
		)

		err := dialer.DialAndSend(mailer)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("Mail sent! to " + transaction.User.Email)
	}
}
func convertResponseTransaction(u models.Transaction) transactiondto.TransactionResponse {
	return transactiondto.TransactionResponse{
		ID:     u.ID,
		UserID: u.UserID,
		User:   models.User(u.User),
		Status: u.Status,
	}
}
