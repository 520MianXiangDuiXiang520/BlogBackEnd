package mmgo

import "go.mongodb.org/mongo-driver/mongo/options"

func UpdateOne(colName string, filter, update any, opts ...*options.UpdateOptions) error {
	_, err := _h.Col(colName).UpdateOne(_h.NewCtxTimeout(), filter, update, opts...)
	return err
}

func UpdateMany(colName string, filter, update any, opts ...*options.UpdateOptions) error {
	_, err := _h.Col(colName).UpdateMany(_h.NewCtxTimeout(), filter, update, opts...)
	return err
}

func UpsetOne(colName string, filter, update any, opts ...*options.UpdateOptions) (upsetCnt int64, err error) {
	opts = append(opts, options.Update().SetUpsert(true))
	res, err := _h.Col(colName).UpdateOne(_h.NewCtxTimeout(), filter, update, opts...)
	if err != nil {
		return 0, err
	}
	return res.UpsertedCount, err
}

func UpsetMany(colName string, filter, update any, opts ...*options.UpdateOptions) (upsetCnt int64, err error) {
	opts = append(opts, options.Update().SetUpsert(true))
	res, err := _h.Col(colName).UpdateMany(_h.NewCtxTimeout(), filter, update, opts...)
	if err != nil {
		return 0, err
	}
	return res.UpsertedCount, err
}

func InsertOne(colName string, doc any, opts ...*options.InsertOneOptions) error {
	_, err := _h.Col(colName).InsertOne(_h.NewCtxTimeout(), doc, opts...)
	return err
}

func InsertMany(colName string, doc []any, opts ...*options.InsertManyOptions) error {
	_, err := _h.Col(colName).InsertMany(_h.NewCtxTimeout(), doc, opts...)
	return err
}

func DeleteOne(colName string, filter any, opts ...*options.DeleteOptions) error {
	_, err := _h.Col(colName).DeleteOne(_h.NewCtxTimeout(), filter, opts...)
	return err
}

func DeleteMany(colName string, filter any, opts ...*options.DeleteOptions) error {
	_, err := _h.Col(colName).DeleteMany(_h.NewCtxTimeout(), filter, opts...)
	return err
}
