package consts

const (
	LastBlockKey             string = "block:last"
	SaveBlockKey             string = "block:save"
	FirstSaveBlockKey        string = "block:first"
	ErrInsertLogsKey         string = "err:insert:log:%s:%d"
	ErrQueryReceiptsKey      string = "err:query:receipt:%s"
	ErrInsertTransactionsKey string = "err:insert:transaction:%s"
	ErrQueryBlocksKey        string = "err:query:block:%d"
	ErrInsertBlocksKey       string = "err:insert:block:%d"
	BlockDataKey             string = "block:data:%d"
	TransactionDataKey       string = "block:data:%d:%s"
	LogDataKey               string = "transaction:data:%s:%d"
)
