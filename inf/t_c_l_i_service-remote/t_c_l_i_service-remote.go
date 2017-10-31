// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
        "context"
        "flag"
        "fmt"
        "math"
        "net"
        "net/url"
        "os"
        "strconv"
        "strings"
        "git.apache.org/thrift.git/lib/go/thrift"
        "inf"
)


func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  TOpenSessionResp OpenSession(TOpenSessionReq req)")
  fmt.Fprintln(os.Stderr, "  TCloseSessionResp CloseSession(TCloseSessionReq req)")
  fmt.Fprintln(os.Stderr, "  TGetInfoResp GetInfo(TGetInfoReq req)")
  fmt.Fprintln(os.Stderr, "  TExecuteStatementResp ExecuteStatement(TExecuteStatementReq req)")
  fmt.Fprintln(os.Stderr, "  TGetTypeInfoResp GetTypeInfo(TGetTypeInfoReq req)")
  fmt.Fprintln(os.Stderr, "  TGetCatalogsResp GetCatalogs(TGetCatalogsReq req)")
  fmt.Fprintln(os.Stderr, "  TGetSchemasResp GetSchemas(TGetSchemasReq req)")
  fmt.Fprintln(os.Stderr, "  TGetTablesResp GetTables(TGetTablesReq req)")
  fmt.Fprintln(os.Stderr, "  TGetTableTypesResp GetTableTypes(TGetTableTypesReq req)")
  fmt.Fprintln(os.Stderr, "  TGetColumnsResp GetColumns(TGetColumnsReq req)")
  fmt.Fprintln(os.Stderr, "  TGetFunctionsResp GetFunctions(TGetFunctionsReq req)")
  fmt.Fprintln(os.Stderr, "  TGetPrimaryKeysResp GetPrimaryKeys(TGetPrimaryKeysReq req)")
  fmt.Fprintln(os.Stderr, "  TGetCrossReferenceResp GetCrossReference(TGetCrossReferenceReq req)")
  fmt.Fprintln(os.Stderr, "  TGetOperationStatusResp GetOperationStatus(TGetOperationStatusReq req)")
  fmt.Fprintln(os.Stderr, "  TCancelOperationResp CancelOperation(TCancelOperationReq req)")
  fmt.Fprintln(os.Stderr, "  TCloseOperationResp CloseOperation(TCloseOperationReq req)")
  fmt.Fprintln(os.Stderr, "  TGetResultSetMetadataResp GetResultSetMetadata(TGetResultSetMetadataReq req)")
  fmt.Fprintln(os.Stderr, "  TFetchResultsResp FetchResults(TFetchResultsReq req)")
  fmt.Fprintln(os.Stderr, "  TGetDelegationTokenResp GetDelegationToken(TGetDelegationTokenReq req)")
  fmt.Fprintln(os.Stderr, "  TCancelDelegationTokenResp CancelDelegationToken(TCancelDelegationTokenReq req)")
  fmt.Fprintln(os.Stderr, "  TRenewDelegationTokenResp RenewDelegationToken(TRenewDelegationTokenReq req)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  var parsedUrl url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Parse()
  
  if len(urlString) > 0 {
    parsedUrl, err := url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  client := inf.NewTCLIServiceClientFactory(trans, protocolFactory)
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "OpenSession":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "OpenSession requires 1 args")
      flag.Usage()
    }
    arg73 := flag.Arg(1)
    mbTrans74 := thrift.NewTMemoryBufferLen(len(arg73))
    defer mbTrans74.Close()
    _, err75 := mbTrans74.WriteString(arg73)
    if err75 != nil {
      Usage()
      return
    }
    factory76 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt77 := factory76.GetProtocol(mbTrans74)
    argvalue0 := inf.NewTOpenSessionReq()
    err78 := argvalue0.Read(jsProt77)
    if err78 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.OpenSession(context.Background(), value0))
    fmt.Print("\n")
    break
  case "CloseSession":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CloseSession requires 1 args")
      flag.Usage()
    }
    arg79 := flag.Arg(1)
    mbTrans80 := thrift.NewTMemoryBufferLen(len(arg79))
    defer mbTrans80.Close()
    _, err81 := mbTrans80.WriteString(arg79)
    if err81 != nil {
      Usage()
      return
    }
    factory82 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt83 := factory82.GetProtocol(mbTrans80)
    argvalue0 := inf.NewTCloseSessionReq()
    err84 := argvalue0.Read(jsProt83)
    if err84 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CloseSession(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetInfo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetInfo requires 1 args")
      flag.Usage()
    }
    arg85 := flag.Arg(1)
    mbTrans86 := thrift.NewTMemoryBufferLen(len(arg85))
    defer mbTrans86.Close()
    _, err87 := mbTrans86.WriteString(arg85)
    if err87 != nil {
      Usage()
      return
    }
    factory88 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt89 := factory88.GetProtocol(mbTrans86)
    argvalue0 := inf.NewTGetInfoReq()
    err90 := argvalue0.Read(jsProt89)
    if err90 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetInfo(context.Background(), value0))
    fmt.Print("\n")
    break
  case "ExecuteStatement":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "ExecuteStatement requires 1 args")
      flag.Usage()
    }
    arg91 := flag.Arg(1)
    mbTrans92 := thrift.NewTMemoryBufferLen(len(arg91))
    defer mbTrans92.Close()
    _, err93 := mbTrans92.WriteString(arg91)
    if err93 != nil {
      Usage()
      return
    }
    factory94 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt95 := factory94.GetProtocol(mbTrans92)
    argvalue0 := inf.NewTExecuteStatementReq()
    err96 := argvalue0.Read(jsProt95)
    if err96 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.ExecuteStatement(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetTypeInfo":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetTypeInfo requires 1 args")
      flag.Usage()
    }
    arg97 := flag.Arg(1)
    mbTrans98 := thrift.NewTMemoryBufferLen(len(arg97))
    defer mbTrans98.Close()
    _, err99 := mbTrans98.WriteString(arg97)
    if err99 != nil {
      Usage()
      return
    }
    factory100 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt101 := factory100.GetProtocol(mbTrans98)
    argvalue0 := inf.NewTGetTypeInfoReq()
    err102 := argvalue0.Read(jsProt101)
    if err102 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetTypeInfo(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetCatalogs":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetCatalogs requires 1 args")
      flag.Usage()
    }
    arg103 := flag.Arg(1)
    mbTrans104 := thrift.NewTMemoryBufferLen(len(arg103))
    defer mbTrans104.Close()
    _, err105 := mbTrans104.WriteString(arg103)
    if err105 != nil {
      Usage()
      return
    }
    factory106 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt107 := factory106.GetProtocol(mbTrans104)
    argvalue0 := inf.NewTGetCatalogsReq()
    err108 := argvalue0.Read(jsProt107)
    if err108 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetCatalogs(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetSchemas":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetSchemas requires 1 args")
      flag.Usage()
    }
    arg109 := flag.Arg(1)
    mbTrans110 := thrift.NewTMemoryBufferLen(len(arg109))
    defer mbTrans110.Close()
    _, err111 := mbTrans110.WriteString(arg109)
    if err111 != nil {
      Usage()
      return
    }
    factory112 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt113 := factory112.GetProtocol(mbTrans110)
    argvalue0 := inf.NewTGetSchemasReq()
    err114 := argvalue0.Read(jsProt113)
    if err114 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetSchemas(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetTables":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetTables requires 1 args")
      flag.Usage()
    }
    arg115 := flag.Arg(1)
    mbTrans116 := thrift.NewTMemoryBufferLen(len(arg115))
    defer mbTrans116.Close()
    _, err117 := mbTrans116.WriteString(arg115)
    if err117 != nil {
      Usage()
      return
    }
    factory118 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt119 := factory118.GetProtocol(mbTrans116)
    argvalue0 := inf.NewTGetTablesReq()
    err120 := argvalue0.Read(jsProt119)
    if err120 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetTables(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetTableTypes":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetTableTypes requires 1 args")
      flag.Usage()
    }
    arg121 := flag.Arg(1)
    mbTrans122 := thrift.NewTMemoryBufferLen(len(arg121))
    defer mbTrans122.Close()
    _, err123 := mbTrans122.WriteString(arg121)
    if err123 != nil {
      Usage()
      return
    }
    factory124 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt125 := factory124.GetProtocol(mbTrans122)
    argvalue0 := inf.NewTGetTableTypesReq()
    err126 := argvalue0.Read(jsProt125)
    if err126 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetTableTypes(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetColumns":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetColumns requires 1 args")
      flag.Usage()
    }
    arg127 := flag.Arg(1)
    mbTrans128 := thrift.NewTMemoryBufferLen(len(arg127))
    defer mbTrans128.Close()
    _, err129 := mbTrans128.WriteString(arg127)
    if err129 != nil {
      Usage()
      return
    }
    factory130 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt131 := factory130.GetProtocol(mbTrans128)
    argvalue0 := inf.NewTGetColumnsReq()
    err132 := argvalue0.Read(jsProt131)
    if err132 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetColumns(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetFunctions":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetFunctions requires 1 args")
      flag.Usage()
    }
    arg133 := flag.Arg(1)
    mbTrans134 := thrift.NewTMemoryBufferLen(len(arg133))
    defer mbTrans134.Close()
    _, err135 := mbTrans134.WriteString(arg133)
    if err135 != nil {
      Usage()
      return
    }
    factory136 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt137 := factory136.GetProtocol(mbTrans134)
    argvalue0 := inf.NewTGetFunctionsReq()
    err138 := argvalue0.Read(jsProt137)
    if err138 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetFunctions(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetPrimaryKeys":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetPrimaryKeys requires 1 args")
      flag.Usage()
    }
    arg139 := flag.Arg(1)
    mbTrans140 := thrift.NewTMemoryBufferLen(len(arg139))
    defer mbTrans140.Close()
    _, err141 := mbTrans140.WriteString(arg139)
    if err141 != nil {
      Usage()
      return
    }
    factory142 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt143 := factory142.GetProtocol(mbTrans140)
    argvalue0 := inf.NewTGetPrimaryKeysReq()
    err144 := argvalue0.Read(jsProt143)
    if err144 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetPrimaryKeys(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetCrossReference":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetCrossReference requires 1 args")
      flag.Usage()
    }
    arg145 := flag.Arg(1)
    mbTrans146 := thrift.NewTMemoryBufferLen(len(arg145))
    defer mbTrans146.Close()
    _, err147 := mbTrans146.WriteString(arg145)
    if err147 != nil {
      Usage()
      return
    }
    factory148 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt149 := factory148.GetProtocol(mbTrans146)
    argvalue0 := inf.NewTGetCrossReferenceReq()
    err150 := argvalue0.Read(jsProt149)
    if err150 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetCrossReference(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetOperationStatus":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetOperationStatus requires 1 args")
      flag.Usage()
    }
    arg151 := flag.Arg(1)
    mbTrans152 := thrift.NewTMemoryBufferLen(len(arg151))
    defer mbTrans152.Close()
    _, err153 := mbTrans152.WriteString(arg151)
    if err153 != nil {
      Usage()
      return
    }
    factory154 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt155 := factory154.GetProtocol(mbTrans152)
    argvalue0 := inf.NewTGetOperationStatusReq()
    err156 := argvalue0.Read(jsProt155)
    if err156 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetOperationStatus(context.Background(), value0))
    fmt.Print("\n")
    break
  case "CancelOperation":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CancelOperation requires 1 args")
      flag.Usage()
    }
    arg157 := flag.Arg(1)
    mbTrans158 := thrift.NewTMemoryBufferLen(len(arg157))
    defer mbTrans158.Close()
    _, err159 := mbTrans158.WriteString(arg157)
    if err159 != nil {
      Usage()
      return
    }
    factory160 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt161 := factory160.GetProtocol(mbTrans158)
    argvalue0 := inf.NewTCancelOperationReq()
    err162 := argvalue0.Read(jsProt161)
    if err162 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CancelOperation(context.Background(), value0))
    fmt.Print("\n")
    break
  case "CloseOperation":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CloseOperation requires 1 args")
      flag.Usage()
    }
    arg163 := flag.Arg(1)
    mbTrans164 := thrift.NewTMemoryBufferLen(len(arg163))
    defer mbTrans164.Close()
    _, err165 := mbTrans164.WriteString(arg163)
    if err165 != nil {
      Usage()
      return
    }
    factory166 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt167 := factory166.GetProtocol(mbTrans164)
    argvalue0 := inf.NewTCloseOperationReq()
    err168 := argvalue0.Read(jsProt167)
    if err168 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CloseOperation(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetResultSetMetadata":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetResultSetMetadata requires 1 args")
      flag.Usage()
    }
    arg169 := flag.Arg(1)
    mbTrans170 := thrift.NewTMemoryBufferLen(len(arg169))
    defer mbTrans170.Close()
    _, err171 := mbTrans170.WriteString(arg169)
    if err171 != nil {
      Usage()
      return
    }
    factory172 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt173 := factory172.GetProtocol(mbTrans170)
    argvalue0 := inf.NewTGetResultSetMetadataReq()
    err174 := argvalue0.Read(jsProt173)
    if err174 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetResultSetMetadata(context.Background(), value0))
    fmt.Print("\n")
    break
  case "FetchResults":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "FetchResults requires 1 args")
      flag.Usage()
    }
    arg175 := flag.Arg(1)
    mbTrans176 := thrift.NewTMemoryBufferLen(len(arg175))
    defer mbTrans176.Close()
    _, err177 := mbTrans176.WriteString(arg175)
    if err177 != nil {
      Usage()
      return
    }
    factory178 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt179 := factory178.GetProtocol(mbTrans176)
    argvalue0 := inf.NewTFetchResultsReq()
    err180 := argvalue0.Read(jsProt179)
    if err180 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.FetchResults(context.Background(), value0))
    fmt.Print("\n")
    break
  case "GetDelegationToken":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetDelegationToken requires 1 args")
      flag.Usage()
    }
    arg181 := flag.Arg(1)
    mbTrans182 := thrift.NewTMemoryBufferLen(len(arg181))
    defer mbTrans182.Close()
    _, err183 := mbTrans182.WriteString(arg181)
    if err183 != nil {
      Usage()
      return
    }
    factory184 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt185 := factory184.GetProtocol(mbTrans182)
    argvalue0 := inf.NewTGetDelegationTokenReq()
    err186 := argvalue0.Read(jsProt185)
    if err186 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.GetDelegationToken(context.Background(), value0))
    fmt.Print("\n")
    break
  case "CancelDelegationToken":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "CancelDelegationToken requires 1 args")
      flag.Usage()
    }
    arg187 := flag.Arg(1)
    mbTrans188 := thrift.NewTMemoryBufferLen(len(arg187))
    defer mbTrans188.Close()
    _, err189 := mbTrans188.WriteString(arg187)
    if err189 != nil {
      Usage()
      return
    }
    factory190 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt191 := factory190.GetProtocol(mbTrans188)
    argvalue0 := inf.NewTCancelDelegationTokenReq()
    err192 := argvalue0.Read(jsProt191)
    if err192 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.CancelDelegationToken(context.Background(), value0))
    fmt.Print("\n")
    break
  case "RenewDelegationToken":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "RenewDelegationToken requires 1 args")
      flag.Usage()
    }
    arg193 := flag.Arg(1)
    mbTrans194 := thrift.NewTMemoryBufferLen(len(arg193))
    defer mbTrans194.Close()
    _, err195 := mbTrans194.WriteString(arg193)
    if err195 != nil {
      Usage()
      return
    }
    factory196 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt197 := factory196.GetProtocol(mbTrans194)
    argvalue0 := inf.NewTRenewDelegationTokenReq()
    err198 := argvalue0.Read(jsProt197)
    if err198 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.RenewDelegationToken(context.Background(), value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
