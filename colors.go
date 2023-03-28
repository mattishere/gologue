package gologue

// The Color type is basically just a string, used for coloring in your terminal.
type Color string

const (
    // Equivalent to "\x1b[34m"
    ColorBlue Color = "\x1b[34m"
    // Equivalent to "\x1b[33m"
    ColorYellow Color = "\x1b[33m"
    // Equivalent to "\x1b[32m"
    ColorGreen Color = "\x1b[32m"
    // Equivalent to "\x1b[35m"
    ColorMagenta Color = "\x1b[35m"
    // Equivalent to "\x1b[31m"
    ColorRed Color = "\x1b[31m"
    // Equivalent to "\x1b[30m"
    ColorBlack Color = "\x1b[30m"
    // Equivalent to "\x1b[37m"
    ColorWhite Color = "\x1b[37m"
    // Equivalent to "\x1b[36m"
    ColorCyan Color = "\x1b[36m"
    // Equivalent to "\x1b[0m"
    ColorReset Color = "\x1b[0m"
)
