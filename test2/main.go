package test2

import (
	"context"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/ratify-project/ratify-go"
)

type Verifier struct{}

func init() {
	ratify.RegisterVerifier("test2", newVerifier)
}

func newVerifier(_ ratify.NewVerifierOptions) (ratify.Verifier, error) {
	return &Verifier{}, nil
}

func (v *Verifier) Name() string {
	return "test2"
}

func (v *Verifier) Type() string {
	return "test2"
}

func (v *Verifier) Verifiable(artifact ocispec.Descriptor) bool {
	return true
}

func (v *Verifier) Verify(ctx context.Context, opts *ratify.VerifyOptions) (*ratify.VerificationResult, error) {
	return &ratify.VerificationResult{}, nil
}
