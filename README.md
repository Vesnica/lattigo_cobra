# STARK-Lattigo

This repo is a part of a test project which trying to combine the STARK and HE(Homomoriphic Encryption) tech.

This project uses lattigo library to encrypt the plain text and output the cipher text to a file. and then, 
the STARK prover can consume the file and perform the compute, the result(which is also a cipher text) and
proof will be produced, which can be verified by STARK verifier. Lastly, the cipher result can be decrypted
by this project and get the plain text result.
