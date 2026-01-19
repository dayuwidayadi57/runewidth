A lightweight, zero-dependency Go library to calculate the visual width of strings containing UTF-8 characters and emojis.

## Features

- **Zero Dependency**: Only uses the Go standard library.
- **Emoji Support**: Correctly handles double-width characters (Emojis, CJK).
- **Truncate Support**: Safely truncate strings without breaking UTF-8 runes.
- **Fast & Clean**: Minimalist implementation optimized for performance.

## Installation

```bash
go get [github.com/dayuwidayadi57/runewidth](https://github.com/dayuwidayadi57/runewidth)

Usage
Get Visual Width of a Character
import "[github.com/dayuwidayadi57/runewidth](https://github.com/dayuwidayadi57/runewidth)"

width := runewidth.GetVisualWidth('ð') // returns 2
width := runewidth.GetVisualWidth('A') // returns 1

Get Total Visual Width of a String
s := "Hello ð"
width := runewidth.StringVisualWidth(s) // returns 8

Truncate String Safely
s := "This is a very long text with emojis ð"
truncated := runewidth.Truncate(s, 10, "...") 
// returns "This is..."

License
MIT License

---
