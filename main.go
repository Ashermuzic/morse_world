package main

import (
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
	' ': "/",
}

func stringToMorse(input string) string {
	var morseString strings.Builder

	for _, c := range input {
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
	var input struct {
		Text string `json:"text"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	morse := stringToMorse(input.Text)

	c.JSON(200, gin.H{"morse": morse})
}

func main() {
	r := gin.Default()
	r.POST("/toMorse", toMoresHandler)

	r.Run(":8800")
}
