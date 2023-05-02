package main

import "errors"

// Factorial .
//
// Implemente una función que calcule el factorial de un número
// El factorial de un númenro se define como la multiplicación sucesiva de los antecesores de un número hasta el número dado.
// Ejemplo: f!(5) =  1 * 2 * 3 * 4 * 5 = 120
//
// Esta implementación debe hacerse con un loop
func Factorial(n int64) int64 {
	panic("you should implement this function")
}

// FactorialRecursivo .
//
// El mismo problema del ejercicio anterior.
// Esta implementación debe hacerse aplicando recursión
func FactorialRecursivo(n int64) int64 {
	panic("you should implement this function")
}

// FactorialWithError .
//
// Extienda la función Factorial para que además del resultado devuelva el posible error de la operación.
//
// Si el argumento es menor a 0, devuelva 0 y un error indicando que solo se soportan números positivos.
// Si el argumento es mayor a 10 devuelva 0 y un error indicando que solo se soporta el calculo del factorial hasta el número 10
// Si la operación tiene éxito, devuelva el resultado y nil como el valor de error
//
// Esta implementación debe hacerse con un loop
func FactorialWithError(n int64) (int64, error) {
	// rr := errors.New("este es un error")
	if n < 0 {
		return 0, errors.New("solo se soporta el calculo del factorial para números positivos")
	}
	panic("you should implement this function")
}

// FactorialRecursiveWithError .
//
// Extienda la función Factorial para que además del resultado devuelva el posible error de la operación.
//
// Si el argumento es menor a 0, devuelva 0 y un error indicando que solo se soportan números positivos.
// Si el argumento es mayor a 10 devuelva 0 y un error indicando que solo se soporta el calculo del factorial hasta el número 10
// Si la operación tiene éxito, devuelva el resultado y nil como el valor de error
//
// Esta implementación debe hacerse con un loop
func FactorialRecursiveWithError(n int64) (int64, error) {
	panic("you should implement this function")
}

// SortSliceBubble
//
// En este ejercicio ud debe implementar una función que reciba un slice  de enteros como argumento y devuelva uns lice con sus elementos
// ordenados de menor a mayor
// Busque y use el algoritmo de la burbuja
//
// Ejemplo  unordered = [8,1,0,2,8,9]  salida = [0,1,2,8,8,9]
//
// # Algoritmo de la burbuja
//
// El método de ordenamiento burbuja consiste en comparar cada elemento de la estructura con el siguiente e intercambiándolos
// si corresponde. El proceso se repite hasta que la estructura esté ordenada.
func SortSliceBubble(unordered []int) []int64 {
	panic("you should implement this function")
}
