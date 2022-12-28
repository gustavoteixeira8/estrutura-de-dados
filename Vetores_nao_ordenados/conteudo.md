# Vetores não ordenados

- Controlar quais jogadores estão presentes no campo de treino (estrutura de dados)

- Vetor de referência:
--- --- --- --- --- --- ---
 4 | 2 | 1 | 8 | 5 |   |   |
--- --- --- --- --- --- ---

- Várias ações poderiam ser executadas:
    - Inserir um jogador na estrutura de dados quando ele chegar ao campo
    - Verificar se um determinado jogador está presente, pesquisando o número do jogador na estrutura
    - Remover um jogador da estrutura de dados quando ele for para casa

- Inserção
    - Único passo (inserido no primeira célula vaga do vetor)
    - O algoritmo já conhece essa localização porque ele já sabe quantos itens já estão no vetor
    - O novo item é simplesmente inserido no próximo espaço disponível
    - Big-O constante - O(1)

- Pesquisa linear
    - Percorrer cada posição do vetor 
    - Melhor caso: 4
    - Pior caso: 5 ou número que não existe
    - Em média, metade dos itens devem ser examinados (n/2)
    - Big-O linear - O(n)

- Exclusão
    - Pesquisar uma média de n/2 elementos (pesquisa linear)
        - Pior caso: N
    - Mover os elementos restantes (n/2 passos)
        - Pior caso: N
    - Big-O - O(2n) = O(n)
    
- Duplicatas
    - Deve-se decidir se itens com chaves duplicadas serão permitidos
    - Exemplo de um arquivo de funcionários
        - Se a chave for o número de registo
        - Se a chave for o sobrenome
    - Pesquisa: Mesmo se encontrar o valor, o algoritmo terá que continuar procurando até a última célula (N passos)
    - Inserção: Verificar cada item antes de fazer uma inserção (N passos)
    - Exclusão do primeiro item: N/2 comparações e N/2 movimentos
    - Exclusão de mais itens: verificar N célular e mais de N/2 células