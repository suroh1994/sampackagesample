# What was the problem?

The binaries built by ```sam build``` used the default settings for ```go build```, which results in the binaries including absolute paths in them. Therefore binaries built from different directories are not identical, even though the source code is.

# How was it resolved?

Passing ```-trimpath``` to ```go build``` removes these absolute path references from the generated binaries. This way all binaries have the same hash values and ```sam package``` no longer uploads the binary more than once.

To implement this, the template file has been adjusted and now uses ```Makefile``` to build the binaries. See the [documentation](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-cli-command-reference-sam-build.html) for more information on how to do this and what else might be possible.