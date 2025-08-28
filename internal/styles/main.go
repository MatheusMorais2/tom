package styles

type ColorPallete struct {
    Primary string
    Secondary string
    Text string
}

func Pallete() *ColorPallete {
    pallete := ColorPallete{
    Primary: "#20283b",
    Secondary: "#CF7539",
    Text: "#D9D9D9",}

    return &pallete
}
