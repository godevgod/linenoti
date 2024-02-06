# Simple Go LINE Notify

A minimalistic Go package for sending notifications to LINE using the LINE Notify API.

## Features

- Send text notifications
- Send notifications with images from a URL
- Send notifications with local images

## Installation

First, you need to install the package using `go get`:

```arduino
go get github.com/juunini/simple-go-line-notify/notify


Send a Text Notification
go
Copy code
func noti(accessToken, message string) {
	if err := notify.SendText(accessToken, message); err != nil {
		fmt.Println("Error sending notification:", err)
	}
}
Send a Notification with an Image from a URL
go
Copy code
func notiWithImage(accessToken, message, imageURL string) {
	if err := notify.SendImage(accessToken, message, imageURL); err != nil {
		fmt.Println("Error sending image notification:", err)
	}
}
Send a Notification with a Local Image
go
Copy code
func notiWithLocalImage(accessToken, message, imagePath string) {
	if err := notify.SendLocalImage(accessToken, message, imagePath); err != nil {
		fmt.Println("Error sending local image notification:", err)
	}
}
Example
Here's a simple example of how to use the package:

go
Copy code
package main

import (
	"fmt"
	"github.com/juunini/simple-go-line-notify/notify"
)

var token string = "YOUR_LINE_NOTIFY_TOKEN" // Replace with your LINE Notify token

func main() {
	accessToken := token
	message := "Hello, this is a test message."
	imageURL := "https://placebear.com/g/200/200" // Replace with your image URL
	imagePath := "test.png"                       // Replace with your local image path

	// Send a text notification
	noti(accessToken, message)

	// Send a notification with an image from URL
	notiWithImage(accessToken, message, imageURL)

	// Send a notification with a local image
	notiWithLocalImage(accessToken, message, imagePath)
}
Security Note
Make sure to keep your LINE Notify token secure. Do not hard-code it in your source code. Instead, consider using environment variables or a secure vault service.

Contribution
Contributions are welcome! Please feel free to submit a Pull Request.

License
Distributed under the MIT License. See LICENSE for more information.

Remember to replace YOUR_LINE_NOTIFY_TOKEN, https://example.com/image.jpg, and test.png with your actual LINE Notify token, image URL, and local image path respectively. Also, ensure you handle the security of your LINE Notify token properly.
