/* eslint:disable */
// hello commonjs - we need some imports - sorted in alphabetical order, by go package
import * as github_com_foomo_gotsrpc_demo from './demo'; // demo/output-commonjs/demo.ts to demo/output-commonjs/demo.ts
import * as github_com_foomo_gotsrpc_demo_nested from './demo-nested'; // demo/output-commonjs/demo.ts to demo/output-commonjs/demo-nested.ts
// github.com/foomo/gotsrpc/demo.Address
export interface Address {
	city?:string;
	signs?:string[];
	PeoplePtr:github_com_foomo_gotsrpc_demo.Person[];
	ArrayOfMaps:{[index:string]:boolean}[];
	ArrayArrayAddress:github_com_foomo_gotsrpc_demo.Address[][];
	People:github_com_foomo_gotsrpc_demo.Person[];
	MapCrap:{[index:string]:{[index:number]:boolean}};
	NestedPtr?:github_com_foomo_gotsrpc_demo_nested.Nested;
	NestedStruct:github_com_foomo_gotsrpc_demo_nested.Nested;
}
// github.com/foomo/gotsrpc/demo.Bar
export type Bar = any
// github.com/foomo/gotsrpc/demo.Check
export interface Check {
	Foo:string;
}
// github.com/foomo/gotsrpc/demo.CustomTypeFoo
export type CustomTypeFoo = string
// github.com/foomo/gotsrpc/demo.CustomTypeInt
export type CustomTypeInt = 1 | 2 | 3
// github.com/foomo/gotsrpc/demo.CustomTypeString
export enum CustomTypeString {
	one = "one",
	two = "two",
	three = "three",
}
// github.com/foomo/gotsrpc/demo.CustomTypeStruct
export interface CustomTypeStruct {
	CustomTypeFoo:github_com_foomo_gotsrpc_demo.CustomTypeFoo;
	CustomTypeInt:github_com_foomo_gotsrpc_demo.CustomTypeInt;
	CustomTypeString:github_com_foomo_gotsrpc_demo.CustomTypeString;
	CustomTypeNested:github_com_foomo_gotsrpc_demo_nested.CustomTypeNested;
	Check:github_com_foomo_gotsrpc_demo.Check;
}
// github.com/foomo/gotsrpc/demo.Err
export interface Err {
	message:string;
}
// github.com/foomo/gotsrpc/demo.Person
export interface Person {
	Name:string;
	address?:github_com_foomo_gotsrpc_demo.Address;
	AddressStruct:github_com_foomo_gotsrpc_demo.Address;
	Addresses:{[index:string]:github_com_foomo_gotsrpc_demo.Address};
	InlinePtr?:{
		Foo:boolean;
	};
	InlineStruct:{
		Bar:string;
	};
	DNA:string;
}
// github.com/foomo/gotsrpc/demo.ScalarError
export type ScalarError = string
// github.com/foomo/gotsrpc/demo.ScalarInPlace
export type ScalarInPlace = string
// end of common js