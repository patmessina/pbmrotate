# PBMRotate

An application that accepts a filename of a PBM image and the number of degrees
to rotate, and will print to stdout or write the image to disk.

> Currently only supports P1 formatted PBM images.

# Installation

TODO

## Build From Source

TODO

# Example Usage

```bash
pbmrotate -i example.pbm -d 180 -o example_flip.pbm
```

# PBM Image

[PBM](https://en.wikipedia.org/wiki/Netpbm) is an image format that can be
represented in ASCII and is easily manipulated.

An example PBM file could look like this

```
P1
# This is an example bitmap of the letter "J"
6 10
0 0 0 0 1 0
0 0 0 0 1 0
0 0 0 0 1 0
0 0 0 0 1 0
0 0 0 0 1 0
0 0 0 0 1 0
1 0 0 0 1 0
0 1 1 1 0 0
0 0 0 0 0 0
0 0 0 0 0 0
```
