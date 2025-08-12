package translation_test

import (
	"path/filepath"
	"runtime"
	"testing"

	"com.github/Bmlla/MorseCodeTranslator/pkg/domain/entities/types"
	"com.github/Bmlla/MorseCodeTranslator/pkg/domain/translation"
)

func mountFilePath(path string) string {
	_, currentFile, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(currentFile)

	absPath := filepath.Join(baseDir, "..", "..", "..", path)

	return absPath
}

func TestNew_WhenACustomDictionaryIsNotGave_shouldReturnError(t *testing.T) {
	_, err := translation.New("")
	if err == nil {
		t.Errorf("Expected error for empty dictionary path, got nil")
	}
}

func TestNew_WhenACustomDictionaryIsGave_shouldReturnErrorIfDictionaryIsNotFound(t *testing.T) {
	_, err := translation.New("fake_path.json")
	if err == nil {
		t.Errorf("Expected error for non-existent dictionary, got nil")
	}
}

func TestNew_WhenACustomDictionaryIsGave_shouldReturnErrorIfFileIsNotInTheRightFormat(t *testing.T) {
	_, err := translation.New(mountFilePath("pkg/dictionaries/invalid.json"))
	if err == nil {
		t.Errorf("Expected error for invalid dictionary format, got nil")
	}
}

func TestFromMorse_WhenEmptyInput_ShouldReturnError(t *testing.T) {
	translator, _ := translation.New(mountFilePath(mountFilePath(types.LATIN)))

	phrase, err := translator.FromMorse("")
	if err == nil {
		t.Errorf("Expected error for empty input, got nil")
	}

	if phrase != "" {
		t.Errorf("Expected correct single letter, got %v", err)
	}
}

func TestFromMorse_WhenInputIsAValidMorseCodeSingleLetter_ShouldReturnAtranslationdLetter(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.FromMorse(".")
	if err != nil {
		t.Errorf("Expected no error for valid input, got %v", err)
	}

	if phrase != "E" {
		t.Errorf("Expected correct single letter, got %v", err)
	}
}

func TestFromMorse_WhenInputIsAValidMorseCodeSingleWord_ShouldReturnAtranslationdWord(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.FromMorse(".... . .-.. .-.. ---")
	if err != nil {
		t.Errorf("Expected no error for valid input, got %v", err)
	}

	if phrase != "HELLO" {
		t.Errorf("Expected correct single letter, got %v", phrase)
	}
}

func TestFromMorse_WhenInputIsAValidPhrase_ShouldReturnAtranslationdPhrase(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.FromMorse(".... . .-.. .-.. ---   .-- --- .-. .-.. -..")
	if err != nil {
		t.Errorf("Expected no error for valid input, got %v", err)
	}

	if phrase != "HELLO WORLD" {
		t.Errorf("Expected correct single letter, got %v", phrase)
	}
}

func TestFromMorse_WhenInputIsAValidPhraseWithMoreThanTwoWords_ShouldReturnAtranslationdPhrase(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.FromMorse(".... . .-.. .-.. ---   .-- --- .-. .-.. -..   -... ..- -.. -.. -.--")
	if err != nil {
		t.Errorf("Expected no error for valid input, got %v", err)
	}

	if phrase != "HELLO WORLD BUDDY" {
		t.Errorf("Expected correct single letter, got %v", phrase)
	}
}

func TestFromMorse_WhenInputIsAValidPhraseWithSpecialCharacters_ShouldReturnAtranslationdPhrase(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.FromMorse(".... . .-.. .-.. --- --..--   .-- --- .-. .-.. -.. -.-.--")
	if err != nil {
		t.Errorf("Expected no error for valid input, got %v", err)
	}

	if phrase != "HELLO, WORLD!" {
		t.Errorf("Expected correct single letter, got %v", phrase)
	}
}

func TestFromMorse_WhenInputIsAValidPhraseWithNumbers_ShouldReturnAtranslationdPhrase(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.FromMorse(".... . .-.. .-.. ---   .-- --- .-. .-.. -..   ....-   -.-- --- ..-   ..---")
	if err != nil {
		t.Errorf("Expected no error for valid input, got %v", err)
	}

	if phrase != "HELLO WORLD 4 YOU 2" {
		t.Errorf("Expected correct single letter, got %v", phrase)
	}
}

func TestFromMorse_WhenInputIsAValidPhraseWithNumbersAndSpecialChars_ShouldReturnAtranslationdPhrase(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.FromMorse(".... . .-.. .-.. ---   .-- --- .-. .-.. -.. --..--   ....-   -.-- --- ..-   ..--- -.-.--")
	if err != nil {
		t.Errorf("Expected no error for valid input, got %v", err)
	}

	if phrase != "HELLO WORLD, 4 YOU 2!" {
		t.Errorf("Expected correct single letter, got %v", phrase)
	}
}

func TestFromMorse_WhenInputIsAValidPhraseAndSpacesBetweenWordsIsHigherThanThree_ShouldReturnAtranslationdPhraseWithASingleSpace(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.FromMorse(".... . .-.. .-.. --- --..--        .-- --- .-. .-.. -.. -.-.--")
	if err != nil {
		t.Errorf("Expected no error for valid input, got %v", err)
	}

	if phrase != "HELLO, WORLD!" {
		t.Errorf("Expected correct single letter, got %v", phrase)
	}
}

func TestFromMorse_WhenInputIsAValidPhraseAndSpacesBetweenWordsAreTwo_ShouldReturnAtranslationdPhraseWithASingleSpace(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.FromMorse(".... . .-.. .-.. --- --..--  .-- --- .-. .-.. -.. -.-.--")
	if err != nil {
		t.Errorf("Expected no error for valid input, got %v", err)
	}

	if phrase != "HELLO, WORLD!" {
		t.Errorf("Expected correct single letter, got %v", phrase)
	}
}

func TestFromMorse_WhenInputIsAValidPhraseWithSingleSpaceInStart_ShouldReturnAtranslationdPhraseDesconsideringStartSpace(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.FromMorse(" .... . .-.. .-.. --- --..--   .-- --- .-. .-.. -.. -.-.--")
	if err != nil {
		t.Errorf("Expected no error for valid input, got %v", err)
	}

	if phrase != "HELLO, WORLD!" {
		t.Errorf("Expected correct single letter, got %v", phrase)
	}
}

func TestFromMorse_WhenInputIsAValidPhraseWithManySpacesInStart_ShouldReturnAtranslationdPhraseDesconsideringStartSpace(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.FromMorse("    .... . .-.. .-.. --- --..--   .-- --- .-. .-.. -.. -.-.--")
	if err != nil {
		t.Errorf("Expected no error for valid input, got %v", err)
	}

	if phrase != "HELLO, WORLD!" {
		t.Errorf("Expected correct single letter, got %v", phrase)
	}
}

func TestFromMorse_WhenInputIsAValidPhraseWithSingleSpaceInTheEnd_ShouldReturnAtranslationdPhraseDesconsideringEndSpace(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.FromMorse(".... . .-.. .-.. --- --..--   .-- --- .-. .-.. -.. -.-.-- ")
	if err != nil {
		t.Errorf("Expected no error for valid input, got %v", err)
	}

	if phrase != "HELLO, WORLD!" {
		t.Errorf("Expected correct single letter, got %v", phrase)
	}
}

func TestFromMorse_WhenInputIsAValidPhraseWithManySpacesInTheEnd_ShouldReturnAtranslationdPhraseDesconsideringEndSpaces(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.FromMorse(".... . .-.. .-.. --- --..--   .-- --- .-. .-.. -.. -.-.--     ")
	if err != nil {
		t.Errorf("Expected no error for valid input, got %v", err)
	}

	if phrase != "HELLO, WORLD!" {
		t.Errorf("Expected correct single letter, got %v", phrase)
	}
}

func TestFromMorse_WhenInputIsAValidPhraseWithSingleBorderSpaces_ShouldReturnAtranslationdPhraseDesconsideringBorderSpaces(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.FromMorse(" .... . .-.. .-.. --- --..--   .-- --- .-. .-.. -.. -.-.-- ")
	if err != nil {
		t.Errorf("Expected no error for valid input, got %v", err)
	}

	if phrase != "HELLO, WORLD!" {
		t.Errorf("Expected correct single letter, got %v", phrase)
	}
}

func TestFromMorse_WhenInputIsAValidPhraseWithManyBorderSpaces_ShouldReturnAtranslationdPhraseDesconsideringBorderSpaces(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.FromMorse("    .... . .-.. .-.. --- --..--   .-- --- .-. .-.. -.. -.-.--    ")
	if err != nil {
		t.Errorf("Expected no error for valid input, got %v", err)
	}

	if phrase != "HELLO, WORLD!" {
		t.Errorf("Expected correct single letter, got %v", phrase)
	}
}

func TestFromMorse_WhenInputHasInvalidLetter_ShouldReturnError(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.FromMorse("........ . .-.. .-.. --- --..--   .-- --- .-. .-.. -.. -.-.--")
	if err == nil {
		t.Errorf("Expected error for invalid input")
	}

	if phrase != "" {
		t.Errorf("Expected empty output for invalid translation")
	}
}

func TestFromMorse_WhenInputHasInvalidLetterAtTheEnd_ShouldReturnError(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.FromMorse(".... . .-.. .-.. --- --..--   .-- --- .-. .-.. -.. ........")
	if err == nil {
		t.Errorf("Expected error for invalid input")
	}

	if phrase != "" {
		t.Errorf("Expected empty output for invalid translation")
	}
}

func TestFromMorse_WhenInputIHasElementInStartThatIsNotAMorseCode_ShouldReturnError(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.FromMorse("test . .-.. .-.. --- --..--   .-- --- .-. .-.. -.. -.-.--")
	if err == nil {
		t.Errorf("Expected error for invalid input")
	}

	if phrase != "" {
		t.Errorf("Expected empty output for invalid translation")
	}
}

func TestFromMorse_WhenInputHasElementInMiddleThatIsNotAMorseCode_ShouldReturnError(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.FromMorse(".... . .-.. .-.. test   .-- --- .-. .-.. -.. -.-.--")
	if err == nil {
		t.Errorf("Expected error for invalid input")
	}

	if phrase != "" {
		t.Errorf("Expected empty output for invalid translation")
	}
}

func TestFromMorse_WhenInputIHasElementInEndThatIsNotAMorseCode_ShouldReturnError(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.FromMorse(".... . .-.. .-.. --- --..--   .-- --- .-. .-.. -.. test")
	if err == nil {
		t.Errorf("Expected error for invalid input")
	}

	if phrase != "" {
		t.Errorf("Expected empty output for invalid translation")
	}
}

// ToMorse Tests
func TestToMorse_WhenEmptyInput_ShouldReturnError(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.ToMorse("")
	if err == nil {
		t.Errorf("Expected error for empty input, got nil")
	}

	if phrase != "" {
		t.Errorf("Expected empty result, got %v", phrase)
	}
}

func TestToMorse_WhenInputIsASingleLetter_ShouldReturnMorseCode(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.ToMorse("E")
	if err != nil {
		t.Errorf("Expected no error for valid input, got %v", err)
	}

	if phrase != "." {
		t.Errorf("Expected correct morse code, got %v", phrase)
	}
}

func TestToMorse_WhenInputIsASingleWord_ShouldReturnMorseCode(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.ToMorse("HELLO")
	if err != nil {
		t.Errorf("Expected no error for valid input, got %v", err)
	}

	if phrase != ".... . .-.. .-.. ---" {
		t.Errorf("Expected correct morse code, got %v", phrase)
	}
}

func TestToMorse_WhenInputIsAPhrase_ShouldReturnMorseCode(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.ToMorse("HELLO WORLD")
	if err != nil {
		t.Errorf("Expected no error for valid input, got %v", err)
	}

	if phrase != ".... . .-.. .-.. ---  .-- --- .-. .-.. -.." {
		t.Errorf("Expected correct morse code, got %v", phrase)
	}
}

func TestToMorse_WhenInputIsLowerCase_ShouldReturnMorseCode(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.ToMorse("hello world")
	if err != nil {
		t.Errorf("Expected no error for valid input, got %v", err)
	}

	if phrase != ".... . .-.. .-.. ---  .-- --- .-. .-.. -.." {
		t.Errorf("Expected correct morse code, got %v", phrase)
	}
}

func TestToMorse_WhenInputHasSpecialCharacters_ShouldReturnMorseCode(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.ToMorse("HELLO, WORLD!")
	if err != nil {
		t.Errorf("Expected no error for valid input, got %v", err)
	}

	if phrase != ".... . .-.. .-.. --- --..--  .-- --- .-. .-.. -.. -.-.--" {
		t.Errorf("Expected correct morse code, got %v", phrase)
	}
}

func TestToMorse_WhenInputHasNumbers_ShouldReturnMorseCode(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.ToMorse("HELLO WORLD 4 YOU 2")
	if err != nil {
		t.Errorf("Expected no error for valid input, got %v", err)
	}

	if phrase != ".... . .-.. .-.. ---  .-- --- .-. .-.. -..  ....-  -.-- --- ..-  ..---" {
		t.Errorf("Expected correct morse code, got %v", phrase)
	}
}

func TestToMorse_WhenInputHasInvalidCharacter_ShouldReturnError(t *testing.T) {
	translator, _ := translation.New(mountFilePath(types.LATIN))

	phrase, err := translator.ToMorse("HELLO£WORLD")
	if err == nil {
		t.Errorf("Expected error for invalid character")
	}

	if phrase != "" {
		t.Errorf("Expected empty output for invalid translation, got %v", phrase)
	}
}
