package main

import (
	"flag"
	"fmt"
	"log"

	"com.github/Bmlla/MorseCodeTranslator/pkg/domain/entities/types"
	"com.github/Bmlla/MorseCodeTranslator/pkg/domain/translation"
)

func main() {
	var text, customDictionaryPath, mode string

	flag.StringVar(&text, "text", "", "Text to be translated")
	flag.StringVar(&customDictionaryPath, "dict", types.LATIN, "custom dictionary path")
	flag.StringVar(&mode, "mode", "", "Translation mode: 'from' or 'to' Morse Code")

	flag.Parse()

	translator, err := translation.New(customDictionaryPath)
	if err != nil {
		log.Fatal()
	}

	if text == "" {
		fmt.Println("Please provide text to translate using the -text flag")

		return
	}

	if mode == "to" {
		r, err := translator.ToMorse(text)
		if err != nil {
			fmt.Print(err)
			return
		}

		fmt.Print(r)
		return
	}

	if mode == "from" {
		r, err := translator.FromMorse(text)
		if err != nil {
			fmt.Print(err)
			return
		}

		fmt.Print(r)
		return
	}

	fmt.Println("Please provide a valid mode")
}
