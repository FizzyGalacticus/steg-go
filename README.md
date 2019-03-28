# steg-go

This project is an (yet another) attempt at recreating an old Digital Forensics project of hiding files in image, and possibly more. [Original Python script here](https://github.com/FizzyGalacticus/Steganography/blob/master/Stegonography.py).

## What is steganography?

You can always [read the wiki](https://en.wikipedia.org/wiki/Steganography). But long story short, it's basically the idea of hiding things in plain sight (like files inside of images).

## How does hiding files in images work?

Basically, you spread out your input files into the individual bits, and you insert those bits into the pixels of an image. In this particular case, we are using the method of "least significant bit" (LSB) steganography, which means that we put the bits into the bit of the pixel value that will have the least significant impact on the image itself.

## Building

I have created a [Makefile](https://en.wikipedia.org/wiki/Makefile) to help with the building process. To build, you only need to run `make build` and it will output the executable binary to `bin/steg`
