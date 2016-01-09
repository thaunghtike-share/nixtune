# Signatures

These files contain the various tunings for different services.

All the signatures follow the
`[SystemConfiger](https://godoc.org/github.com/anatma/autotune/signatures#SystemConfiger)`
Golang interface.

## Adding a New Signature

1. Create a new file in signatures.
2. Make sure that it follows the `SystemConfiger` interface.
3. Add the signature map to `NewSignature` in api.go
4. Make sure to describe each kernel change and describe what each change does.
