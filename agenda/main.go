qpackage main

import (
	"fmt"
)

type Contato struct {
	Nome string
	Telefone string
	Email string
}

func (c Contato) VerContato()  {
	
	fmt.Println(c.Nome)
	
}
//  a definião do objeto seria em um arquivo main separado, e oque esta abaixo é colocado emoutro arquivo
func main(){
	thiago := Contato{
		Nome : "thiago",
		Telefone : "9999-9999",
		Email : "tgazaroli@gmail.com",
	}

	fmt.Println(thiago.Telefone)
}




/*
reader := bufio.NewReader(os.Stdin)
text, _ := reader.ReadString('\n')
*/