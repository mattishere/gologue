package gologue

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/term"
)

// The Character struct is the core for using Gologue, as most events are passed directly via it.
type Character struct {
	Name      string
	NameColor Color
}

// Creates a new character (uses the most generic options, make your own Character struct for finer tuning).
func NewCharacter(name string) Character {
    return Character{
        Name: name,
        NameColor: ColorWhite,
    }
}

// Settings is a struct used for setting things like speed, colors, and the cutoff. 
type Settings struct {
	// Speed is in Milliseconds
	Speed       time.Duration
	BorderColor Color
	TextColor   Color
    // Cutoff controls how soon you'll go to a new line. E.g., if Cutoff is 2, you'll go to a new line 6 characters sooner (4 + 2)
	Cutoff      int
}
// Creates a new Settings struct (uses the most generic options, make your own Settings struct for finer tuning).
func NewSettings(speed time.Duration) Settings {
    return Settings{
        Speed: speed,
        BorderColor: ColorWhite,
        TextColor: ColorWhite,
        Cutoff: 0,
    }
}

/*
    Talk starts the talking event of a character.
*/
func (c Character) Talk(text string, settings Settings) error {
	width, _, err := term.GetSize(0)
	if err != nil {
		return err
	}

	var words []string
	for _, element := range strings.Split(text, " ") {
		if strings.Contains(element, "\n") {
			splitText := strings.Split(element, "\n")
			for i, splitElement := range splitText {
				if i < len(splitText)-1 {
					words = append(words, splitElement+"\n")
				} else {
					words = append(words, splitElement)
				}
			}
		} else {
			words = append(words, element+" ")
		}
	}

	fmt.Println(string(settings.BorderColor) + "+-----{ " + string(c.NameColor) + c.Name + string(settings.BorderColor) + " }-----")

	lines := getLines(words, width-4-settings.Cutoff, settings)
	for _, line := range lines {
		fmt.Print(string(settings.BorderColor) + "| ")
		for i := range line {
			fmt.Print(string(settings.TextColor), line[i])
			time.Sleep(settings.Speed * time.Millisecond)
		}
		fmt.Println(string(ColorReset))
	}
	return nil
}


func getLines(text []string, width int, settings Settings) [][]string {
	var lines [][]string

	lineWidth := 0
	for _, element := range text {
		if lineWidth+len(element) >= width {
			lines = append(lines, []string{element})
			lineWidth = 0
		} else if strings.Contains(element, "\n") {
			strippedElement := strings.ReplaceAll(element, "\n", "")
			if len(lines) == 0 {
				lines = append(lines, []string{strippedElement})
			} else {
				lines[len(lines)-1] = append(lines[len(lines)-1], strippedElement)
				lineWidth += len(strippedElement)
			}
			lines = append(lines, []string{})
		} else {
			if len(lines) == 0 {
				lines = append(lines, []string{element})
			} else {
				lines[len(lines)-1] = append(lines[len(lines)-1], element)
				lineWidth += len(element)
			}
		}
	}

	return lines
} 
