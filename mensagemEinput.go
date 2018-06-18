package main

import "fmt"

func mensagemInicial() {
  fmt.Println("\n\n\n")
  fmt.Println("///////////////////////////////////////////////////")
  fmt.Println("Registro de ID, Aluno, Notas dos graus e nota final")
  fmt.Printf("///////////////////////////////////////////////////\n\n")
}

func inputOptionCrud(p1 int) (answeroptioncrud int){
  fmt.Println("O que você deseja fazer?")
  fmt.Println("1 - Adicionar Aluno")
  fmt.Println("2 - Deletar Aluno")
  fmt.Println("3 - Atualizar Aluno")
  fmt.Println("4 - Procurar Aluno")
  fmt.Println("\nOpção:")
  fmt.Scanln(&p1)
  fmt.Printf("\n")

  return p1
}

func mensagemEscolha(p1 int){
  for p1 < 1 || p1 > 4{
    fmt.Println("*****************")
    fmt.Println("Erro de seleção")
    fmt.Println("*****************")
    p1 = inputOptionCrud(p1)

  }

  switch p1 {
  case 1:
    fmt.Println("Criar Aluno\n")
  case 2:
    fmt.Println("Deletar Aluno\n")
  case 3:
    fmt.Println("Atualizar Aluno\n")
  case 4:
    fmt.Println("Deletar Aluno\n")
  }


  return
}
