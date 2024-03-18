package inits

import (
	"time"

	"github.com/adshao/go-binance/v2"
)

const (
	ExpireTime = "2024-12-01 00:00:00"
)

const (
	// endpoints
	EndpointFutureWS string = "wss://fstream.binance.com/ws"
	//
	TimeoutBinanceAPI time.Duration = 5 * time.Second // 5s
	ReconnectInterval time.Duration = 2 * time.Second // 2s

	//
	MaxNetworkDelayCnt int = 5

	// binance host
	BinanceHost  string = "api.binance.com"
	BinanceHost1 string = "api1.binance.com"
	BinanceHost2 string = "api2.binance.com"
	BinanceHost3 string = "api3.binance.com"
	BinanceHost4 string = "api4.binance.com"

	// global quote coins
	CoinFDUSD string = "FDUSD"
	CoinUSDT  string = "USDT"
	CoinTUSD  string = "TUSD"
	CoinBUSD  string = "BUSD"
	CoinUSDC  string = "USDC"
	CoinAEUR  string = "AEUR"
	CoinEUR   string = "EUR"
	CoinBNB   string = "BNB"
	CoinBTC   string = "BTC"
	CoinETH   string = "ETH"
	CoinTRY   string = "TRY"
	CoinRUB   string = "RUB"
	CoinGBP   string = "GBP"
	CoinDAI   string = "DAI"

	// global symbols
	// with free charge
	SymbolBTCTUSD   string = "BTCTUSD"
	SymbolTUSDUSDT  string = "TUSDUSDT"
	SymbolTUSDBUSD  string = "TUSDBUSD"
	SymbolBUSDUSDT  string = "BUSDUSDT"
	SymbolUSDCUSDT  string = "USDCUSDT"
	SymbolBTCFDUSD  string = "BTCFDUSD"
	SymbolFDUSDUSDT string = "FDUSDUSDT"
	SymbolFDUSDBUSD string = "FDUSDBUSD"
	SymbolUSDPUSDT  string = "USDPUSDT"
	// with charge
	SymbolBTCBUSD string = "BTCBUSD"
	SymbolBTCUSDT string = "BTCUSDT"
	SymbolAPTETH  string = "APTETH"

	// buy BNB with USDT
	SymbolBNBUSDT string = "BNBUSDT"

	StatusTrading string = "TRADING"

	// asset
	AssetBNB string  = "BNB"
	USDT10   float64 = 10
	USDT15   float64 = 15
)

var (
	Float001        float64 = 0.01
	Float098        float64 = 0.98
	Float099        float64 = 0.99
	Float0999       float64 = 0.999
	Float1          float64 = 1
	Float1Point001  float64 = 1.001
	Float1Point02   float64 = 1.02
	Float1Point1    float64 = 1.1
	Float1Point8    float64 = 1.8
	Float1Point0013 float64 = 1.0013
	Float2          float64 = 2
	Float3          float64 = 3
	Float13         float64 = 13
	Float30         float64 = 30
	Float10000      float64 = 10000
)

const (
	ZeroFloat float64 = 1e-8

	MaxRetryCnt   int    = 3
	QuanlityConst string = "333333"

	Mod1000 int32 = 1000

	ExitMsgTransferFailure        string = "failed to transfer from funding to main"
	ExitMsgReachPlanQuoteQuantity string = "cumulative quantity reaches the plan quote quantity"
	ExitMsgBNBFree                string = "the progress will exit because of BNB"
	ExitMsgTotalValueThrwhold     string = "the progress will exit because total value is less than threshold"
	ExitMsgAssetZero              string = "one of the asset is zero"
	ExitMsgReachingExitingTime    string = "the progress will exit because of reaching exiting time"
	ExitMsgUnexpectedCostFee      string = "the progress will exit because of unexpected cost fee"
)

const (
	EventAccountUpdate binance.UserDataEventType = "outboundAccountPosition"
	EventBalanceUpdate binance.UserDataEventType = "balanceUpdate"
	EventOrderUpdate   binance.UserDataEventType = "executionReport"
)

var (
	HoldCoinMap = map[string]bool{}
)

const (
	SepEmpty           string = ""
	SepDoubleUnderline string = "__"
	SepDoubleEnter     string = "\n\n"

	DecimalPlaces3 int32 = 3
	DecimalPlaces8 int32 = 8
)

const (
	WSSourceBinance string = "binance"
)

const (
	ErrMsgAskOrBidValue string = "ask or bid value is too small"
)
