package greetings

import (
   "errors"
   "fmt"

)

func Hello(name string) (string, error) {

   if name == "" {
      return "", errors.New("name must not be empty")
   }

   message := fmt.Sprintf("Hii, %v. Welcome to go", name)
   return message, nil
}