package main

import (
	"bitfinex/bot/subscriptions"
	"bitfinex/bot/timescale"
)

var tickers = []string{"tBTCUSD", "tLTCUSD", "tLTCBTC", "tETHUSD", "tETHBTC",
	"tETCBTC", "tETCUSD", "tRRTUSD", "tRRTBTC", "tZECUSD", "tZECBTC", "tXMRUSD",
	"tXMRBTC", "tDSHUSD", "tDSHBTC", "tBTCEUR", "tBTCJPY", "tXRPUSD", "tXRPBTC",
	"tIOTUSD", "tIOTBTC", "tIOTETH", "tEOSUSD", "tEOSBTC", "tEOSETH", "tSANUSD",
	"tSANBTC", "tSANETH", "tOMGUSD", "tOMGBTC", "tOMGETH", "tNEOUSD", "tNEOBTC",
	"tNEOETH", "tETPUSD", "tETPBTC", "tETPETH", "tQTMUSD", "tQTMBTC", "tQTMETH",
	"tAVTUSD", "tAVTBTC", "tAVTETH", "tEDOUSD", "tEDOBTC", "tEDOETH", "tBTGUSD",
	"tBTGBTC", "tDATUSD", "tDATBTC", "tDATETH", "tQSHUSD", "tQSHBTC", "tQSHETH",
	"tYYWUSD", "tYYWBTC", "tYYWETH", "tGNTUSD", "tGNTBTC", "tGNTETH", "tSNTUSD",
	"tSNTBTC", "tSNTETH", "tIOTEUR", "tBATUSD", "tBATBTC", "tBATETH", "tMNAUSD",
	"tMNABTC", "tMNAETH", "tFUNUSD", "tFUNBTC", "tFUNETH", "tZRXUSD", "tZRXBTC",
	"tZRXETH", "tTNBUSD", "tTNBBTC", "tTNBETH", "tSPKUSD", "tSPKBTC", "tSPKETH",
	"tTRXUSD", "tTRXBTC", "tTRXETH", "tRCNUSD", "tRCNBTC", "tRCNETH", "tRLCUSD",
	"tRLCBTC", "tRLCETH", "tAIDUSD", "tAIDBTC", "tAIDETH", "tSNGUSD", "tSNGBTC",
	"tSNGETH", "tREPUSD", "tREPBTC", "tREPETH", "tELFUSD", "tELFBTC", "tELFETH",
	"tNECUSD", "tNECBTC", "tNECETH", "tBTCGBP", "tETHEUR", "tETHJPY", "tETHGBP",
	"tNEOEUR", "tNEOJPY", "tNEOGBP", "tEOSEUR", "tEOSJPY", "tEOSGBP", "tIOTJPY",
	"tIOTGBP", "tIOSUSD", "tIOSBTC", "tIOSETH", "tAIOUSD", "tAIOBTC", "tAIOETH",
	"tREQUSD", "tREQBTC", "tREQETH", "tRDNUSD", "tRDNBTC", "tRDNETH", "tLRCUSD",
	"tLRCBTC", "tLRCETH", "tWAXUSD", "tWAXBTC", "tWAXETH", "tDAIUSD", "tDAIBTC",
	"tDAIETH", "tAGIUSD", "tAGIBTC", "tAGIETH", "tBFTUSD", "tBFTBTC", "tBFTETH",
	"tMTNUSD", "tMTNBTC", "tMTNETH", "tODEUSD", "tODEBTC", "tODEETH", "tANTUSD",
	"tANTBTC", "tANTETH", "tDTHUSD", "tDTHBTC", "tDTHETH", "tMITUSD", "tMITBTC",
	"tMITETH", "tSTJUSD", "tSTJBTC", "tSTJETH", "tXLMUSD", "tXLMEUR", "tXLMJPY",
	"tXLMGBP", "tXLMBTC", "tXLMETH", "tXVGUSD", "tXVGEUR", "tXVGJPY", "tXVGGBP",
	"tXVGBTC", "tXVGETH", "tBCIUSD", "tBCIBTC", "tMKRUSD", "tMKRBTC", "tMKRETH",
	"tKNCUSD", "tKNCBTC", "tKNCETH", "tPOAUSD", "tPOABTC", "tPOAETH", "tEVTUSD",
	"tLYMUSD", "tLYMBTC", "tLYMETH", "tUTKUSD", "tUTKBTC", "tUTKETH", "tVEEUSD",
	"tVEEBTC", "tVEEETH", "tDADUSD", "tDADBTC", "tDADETH", "tORSUSD", "tORSBTC",
	"tORSETH", "tAUCUSD", "tAUCBTC", "tAUCETH", "tPOYUSD", "tPOYBTC", "tPOYETH",
	"tFSNUSD", "tFSNBTC", "tFSNETH", "tCBTUSD", "tCBTBTC", "tCBTETH", "tZCNUSD",
	"tZCNBTC", "tZCNETH", "tSENUSD", "tSENBTC", "tSENETH", "tNCAUSD", "tNCABTC",
	"tNCAETH", "tCNDUSD", "tCNDBTC", "tCNDETH", "tCTXUSD", "tCTXBTC", "tCTXETH",
	"tPAIUSD", "tPAIBTC", "tSEEUSD", "tSEEBTC", "tSEEETH", "tESSUSD", "tESSBTC",
	"tESSETH", "tATMUSD", "tATMBTC", "tATMETH", "tHOTUSD", "tHOTBTC", "tHOTETH",
	"tDTAUSD", "tDTABTC", "tDTAETH", "tIQXUSD", "tIQXBTC", "tIQXEOS", "tWPRUSD",
	"tWPRBTC", "tWPRETH", "tZILUSD", "tZILBTC", "tZILETH", "tBNTUSD", "tBNTBTC",
	"tBNTETH", "tABSUSD", "tABSETH", "tXRAUSD", "tXRAETH", "tMANUSD", "tMANETH",
	"tBBNUSD", "tBBNETH", "tNIOUSD", "tNIOETH", "tDGXUSD", "tDGXETH", "tVETUSD",
	"tVETBTC", "tVETETH", "tUTNUSD", "tUTNETH", "tTKNUSD", "tTKNETH", "tGOTUSD",
	"tGOTEUR", "tGOTETH", "tXTZUSD", "tXTZBTC", "tCNNUSD", "tCNNETH", "tBOXUSD",
	"tBOXETH", "tTRXEUR", "tTRXGBP", "tTRXJPY", "tMGOUSD", "tMGOETH", "tRTEUSD",
	"tRTEETH", "tYGGUSD", "tYGGETH", "tMLNUSD", "tMLNETH", "tWTCUSD", "tWTCETH",
	"tCSXUSD", "tCSXETH", "tOMNUSD", "tOMNBTC", "tINTUSD", "tINTETH", "tDRNUSD",
	"tDRNETH", "tPNKUSD", "tPNKETH", "tDGBUSD", "tDGBBTC", "tBSVUSD", "tBSVBTC",
	"tBABUSD", "tBABBTC", "tWLOUSD", "tWLOXLM", "tVLDUSD", "tVLDETH", "tENJUSD",
	"tENJETH", "tONLUSD", "tONLETH", "tRBTUSD", "tRBTBTC", "tUSTUSD", "tEUTEUR",
	"tEUTUSD", "tGSDUSD", "tUDCUSD", "tTSDUSD", "tPAXUSD", "tRIFUSD", "tRIFBTC",
	"tPASUSD", "tPASETH", "tVSYUSD", "tVSYBTC", "tZRXDAI", "tMKRDAI", "tOMGDAI",
	"tBTTUSD", "tBTTBTC", "tBTCUST", "tETHUST", "tCLOUSD", "tCLOBTC", "tIMPUSD",
	"tIMPETH", "tLTCUST", "tEOSUST", "tBABUST", "tSCRUSD", "tSCRETH", "tGNOUSD",
	"tGNOETH", "tGENUSD", "tGENETH", "tATOUSD", "tATOBTC", "tATOETH", "tWBTUSD",
	"tXCHUSD", "tEUSUSD", "tWBTETH", "tXCHETH", "tEUSETH", "tLEOUSD", "tLEOBTC",
	"tLEOUST", "tLEOEOS", "tLEOETH", "tASTUSD", "tASTETH", "tFOAUSD", "tFOAETH",
	"tUFRUSD", "tUFRETH", "tZBTUSD", "tZBTUST", "tOKBUSD", "tUSKUSD", "tGTXUSD",
	"tKANUSD", "tOKBUST", "tOKBETH", "tOKBBTC", "tUSKUST", "tUSKETH", "tUSKBTC",
	"tUSKEOS", "tGTXUST", "tKANUST", "tAMPUSD", "tALGUSD", "tALGBTC", "tALGUST",
	"tBTCXCH", "tSWMUSD", "tSWMETH", "tTRIUSD", "tTRIETH", "tLOOUSD", "tLOOETH",
	"tAMPUST", "tDUSK:USD", "tDUSK:BTC", "tUOSUSD", "tUOSBTC", "tBTCF0:USTF0",
	"tETHF0:USTF0"}

func main() {
	subs := subscriptions.BitfinexSubscriptions{Tickers: tickers, Position: 0, Initialised: false}
	timescale.SetupDb()
	subs.NewConnection()
}
