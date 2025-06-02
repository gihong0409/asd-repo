package dmrsclient

import (
	"git.datau.co.kr/ferrari/ferrari-common/commonutils"
	"git.datau.co.kr/ferrari/ferrari-common/dmrsapi/dmrsformats"
)

func SelectData(DMRSURL string, CallApp string, XMLName string, QueryName string, reqData []interface{}, reqID string) []byte {
	recvBody := dmrsformats.ReqBody{Data: reqData}
	dmrsHeader := dmrsformats.ReqHeader{
		CmdType: "DBMW_00010",
		CallApp: CallApp,
		XMLName: XMLName,
		Query: QueryName,
		RequestID: reqID,
	}

	return commonutils.RestfulSendData(
		DMRSURL,
		commonutils.MakeJsonData(dmrsHeader, recvBody),
	)
}

func UpdateData(DMRSURL string, CallApp string, XMLName string, QueryName string, reqData []interface{}, trid string) []byte {
	recvBody := dmrsformats.ReqBody{Data: reqData}
	dmrsHeader := dmrsformats.ReqHeader{
		CmdType: "DBMW_00020",
		CallApp: CallApp, 
		XMLName: XMLName,
		Query: QueryName,
		RequestID: trid,
	}

	return commonutils.RestfulSendData(
		DMRSURL,
		commonutils.MakeJsonData(dmrsHeader, recvBody),
	)
}

func DBMCall(dmrsInfo dmrsformats.DMRSInfo, queryType int, queryName string, param []interface{}, recvBody interface{}, requestID string) dmrsformats.RspHeader {
	var dbData []byte
	var recvHeader dmrsformats.RspHeader

	if queryType == dmrsformats.SELECTQUERY {
		dbData = SelectData(
			dmrsInfo.DMRSURL,
			dmrsInfo.APPNAME,
			dmrsInfo.XMLNAME,
			queryName,
			param,
			requestID,
		)
	} else {
		dbData = UpdateData(
			dmrsInfo.DMRSURL,
			dmrsInfo.APPNAME,
			dmrsInfo.XMLNAME,
			queryName,
			param,
			requestID,
		)
	}
	if recvBody == nil {
		commonutils.JsonToHader(dbData, &recvHeader)
	} else {
		commonutils.JsonToHaderBody(dbData, &recvHeader, &recvBody)
	}

	// self.Print(requestID, "DBRS RESULT", string(dbData))

	return recvHeader
}
