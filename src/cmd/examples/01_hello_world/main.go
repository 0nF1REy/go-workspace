/*
-------------------------------
Tabela de termos importantes
-------------------------------
Palavra       : Significado
-------------------------------
array         : Estrutura de tamanho fixo, todos os elementos do mesmo tipo.
slice         : "Fatia" dinâmica de um array; tamanho pode crescer/diminuir.
append        : Função usada com slices para adicionar elementos dinamicamente.
for           : Loop usado para repetir blocos de código.
range         : Forma de iterar sobre arrays, slices, strings, maps ou canais.
continue      : Pula a iteração atual do loop e vai para a próxima.
break         : Interrompe completamente o loop.
switch        : Estrutura de seleção para executar blocos diferentes dependendo do valor.
len           : Função que retorna o tamanho de arrays, slices, strings ou mapas.
rune          : Tipo de dado que representa um caractere Unicode.
byte          : Tipo de dado que representa um byte (8 bits), usado em strings ASCII.
map           : Estrutura chave-valor dinâmica.
const         : Constante, valor fixo que não pode ser alterado durante a execução.
var           : Declaração de variável.
:=            : Atalho para declarar e inicializar variáveis locais.
fmt.Println   : Função para imprimir valores no console.
unicode       : Biblioteca para manipulação de caracteres Unicode (opcional).
-------------------------------
*/

package main

import "fmt"

func helloWorld1() {
	fmt.Println("Hello World!")
}

func hello() {
	fmt.Print("Hello" + " ")
}

func world() {
	fmt.Print("World")
}

func helloWorld2() {
	hello()
	world()
	fmt.Println("!")
}

func content() {
	fmt.Println("Chamando: 'helloWorld1()'...")
	helloWorld1()
	fmt.Println("Chamando: 'helloWorld2()'...")
	helloWorld2()
}

func main() {
	content()
}
