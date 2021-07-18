package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
)

const (
	ModuleName   = "subscription"
	QuerierRoute = ModuleName
)

var (
	ParamsSubspace = ModuleName
	RouterKey      = ModuleName
	StoreKey       = ModuleName
)

var (
	TypeMsgSubscribeToNodeRequest = ModuleName + ":subscribe_to_node"
	TypeMsgSubscribeToPlanRequest = ModuleName + ":subscribe_to_plan"
	TypeMsgCancelRequest          = ModuleName + ":cancel"
	TypeMsgAddQuotaRequest        = ModuleName + ":add_quota"
	TypeMsgUpdateQuotaRequest     = ModuleName + ":update_quota"
)

var (
	EventModuleName = EventModule{Name: ModuleName}
)

var (
	CountKey                                = []byte{0x00}
	SubscriptionKeyPrefix                   = []byte{0x10}
	SubscriptionForNodeKeyPrefix            = []byte{0x11}
	SubscriptionForPlanKeyPrefix            = []byte{0x12}
	ActiveSubscriptionForAddressKeyPrefix   = []byte{0x20}
	InactiveSubscriptionForAddressKeyPrefix = []byte{0x21}
	InactiveSubscriptionAtKeyPrefix         = []byte{0x30}
	QuotaKeyPrefix                          = []byte{0x40}
)

func SubscriptionKey(id uint64) []byte {
	v := append(SubscriptionKeyPrefix, sdk.Uint64ToBigEndian(id)...)
	if len(v) != 1+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+8))
	}

	return v
}

func GetSubscriptionForNodeKeyPrefix(address hubtypes.NodeAddress) []byte {
	v := append(SubscriptionForNodeKeyPrefix, address.Bytes()...)
	if len(v) != 1+sdk.AddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+sdk.AddrLen))
	}

	return v
}

func SubscriptionForNodeKey(address hubtypes.NodeAddress, id uint64) []byte {
	v := append(GetSubscriptionForNodeKeyPrefix(address), sdk.Uint64ToBigEndian(id)...)
	if len(v) != 1+sdk.AddrLen+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+sdk.AddrLen+8))
	}

	return v
}

func GetSubscriptionForPlanKeyPrefix(id uint64) []byte {
	v := append(SubscriptionForPlanKeyPrefix, sdk.Uint64ToBigEndian(id)...)
	if len(v) != 1+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+8))
	}

	return v
}

func SubscriptionForPlanKey(p, s uint64) []byte {
	v := append(GetSubscriptionForPlanKeyPrefix(p), sdk.Uint64ToBigEndian(s)...)
	if len(v) != 1+2*8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+2*8))
	}

	return v
}

func GetActiveSubscriptionForAddressKeyPrefix(address sdk.AccAddress) []byte {
	v := append(ActiveSubscriptionForAddressKeyPrefix, address.Bytes()...)
	if len(v) != 1+sdk.AddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+sdk.AddrLen))
	}

	return v
}

func ActiveSubscriptionForAddressKey(address sdk.AccAddress, i uint64) []byte {
	v := append(GetActiveSubscriptionForAddressKeyPrefix(address), sdk.Uint64ToBigEndian(i)...)
	if len(v) != 1+sdk.AddrLen+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+sdk.AddrLen+8))
	}

	return v
}

func GetInactiveSubscriptionForAddressKeyPrefix(address sdk.AccAddress) []byte {
	v := append(InactiveSubscriptionForAddressKeyPrefix, address.Bytes()...)
	if len(v) != 1+sdk.AddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+sdk.AddrLen))
	}

	return v
}

func InactiveSubscriptionForAddressKey(address sdk.AccAddress, i uint64) []byte {
	v := append(GetInactiveSubscriptionForAddressKeyPrefix(address), sdk.Uint64ToBigEndian(i)...)
	if len(v) != 1+sdk.AddrLen+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+sdk.AddrLen+8))
	}

	return v
}

func GetInactiveSubscriptionAtKeyPrefix(at time.Time) []byte {
	v := append(InactiveSubscriptionAtKeyPrefix, sdk.FormatTimeBytes(at)...)
	if len(v) != 1+29 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+29))
	}

	return v
}

func InactiveSubscriptionAtKey(at time.Time, id uint64) []byte {
	v := append(GetInactiveSubscriptionAtKeyPrefix(at), sdk.Uint64ToBigEndian(id)...)
	if len(v) != 1+29+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+29+8))
	}

	return v
}

func GetQuotaKeyPrefix(id uint64) []byte {
	v := append(QuotaKeyPrefix, sdk.Uint64ToBigEndian(id)...)
	if len(v) != 1+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+8))
	}

	return v
}

func QuotaKey(id uint64, address sdk.AccAddress) []byte {
	v := append(GetQuotaKeyPrefix(id), address.Bytes()...)
	if len(v) != 1+8+sdk.AddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+8+sdk.AddrLen))
	}

	return v
}

func IDFromSubscriptionForNodeKey(key []byte) uint64 {
	if len(key) != 1+sdk.AddrLen+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 1+sdk.AddrLen+8))
	}

	return sdk.BigEndianToUint64(key[1+sdk.AddrLen:])
}

func IDFromSubscriptionForPlanKey(key []byte) uint64 {
	if len(key) != 1+2*8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 1+2*8))
	}

	return sdk.BigEndianToUint64(key[1+8:])
}

func IDFromStatusSubscriptionForAddressKey(key []byte) uint64 {
	if len(key) != 1+sdk.AddrLen+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 1+sdk.AddrLen+8))
	}

	return sdk.BigEndianToUint64(key[1+sdk.AddrLen:])
}

func IDFromInactiveSubscriptionAtKey(key []byte) uint64 {
	if len(key) != 1+29+8 {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 1+29+8))
	}

	return sdk.BigEndianToUint64(key[1+29:])
}
