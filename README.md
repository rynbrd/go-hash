Hash
====
A simple library to generate hashes for data comparison.

Usage
-----
All functions return a string containing the base64 encoded sha1 hash of the
input.

To hash some bytes:

    val := hash.Bytes(data)

To hash a file:

    val := hash.File("/tmp/somefile")

License
-------
Copyright (c) 2014 Ryan Bourgeois. Licensed under BSD-Modified. See the LICENSE
file for a copy of the license.
