package console

import (
	"github.com/fatih/color"
)

type Printer interface {
	Println(a ...any)
	Print(a ...any)
	Printf(format string, a ...any)
}

func Errorf(p Printer, format string, args ...any) {
	p.Print(color.New(color.FgRed).Add(color.Bold).Print("[-] "))
	p.Printf(format, args...)
}

func Errorln(p Printer, s string) {
	p.Print(color.New(color.FgRed).Add(color.Bold).Sprint("[-] "))
	p.Println(s)
}

func Warningf(p Printer, format string, args ...any) {
	p.Print(color.New(color.FgYellow).Add(color.Bold).Sprintf("[!] "))
	p.Printf(format, args...)
}

func Warningln(p Printer, s string) {
	p.Print(color.New(color.FgYellow).Add(color.Bold).Sprintf("[!] "))
	p.Println(s)
}

func Printf(p Printer, format string, args ...any) {
	p.Print(color.New(color.Reset).Printf(format, args...))
}

func InfoComponent(p Printer, component, message string) {
	p.Print(color.New(color.FgGreen).Add(color.Bold).Sprint("[+] "))
	p.Printf("%s\t: ", component)
	p.Println(message)
}

func InfoComponentf(p Printer, component, format string, args ...any) {
	p.Print(color.New(color.FgGreen).Add(color.Bold).Sprint("[+] "))
	p.Printf("%s\t: ", component)
	p.Printf(format, args...)
}

func ErrorComponent(p Printer, component, message string) {
	p.Print(color.New(color.FgRed).Add(color.Bold).Sprint("[-] "))
	p.Printf("%s\t: ", component)
	p.Println(message)
}

func ErrorComponentf(p Printer, component, format string, args ...any) {
	p.Print(color.New(color.FgRed).Add(color.Bold).Sprint("[-] "))
	p.Printf("%s\t: ", component)
	p.Printf(format, args...)
}
