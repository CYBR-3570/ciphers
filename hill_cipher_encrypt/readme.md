# Hill cipher encryption  
```
>> go run hill_cipher_encrypt.go 3 7 8 9
[3 7 8 9]

WASHINGTON
OUZZLZVLDV
```  

```
>> go run hill_cipher_encrypt.go 1 0 2 3 3 2 1 1 1
[1 0 2 3 3 2 1 1 1]

AUTUMNLEAVES
MUNUSTLTPFHR
```  

running `hill_cipher_encyption.go` with the arg 3 7 8 9 will use the matrix:  
```
3 7
8 9
```  
to encyrpt the text in the file `text1`
