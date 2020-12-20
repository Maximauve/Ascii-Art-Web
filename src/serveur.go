package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// Params :
type Params struct {
	OUT        string
	Texte      string
	Couleur    string
	Alignement string
	Left       string
	Center     string
	Right      string
	Standard   string
	Shadow     string
	Thinkertoy string
	Htag 	   string 
	Default    string
	White      string
	Red        string
	Yellow     string
	Green      string
	Cyan       string
	Blue       string
	Purple     string
	Black      string
}

// Profil :
type Profil struct {
	Name string
}

func main() {
	fs := http.FileServer(http.Dir("./template/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", serveHome)

	http.HandleFunc("/ascii-art", ascii)
	http.HandleFunc("/Costa", ProfileCosta)
	http.HandleFunc("/Maxime", ProfileMcSim)
	http.HandleFunc("/Mattéo", ProfileMathehoh)
	fmt.Println("Start... ")

	http.ListenAndServe(":8080", nil)
}

/*Vers Profil des programmeurs*/
func ProfileCosta(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/Costa.html")
	if err != nil {
		fmt.Println(err)
	}
	data := &Profil{
		Name: "Costa Reype",
	}
	tmpl.Execute(w, data)
}
func ProfileMcSim(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/Maxime.html")
	if err != nil {
		fmt.Println(err)
	}
	data := &Profil{
		Name: "Maxime Mourgues",
	}
	tmpl.Execute(w, data)
}
func ProfileMathehoh(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/Mattéo.html")
	if err != nil {
		fmt.Println(err)
	}
	data := &Profil{
		Name: "Mattéo Ferreira",
	}
	tmpl.Execute(w, data)
}

/*Page d'accueil*/
func serveHome(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("template/Home.html")
	if err != nil {
		fmt.Println(err)
	}

	texte := "ça marche pas"
	out := ""

	data := &Params{
		OUT:        out,
		Texte:      "Enter your text",
		Couleur:    "---Choisir une couleur---",
		Alignement: "Center",
		Left:       "",
		Center:     "checked",
		Right:      "",
		Standard:   "checked",
		Shadow:     "",
		Thinkertoy: "",
		Htag: 		"",
		Default:    "selected",
		White:      "",
		Red:        "",
		Yellow:     "",
		Green:      "",
		Cyan:       "",
		Blue:       "",
		Purple:     "",
		Black:      "",
	}

	if r.Method == "POST" {
		font := r.FormValue("Police")
		color := r.FormValue("Color")
		align := r.FormValue("Align")
		texte := r.PostFormValue("TEXT")
		fmt.Println(texte)
		fmt.Println(font)
		fmt.Println(color)
		fmt.Println(align)
		fmt.Println(out)

		fmt.Fprintf(w, "%s", out)
	} else {
		texte = "Pas POST"
		fmt.Println(texte)
	}
	fmt.Println(texte)
	tmpl.Execute(w, data)
}

/*Transformation Ascii*/
func ascii(w http.ResponseWriter, r *http.Request) {
	font := r.FormValue("font")
	fmt.Println("font: ", font)
	couleur := r.FormValue("couleur")
	fmt.Println("couleur: ", couleur)
	align := r.FormValue("align")
	fmt.Println("align: ", align)
	texte := r.PostFormValue("inputText")
	fmt.Println(texte)
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")
	var out string

	tab, position := Split(texte, string(rune(10)))
	runes, text := getChar(font)

	fmt.Println(position)
	part := len(position) - 2
	for word := 0; word <= part; word++ {
		print := [9][]string{}

		for _, i := range tab[word] {
			for j := 0; j < len(runes); j++ {
				if i == (runes[j]) {
					//9*j -7
					for ligne := 0; ligne < 9; ligne++ {
						pos := 0
						countn := 0
						//cherche la position du texte avec l'aide de l'ascii table
						for countn != j*9+ligne{
							pos++
							if text[pos] == '\n' {
								countn++
							}
						}
						point := []string{""}
						test := []string{""}
						count := 0
						for test[count] != "\n" {
							count++
							result := string(text[count+pos])
							test = append(test, result)
						}
						for copie := 0; copie < len(test)-2; copie++ {
							point = append(point, test[copie])
						}
						for l := 1; l < len(point); l++ {
							if point[l] != "\n" {

								print[ligne] = append(print[ligne], point[l])
							}
						}
					}
				}
			}
		}

		for ligne := 0; ligne < len(print); ligne++ {
			for i := 0; i < len(print[ligne]); i++ {
				if print[ligne][i] != "\n" && i != len(print[ligne]) {
					out = out + (print[ligne][i])
				}
			}
			out += "\n"
		}
		out += "\n"

	}
	/*checked*/
	LeftChecked := ""
	CenterChecked := ""
	RightChecked := ""
	switch align {
	case "left":
		LeftChecked = "checked"
	case "center":
		CenterChecked = "checked"
	case "right":
		RightChecked = "checked"
	default:
		CenterChecked = "checked"
	}

	StandardChecked := ""
	ShadowChecked := ""
	ThinkertoyChecked := ""
	HtagChecked := ""
	switch font {
	case "standard":
		StandardChecked = "checked"
	case "shadow":
		ShadowChecked = "checked"
	case "thinkertoy":
		ThinkertoyChecked = "checked"
	case "htag":
		HtagChecked = "checked"
	default:
		StandardChecked = "checked"
	}
	WhiteSelect := ""
	RedSelect := ""
	YellowSelect := ""
	GreenSelect := ""
	CyanSelect := ""
	BlueSelect := ""
	PurpleSelect := ""
	BlackSelect := ""
	DefaultSelect := ""

	switch couleur {
	case "white":
		WhiteSelect = "selected"
		DefaultSelect = ""
	case "red":
		RedSelect = "selected"
		DefaultSelect = ""
	case "yellow":
		YellowSelect = "selected"
		DefaultSelect = ""
	case "green":
		GreenSelect = "selected"
		DefaultSelect = ""
	case "cyan":
		CyanSelect = "selected"
		DefaultSelect = ""
	case "blue":
		BlueSelect = "selected"
		DefaultSelect = ""
	case "purple":
		PurpleSelect = "selected"
		DefaultSelect = ""
	case "black":
		BlackSelect = "selected"
		DefaultSelect = ""
	default:
		DefaultSelect = "selected"
	}

	/*Params envoyé à la page HTML*/
	data := &Params{
		OUT:        out,
		Texte:      texte,
		Couleur:    couleur,
		Alignement: align,
		Left:       LeftChecked,
		Center:     CenterChecked,
		Right:      RightChecked,
		Standard:   StandardChecked,
		Shadow:     ShadowChecked,
		Thinkertoy: ThinkertoyChecked,
		Htag: 		HtagChecked,
		White:      WhiteSelect,
		Red:        RedSelect,
		Yellow:     YellowSelect,
		Green:      GreenSelect,
		Cyan:       CyanSelect,
		Blue:       BlueSelect,
		Purple:     PurpleSelect,
		Black:      BlackSelect,
		Default:    DefaultSelect,
	}

	tmpl, err := template.ParseFiles("template/Home.html")
	if err != nil {
		print(err)
		fmt.Println("Erreur dans la template")
	}

	tmpl.Execute(w, data)
}

// getChar : fontion qui va chercher le caractère
func getChar(police string) ([]rune, []rune) {
	filename := "Police/ascii-table-template.txt"
	file, _ := os.Open(filename)
	char := make([]byte, 283)
	file.Read(char)
	var runes []rune
	var ephemere []rune
	for i := 0; i < len(char); i++ {
		ephemere = append(ephemere, rune(char[i]))
	}
	for i := 0; i < len(ephemere); i++ {
		if ephemere[i] != 10 && ephemere[i] != 13 {
			runes = append(runes, rune(char[i]))
		}
	}
	file.Close()

	var text []rune
	switch police {
	case "shadow":
		filename = "Police/shadow.txt"
		file, _ = os.Open(filename)
		/*if err != nil {
			fmt.Println("open " + filename + ": no such file or directory")
		}*/
		arr := make([]byte, 8317)
		file.Read(arr)
		for i := 0; i < len(arr); i++ {
			text = append(text, rune(arr[i]))
		}
		file.Close()
	case "thinkertoy":
		filename = "Police/thinkertoy.txt"
		file, _ = os.Open(filename)
		/*if err != nil {
			fmt.Println("open " + filename + ": no such file or directory")
		}*/
		arr := make([]byte, 5556)
		file.Read(arr)
		for i := 0; i < len(arr); i++ {
			text = append(text, rune(arr[i]))
		}
		file.Close()
	case "standard":
		filename = "Police/standard.txt"
		file, _ = os.Open(filename)
		/*if err != nil {
			fmt.Println("open " + filename + ": no such file or directory")
		}*/
		arr := make([]byte, 7475)
		file.Read(arr)
		for i := 0; i < len(arr); i++ {
			text = append(text, rune(arr[i]))
		}
		file.Close()
	case "htag":
		filename = "Police/htagtext.txt"
		file, _ = os.Open(filename)
		/*if err != nil {
			fmt.Println("open " + filename + ": no such file or directory")
		}*/
		arr := make([]byte, 8316)
		file.Read(arr)
		for i := 0; i < len(arr); i++ {
			text = append(text, rune(arr[i]))
		}
		file.Close()
	default:
		filename = "Police/standard.txt"
		file, _ = os.Open(filename)
		/*if err != nil {
			fmt.Println("open " + filename + ": no such file or directory")
		}*/
		arr := make([]byte, 7474)
		file.Read(arr)
		for i := 0; i < len(arr); i++ {
			text = append(text, rune(arr[i]))
		}
		file.Close()
	}
	return runes, text
}

// Split : fonction qui sépare une string et les stockes dans un tableau de string
func Split(s, sep string) ([]string, []int) {
	runes := []rune(s)
	seprunes := []rune(sep)
	position := []int{}
	position = append(position, 0)
	var tab []string
	lastpos := 0
	i := 0
	for i < len(runes) {
		if runes[i] == seprunes[0] {
			if string(runes[i:i+len(seprunes)]) == sep {
				position = append(position, i+2)
				tab = append(tab, string(runes[lastpos:i]))
				lastpos = i + len(seprunes)
			}
		}
		i++
	}
	if lastpos != i {
		tab = append(tab, string(runes[lastpos:i]))
	}
	position = append(position, len(runes))
	return tab, position
}
