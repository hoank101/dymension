package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dymensionxyz/dymension/v3/x/rollapp/types"
)

func (k *Keeper) VerifyNonAvailableBatch(ctx sdk.Context, msg *types.MsgNonAvailableBatch) error {

	/*stateInfo, found := k.GetStateInfo(ctx, msg.GetRollappId(), msg.GetSlIndex())
	if !found {
		return nil
	}
	DAPath := stateInfo.GetDAPath()
	DAMetaData, err := types.NewDAMetaData(DAPath)
	if err != nil {
		return nil
	}
	//var namespace []byte
	blob, _, err := k.blobsAndCommitments(DAMetaData.GetNameSpace(), msg.GetBlob())
	if err != nil {
		return err
	}

	err = k.verifyBlobNonInclusion(ctx, DAMetaData.GetNameSpace(), msg.GetRproofs(), msg.GetDataroot())
	if err != nil {
		return err
	}*/

	return nil
}

/*func (k *Keeper) verifyBlobNonInclusion(ctx sdk.Context, namespace []byte, rProofs [][]byte, dataRoot []byte) error {
	//TODO (srene): Implement non-inclusion proof validation
	return types.ErrUnableToVerifyProof
}*/