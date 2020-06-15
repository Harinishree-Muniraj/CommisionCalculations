// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	//"fmt"
	//"log"
	"testing"
	//"github.com/golang/protobuf/proto"
	cel "github.com/google/cel-go/cel"
	//"github.com/google/cel-go/common/types/ref"
	"github.com/google/cel-go/checker/decls"
	//"github.com/google/cel-go/interpreter/functions"
	//"github.com/google/cel-go/common/types"
	//exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
)



// func main() {
// 	// Create the CEL environment with declarations for the input attributes and
// 	// the desired extension functions. In many cases the desired functionality will
// 	// be present in a built-in function.
// 	decls := cel.Declarations(
// 		// Identifiers used within this expression.
// 		decls.NewVar("i", decls.String),
// 		decls.NewVar("you", decls.String),
// 		// Function to generate a greeting from one person to another.
// 		//    i.greet(you)
// 		decls.NewFunction("greet",
// 			decls.NewInstanceOverload("string_greet_string",
// 				[]*exprpb.Type{decls.String, decls.String},
// 				decls.String)))
// 	e, err := cel.NewEnv(decls)
// 	if err != nil {
// 		log.Fatalf("environment creation error: %s\n", err)
// 	}
	
// 	//fmt.Println("hello",e)
	
// 	// Compile the expression.
// 	ast, iss := e.Compile("i.greet(you)")
// 	if iss.Err() != nil {
// 		log.Fatalln(iss.Err())
// 	}
	
// 	// Create the program.
// 	funcs := cel.Functions(
// 		&functions.Overload{
// 			Operator: "string_greet_string",
// 			Binary: func(lhs ref.Val, rhs ref.Val) ref.Val {
// 				return types.String(
// 					fmt.Sprintf("Hello %s! Nice to meet you, I'm %s.\n", rhs, lhs))
// 			}})
			
// 	prg, err := e.Program(ast, funcs)
// 	if err != nil {
// 		log.Fatalf("program creation error: %s\n", err)
// 	}
	
// 	// Evaluate the program against some inputs. Note: the details return is not used.
// 	out, _, err := prg.Eval(map[string]interface{}{
// 		// Native values are converted to CEL values under the covers.
// 		"i": "CEL",
// 		// Values may also be lazily supplied.
// 		"you": func() ref.Val { return types.String("world") },
// 	})
// 	if err != nil {
// 		log.Fatalf("runtime error: %s\n", err)
// 	}
	
// 	fmt.Println(out)
// 	Example_globalOverload()
// 	Benchmark_EvalOptions()

// }


// func Example_globalOverload() {
// 	// Create the CEL environment with declarations for the input attributes and
// 	// the desired extension functions. In many cases the desired functionality will
// 	// be present in a built-in function.
// 	decls := cel.Declarations(
// 		// Identifiers used within this expression.
// 		decls.NewVar("i", decls.String),
// 		decls.NewVar("you", decls.String),
// 		// Function to generate shake_hands between two people.
// 		//    shake_hands(i,you)
// 		decls.NewFunction("shake_hands",
// 			decls.NewOverload("shake_hands_string_string",
// 				[]*exprpb.Type{decls.String, decls.String},
// 				decls.String)))
// 				fmt.Println(decls)
// 	e, err := cel.NewEnv(decls)
// 	if err != nil {
// 		log.Fatalf("environment creation error: %s\n", err)
// 	}
// 	fmt.Println(e,"e")

// 	// Compile the expression.
// 	ast, iss := e.Compile(`shake_hands(i,you)`)
// 	if iss.Err() != nil {
// 		log.Fatalln(iss.Err())
// 	}

// 	// Create the program.
// 	funcs := cel.Functions(
// 		&functions.Overload{
// 			Operator: "shake_hands_string_string",
// 			Binary: func(lhs ref.Val, rhs ref.Val) ref.Val {
// 				s1, ok := lhs.(types.String)
// 				if !ok {
// 					return types.ValOrErr(lhs, "unexpected type '%v' passed to shake_hands", lhs.Type())
// 				}
// 				s2, ok := rhs.(types.String)
// 				if !ok {
// 					return types.ValOrErr(rhs, "unexpected type '%v' passed to shake_hands", rhs.Type())
// 				}
// 				return types.String(
// 					fmt.Sprintf("%s and %s are shaking hands.\n", s1, s2))
// 			}})
// 	prg, err := e.Program(ast, funcs)
// 	if err != nil {
// 		log.Fatalf("program creation error: %s\n", err)
// 	}

// 	// Evaluate the program against some inputs. Note: the details return is not used.
// 	out, _, err := prg.Eval(map[string]interface{}{
// 		"i":   "CEL",
// 		"you": func() ref.Val { return types.String("world") },
// 	})
// 	if err != nil {
// 		log.Fatalf("runtime error: %s\n", err)
// 	}

// 	fmt.Println(out)
// 	// Output:CEL and world are shaking hands.
// }




// func Test_CustomTypes(t *testing.T) {
// 	exprType := decls.NewObjectType("google.api.expr.v1alpha1.Expr")
// 	reg := types.NewEmptyRegistry()
// 	e, _ := cel.NewEnv(
// 		cel.CustomTypeAdapter(reg),
// 		cel.CustomTypeProvider(reg),
// 		cel.Container("google.api.expr.v1alpha1"),
// 		cel.Types(
// 			&exprpb.Expr{},
// 			types.BoolType,
// 			types.IntType,
// 			types.StringType),
// 			cel.Declarations(
// 			decls.NewVar("expr", exprType)))

// 	ast, _ := e.Compile(`
// 		expr == Expr{id: 2,
// 			call_expr: Expr.Call{
// 				function: "_==_",
// 				args: [
// 					Expr{id: 1, ident_expr: Expr.Ident{ name: "a" }},
// 					Expr{id: 3, ident_expr: Expr.Ident{ name: "b" }}]
// 			}}`)
// 	if !proto.Equal(ast.ResultType(), decls.Bool) {
// 		t.Fatalf("got %v, wanted type bool", ast.ResultType())
// 	}
// 	prg, _ := e.Program(ast)
// 	vars := map[string]interface{}{"expr": &exprpb.Expr{
// 		Id: 2,
// 		ExprKind: &exprpb.Expr_CallExpr{
// 			CallExpr: &exprpb.Expr_Call{
// 				Function: "_==_",
// 				Args: []*exprpb.Expr{
// 					{
// 						Id: 1,
// 						ExprKind: &exprpb.Expr_IdentExpr{
// 							IdentExpr: &exprpb.Expr_Ident{Name: "a"},
// 						},
// 					},
// 					{
// 						Id: 3,
// 						ExprKind: &exprpb.Expr_IdentExpr{
// 							IdentExpr: &exprpb.Expr_Ident{Name: "b"},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}}
// 	out, _, _ := prg.Eval(vars)
// 	if out != types.True {
// 		t.Errorf("got '%v', wanted 'true'", out.Value())
// 	}
// }

func Benchmark_EvalOptions(b *testing.B) {
	e, _ := cel.NewEnv(
		cel.Declarations(
			decls.NewVar("ai", decls.Int),
			decls.NewVar("ar", decls.NewMapType(decls.String, decls.String)),
		),
	)
	ast, _ := e.Compile("ai == 20 || ar['foo'] == 'bar'")
	vars := map[string]interface{}{
		"ai": 2,
		"ar": map[string]string{
			"foo": "bar",
		},
	}

	opts := map[string]cel.EvalOption{
		"track-state":     cel.OptTrackState,
		"exhaustive-eval": cel.OptExhaustiveEval,
		"optimize":        cel.OptOptimize,
	}
	for k, opt := range opts {
		b.Run(k, func(bb *testing.B) {
			prg, _ := e.Program(ast, cel.EvalOptions(opt))
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < bb.N; i++ {
				prg.Eval(vars)
				
			}
		})
	}
}