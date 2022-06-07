package controllers

import (
	"reflect"
	"testing"
)

func TestObterListaRomanos(t *testing.T) {
	texto := "AXXBLX"
	esperado := []string{"XX", "LX"}
	sliceRomanos := []string{}
	var cont uint64
	var numeroRomano, algarismoAnterior string

	for i := 0; i < len(texto); i++ {
		if _, valor := mapeamento[texto[i]]; valor {
			numeroRomano += string(texto[i])
			if string(texto[i]) != algarismoAnterior {
				algarismoAnterior = string(texto[i])
				cont = 0
			} else {
				cont++
			}
			if cont == 3 || i == len(texto)-1 {
				sliceRomanos = append(sliceRomanos, numeroRomano)
				cont = 0
				numeroRomano = ""
			}
		} else {
			if numeroRomano != "" {
				sliceRomanos = append(sliceRomanos, numeroRomano)
			}
			cont = 0
			numeroRomano = ""
		}
	}
	if !reflect.DeepEqual(sliceRomanos, esperado) {
		t.Fatalf("Valor esperado: %#v. Valor retornardo %#v: ", esperado, sliceRomanos)
	}

}

func TestConverterNumero(t *testing.T) {

	esperado := 60
	romano := "LX"
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
	if resultado != esperado {
		t.Fatalf("Valor esperado: %d. Valor retornado: %d", esperado, resultado)
	}
}

func TestObterMaiorNum(t *testing.T) {
	lista := []string{"XX", "LX"}
	esperadoRom := "LX"
	esperadoInt := 60

	maiorNumRom := ""
	maiorNumInt := 0
	for _, num := range lista {
		numConvertido := ConverterNumero(num)
		if numConvertido > maiorNumInt {
			maiorNumInt = numConvertido
			maiorNumRom = num
		}
	}

	if maiorNumRom != esperadoRom {
		t.Fatalf("Valor esperado: %s. Valor retornado: %s", esperadoRom, maiorNumRom)
	}
	if maiorNumInt != esperadoInt {
		t.Fatalf("Valor esperado: %d. Valor retornado: %d", esperadoInt, maiorNumInt)
	}
}
