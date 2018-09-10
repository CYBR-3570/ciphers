The keywords PRIME MINISTER are to be used to construct a mixed cipher alphabet by columnar transposition:  
  a. Obtain the cipher alphabet  
    ABCDEFGHIJKLMNOPQRSTUVWXYZ  
    PAKYRBLZICOMDQEFUNGVSHWTJX  
  b. Use it to encipher  
  ITISM UCHEA SIERT OBECR ITICA LTHAN TOBEC ORREC T  
    IVIGDSKZRPGIRNVEARKNIVIKPMVZPQVEARKENNRKV  
  c. Use it to decipher  
  PFNRK RYRQV RDAPM DGPFN IQKIF MR  
    A PRECEDENT EMBALMS A PRINCIPLE  


```
>> go run mixed_columnar_transposition.go --file text2 --encipher
original alphabet: ABCDEFGHIJKLMNOPQRSTUVWXYZ
new alphabet:      PAKYRBLZICOMDQEFUNGVSHWTJX

original message:  ITISMUCHEASIERTOBECRITICALTHANTOBECORRECT
enciphered: IVIGDSKZRPGIRNVEARKNIVIKPMVZPQVEARKENNRKV


>> go run mixed_columnar_transposition.go --file text3 --decipher
original alphabet: ABCDEFGHIJKLMNOPQRSTUVWXYZ
new alphabet:      PAKYRBLZICOMDQEFUNGVSHWTJX

original message:  PFNRKRYRQVRDAPMDGPFNIQKIFMR
deciphered: APRECEDENTEMBALMSAPRINCIPLE
```
