// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/session/types
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/session/keeper
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/session/querier
package session

import (
	"github.com/sentinel-official/hub/x/dvpn/session/keeper"
	"github.com/sentinel-official/hub/x/dvpn/session/types"
)

const (
	Codespace                    = types.Codespace
	EventTypeSetSessionsCount    = types.EventTypeSetSessionsCount
	EventTypeSetActiveSession    = types.EventTypeSetActiveSession
	EventTypeUpdateSession       = types.EventTypeUpdateSession
	AttributeKeyCount            = types.AttributeKeyCount
	AttributeKeyID               = types.AttributeKeyID
	AttributeKeySubscription     = types.AttributeKeySubscription
	AttributeKeyNode             = types.AttributeKeyNode
	AttributeKeyAddress          = types.AttributeKeyAddress
	ModuleName                   = types.ModuleName
	QuerierRoute                 = types.QuerierRoute
	QuerySession                 = types.QuerySession
	QuerySessions                = types.QuerySessions
	QuerySessionsForSubscription = types.QuerySessionsForSubscription
	QuerySessionsForNode         = types.QuerySessionsForNode
	QuerySessionsForAddress      = types.QuerySessionsForAddress
)

var (
	// functions aliases
	RegisterCodec                         = types.RegisterCodec
	ErrorMarshal                          = types.ErrorMarshal
	ErrorUnmarshal                        = types.ErrorUnmarshal
	ErrorUnknownMsgType                   = types.ErrorUnknownMsgType
	ErrorUnknownQueryType                 = types.ErrorUnknownQueryType
	ErrorInvalidField                     = types.ErrorInvalidField
	ErrorSubscriptionDoesNotExit          = types.ErrorSubscriptionDoesNotExit
	ErrorInvalidSubscriptionStatus        = types.ErrorInvalidSubscriptionStatus
	ErrorUnauthorized                     = types.ErrorUnauthorized
	ErrorAddressWasNotAdded               = types.ErrorAddressWasNotAdded
	ErrorInvalidBandwidth                 = types.ErrorInvalidBandwidth
	NewGenesisState                       = types.NewGenesisState
	DefaultGenesisState                   = types.DefaultGenesisState
	SessionKey                            = types.SessionKey
	ActiveSessionIDKey                    = types.ActiveSessionIDKey
	NewMsgUpdateSession                   = types.NewMsgUpdateSession
	NewQuerySessionParams                 = types.NewQuerySessionParams
	NewQuerySessionsForSubscriptionParams = types.NewQuerySessionsForSubscriptionParams
	NewQuerySessionsForNodeParams         = types.NewQuerySessionsForNodeParams
	NewQuerySessionsForAddressParams      = types.NewQuerySessionsForAddressParams
	NewKeeper                             = keeper.NewKeeper

	// variable aliases
	ModuleCdc        = types.ModuleCdc
	RouterKey        = types.RouterKey
	StoreKey         = types.StoreKey
	SessionsCountKey = types.SessionsCountKey
	SessionKeyPrefix = types.SessionKeyPrefix
)

type (
	GenesisState                       = types.GenesisState
	MsgUpdateSession                   = types.MsgUpdateSession
	QuerySessionParams                 = types.QuerySessionParams
	QuerySessionsForSubscriptionParams = types.QuerySessionsForSubscriptionParams
	QuerySessionsForNodeParams         = types.QuerySessionsForNodeParams
	QuerySessionsForAddressParams      = types.QuerySessionsForAddressParams
	Session                            = types.Session
	Sessions                           = types.Sessions
	Keeper                             = keeper.Keeper
)
