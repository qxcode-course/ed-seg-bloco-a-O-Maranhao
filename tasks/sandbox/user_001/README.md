# apagando

Escreva aqui as informações que você quer salvar, esse é o seu rascunho.
O texto abaixo é informativo e você pode apagar depois de aprender como usar os rascunhos.

## Como usar os rascunhos

- A chave do seu rascunho é o nome da pasta.
- O título do seu rascunho é carregado a partir da primeira linha desse arquivo \#
- Tudo que você fizer nos rascunhos também será rastreado pelo tko.

## Como criar seus próprios testes

Um teste é composto de um `input` (o texto que será fornecido para o programa) e um `output` (o texto que o programa deve retornar para esse input) e pode ter opcionalmente um `label` para facilitar a identificação do teste.

Seus casos de teste personalizados podem ser escritos diretamente aqui na descrição do problema dentro de um fenced code block com a linguagem `toml` ou em um arquivo `tests.toml` na pasta do problema. O TKO irá carregar automaticamente os testes quando a tarefa for aberta ou executada novamente.

Exemplo de teste para ler dois números, um por linha, e imprimir a soma e a subtração deles.

Se quiser habilitar esses casos de teste e ver funcionando, insira algo no input e no output.

```toml
# Exemplo de entrada em uma linha
[[tests]]
input = ''
output = ''

# Exemplo de entrada em múltiplas linhas
# [[tests]]
# input = '''
# 1
# 2
# '''
# output = '''
# 3
# 4
# '''
```

