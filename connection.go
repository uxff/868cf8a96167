// Package hivething wraps the hiveserver2 thrift interface in a few
// related interfaces for more convenient use.
package gohive

import (
	//"context"
	//"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"

	inf "git.dev.acewill.net/rpc/gohive/tcliservice"
	//inf "github.com/dazheng/gohive/tcliservice"

	"git.apache.org/thrift.git/lib/go/thrift"
)

// Options for opened Hive sessions.
type Options struct {
	PollIntervalSeconds int64
	BatchSize           int64
	QueryTimeout        int64
}

var (
	DefaultOptions = Options{PollIntervalSeconds: 2, BatchSize: 8, QueryTimeout: 100}
)

type Connection struct {
	thrift  *inf.TCLIServiceClient
	session *inf.TSessionHandle
	options Options
}

func Connect(host string, options Options) (*Connection, error) {
	transport, err := thrift.NewTSocket(host)
	if err != nil {
		return nil, err
	}

	if err := transport.Open(); err != nil {
		return nil, err
	}

	if transport == nil {
		return nil, errors.New("nil thrift transport")
	}

	/*
		NB: hive 0.13's default is a TSaslProtocol, but
		there isn't a golang implementation in apache thrift as
		of this writing.
	*/
	protocol := thrift.NewTBinaryProtocolFactoryDefault()
	client := inf.NewTCLIServiceClientFactory(transport, protocol)
	s := inf.NewTOpenSessionReq()
	s.ClientProtocol = 6
	//	session, err := client.OpenSession(inf.NewTOpenSessionReq())
	session, err := client.OpenSession(context.Background(), s)
	if err != nil {
		return nil, err
	}

	return &Connection{client, session.SessionHandle, options}, nil
}

func (c *Connection) isOpen() bool {
	return c.session != nil
}

// Closes an open hive session. After using this, the
// connection is invalid for other use.
func (c *Connection) Close() error {
	if c.isOpen() {
		closeReq := inf.NewTCloseSessionReq()
		closeReq.SessionHandle = c.session
		resp, err := c.thrift.CloseSession(context.Background(), closeReq)
		if err != nil {
			return fmt.Errorf("Error closing session: ", resp, err)
		}

		c.session = nil
	}

	return nil
}

// Issue a query on an open connection, returning a RowSet, which
// can be later used to query the operation's status.
func (c *Connection) Query(query string) (RowSet, error) {
	executeReq := inf.NewTExecuteStatementReq()
	executeReq.SessionHandle = c.session
	executeReq.Statement = query
	executeReq.RunAsync = true

	resp, err := c.thrift.ExecuteStatement(context.Background(), executeReq)
	if err != nil {
		return nil, fmt.Errorf("Error in ExecuteStatement: %+v, %v", resp, err)
	}

	if !isSuccessStatus(resp.Status) {
		return nil, fmt.Errorf("Error from server: %s", resp.Status.String())
	}

	fmt.Println("push query ok:", query)

	return newRowSet(c.thrift, resp.OperationHandle, c.options), nil
}

func (c *Connection) ExecMode(query string, isAsync bool) (*inf.TExecuteStatementResp, error) {
	executeReq := inf.NewTExecuteStatementReq()
	executeReq.SessionHandle = c.session
	executeReq.Statement = query
	executeReq.RunAsync = isAsync
	executeReq.QueryTimeout = DefaultOptions.QueryTimeout

	resp, err := c.thrift.ExecuteStatement(context.Background(), executeReq)
	if err != nil {
		return nil, fmt.Errorf("Error in ExecuteStatement: %+v, %v", resp, err)
	}

	if !isSuccessStatus(resp.Status) {
		return nil, fmt.Errorf("Error from server: %s", resp.Status.String())
	}

	return resp, err
}

func (c *Connection) Exec(query string) (*inf.TExecuteStatementResp, error) {
	return c.ExecMode(query, false)
}

func isSuccessStatus(p *inf.TStatus) bool {
	status := p.GetStatusCode()
	return status == inf.TStatusCode_SUCCESS_STATUS || status == inf.TStatusCode_SUCCESS_WITH_INFO_STATUS
}

func (c *Connection) FetchOne(r *inf.TExecuteStatementResp) (rows *inf.TRowSet, hasMoreRows bool, e error) {
	fetchReq := inf.NewTFetchResultsReq()
	fetchReq.OperationHandle = r.OperationHandle
	fetchReq.Orientation = inf.TFetchOrientation_FETCH_NEXT
	fetchReq.MaxRows = DefaultOptions.BatchSize

	resp, err := c.thrift.FetchResults(context.Background(), fetchReq)
	if err != nil {
		fmt.Printf("FetchResults failed: %v\n", err)
		return nil, false, err
	}

	if !isSuccessStatus(resp.Status) {
		fmt.Printf("FetchResults failed: %s\n", resp.Status.String())
		return nil, false, errors.New("FetchResult failed, status not ok: " + resp.Status.String())
	}

	rows = resp.GetResults() // return *TRowSet{StartRowOffset int64, Rows []*TRow, Columns []*TColumn, BinaryColumns []byte, ColumnCount *int32}

	//fmt.Println("the fetch rows=", rows)

	// for json debug
	//jret, jerr := json.Marshal(rows)
	//fmt.Println("json rows=", string(jret), jerr)

	// GetHasMoreRow()没生效，返回总是false
	return rows, resp.GetHasMoreRows(), nil
}

func (c *Connection) GetMetadata(r *inf.TExecuteStatementResp) (*inf.TTableSchema, error) {
	req := inf.NewTGetResultSetMetadataReq()
	req.OperationHandle = r.OperationHandle

	resp, err := c.thrift.GetResultSetMetadata(context.Background(), req)

	if err != nil {
		fmt.Println("GetMetadata failed:", err)
		return nil, err
	}

	schema := resp.GetSchema()

	// for json debug
	//jret, jerr := json.Marshal(schema)
	//fmt.Println("schema=", string(jret), jerr)

	return schema, nil
}

/*
	Simple Query
	1. Exec
	2. FetchResult
	3. GetMetadata
	4. Convert to map
*/
func (c *Connection) SimpleQuery(sql string) (rets []map[string]interface{}, err error) {
	resp, err := c.ExecMode(sql, true)
	if err != nil {
		return nil, err
	}

	// wait for ok
	status, err := c.WaitForOk(resp.GetOperationHandle())
	if err != nil {
		fmt.Println("when waiting occur error:", err, " status=", status.String())
		return nil, err
	}

	schema, err := c.GetMetadata(resp)
	if err != nil {
		return nil, err
	}

	/*
		"columns": [
		       {
		           "i64Val": {
		               "values": [
		                   1,
		                   2
		               ],
		               "nulls": "AA=="
		           }
		       },
		       {
		           "stringVal": {
		               "values": [
		                   "f14581122165221",
		                   "t14581122175212"
		               ],
		               "nulls": "AA=="
		           }
		       },
			...
	*/

	// multiple fetch til all result have got
	var recvLen int
	for {
		var rowLen int

		rows, hasMore, err := c.FetchOne(resp)
		if rows == nil || err != nil {
			fmt.Println("the FetchResult is nil")
			return nil, err
		}

		//var colNames []string
		var colValues = make(map[string]interface{}, 0)

		for cpos, tcol := range rows.Columns {
			// 此循环内遍历列名 取出所有列下的结果

			colName := schema.Columns[cpos].ColumnName

			//fmt.Println("cpos=", cpos, "tcol=", tcol)

			switch true {
			case tcol.IsSetBinaryVal():
				//fmt.Println("IsSetBinaryVal")
				colValues[colName] = tcol.GetBinaryVal().GetValues()
				rowLen = len(tcol.GetBinaryVal().GetValues())
			case tcol.IsSetBoolVal():
				//fmt.Println("IsSetBoolVal")
				colValues[colName] = tcol.GetBoolVal().GetValues()
				rowLen = len(tcol.GetBoolVal().GetValues())
			case tcol.IsSetByteVal():
				//fmt.Println("IsSetByteVal")
				colValues[colName] = tcol.GetByteVal().GetValues()
				rowLen = len(tcol.GetByteVal().GetValues())
			case tcol.IsSetDoubleVal():
				//fmt.Println("IsSetDoubleVal", tcol.GetDoubleVal().GetValues())
				colValues[colName] = tcol.GetDoubleVal().GetValues()
				rowLen = len(tcol.GetDoubleVal().GetValues())
			case tcol.IsSetI16Val():
				//fmt.Println("IsSetI16Val", tcol.GetI16Val().GetValues())
				colValues[colName] = tcol.GetI16Val().GetValues()
				rowLen = len(tcol.GetI16Val().GetValues())
			case tcol.IsSetI32Val():
				//fmt.Println("IsSetI32Val", tcol.GetI32Val().GetValues())
				colValues[colName] = tcol.GetI32Val().GetValues()
				rowLen = len(tcol.GetI32Val().GetValues())
			case tcol.IsSetI64Val():
				//fmt.Println("IsSetI64Val", tcol.GetI64Val().GetValues())
				colValues[colName] = tcol.GetI64Val().GetValues()
				rowLen = len(tcol.GetI64Val().GetValues())
			case tcol.IsSetStringVal():
				//fmt.Println("IsSetStringVal")
				colValues[colName] = tcol.GetStringVal().GetValues()
				rowLen = len(tcol.GetStringVal().GetValues())
			}
		}

		// 将列结构转换成行结构
		for i := 0; i < rowLen; i++ {
			formatedRow := make(map[string]interface{}, 0)
			for colName, colValueList := range colValues {
				// column => [v1, v2, v3, ...]
				formatedRow[colName] = reflect.ValueOf(colValueList).Index(i).Interface()
			}

			rets = append(rets, formatedRow)
		}

		recvLen += rowLen

		// hasMoreRow 没生效
		if !hasMore {
			fmt.Println("now more rows, this time rowlen=", rowLen, "StartRowOffset=", rows.StartRowOffset)
			//break
		} else {
			fmt.Println("has more rows, this time rowlen=", rowLen, "StartRowOffset=", rows.StartRowOffset)
		}
		// 需要从返回的数量里判断任务有没有进行完
		if rowLen <= 0 {
			fmt.Println("no more rows find, rowlen=", rowLen)
			break
		}
	}

	return rets, nil
}

// Issue a thrift call to check for the job's current status.
func (c *Connection) CheckStatus(operation *inf.TOperationHandle) (*Status, error) {
	req := inf.NewTGetOperationStatusReq()
	req.OperationHandle = operation

	fmt.Println("will request GetOperationStatus")

	resp, err := c.thrift.GetOperationStatus(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("Error getting status: %+v, %v", resp, err)
	}

	if !isSuccessStatus(resp.Status) {
		return nil, fmt.Errorf("GetStatus call failed: %s", resp.Status.String())
	}

	if resp.OperationState == nil {
		return nil, errors.New("No error from GetStatus, but nil status!")
	}

	fmt.Println("OperationStatus", resp.GetOperationState(), "ProgressUpdate=", resp.GetProgressUpdateResponse())

	return &Status{resp.OperationState, nil, time.Now()}, nil
}

// Wait until the job is complete, one way or another, returning Status and error.
func (c *Connection) WaitForOk(operation *inf.TOperationHandle) (*Status, error) {
	for {
		status, err := c.CheckStatus(operation)

		if err != nil {
			return nil, err
		}

		if status.IsComplete() {
			if status.IsSuccess() {
				return status, nil
			}
			return nil, fmt.Errorf("Query failed execution: %s", status.state.String())
		}

		time.Sleep(time.Duration(DefaultOptions.PollIntervalSeconds) * time.Second)
	}
	return nil, errors.New("Cannot run here when wait for operation ok")
}
