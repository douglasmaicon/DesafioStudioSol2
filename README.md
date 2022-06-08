# Desafio Backend StudioSOL

# Maior número romano em uma palavra

Joaquim é um jovem estudante que pretende ajudar seus colegas de classe a resolver um problema de
busca. Ele pretende construir um API web que resolva o problema passado pela professora e que
retorne o resultado para os seus colegas utilizarem. A sua tarefa será ajudar Joaquim nessa missão.

# O Problema
Dada uma palavra contínua, Joaquim precisa localizar o maior número romano dentro da palavra. Os
algarismos romanos são representados por letras maiúsculas, num total de 7 numerações: I (1), V (5), X
(10), L (50), C (100), D (500), M (1000).

Por exemplo, na palavra apresentada abaixo, nós encontramos dois pares de números romanos XX (20)
e LX (60). Logo, o maior número romano presente na sequência é LX (60)

![image](https://user-images.githubusercontent.com/14845446/172515844-74efe16d-1a22-401a-af5a-1008eab6233c.png)

# Entrada / Saída

Para facilitar a consulta dos amigos de Joaquim, você deve receber a string por meio de uma API Web.
Ou seja, cada consulta é feita através de uma requisição HTTP que deve retornar o resultado em
formato JSON. A sua API pode ser feita seguindo os padrões REST ou GraphQL. Você pode escolher o
formato que se sentir mais confortável. A entrada é sempre válida e possui apenas valores
alfanuméricos de A a Z.
A professora do Joaquim costuma dar pontos extras para APIs em formato GraphQL :)

# API REST
A API deve possuir uma única rota /search que recebe uma requisição REST em formato JSON
contendo a lista de números no corpo da requisição. Por exemplo, para a entrada apresentada
anteriormente, a requisição seria:

![image](https://user-images.githubusercontent.com/14845446/172516029-043a2127-9ce3-46ef-abed-9b8d5be46713.png)

A resposta deve ser feita também em formato JSON, retornando o maior número romano encontrado e
também o seu respectivo valor em formato decimal. No exemplo anteriormente apresentado, o resultado
seria:

![image](https://user-images.githubusercontent.com/14845446/172516198-63607621-81fd-49eb-876f-39483300de71.png)

# Explicando a Lógica da solução proposta

Tendo realizado análise do enunciado, inicie o planejamento da solução da seguinte forma.

-> Utilizar GIN para criar a API de forma simples usando as configurações default

-> Utilizar o Swaggo com gin-swagger para documentar e testar a API

Quanto à implementação da solução, separei as reponsabilidades considerando as seguintes etapas.

1- Mapear cada algarismo romano com seu respectivo valor decimal;

2- Criar um método para identificar e devolver um slice com todos os possíveis algarismos romanos;

3- Criar um método para converter o algarismo romano em decimal;

4- Criar um método para comparar todos os números convertidos identificando o maior deles, devolvendo tanto a representação romana como a decimal;

5- Criar um método para devolver o resultado de todo esse processamento para o endpoint da API.


# Implementação de Testes

Para validar os métodos criados nas etapas de 2 a 4 foram implementados testes unitários usando a biblioteca de testes nativa da linguagem Go. É possível executar os teste com o comando: 
> $go test -v ./controllers

![image](https://user-images.githubusercontent.com/14845446/172519400-8c9b72f9-11e1-4efe-83eb-e4018df5c2c7.png)


# Testando a API

Considerando o que foi solicitado a API pode ser executada rodando o comando 
> $go run main.go 

Em seguida utilizar uma ferramenta de suporte a APIs de sua preferência como Postman ou Insomnia, bastando enviar uma requisição usando o método http POST para http://localhost:8080/search preenchendo o parâmetro de entrada no corpo da requisição:

![image](https://user-images.githubusercontent.com/14845446/172517122-46475b2f-3fbe-44fd-a8ae-c27eb7cc7a2a.png)

Tomei a liberdade de implemtnar uma outra forma de testar a API é utlizando o Swagger neste caso acessando a a url http://localhost:8080/swagger/index.html#/romanos/post_search

![image](https://user-images.githubusercontent.com/14845446/172517484-af727918-32b4-4b85-8971-619b5b785252.png)




## Informação Importante!

Gostaria de ressaltar que apesar de possuir uma experiência considerável com desenvolvimento de software (aproximadamente 10 anos), iniciei meus estudos com Go recentemente, e ainda estou me aventurando e me apaixonando cada vez mais pela linguagem. Dito isso, tenho plena conciência de que é porvavel que existam soluções mais simples e otimizadas para esse exercício proposto e espero poder refatorar esta API futuramente. Agradeço pela oportunidade e me coloco à disposição para quaisquer assuntos. Deixo abaixo meus contatos

## Douglas Maicon do Nascimento:

<div>
  <a target="_blank" href="https://www.linkedin.com/in/douglas-maicon-2b464157/"><img src="https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white" /></a>  
  <a href="mailto:douglas.dmn@gmail.com"><img src="https://img.shields.io/badge/Gmail-D14836?style=for-the-badge&logo=gmail&logoColor=white"/></a>
  <a target="_blank" href="https://wa.me/5532991103317"><img src="https://img.shields.io/badge/WhatsApp-25D366?style=for-the-badge&logo=whatsapp&logoColor=white" /></a>  
  <a target="_blank" href="https://t.me/douglasmaicon"><img src="https://img.shields.io/badge/Telegram-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white" /></a>   
  <a target="_blank" href="https://www.instagram.com/invites/contact/?i=1wgvs5ud4skwu&utm_content=11wf84k"><img src="https://img.shields.io/badge/Instagram-E4405F?style=for-the-badge&logo=instagram&logoColor=white" /></a>
</div>
