package q2

//Um dia, três melhores amigos - Pedro, Vanessa e Tônia - decidiram formar uma equipe e participar de concursos de
//programação. Os participantes geralmente recebem vários problemas durante esses concursos. Muito antes do início, os
//amigos decidiram que implementariam um problema somente se pelo menos dois deles tivessem certeza da solução. Caso
//contrário, os amigos não escreveriam a solução do problema.
//
//Este concurso oferece n problemas aos participantes. Para cada problema, sabemos qual amigo tem certeza da solução. Você
//receberá uma matriz booleana de n linhas e 3 colunas, em que a i-ésima linha representa as opiniões de Pedro, Vanessa e
//Tônia, respectivamente, sobre o i-ésimo problema. O valor "true" indica que o amigo tem certeza da solução, e "false"
//indica o contrário.
//
//Ajude os amigos a encontrar o número de problemas para os quais eles escreverão uma solução.

func ProblemsSolved(answers [][3]bool) int {

	//questoes respondidas
	questoesRespondidas := 0

	//para percorrer as linhas da matriz
	for i := 0; i < len(answers); i++ {

		//quantidade de respostas
		n := 0

		//para percorrer as colunas da matriz
		for j := 0; j < 3; j++ {

			//se o valor na posição answers[i][j] for igual a falso
			if !answers[i][j] {

				continue

			} else {

				//adiciona 1 as respostas para linha
				n++

			}

		}

		if n >= 2 {

			questoesRespondidas++

		}

	}

	return questoesRespondidas

}
