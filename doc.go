// farmhash project doc.go

/*
farmhash document

Converted from the original C++ source code by building it (on a Ubuntu
based system with an Intel CPU). Then copying the build command and using it to
generate the output of the cpp stage. This showed a version of what code was
required. Note that I copied code from the original files to convert to Go
in order to preserve original comments.

NOTE: When building its important to build with -DFARMHASH_DEBUG=0
(or edit src/farmhash.cc and add a #define) otherwise the results are byteswapped!

Pointer arithmatic in the C++ was mostly simple offsets so instead slices were
sliced to the same offset.

To obey Go export rules some functions had their first character case changed

*/
package farmhash
