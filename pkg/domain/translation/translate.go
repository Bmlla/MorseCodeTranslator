package translation

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"com.github/Bmlla/MorseCodeTranslator/pkg/domain/entities"
)

type MorseCodeTranslator struct {
	dictionary *entities.Dictionary
}

func New(customDictionaryPath string) (*MorseCodeTranslator, error) {
	if customDictionaryPath == "" {
		return nil, fmt.Errorf("dictionary path cannot be empty")
	}

	dictionary, err := loadDictionary(customDictionaryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load dictionary: %w", err)
	}

	return &MorseCodeTranslator{
		dictionary: dictionary,
	}, nil
}

func loadDictionary(path string) (*entities.Dictionary, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read dictionary file: %w", err)
	}

	var dictionary *entities.Dictionary
	if err := json.Unmarshal(bytes, &dictionary); err != nil {
		return nil, fmt.Errorf("failed to unmarshal dictionary: %w", err)
	}

	return dictionary, nil
}

func (mct *MorseCodeTranslator) FromMorse(input string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("input cannot be empty")
	}

	trimmedInput := strings.TrimSpace(input)

	if validInput := validateMorseCode(trimmedInput); !validInput {
		return "", fmt.Errorf("input is not a valid morse code")
	}

	splitted := strings.Split(trimmedInput, "")

	var seemsToBeAnotherWord bool
	var buffer, output strings.Builder

	for k, v := range splitted {
		if v == " " {
			if buffer.Len() > 0 {
				translated, err := findReferenceInDictionary(mct.dictionary.From, buffer.String())
				if err != nil {
					return "", err
				}

				output.WriteString(translated)
				buffer.Reset()

				continue
			}

			seemsToBeAnotherWord = true
			continue
		}

		if seemsToBeAnotherWord {
			output.WriteString(" ")
			seemsToBeAnotherWord = false
		}

		buffer.WriteString(v)

		if k == len(splitted)-1 {
			translated, err := findReferenceInDictionary(mct.dictionary.From, buffer.String())
			if err != nil {
				return "", err
			}

			output.WriteString(translated)
		}
	}

	return output.String(), nil
}

func (mct *MorseCodeTranslator) ToMorse(input string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("input cannot be empty")
	}

	input = strings.ToUpper(strings.TrimSpace(input))
	chars := strings.Split(input, "")

	var output strings.Builder

	for i, char := range chars {
		if char == " " {
			output.WriteString("  ")
			continue
		}

		morse, err := findReferenceInDictionary(mct.dictionary.To, char)
		if err != nil {
			return "", fmt.Errorf("character '%s' cannot be translated: %w", char, err)
		}

		output.WriteString(morse)

		if i < len(chars)-1 && chars[i+1] != " " {
			output.WriteString(" ")
		}
	}

	return output.String(), nil
}

func validateMorseCode(input string) bool {
	re, _ := regexp.Compile(`^(\.|-)+( {1,}(\.|-)+)*$`)

	return re.MatchString(input)
}

func findReferenceInDictionary(dictionary map[string]string, key string) (string, error) {
	translated, ok := dictionary[key]
	if !ok {
		return "", fmt.Errorf("could not translate word %s. This word is not present in dictionary", key)
	}

	return translated, nil
}
