package main

import (
	"net/http"
	"io"
	"fmt"
	"encoding/json"
	"strings"
	"github.com/fatih/color"
	//"strconv"
)

type Sismos struct {
	Fecha string `json:"fecha"`
	Profundidad string `json:"profundidad"`
	Magnitud string `json:"magnitud"`
	RefGeografica string `json:"refGeografica"`
	FechaUpdate string `json:"fechaUpdate"`
}

func main() {
	res, err := http.Get("https://api.gael.cloud/general/public/sismos")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("API Sismos No Disponible!!!")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	content := string(body)
	sismos := sismosFromJson(content)
	fmt.Println(sismos)

	for _, sismo := range sismos {
		color.Red("Sismos!!!")
		// miInt, err := strconv.Atoi(sismo.Profundidad)
		// if err != nil {
		// 	panic(err)
		// }
		fmt.Printf("%T\n",sismo.Profundidad)
		fmt.Println(sismo.Fecha, " ", sismo.Profundidad, " ", sismo.Magnitud, " ", sismo.RefGeografica)
	}

}

func sismosFromJson(content string) []Sismos {
	sismos := make([]Sismos, 0, 20)
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}

	var sismo Sismos 
	for decoder.More() {
		err := decoder.Decode(&sismo)
		if err != nil {
			panic(err)
		}
		sismos = append(sismos, sismo)
	}
	return sismos
}