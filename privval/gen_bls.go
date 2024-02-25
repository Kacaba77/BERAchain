//go:build ((linux && amd64) || (linux && arm64) || (darwin && amd64) || (darwin && arm64) || (windows && amd64)) && blst

package privval

import (
	bls "github.com/berachain/comet-bls12-381"
	cmtos "github.com/cometbft/cometbft/internal/os"
)

// GenFilePV generates a new validator with randomly generated private key
// and sets the filePaths, but does not call Save().
func GenFilePV(keyFilePath, stateFilePath string) *FilePV {
	key, err := bls.GenPrivKey()
	if err != nil {
		cmtos.Exit(err.Error())
	}
	return NewFilePV(key, keyFilePath, stateFilePath)
}
