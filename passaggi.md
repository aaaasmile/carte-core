# Carte-core
Ho creato questa repository per provare a modellare il mazzo delle carte
ed effettuare esperimenti con golang e strutture varie. 

## Esperimenti
Dopo una breve introduzione in Julia, mi interessava avere una struttura Deck
con la possibilità di avere lo stesso metodo (per esempio initialize) ma in package
differenti (per esempio Ramino e Briscola). Invece in Golang non funziona, il metodo 
deve essere nello stesso package. In questo caso, però, il nome non può essere uguale.
Quindi se voglio un mazzo per le carte da invido, devo dichiarare un initialize con un nome
univoco. La specializzazione la ottengo poi con puntatori a funzioni. 


## Unit Test
Uso il file di test per i miei esperimenti anziché usare il main.go. Per
avere l'output del test in Visual Code anziché il semplice OK, bisogna cambiare i settings
nel workspace (vedi file settings.json)