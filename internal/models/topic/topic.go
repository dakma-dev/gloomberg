package topic

type Topic string

const (
	OrdersMatched  Topic = "0xc4109843e0b7d514e4c093114b863f8e7d8d9a458c372cd51bfe526b588006c9"
	Transfer       Topic = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	TransferSingle Topic = "0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"
	ApprovalForAll Topic = "0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31"
	OrderFulfilled Topic = "0x9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f31"
)

func (t Topic) String() string {
	return map[Topic]string{
		OrdersMatched: "OrdersMatched", Transfer: "Transfer", TransferSingle: "TransferSingle", ApprovalForAll: "ApprovalForAll", OrderFulfilled: "OrderFulfilled",
	}[t]
}
