// Some simple examples of:
// * defining and initializing a struct
// * marshalling and unmarshalling JSON
// * using "reflect" to compare non-trivial structs
// * unmarshalling an arbitrary blob of JSON
//
// For more, see: https://gobyexample.com/json
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Email is our sample object we wish to convert to/from JSON
type Attachment struct {
	MimeType string
	Content  []byte
}

type Email struct {
	To          []string
	Cc          []string
	Body        string
	Attachments []Attachment
}

func main() {
	fmt.Println("Begin JSON demo...")
	// create a simple email message (one recipient, no attachments)
	msg1 := Email{
		To:          []string{"Bill Gates"},
		Cc:          []string{},
		Body:        "Hi, I'm an email message!",
		Attachments: nil,
	}
	fmt.Printf("\noriginal message: %v\n\tObject type: %T\n", msg1, msg1)
	jsonMsg, err := json.Marshal(msg1)
	if err != nil {
		fmt.Println("Failed while converting to JSON: ", err)
		return
	}
	fmt.Println("\nmessage as JSON :", string(jsonMsg))

	// convert JSON back to Go object

	// Method 1: When you know the structure represented by the JSON
	msg2 := Email{}
	err = json.Unmarshal(jsonMsg, &msg2)
	if err != nil {
		fmt.Println("Failed while converting from JSON: ", err)
		return
	}
	fmt.Printf("\nunmarshalled message: %v\n\tObject type: %T\n", msg2, msg2)

	if reflect.DeepEqual(msg1, msg2) {
		fmt.Println("The unmarshalled message equals our original message.")
	} else {
		fmt.Println("Ruh-roh! Something was lost in translation!")
	}

	// Method 2: When you do NOT know the structure represented by the JSON
	var arbitraryStruct map[string]interface{}
	err = json.Unmarshal(jsonMsg, &arbitraryStruct)
	if err != nil {
		fmt.Println("Failed while converting from JSON: ", err)
		return
	}
	// iterate over the map to see what's inside
	fmt.Println("\nContents of our 'arbitraryStruct'")
	for key, value := range arbitraryStruct {
		fmt.Printf("\t[%v] = %v\n", key, value)
	}

	fmt.Println("End JSON demo.  Wasn't that fun?")
}
