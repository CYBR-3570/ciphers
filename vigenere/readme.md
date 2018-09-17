# Vigenere Cipher

## Usage  
```
go run vigenere.go \
  --file text1 \
  --key PRUDENT \
  --encipher
```  
or  
```
go run viginere.go --file text1 --key PRUDENT --encipher
```  
the above commands will run the vigenere cipher on the file "text1" with the key "PRUDENT" and encipher it.  
you can also decipher with the --decipher flag instead.  


```
>> go run vigenere.go
ONE FATHER IS MORE THAN A HUNDRED SCHOOLMASTERS

enciphered message:
DEYIEGATICVQBKTKBDRNAJEXUIQLRYIRPZTHKYUW
```
### code
used some of odysseus/vigenere's functions
