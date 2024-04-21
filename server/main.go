package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// string to morse
var stm = map[rune]string{
	'A': ".-",
	'B': "-...",
	'C': "-.-.",
	'D': "-..",
	'E': ".",
	'F': "..-.",
	'G': "--.",
	'H': "....",
	'I': "..",
	'J': ".---",
	'K': "-.-",
	'L': ".-..",
	'M': "--",
	'N': "-.",
	'O': "---",
	'P': ".--.",
	'Q': "--.-",
	'R': ".-.",
	'S': "...",
	'T': "-",
	'U': "..-",
	'V': "...-",
	'W': ".--",
	'X': "-..-",
	'Y': "-.--",
	'Z': "--..",
	'0': "-----",
	'1': ".----",
	'2': "..---",
	'3': "...--",
	'4': "....-",
	'5': ".....",
	'6': "-....",
	'7': "--...",
	'8': "---..",
	'9': "----.",
}

func stringToMorse(input string) string {
	input = strings.ToUpper(input)
	var morseString strings.Builder

	for _, c := range input {
		if c == ' ' {
			continue
		}

		code, ok := stm[c]
		if !ok {
			morseString.WriteString("? ")
			continue
		}

		morseString.WriteString(code + " ")
	}

	return morseString.String()
}

func toMoresHandler(c *gin.Context) {
	input := c.Param("morse")

	morse := stringToMorse(input)
	c.JSON(200, gin.H{"morse": morse})
}

func main() {
	r := gin.Default()
	r.Use(corsMiddleware())
	r.GET("/toMorse/:morse", toMoresHandler)
	r.Run(":8800")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
