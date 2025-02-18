/* eslint:disable */
// hello commonjs - we need some imports - sorted in alphabetical order, by go package
import * as github_com_foomo_gotsrpc_demo from './demo'; // demo/output-commonjs-async/demo-nested.ts to demo/output-commonjs-async/demo.ts
import * as github_com_foomo_gotsrpc_demo_nested from './demo-nested'; // demo/output-commonjs-async/demo-nested.ts to demo/output-commonjs-async/demo-nested.ts
// github.com/foomo/gotsrpc/demo/nested.Nested
export interface Nested {
	Name:string;
	Any:any;
	AnyMap:{[index:string]:any};
	AnyList:any[];
	SuperNestedString:{
		Ha:number;
	};
	SuperNestedPtr?:{
		Bla:string;
	};
}
// end of common js