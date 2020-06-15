
package main

import (
	"fmt"
	"log"
	"encoding/json"
	//"testing"
	//"github.com/golang/protobuf/proto"
	cel "github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types/ref"
	"github.com/google/cel-go/checker/decls"
	"github.com/google/cel-go/interpreter/functions"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/traits"
	"github.com/google/cel-go/common/operators"
	"github.com/google/cel-go/common/overloads"
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
)

type Datadeal struct {
	Id int
	Name, Value,Owner,Pipeline,Currency,CloseDate,CreatedDate,CreatedBy string
	ProductLineItem map[string]string
	Title []string
	
}
type Result struct {
    Id   string        `json:"Id"`
    Owner string        `json:"Owner"`
	ProductLineItem map[string]string `json:"ProductLineItem"`
	Title []string `json:"Title"`
}



// func main() {
// 	e, _ := cel.NewEnv(
// 		cel.Declarations(
// 			decls.NewVar("ai", decls.Int),
// 			decls.NewVar("ar", decls.NewMapType(decls.String, decls.String)),
// 		),
// 	)
// 	ast, _ := e.Compile("ai == 20 && ar['foo'] == 'bar'")
// 	vars := map[string]interface{}{
// 		"ai": 20,
// 		"ar": map[string]string{
// 			"foo": "bar",
// 		},
// 	}

// 	opts := map[string]cel.EvalOption{
// 		"track-state":     cel.OptTrackState,
// 		"exhaustive-eval": cel.OptExhaustiveEval,
// 		"optimize":        cel.OptOptimize,
// 	}
// 	for _, opt := range opts {
		 
// 			prg, _ := e.Program(ast, cel.EvalOptions(opt))
			
			
// 				fmt.Println(vars)
// 				//prg.Eval(vars)
// 				out, _, _ := prg.Eval(vars)
				
// 				fmt.Println(out)		
	
// }
// }


func main() {


	var Data Datadeal
	Data=Datadeal{
		 Id:123,
		Name: "John",
		Value:"23",
		Owner:"daniel",
		Pipeline:"value",
		ProductLineItem:map[string]string{
			"Quantity": "johndoe91",
			"ProductType":"CustomerPreference",
			"Price":"2500",
			"Currency":"$25",
		},
		Title:[]string{"Account Manager"},
		Currency:"$25",
		CloseDate:"25.06.2020",
		CreatedDate:"21.01.2020",
		CreatedBy:"John",
	}


	DatadealJSON, err := json.MarshalIndent( Data, "", "  " )

	fmt.Println( string(DatadealJSON), err )
	deal := (DatadealJSON)
	//fmt.Println(deal)
	var result Result 
	json.Unmarshal(deal, &result)
	fmt.Println(result)
	//var owner =result.Owner
	//var Title=result.Title

	e, _ := cel.NewCustomEnv(
		cel.Declarations(
			decls.NewVar("title", decls.String),
			decls.NewFunction(
				operators.In,
				decls.NewOverload(overloads.InList, []*exprpb.Type{
					decls.String, decls.NewListType(decls.String),
				}, decls.Bool),
				)),
		cel.HomogeneousAggregateLiterals())


		_, iss := e.Compile("title in ['Account Manager', 'Custom Success Manager',0]")
		if iss == nil || iss.Err() == nil {
			log.Fatalf("got successful compile, expected error for mixed list entry types.")
		}


	
	funcs := cel.Functions(&functions.Overload{
		Operator: operators.In,
		Binary: func(lhs ref.Val, rhs ref.Val) ref.Val {
			if rhs.Type().HasTrait(traits.ContainerType) {
				return rhs.(traits.Container).Contains(lhs)
			}
			return types.ValOrErr(rhs, "no such overload")
		},
	})

		ast, iss := e.Compile("title in ['Account Manager', 'Custom Success Manager']")
		if iss.Err() != nil {
			log.Fatalf("got issue: %v, expected successful compile.", iss.Err())
		}
		prg, _ := e.Program(ast, funcs)
		out, _, err := prg.Eval(map[string]interface{}{"title": "Account Manager"})
		if err != nil {
			log.Fatalf("got err: %v, wanted result", err)
		}
		if out != types.True {
			log.Fatalf("got '%v', wanted 'true'", out)
		}
		fmt.Println(out)
	}