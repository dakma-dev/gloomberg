package debug

//
// usage example:
//
// // debug a ttx at the last possible moment to have as much data as possible
// if ttx.IsItemBid() && ttx.TotalTokens > 1 && ttx.Marketplace == &marketplace.Blur && ttx.GetTransfersByContract()[common.HexToAddress("0xd564c25b760cb278a55bdd98831f4ff4b6c97b38")] != nil {
// 	debug.DebugIt(ttx)
// }

// PrintIt prints information about a object for debugging purposes.
