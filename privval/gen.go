//go:build !blst

package privval

import (
	"github.com/cometbft/cometbft/crypto/ed25519"
)

// GenFilePV generates a new validator with randomly generated private key
// and sets the filePaths, but does not call Save().
func GenFilePV(keyFilePath, stateFilePath string) *FilePV {
	return NewFilePV(ed25519.GenPrivKey(), keyFilePath, stateFilePath)
}
