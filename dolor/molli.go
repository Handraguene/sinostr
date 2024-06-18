	// transactionBlock: moveTransaction,
	moveTransaction := &pb.TransactionBlock{
		Transactions: []*pb.Transaction{
			{
				Type: &pb.Transaction_MoveTransaction{
					MoveTransaction: &pb.MoveTransaction{
						Source:   common.MustBytesFromHex("0x1"),
						Dest:     common.MustBytesFromHex("0x2"),
						Amount:   common.MustBytesFromHex("0x1"),
						Currency: common.MustBytesFromHex("0x1"),
					},
				},
			},
		},
	}  
