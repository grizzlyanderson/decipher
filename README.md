# decipher

A Project to (re-)learn GO by creating tools to help solve relatively simple cipher challenges
set by a co-worker.

## Running
* there is no makefile or anything so you'll need to generate the embedded resource first
by running ` go generate` in the project root
* you must supply a ciphertext file ` go run decipher -f $PATH_TO_FIILE`

The ciphertext file currently only supports ASCII for sure, and should not contain any
punctuation. It may contain spaces and line breaks. E.g.
```$xslt
AJKFL BEWUA MEMEU
SNAFU FUBAR UBFUM
NATOX CIAFB IIRSX
```

## Things I found intersting
* *embedding resources* into the binary. In this case ngram data for the english language.
I chose not to use the go-bindata library as it has been forked like crazy since being abandomed.
Instead I'm using vfsgen, but also considered fileb0x. 
**TODO** - I may have shot myself in the foot there by moving the generator to the assets folder.

## Source for actual cryptanalsis stuff
I definitely don't have this stuff down cold.  

* [Practical Cryptography] (http://practicalcryptography.com/)

   This site was sometimes very slow or down. The [Internet Archive](https://web.archive.org/web/sitemap/practicalcryptography.com) has a good copy as well.
   
* Elementary Cryptanalysis by Helen Fouché Gaines, c. 1939

   Nope, not the Red Book. That's way too up-to-date for this old-school stuff.


