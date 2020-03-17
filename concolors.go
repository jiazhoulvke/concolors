package concolors

func Black(str string) string {
	return "\033[30m" + str + "\033[0m"
}

func Red(str string) string {
	return "\033[31m" + str + "\033[0m"
}

func Green(str string) string {
	return "\033[32m" + str + "\033[0m"
}

func Yellow(str string) string {
	return "\033[33m" + str + "\033[0m"
}

func Blue(str string) string {
	return "\033[34m" + str + "\033[0m"
}

func Purple(str string) string {
	return "\033[35m" + str + "\033[0m"
}

func SkyBlue(str string) string {
	return "\033[36m" + str + "\033[0m"
}

func White(str string) string {
	return "\033[37m" + str + "\033[0m"
}

func WhiteOnBlack(str string) string {
	return "\033[40;37m" + str + "\033[0m"
}

func WhiteOnRed(str string) string {
	return "\033[41;37m" + str + "\033[0m"
}

func WhiteOnGreen(str string) string {
	return "\033[42;37m" + str + "\033[0m"
}

func WhiteOnYellow(str string) string {
	return "\033[43;37m" + str + "\033[0m"
}

func WhiteOnBlue(str string) string {
	return "\033[44;37m" + str + "\033[0m"
}

func WhiteOnPurple(str string) string {
	return "\033[45;37m" + str + "\033[0m"
}

func WhiteOnSkyBlue(str string) string {
	return "\033[46;37m" + str + "\033[0m"
}

func BlackOnWhite(str string) string {
	return "\033[47;30m" + str + "\033[0m"
}
