# Vetores ordenados

- Os dados estão organizados na ordem ascendente (crescente) de valores-chave, ou seja, com o menor valor no índice 0 e cada célula mantendo um valor maior que a célula abaixo

- Vantagem: agiliza os tempos de pesquisa

- Inserção
    - Pesquisa linear - O(n)
    - Mover os elementos restantes - O(n)
    - Big-O - O(2n) = O(n)

- Pesquisa linear
    - A pesquisa termina quando o primeiro item maior que o valor de pesquisa é atingido
    - Como o vetor está ordenado, o algoritmo sabe que não há necessidade de procurar mais
    - Pior caso: se o elemento não estiver no vetor ou na última posição
    - Big-O - O(n)

- Exclusão
    - O algoritmo pode terminar na metade do caminho se não encontrar o item
    - Pesquisa linear
        - Pior caso: N
    - Mover os elementos restantes
        - Pior caso: N
    - Big-O - O(2n) = O(n)
