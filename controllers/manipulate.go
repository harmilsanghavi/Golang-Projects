package controllers

import (
	"encoding/json"
	"first-projectfinal/initializers"
	"firstproject/models"
	"firstproject/services"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/twilio/twilio-go"

	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	conversations "github.com/twilio/twilio-go/rest/conversations/v1"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

func init() {
	initializers.Loadenv()
}

func DeleteData(c *gin.Context) {
	_, err := c.Cookie("mycookie")
	if err != nil {
		c.Redirect(302, "/login")
	} else {
		id := c.Param("id")
		fmt.Println("delete :---", id)
		id1, _ := strconv.Atoi(id)
		services.DeleteDataID(id1)
		c.Redirect(302, "/page2")
	}
}

func EmailOtpGet(c *gin.Context) {
	c.HTML(http.StatusOK, "EmailOtp.html", nil)
}
func ForgetPasswordGet(c *gin.Context) {
	c.HTML(http.StatusOK, "forgetPass.html", nil)
}
func GetFileSMS(c *gin.Context) {
	c.HTML(http.StatusOK, "sendSMS.html", nil)
}

var OTP int

func GetOtp(c *gin.Context) {

	OTP = rand.Intn(900000) + 100000

	fmt.Println(OTP)

	msg := gomail.NewMessage()
	femail := c.PostForm("femail")
	msg.SetHeader("From", "harmil.sanghavi.23@gmail.com")
	msg.SetHeader("To", femail)
	msg.SetHeader("Subject", "Forget Password OTP")
	msg.SetBody("text/html", fmt.Sprintf("Your OTP is:%v\t", OTP))

	SendMail := gomail.NewDialer("smtp.gmail.com", 587, "harmil.sanghavi.23@gmail.com", "isemxvskmcpsftrub")

	// Send the email
	if err := SendMail.DialAndSend(msg); err != nil {
		panic(err)
	}
	c.Redirect(302, "/forgetpass")
	fmt.Println(femail)

}
func SetPassword(c *gin.Context) {
	emailotp := OTP
	checkotp := c.PostForm("checkotp")
	Cotp, _ := strconv.ParseInt(checkotp, 10, 64)
	epass := c.PostForm("epass")
	cpass := c.PostForm("cpass")

	if int(Cotp) != emailotp {
		fmt.Println("otp not matched")
	} else {
		fmt.Println("otp matched")
	}

	if cpass != epass {
		fmt.Println("password not matched")
	} else {
		fmt.Println("password matched")
	}

	fmt.Println("email otp", emailotp)
	fmt.Println("check otp", checkotp)
}
func SetChangePass(c *gin.Context) {
	oldpass := c.PostForm("oldpass")
	newpass := c.PostForm("newpass")
	confpass := c.PostForm("confpass")
	fmt.Println(oldpass)
	fmt.Println(newpass)
	fmt.Println(confpass)

	var Store []models.Demo
	initializers.DB.Select("password").Find(&Store)

	for _, dbpass := range Store {
		fmt.Println("pass", dbpass.Password)
		err := bcrypt.CompareHashAndPassword([]byte(dbpass.Password), []byte(oldpass))
		if err == nil {
			fmt.Println("matched")
		} else {
			fmt.Println("not matched")
		}
	}

	if confpass != newpass {
		fmt.Println("password not matched")
	} else {
		fmt.Println("password matched")
	}
}

func SendSMS(c *gin.Context) {
	accountSid := os.Getenv("ACCOUNTSID")
	authToken := os.Getenv("AUTHTOKEN")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(os.Getenv("TO_PHONE"))
	params.SetFrom(os.Getenv("FROM_PHONE"))
	params.SetBody("Hello Welcome In Twillio !")

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
		c.JSON(200, gin.H{
			"message": "Not Send",
		})
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
		c.JSON(200, gin.H{
			"message": "successfully Send",
		})
	}

}

func ReceiveSMS(c *gin.Context) {
	to := c.PostForm("sendnum")
	message := c.PostForm("msg")
	from := os.Getenv("FROM_PHONE")
	// Do something with the message (e.g. store it in a database, send a response)
	fmt.Printf("Received message from %s %s: %s\n", from, to, message)

	// Send a response back to Twilio
	accountSid := os.Getenv("ACCOUNTSID")
	authToken := os.Getenv("AUTHTOKEN")
	// TwillioNum := os.Getenv("FROM_PHONE")
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &conversations.CreateConversationParams{}

	// params.SetFriendlyName("Friendly Conversation")
	//resp, err := client.ConversationsV1.CreateConversation(params)
	// Create a conversation
	// convParams := &Conversations.CreateConversationParams{}

	conv, err := client.ConversationsV1.CreateConversation(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create conversation"})
		return
	}

	// Send a message to the conversation
	// messageParams := &conversations.CreateParams{
	// 	Body: &message,
	// 	From: &from,
	// 	To:   &to,
	// }

	// msg, err := client.ConversationsV1.RetrieveConversation(conv.Sid).Messages.Create(messageParams)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"conversation_sid": conv.Sid,
	})
	// Return a response to Twilio
	c.String(http.StatusOK, "")
}
