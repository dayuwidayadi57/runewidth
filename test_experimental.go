package runewidth

import "fmt"

func TestWidth() {
	// Tes karakter yang tadi Abang tambahkan
	emoji := "🚀" // Harusnya lebar 2
	kanji := "本" // Harusnya lebar 2
	ascii := "A"  // Harusnya lebar 1

	fmt.Printf("Visual Width %s: %d\n", emoji, GetVisualWidth([]rune(emoji)[0]))
	fmt.Printf("Visual Width %s: %d\n", kanji, GetVisualWidth([]rune(kanji)[0]))
	fmt.Printf("Visual Width %s: %d\n", ascii, GetVisualWidth([]rune(ascii)[0]))
}

