// nolint
package staking

import (
	"github.com/cosmos/cosmos-sdk/x/staking"
)

const (
	DefaultParamspace                  = staking.DefaultParamspace
	DefaultCodespace                   = staking.DefaultCodespace
	CodeInvalidValidator               = staking.CodeInvalidValidator
	CodeInvalidDelegation              = staking.CodeInvalidDelegation
	CodeInvalidInput                   = staking.CodeInvalidInput
	CodeValidatorJailed                = staking.CodeValidatorJailed
	CodeInvalidAddress                 = staking.CodeInvalidAddress
	CodeUnauthorized                   = staking.CodeUnauthorized
	CodeInternal                       = staking.CodeInternal
	CodeUnknownRequest                 = staking.CodeUnknownRequest
	ModuleName                         = staking.ModuleName
	StoreKey                           = staking.StoreKey
	TStoreKey                          = staking.TStoreKey
	QuerierRoute                       = staking.QuerierRoute
	RouterKey                          = staking.RouterKey
	DefaultUnbondingTime               = staking.DefaultUnbondingTime
	DefaultMaxValidators               = staking.DefaultMaxValidators
	DefaultMaxEntries                  = staking.DefaultMaxEntries
	NotBondedPoolName                  = staking.NotBondedPoolName
	BondedPoolName                     = staking.BondedPoolName
	QueryValidators                    = staking.QueryValidators
	QueryValidator                     = staking.QueryValidator
	QueryDelegatorDelegations          = staking.QueryDelegatorDelegations
	QueryDelegatorUnbondingDelegations = staking.QueryDelegatorUnbondingDelegations
	QueryRedelegations                 = staking.QueryRedelegations
	QueryValidatorDelegations          = staking.QueryValidatorDelegations
	QueryValidatorRedelegations        = staking.QueryValidatorRedelegations
	QueryValidatorUnbondingDelegations = staking.QueryValidatorUnbondingDelegations
	QueryDelegation                    = staking.QueryDelegation
	QueryUnbondingDelegation           = staking.QueryUnbondingDelegation
	QueryDelegatorValidators           = staking.QueryDelegatorValidators
	QueryDelegatorValidator            = staking.QueryDelegatorValidator
	QueryPool                          = staking.QueryPool
	QueryParameters                    = staking.QueryParameters
	MaxMonikerLength                   = staking.MaxMonikerLength
	MaxIdentityLength                  = staking.MaxIdentityLength
	MaxWebsiteLength                   = staking.MaxWebsiteLength
	MaxDetailsLength                   = staking.MaxDetailsLength
	DoNotModifyDesc                    = staking.DoNotModifyDesc
)

var (
	// functions aliases
	RegisterInvariants                 = staking.RegisterInvariants
	AllInvariants                      = staking.AllInvariants
	ModuleAccountInvariants            = staking.ModuleAccountInvariants
	NonNegativePowerInvariant          = staking.NonNegativePowerInvariant
	PositiveDelegationInvariant        = staking.PositiveDelegationInvariant
	DelegatorSharesInvariant           = staking.DelegatorSharesInvariant
	NewKeeper                          = staking.NewKeeper
	ParamKeyTable                      = staking.ParamKeyTable
	NewQuerier                         = staking.NewQuerier
	NewCommissionRates                 = staking.NewCommissionRates
	NewCommission                      = staking.NewCommission
	NewCommissionWithTime              = staking.NewCommissionWithTime
	NewDelegation                      = staking.NewDelegation
	MustMarshalDelegation              = staking.MustMarshalDelegation
	MustUnmarshalDelegation            = staking.MustUnmarshalDelegation
	UnmarshalDelegation                = staking.UnmarshalDelegation
	NewUnbondingDelegation             = staking.NewUnbondingDelegation
	NewUnbondingDelegationEntry        = staking.NewUnbondingDelegationEntry
	MustMarshalUBD                     = staking.MustMarshalUBD
	MustUnmarshalUBD                   = staking.MustUnmarshalUBD
	UnmarshalUBD                       = staking.UnmarshalUBD
	NewRedelegation                    = staking.NewRedelegation
	NewRedelegationEntry               = staking.NewRedelegationEntry
	MustMarshalRED                     = staking.MustMarshalRED
	MustUnmarshalRED                   = staking.MustUnmarshalRED
	UnmarshalRED                       = staking.UnmarshalRED
	NewDelegationResp                  = staking.NewDelegationResp
	NewRedelegationResponse            = staking.NewRedelegationResponse
	NewRedelegationEntryResponse       = staking.NewRedelegationEntryResponse
	ErrNilValidatorAddr                = staking.ErrNilValidatorAddr
	ErrBadValidatorAddr                = staking.ErrBadValidatorAddr
	ErrNoValidatorFound                = staking.ErrNoValidatorFound
	ErrValidatorOwnerExists            = staking.ErrValidatorOwnerExists
	ErrValidatorPubKeyExists           = staking.ErrValidatorPubKeyExists
	ErrValidatorPubKeyTypeNotSupported = staking.ErrValidatorPubKeyTypeNotSupported
	ErrValidatorJailed                 = staking.ErrValidatorJailed
	ErrBadRemoveValidator              = staking.ErrBadRemoveValidator
	ErrDescriptionLength               = staking.ErrDescriptionLength
	ErrCommissionNegative              = staking.ErrCommissionNegative
	ErrCommissionHuge                  = staking.ErrCommissionHuge
	ErrCommissionGTMaxRate             = staking.ErrCommissionGTMaxRate
	ErrCommissionUpdateTime            = staking.ErrCommissionUpdateTime
	ErrCommissionChangeRateNegative    = staking.ErrCommissionChangeRateNegative
	ErrCommissionChangeRateGTMaxRate   = staking.ErrCommissionChangeRateGTMaxRate
	ErrCommissionGTMaxChangeRate       = staking.ErrCommissionGTMaxChangeRate
	ErrSelfDelegationBelowMinimum      = staking.ErrSelfDelegationBelowMinimum
	ErrMinSelfDelegationInvalid        = staking.ErrMinSelfDelegationInvalid
	ErrMinSelfDelegationDecreased      = staking.ErrMinSelfDelegationDecreased
	ErrNilDelegatorAddr                = staking.ErrNilDelegatorAddr
	ErrBadDenom                        = staking.ErrBadDenom
	ErrBadDelegationAddr               = staking.ErrBadDelegationAddr
	ErrBadDelegationAmount             = staking.ErrBadDelegationAmount
	ErrNoDelegation                    = staking.ErrNoDelegation
	ErrBadDelegatorAddr                = staking.ErrBadDelegatorAddr
	ErrNoDelegatorForAddress           = staking.ErrNoDelegatorForAddress
	ErrInsufficientShares              = staking.ErrInsufficientShares
	ErrDelegationValidatorEmpty        = staking.ErrDelegationValidatorEmpty
	ErrNotEnoughDelegationShares       = staking.ErrNotEnoughDelegationShares
	ErrBadSharesAmount                 = staking.ErrBadSharesAmount
	ErrBadSharesPercent                = staking.ErrBadSharesPercent
	ErrNotMature                       = staking.ErrNotMature
	ErrNoUnbondingDelegation           = staking.ErrNoUnbondingDelegation
	ErrMaxUnbondingDelegationEntries   = staking.ErrMaxUnbondingDelegationEntries
	ErrBadRedelegationAddr             = staking.ErrBadRedelegationAddr
	ErrNoRedelegation                  = staking.ErrNoRedelegation
	ErrSelfRedelegation                = staking.ErrSelfRedelegation
	ErrVerySmallRedelegation           = staking.ErrVerySmallRedelegation
	ErrBadRedelegationDst              = staking.ErrBadRedelegationDst
	ErrTransitiveRedelegation          = staking.ErrTransitiveRedelegation
	ErrMaxRedelegationEntries          = staking.ErrMaxRedelegationEntries
	ErrDelegatorShareExRateInvalid     = staking.ErrDelegatorShareExRateInvalid
	ErrBothShareMsgsGiven              = staking.ErrBothShareMsgsGiven
	ErrNeitherShareMsgsGiven           = staking.ErrNeitherShareMsgsGiven
	ErrMissingSignature                = staking.ErrMissingSignature
	NewGenesisState                    = staking.NewGenesisState
	DefaultGenesisState                = staking.DefaultGenesisState
	NewMultiStakingHooks               = staking.NewMultiStakingHooks
	GetValidatorKey                    = staking.GetValidatorKey
	GetValidatorByConsAddrKey          = staking.GetValidatorByConsAddrKey
	AddressFromLastValidatorPowerKey   = staking.AddressFromLastValidatorPowerKey
	GetValidatorsByPowerIndexKey       = staking.GetValidatorsByPowerIndexKey
	GetLastValidatorPowerKey           = staking.GetLastValidatorPowerKey
	ParseValidatorPowerRankKey         = staking.ParseValidatorPowerRankKey
	GetValidatorQueueTimeKey           = staking.GetValidatorQueueTimeKey
	GetDelegationKey                   = staking.GetDelegationKey
	GetDelegationsKey                  = staking.GetDelegationsKey
	GetUBDKey                          = staking.GetUBDKey
	GetUBDByValIndexKey                = staking.GetUBDByValIndexKey
	GetUBDKeyFromValIndexKey           = staking.GetUBDKeyFromValIndexKey
	GetUBDsKey                         = staking.GetUBDsKey
	GetUBDsByValIndexKey               = staking.GetUBDsByValIndexKey
	GetUnbondingDelegationTimeKey      = staking.GetUnbondingDelegationTimeKey
	GetREDKey                          = staking.GetREDKey
	GetREDByValSrcIndexKey             = staking.GetREDByValSrcIndexKey
	GetREDByValDstIndexKey             = staking.GetREDByValDstIndexKey
	GetREDKeyFromValSrcIndexKey        = staking.GetREDKeyFromValSrcIndexKey
	GetREDKeyFromValDstIndexKey        = staking.GetREDKeyFromValDstIndexKey
	GetRedelegationTimeKey             = staking.GetRedelegationTimeKey
	GetREDsKey                         = staking.GetREDsKey
	GetREDsFromValSrcIndexKey          = staking.GetREDsFromValSrcIndexKey
	GetREDsToValDstIndexKey            = staking.GetREDsToValDstIndexKey
	GetREDsByDelToValDstIndexKey       = staking.GetREDsByDelToValDstIndexKey
	NewMsgCreateValidator              = staking.NewMsgCreateValidator
	NewMsgEditValidator                = staking.NewMsgEditValidator
	NewMsgDelegate                     = staking.NewMsgDelegate
	NewMsgBeginRedelegate              = staking.NewMsgBeginRedelegate
	NewMsgUndelegate                   = staking.NewMsgUndelegate
	NewParams                          = staking.NewParams
	DefaultParams                      = staking.DefaultParams
	MustUnmarshalParams                = staking.MustUnmarshalParams
	UnmarshalParams                    = staking.UnmarshalParams
	NewPool                            = staking.NewPool
	NewQueryDelegatorParams            = staking.NewQueryDelegatorParams
	NewQueryValidatorParams            = staking.NewQueryValidatorParams
	NewQueryBondsParams                = staking.NewQueryBondsParams
	NewQueryRedelegationParams         = staking.NewQueryRedelegationParams
	NewQueryValidatorsParams           = staking.NewQueryValidatorsParams
	NewValidator                       = staking.NewValidator
	MustMarshalValidator               = staking.MustMarshalValidator
	MustUnmarshalValidator             = staking.MustUnmarshalValidator
	UnmarshalValidator                 = staking.UnmarshalValidator
	NewDescription                     = staking.NewDescription
	WriteValidators                    = staking.WriteValidators
	NewCosmosAppModule                 = staking.NewAppModule

	// variable aliases
	CosmosModuleCdc                  = staking.ModuleCdc
	LastValidatorPowerKey            = staking.LastValidatorPowerKey
	LastTotalPowerKey                = staking.LastTotalPowerKey
	ValidatorsKey                    = staking.ValidatorsKey
	ValidatorsByConsAddrKey          = staking.ValidatorsByConsAddrKey
	ValidatorsByPowerIndexKey        = staking.ValidatorsByPowerIndexKey
	DelegationKey                    = staking.DelegationKey
	UnbondingDelegationKey           = staking.UnbondingDelegationKey
	UnbondingDelegationByValIndexKey = staking.UnbondingDelegationByValIndexKey
	RedelegationKey                  = staking.RedelegationKey
	RedelegationByValSrcIndexKey     = staking.RedelegationByValSrcIndexKey
	RedelegationByValDstIndexKey     = staking.RedelegationByValDstIndexKey
	UnbondingQueueKey                = staking.UnbondingQueueKey
	RedelegationQueueKey             = staking.RedelegationQueueKey
	ValidatorQueueKey                = staking.ValidatorQueueKey
	KeyUnbondingTime                 = staking.KeyUnbondingTime
	KeyMaxValidators                 = staking.KeyMaxValidators
	KeyMaxEntries                    = staking.KeyMaxEntries
	KeyBondDenom                     = staking.KeyBondDenom
)

type (
	Keeper                    = staking.Keeper
	Commission                = staking.Commission
	CommissionRates           = staking.CommissionRates
	DVPair                    = staking.DVPair
	DVVTriplet                = staking.DVVTriplet
	Delegation                = staking.Delegation
	Delegations               = staking.Delegations
	UnbondingDelegation       = staking.UnbondingDelegation
	UnbondingDelegationEntry  = staking.UnbondingDelegationEntry
	UnbondingDelegations      = staking.UnbondingDelegations
	Redelegation              = staking.Redelegation
	RedelegationEntry         = staking.RedelegationEntry
	Redelegations             = staking.Redelegations
	DelegationResponse        = staking.DelegationResponse
	DelegationResponses       = staking.DelegationResponses
	RedelegationResponse      = staking.RedelegationResponse
	RedelegationEntryResponse = staking.RedelegationEntryResponse
	RedelegationResponses     = staking.RedelegationResponses
	CodeType                  = staking.CodeType
	GenesisState              = staking.GenesisState
	LastValidatorPower        = staking.LastValidatorPower
	MultiStakingHooks         = staking.MultiStakingHooks
	MsgCreateValidator        = staking.MsgCreateValidator
	MsgEditValidator          = staking.MsgEditValidator
	MsgDelegate               = staking.MsgDelegate
	MsgBeginRedelegate        = staking.MsgBeginRedelegate
	MsgUndelegate             = staking.MsgUndelegate
	Params                    = staking.Params
	Pool                      = staking.Pool
	QueryDelegatorParams      = staking.QueryDelegatorParams
	QueryValidatorParams      = staking.QueryValidatorParams
	QueryBondsParams          = staking.QueryBondsParams
	QueryRedelegationParams   = staking.QueryRedelegationParams
	QueryValidatorsParams     = staking.QueryValidatorsParams
	Validator                 = staking.Validator
	Validators                = staking.Validators
	Description               = staking.Description
	DelegationI               = staking.DelegationI
	ValidatorI                = staking.ValidatorI
	CosmosAppModule           = staking.AppModule
	CosmosAppModuleBasic      = staking.AppModuleBasic
)