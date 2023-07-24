package eip6551

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func Test_getCreationCode(t *testing.T) {
	type args struct {
		implementation common.Address
		tokenContract  common.Address
		tokenId        *big.Int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Test getCreationCode",
			args: args{
				implementation: TokenboundERC6551AccountProxy,
				tokenContract:  common.HexToAddress("0xa23a9e6002cebb284e1797be6e0cad201d13c2f2"),
				tokenId:        big.NewInt(1925),
			},
			want: []byte{0x3d, 0x60, 0xad, 0x80, 0x60, 0xa, 0x3d, 0x39, 0x81, 0xf3, 0x36, 0x3d, 0x3d, 0x37, 0x3d, 0x3d, 0x3d, 0x36, 0x3d, 0x73, 0x2d, 0x25, 0x60, 0x25, 0x51, 0x48, 0x7c, 0x3f, 0x33, 0x54, 0xdd, 0x80, 0xd7, 0x6d, 0x54, 0x38, 0x3a, 0x24, 0x33, 0x58, 0x5a, 0xf4, 0x3d, 0x82, 0x80, 0x3e, 0x90, 0x3d, 0x91, 0x60, 0x2b, 0x57, 0xfd, 0x5b, 0xf3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa2, 0x3a, 0x9e, 0x60, 0x2, 0xce, 0xbb, 0x28, 0x4e, 0x17, 0x97, 0xbe, 0x6e, 0xc, 0xad, 0x20, 0x1d, 0x13, 0xc2, 0xf2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0x85},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCreationCode(tt.args.implementation, &tt.args.tokenContract, tt.args.tokenId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCreationCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_computeAddress(t *testing.T) {
	type args struct {
		creationCodeHash []byte
		salt             [32]byte
	}
	tests := []struct {
		name string
		args args
		want common.Address
	}{
		{
			name: "Test computeAddress",
			args: args{
				creationCodeHash: []byte{0x9f, 0xa3, 0x5f, 0x40, 0x24, 0x6e, 0x6c, 0x60, 0x4e, 0xa8, 0x39, 0x83, 0xbe, 0xde, 0x2e, 0x86, 0x8a, 0xf3, 0xed, 0xbe, 0x9a, 0x96, 0xfd, 0x7c, 0xd5, 0x81, 0x33, 0x93, 0x84, 0x9b, 0xb0, 0x6f},
				salt:             [32]byte{0x0},
			},
			want: common.HexToAddress("0x952311534da0393ADBe14006d43d830e32218f39"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeAddress(tt.args.creationCodeHash, tt.args.salt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("computeAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTokenboundTokenAddress(t *testing.T) {
	type args struct {
		tokenContract common.Address
		tokenId       *big.Int
	}
	tests := []struct {
		name string
		args args
		want common.Address
	}{
		{
			name: "Test GetTokenboundTokenAddress",
			args: args{
				tokenContract: common.HexToAddress("0xa23a9e6002cebb284e1797be6e0cad201d13c2f2"),
				tokenId:       big.NewInt(1925),
			},
			want: common.HexToAddress("0x952311534da0393ADBe14006d43d830e32218f39"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTokenboundTokenAddress(&tt.args.tokenContract, tt.args.tokenId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTokenboundTokenAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}