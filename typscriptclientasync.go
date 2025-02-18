package gotsrpc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/foomo/gotsrpc/config"
)

func renderTypescriptClientAsync(service *Service, mappings config.TypeScriptMappings, scalars map[string]*Scalar, structs map[string]*Struct, ts *code) error {
	clientName := service.Name + "Client"

	ts.l("export class " + clientName + " {")

	ts.ind(1)
	//ts.l(`static defaultInst = new ` + clientName + `()`)
	//ts.l(`constructor(public endpoint = "` + service.Endpoint + `") {}`)
	ts.l(`public static defaultEndpoint = "` + service.Endpoint + `";`)
	ts.l("constructor(")
	ts.ind(1)
	ts.l("public transport:<T>(method: string, data?: any[]) => Promise<T>")
	ts.ind(-1)
	ts.l(") {}")

	for _, method := range service.Methods {
		ts.app("async " + lcfirst(method.Name) + "(")
		callArgs := []string{}
		argOffset := 0
		for index, arg := range method.Args {
			if index == 0 && arg.Value.isHTTPResponseWriter() {
				trace("skipping first arg is a http.ResponseWriter")
				argOffset = 1
				continue
			}
			if index == 1 && arg.Value.isHTTPRequest() {
				trace("skipping second arg is a *http.Request")
				argOffset = 2
				continue
			}
		}
		argCount := 0
		for index, arg := range method.Args {
			if index < argOffset {
				continue
			}
			if index > argOffset {
				ts.app(", ")
			}
			ts.app(arg.tsName() + ":")
			arg.Value.tsType(mappings, scalars, structs, ts)
			callArgs = append(callArgs, arg.Name)
			argCount++
		}
		ts.app("):")

		throwLastError := false
		//lastErrorName := ""

		returnTypeTS := newCode("	")
		returnTypeTS.app("{")
		innerReturnTypeTS := newCode("	")
		innerReturnTypeTS.app("{")
		firstReturnType := ""
		//firstReturnFieldName := ""
		countReturns := 0
		countInnerReturns := 0
		errIndex := 0
		responseObjectPrefix := ""
		responseObject := "let responseObject = {"

		for index, retField := range method.Return {
			countInnerReturns++
			retArgName := retField.tsName()

			if len(retArgName) == 0 {
				retArgName = "ret"
				if index > 0 {
					retArgName += "_" + fmt.Sprint(index)
				}
			}
			if index > 0 {
				returnTypeTS.app("; ")
				innerReturnTypeTS.app("; ")
			}

			innerReturnTypeTS.app(strconv.Itoa(index) + ":")
			retField.Value.tsType(mappings, scalars, structs, innerReturnTypeTS)

			if index == len(method.Return)-1 && retField.Value.IsError {
				throwLastError = true
				//lastErrorName = retArgName
				errIndex = index
			} else {
				if index == 0 {
					firstReturnTypeTS := newCode("	")
					retField.Value.tsType(mappings, scalars, structs, firstReturnTypeTS)
					firstReturnType = firstReturnTypeTS.string()
					//firstReturnFieldName = retArgName
				}
				countReturns++
				returnTypeTS.app(retArgName + ":")
				responseObject += responseObjectPrefix + retArgName + " : response[" + strconv.Itoa(index) + "]"
				retField.Value.tsType(mappings, scalars, structs, returnTypeTS)
			}
			responseObjectPrefix = ", "
		}
		responseObject += "};"
		returnTypeTS.app("}")
		innerReturnTypeTS.app("}")
		if countReturns == 0 {
			ts.app("Promise<void> {")
		} else if countReturns == 1 {
			ts.app("Promise<" + firstReturnType + "> {")
		} else if countReturns > 1 {
			ts.app("Promise<" + returnTypeTS.string() + "> {")
		}
		ts.nl()

		ts.ind(1)

		innerCallTypeString := "void"
		if countInnerReturns > 0 {
			innerCallTypeString = innerReturnTypeTS.string()
		}

		call := "this.transport<" + innerCallTypeString + ">(\"" + method.Name + "\", [" + strings.Join(callArgs, ", ") + "])"
		if throwLastError {
			ts.l("let response = await " + call)

			ts.l("let err = response[" + strconv.Itoa(errIndex) + "];")
			//ts.l("delete response." + lastErrorName + ";")
			ts.l("if(err) { throw err }")
			if countReturns == 1 {
				ts.l("return response[0]")
			} else if countReturns == 0 {
				//ts.l("return response;")
			} else {
				ts.l(responseObject)
				ts.l("return responseObject;")
			}
		} else {
			if countReturns == 0 {
				ts.l("await " + call)
			} else if countReturns == 1 {
				ts.l("return (await " + call + ")[0]")
			} else {
				ts.l("let response = await " + call)
				ts.l(responseObject)
				ts.l("return responseObject;")
			}
		}

		ts.ind(-1)
		ts.app("}")
		ts.nl()
	}
	ts.ind(-1)
	ts.l("}")
	return nil
}
