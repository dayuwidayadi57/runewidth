package runewidth

func GetVisualWidth(r rune) int {
    // Control characters
    if r < 32 || (r >= 0x7F && r < 0xA0) {
        return 0
    }
    
    // Full-width ranges (East Asian, Emoji, etc.)
    switch {
    // Basic ASCII (single width)
    case r < 0x80:
        return 1
    
    // East Asian Full-width
    case (r >= 0x1100 && r <= 0x115F) || // Hangul Jamo
         (r >= 0x231A && r <= 0x231B) || // Watch, hourglass
         (r >= 0x2329 && r <= 0x232A) || // Angle brackets
         (r >= 0x23E9 && r <= 0x23EC) || // Media controls
         (r >= 0x23F0) && (r <= 0x23F3) || // Alarms
         (r >= 0x25FD && r <= 0x25FE) || // Geometric shapes
         (r >= 0x2614 && r <= 0x2615) || // Umbrella, coffee
         (r >= 0x2648 && r <= 0x2653) || // Zodiac
         (r >= 0x26F1 && r <= 0x26F3) || // Beach, ferry
         (r >= 0x26F5) && (r <= 0x26FA) || // Transport
         (r >= 0x26FD && r <= 0x26FD) || // Fuel pump
         (r >= 0x2705) && (r <= 0x2705) || // Check mark
         (r >= 0x270A && r <= 0x270B) || // Fist, victory hand
         (r >= 0x2728 && r <= 0x2728) || // Sparkles
         (r >= 0x274C && r <= 0x274C) || // Cross mark
         (r >= 0x274E && r <= 0x274E) || // Negative check
         (r >= 0x2753 && r <= 0x2755) || // Question/exclamation
         (r >= 0x2757 && r <= 0x2757) || // Exclamation
         (r >= 0x2795 && r <= 0x2797) || // Math symbols
         (r >= 0x27B0 && r <= 0x27B0) || // Curly loop
         (r >= 0x27BF && r <= 0x27BF) || // Double curly loop
         (r >= 0x2B1B && r <= 0x2B1C) || // Black/white squares
         (r >= 0x2B50 && r <= 0x2B50) || // Star
         (r >= 0x2B55 && r <= 0x2B55) || // Circle
         (r >= 0x1F004 && r <= 0x1F004) || // Mahjong
         (r >= 0x1F0CF && r <= 0x1F0CF) || // Joker
         (r >= 0x1F18E && r <= 0x1F18E) || // Blood types
         (r >= 0x1F191 && r <= 0x1F19A) || // Square letters
         (r >= 0x1F201 && r <= 0x1F202) || // Japanese symbols
         (r >= 0x1F21A && r <= 0x1F21A) || // Japanese symbol
         (r >= 0x1F22F && r <= 0x1F22F) || // Japanese symbol
         (r >= 0x1F232 && r <= 0x1F23A) || // Japanese symbols
         (r >= 0x1F250 && r <= 0x1F251) || // Japanese symbols
         (r >= 0x1F300 && r <= 0x1F321) || // Weather, objects
         (r >= 0x1F324 && r <= 0x1F393) || // Weather, objects
         (r >= 0x1F396 && r <= 0x1F397) || // Military
         (r >= 0x1F399 && r <= 0x1F39B) || // Audio/visual
         (r >= 0x1F39E && r <= 0x1F3F0) || // Audio/visual, maps
         (r >= 0x1F3F3 && r <= 0x1F3F5) || // Flags
         (r >= 0x1F3F7 && r <= 0x1F4FD) || // Objects
         (r >= 0x1F4FF && r <= 0x1F53D) || // Objects, symbols
         (r >= 0x1F549 && r <= 0x1F54E) || // Religious
         (r >= 0x1F550 && r <= 0x1F567) || // Clock faces
         (r >= 0x1F56F && r <= 0x1F570) || // Objects
         (r >= 0x1F573 && r <= 0x1F57A) || // Objects
         (r >= 0x1F587 && r <= 0x1F587) || // Link
         (r >= 0x1F58A && r <= 0x1F58D) || // Writing
         (r >= 0x1F590 && r <= 0x1F590) || // Hand
         (r >= 0x1F595 && r <= 0x1F596) || // Hands
         (r >= 0x1F5A4 && r <= 0x1F5A5) || // Computer
         (r >= 0x1F5A8 && r <= 0x1F5B2) || // Computer
         (r >= 0x1F5B9 && r <= 0x1F5BC) || // Computer
         (r >= 0x1F5C2 && r <= 0x1F5D0) || // Office
         (r >= 0x1F5D2 && r <= 0x1F5D4) || // Office
         (r >= 0x1F5DC && r <= 0x1F5DE) || // Tools
         (r >= 0x1F5E1 && r <= 0x1F5E1) || // Dagger
         (r >= 0x1F5E3 && r <= 0x1F5E3) || // Speech bubble
         (r >= 0x1F5E8 && r <= 0x1F5E8) || // Speech bubbles
         (r >= 0x1F5EF && r <= 0x1F5EF) || // Angry bubble
         (r >= 0x1F5F3 && r <= 0x1F5F3) || // Ballot box
         (r >= 0x1F5FA && r <= 0x1F64F) || // Map, faces
         (r >= 0x1F680 && r <= 0x1F6C5) || // Transport
         (r >= 0x1F6CB && r <= 0x1F6D2) || // Furniture
         (r >= 0x1F6D5 && r <= 0x1F6D7) || // Religious
         (r >= 0x1F6DC && r <= 0x1F6DF) || // Wireless
         (r >= 0x1F6E0 && r <= 0x1F6E5) || // Tools
         (r >= 0x1F6E9 && r <= 0x1F6E9) || // Satellite
         (r >= 0x1F6EB && r <= 0x1F6EC) || // Airplane
         (r >= 0x1F6F0 && r <= 0x1F6F0) || // Satellite
         (r >= 0x1F6F3 && r <= 0x1F6FC) || // Transport
         (r >= 0x1F7E0 && r <= 0x1F7EB) || // Colored circles
         (r >= 0x1F7F0 && r <= 0x1F7F0) || // Heavy equals
         (r >= 0x1F90C && r <= 0x1F93A) || // People, sports
         (r >= 0x1F93C && r <= 0x1F945) || // Sports
         (r >= 0x1F947 && r <= 0x1F9FF) || // Medals, objects
         (r >= 0x1FA70 && r <= 0x1FA7C) || // Body parts
         (r >= 0x1FA80 && r <= 0x1FA88) || // Objects
         (r >= 0x1FA90 && r <= 0x1FABD) || // Nature
         (r >= 0x1FABF && r <= 0x1FAC5) || // Body parts
         (r >= 0x1FACE && r <= 0x1FADB) || // Food
         (r >= 0x1FAE0 && r <= 0x1FAE8) || // Faces
         (r >= 0x1FAF0 && r <= 0x1FAF8) || // Hands
         (r >= 0x13000 && r <= 0x13455) || // Egyptian Hieroglyphs
         (r >= 0x1F0A0 && r <= 0x1F0FF) || // Playing Cards
         (r >= 0x1D400 && r <= 0x1D7FF):   // Mathematical Alphanumeric
        return 2
    
    // Zero-width characters
    case r == 0x200D || // Zero-width joiner
         r == 0xFE0E || // Text variation selector
         r == 0xFE0F || // Emoji variation selector
         (r >= 0x0300 && r <= 0x036F) || // Combining diacritics
         (r >= 0x1AB0 && r <= 0x1AFF) || // Extended combining
         (r >= 0x1DC0 && r <= 0x1DFF) || // More combining
         (r >= 0x20D0 && r <= 0x20FF) || // Combining for symbols
         (r >= 0xFE20 && r <= 0xFE2F):   // Combining half marks
        return 0
    
    // Default: single width for other Unicode
    default:
        return 1
    }
}

func StringVisualWidth(s string) int {
    w := 0
    for _, r := range s {
        w += GetVisualWidth(r)
    }
    return w
}

func Truncate(s string, width int, tail string) string {
    if StringVisualWidth(s) <= width {
        return s
    }
    w := 0
    tw := StringVisualWidth(tail)
    for i, r := range s {
        rw := GetVisualWidth(r)
        if w+rw+tw > width {
            return s[:i] + tail
        }
        w += rw
    }
    return s
}
