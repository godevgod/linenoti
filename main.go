package main

import (
	"fmt"

	"github.com/juunini/simple-go-line-notify/notify"
)

func noti(accessToken, message string) {
	if err := notify.SendText(accessToken, message); err != nil {
		fmt.Println("Error sending notification:", err)
	}
}

func notiWithImage(accessToken, message, imageURL string) {
	if err := notify.SendImage(accessToken, message, imageURL); err != nil {
		fmt.Println("Error sending image notification:", err)
	}
}

func notiWithLocalImage(accessToken, message, imagePath string) {
	if err := notify.SendLocalImage(accessToken, message, imagePath); err != nil {
		fmt.Println("Error sending local image notification:", err)
	}

	fmt.Println("line noti:", message)
}

var token string = "xxx" //line notify token

func main() {
	//accessToken := "YOUR_ACCESS_TOKEN_HERE" // Replace with your LINE access token
	accessToken := token
	message := "Hello, this is a test message."
	imageURL := "https://placebear.com/g/200/200" //"https://example.com/image.jpg" // Replace with your image URL
	imagePath := "test.png"                       // Replace with your local image path

	// Send a text notification
	noti(accessToken, message)

	// Send a notification with an image from URL
	notiWithImage(accessToken, message, imageURL)

	// Send a notification with a local image
	notiWithLocalImage(accessToken, message, imagePath)
}
