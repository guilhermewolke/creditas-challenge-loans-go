# Creditas - Challenge de Golang - Cálculo de opções de empréstimo

Exemplo de código em Golang, utilizando o design pattern "Builder" para a solução do problema.

O problema proposto é o seguinte:

- Existem 3 tipos de empréstimo:
  - Empréstimo pessoal: Com taxa de 4%
  - Empréstimo com garantia: Com taxa de 3%
  - Empréstimo consignado: Com taxa de 2%
- Baseado no salário do solicitante (e outros critérios), o sistema devolve quais são as opções disponíveis de empréstimo para este cliente. As regras são:
  - Para salários iguais ou inferiores à 3000: 
    - Empréstimo pessoal
    - Empréstimo com garantia (somente se o solicitante tiver menos de 30 anos e residir em SP)
  - Para salários entre 3001 e 5000:
    - Empréstimo pessoal
    - Empréstimo com garantia (somente para clientes que residem em SP)
  - Para salários acima de 5000:
    - Empréstimo pessoal
    - Empréstimo com garantia (somente para clientes com menos de 30 anos)
    - Empréstimo consignao
- O sistema receberia os seguintes dados do cliente:
  - Nome
  - CPF
  - Idade
  - Localização
  - Salário
- Após o recebimento e tratamento das informações, o sistema retorna o resultado no seguinte formato:
  - customer: (nome do cliente)
  - loans: (uma lista de opções de empréstimo no seguinte formato: tipo (o tipo do empréstimo) e taxa (valor numérico com a taxa correspondente))
