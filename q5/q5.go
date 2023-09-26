package q5

import (
	"strings"
)

//Pedro começou a frequentar aulas de programação. Na primeira aula, sua tarefa foi escrever um programa simples. O
//programa deveria fazer o seguinte: na sequência de caracteres fornecida, composta por letras latinas maiúsculas e
//minúsculas, ele:
//
//* deleta todas as vogais;
//* insere um caractere "." antes de cada consoante;
//* substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas.
//
//As vogais são as letras "A", "O", "E", "U", "I", e o restante são consoantes. A entrada do programa é exatamente uma
//sequência de caracteres e a saída deve ser uma única sequência de caracteres, resultante após o processamento do
//programa na sequência de caracteres inicial.
//
//Ajude Pedro a lidar com esta tarefa fácil.

func ProcessString(s string) string {

	//deixar tudo minusculo
	s = strings.ToLower(s)

	//remover as vogais
	s = strings.ReplaceAll(s, "a", "")
	s = strings.ReplaceAll(s, "e", "")
	s = strings.ReplaceAll(s, "i", "")
	s = strings.ReplaceAll(s, "o", "")
	s = strings.ReplaceAll(s, "u", "")

	//como só há consoantes, adicionar um ponto no inicio
	s = "." + s

	//loop: i recebe 0; enquanto i for menor do que o tamanho da string menos um (pq se só fosse menor do que o
	//tamanho da string, iria ter um ponto a mais no final); adicionar um ao i
	for i := 0; i < len(s)-1; i++ {

		//se a string de index i for diferente de um ponto (já que só há vogais)...
		if string(s[i]) != "." {

			//...a string recebe até a posição i da string (é usado o s[i+1] pq ela é excludente - vai até a
			//posição s[i]); mais um ponto; mais o final da string (é usado o s[i+1] pq é includente - vai a
			//partir da posição s[i+1])
			s = s[:i+1] + "." + s[i+1:]

		}

	}

	//se a string só tiver um caractere (só havia duas letras A's isoladas)...
	if len(s) == 1 {

		//...retorne nada/string vazia porque as vogais devem, apenas, sumir
		return ""

	}

	return s

}
