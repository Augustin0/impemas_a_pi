package impemas_a_pi

import (
	"io/ioutil"
	"strconv"
)

var (
	CFC string = "Crédito Fiscal"
	CF  string = "Consumidor final"
	RE  string = "Regímen Especial"
	G   string = "Gubernamental"
	B01 string = "B01"
	B02 string = "B02"
	B14 string = "B14"
	B15 string = "B15"
)

func createCsvCopy(data map[string]map[string]Venta) {

	var format string
	for _, ventas := range data {
		for _, venta := range ventas {
			format += getNumFormated(venta.Numero) + "," + venta.Fecha + "," + venta.Cliente + "," + venta.Rnc + "," + venta.SubTotal + "," + strconv.Itoa(venta.ITBIS) + "," + venta.Total + "," + formatTipo(venta.Tipo) + "\n"
		}
	}

	createCsvCopyFile([]byte(format))
}
func getNumFormated(numero string) string {
	var expected int = 11
	result := len(numero)
	var format string
	if result < expected {
		for i := 0; i < expected-result; i++ {
			format += "0"
		}

	}
	return format + numero
}

func formatTipo(tipoFac string) string {

	switch tipoFac {
	case CFC:
		return B01
	case CF:
		return B02
	case RE:
		return B14
	case G:
		return B15
	}
	return ""
}

func createCsvCopyFile(data []byte) {
	R, err := ioutil.ReadFile("DATA/data.csv")
	if err != nil {
		ioutil.WriteFile("DATA/data.csv", data, 0755)
		return
	}
	if len(R) > len(data) {
		return
	}
	if len(R) < len(data) {
		ioutil.WriteFile("DATA/data.csv", data, 0755)
	}

}
