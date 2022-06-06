package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
Realiza um mapeamento entre o byte do caracter que representa o algarismo romano e seu decimal equivalente
*/
var mapeamento = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

/*
Definição de uma struct para representar o corpo da requisição
*/
type RequestBody struct {
	Text string `json:"text"`
}

// type ResponsetBody struct {
// 	Number string `json:"number"`
// 	Value  int    `json:"value`
// }

func obterListaRomanos(texto string) []string {
	sliceRomanos := []string{}
	var cont uint64
	var numeroRomano, algarismoAnterior string

	for i := 0; i < len(texto); i++ {
		//verificando se o valor do texto no índice atual é um algarismo romano
		if _, valor := mapeamento[texto[i]]; valor {
			numeroRomano += string(texto[i])
			if string(texto[i]) != algarismoAnterior {
				algarismoAnterior = string(texto[i])
				cont = 0
			} else {
				cont++
			}
			//se o algarismo se repetiu 3 vezes: NÃO PODERÁ HAVER A QUARTA VEZ {reiniciar os valores e adicionar um novo numero romano}
			if cont == 3 || i == len(texto)-1 {
				sliceRomanos = append(sliceRomanos, numeroRomano)
				cont = 0
				numeroRomano = ""
			}
		} else {
			//não é um algarismo romano (reiniciar os valores)
			//se numeroRomano != ""	adiciona o númeroRomano no slice
			if numeroRomano != "" {
				sliceRomanos = append(sliceRomanos, numeroRomano)
			}
			cont = 0
			numeroRomano = ""
		}
	}
	return sliceRomanos
}

/*
Executa um laço for reverso na stringRomana para
determinar o valor em decimal com base no mapeamento

digitoCorrespondente = equivalência do algarismo romano em decimais
unidadeAcumulada = valor acumulado da operação de soma para determinar se é unidade, dezena, centena, etc...
resultado = numero convertido
*/
func converterNumero(romano string) int {

	var resultado, unidadeAcumulada int

	for i := len(romano) - 1; i >= 0; i-- {
		digitoCorrespondente := mapeamento[romano[i]]
		if digitoCorrespondente < unidadeAcumulada {
			resultado -= digitoCorrespondente
		} else {
			resultado += digitoCorrespondente
			unidadeAcumulada = digitoCorrespondente
		}
	}
	return resultado
}

/*
Itera sobre a lista de numeros romanos, convertendo e retornando o maior deles
*/
func ObterMaiorNum(lista []string) (string, int) {
	maiorNumRom := ""
	maiorNumInt := 0
	for _, num := range lista {
		numConvertido := converterNumero(num)
		if numConvertido > maiorNumInt {
			maiorNumInt = numConvertido
			maiorNumRom = num

		}
	}
	return maiorNumRom, maiorNumInt
}

func Search(c *gin.Context) {
	requestBody := RequestBody{}
	// usando o método BindJson para serializar body com o struct
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// listaRomanos := obterListaRomanos("AXXXBLXZMMDLXXVIII")
	// fmt.Println(ObterMaiorNum(listaRomanos))
	listaRomanos := obterListaRomanos(string(requestBody.Text))

	// respBody := ResponsetBody{}
	// respBody.Number = requestBody.Text
	// respBody.Value = ObterMaiorNum(listaRomanos)
	// fmt.Println(ObterMaiorNum(listaRomanos))

	// fmt.Println(body)
	// c.JSON(http.StatusAccepted, &body)

	number, value := ObterMaiorNum(listaRomanos)

	c.JSON(http.StatusAccepted, gin.H{
		"number": number,
		"value":  value,
	})
}
