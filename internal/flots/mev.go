package flots

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/lmittmann/w3/w3types"
)

type MevBundleTx struct {
	Hash      common.Hash `json:"hash"      mapstructure:"hash"`
	Tx        string      `json:"tx"        mapstructure:"tx"`
	CanRevert bool        `json:"canRevert" mapstructure:"canRevert"`
	// Bundle    MevSendBundleRequest `json:"bundle" mapstructure:"bundle"`
}

type MevParamInclusion struct {
	Block    string `json:"block,omitempty"    mapstructure:"block"`    // hex-encoded number
	MaxBlock string `json:"maxBlock,omitempty" mapstructure:"maxBlock"` // hex-encoded number
}

type MevParamValidity struct {
	Refund []struct {
		BodyIdx int64 `json:"bodyIdx" mapstructure:"bodyIdx"`
		Percent int64 `json:"percent" mapstructure:"percent"`
	} `json:"refund,omitempty" mapstructure:"refund"`
	RefundConfig []struct {
		Address string `json:"address" mapstructure:"address"`
		Percent int64  `json:"percent" mapstructure:"percent"`
	} `json:"refundConfig,omitempty" mapstructure:"refundConfig"`
}

type MevParamPrivacy struct {
	Hints []string `json:"hints,omitempty" mapstructure:"hints"`
}

type MevSendBundleRequest struct {
	Version   string             `json:"version,omitempty"   mapstructure:"version"`
	Inclusion MevParamInclusion  `json:"inclusion,omitempty" mapstructure:"inclusion"`
	Body      []MevBundleTx      `json:"body"                mapstructure:"body"`
	Privacy   MevParamPrivacy    `json:"privacy,omitempty"`
	Validity  []MevParamValidity `json:"validity,omitempty"  mapstructure:"validity"`
}

type mevSendBundleResponse struct {
	BundleHash common.Hash `json:"bundleHash"`
}

// SendBundle sends the bundle to the client's endpoint.
func MevSendBundle(r *MevSendBundleRequest) w3types.CallerFactory[common.Hash] {
	return &mevSendBundleFactory{param: r}
}

type mevSendBundleFactory struct {
	// args
	param *MevSendBundleRequest

	// returns
	result  mevSendBundleResponse
	returns *common.Hash
}

func (f *mevSendBundleFactory) Returns(hash *common.Hash) w3types.Caller {
	f.returns = hash

	return f
}

// CreateRequest implements the [w3types.RequestCreator].
func (f *mevSendBundleFactory) CreateRequest() (rpc.BatchElem, error) {
	return rpc.BatchElem{
		Method: "eth_sendBundle",
		Args:   []any{f.param},
		Result: &f.result,
	}, nil
}

// HandleResponse implements the [w3types.ResponseHandler].
func (f *mevSendBundleFactory) HandleResponse(elem rpc.BatchElem) error {
	if err := elem.Error; err != nil {
		return err
	}

	if f.returns != nil {
		*f.returns = f.result.BundleHash
	}

	return nil
}
