# PBMRotate

An application that accepts a filename of a PBM image and the number of degrees
to rotate, and will print to stdout or write the image to disk.

> Currently only supports P1 formatted PBM images.

# Build From Source

Requires go 1.17. With make run

```bash
make
```

Or to just build with go

```bash
go build -o bin/pbmrotate
```

# Example Usage

The following will rotate the image `example.pbm` 180 degrees and write it to
`example_flip.pbm`.

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
