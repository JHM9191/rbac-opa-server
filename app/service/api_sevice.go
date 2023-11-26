package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/plugins"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/storage"
	"net/http"
	"rbac-opa-server-mariadb/app/constants"
	"rbac-opa-server-mariadb/app/dto"
	"rbac-opa-server-mariadb/app/repository"
	"rbac-opa-server-mariadb/pkg/api"
	"rbac-opa-server-mariadb/pkg/opa"
)

type ApiService interface {
	EvaluateRule(c *gin.Context)
}

type ApiServiceImpl struct {
	dataRepo      repository.DataRepository
	pluginManager *plugins.Manager
	compiler      *ast.Compiler
}

func ApiServiceInit(dataRepository repository.DataRepository) *ApiServiceImpl {
	pm, c := opa.SetRego()
	return &ApiServiceImpl{
		dataRepo:      dataRepository,
		pluginManager: pm,
		compiler:      c,
	}
}

func (a ApiServiceImpl) EvaluateRule(c *gin.Context) {
	defer pkg.PanicHandler(c)

	var input dto.Input

	if err := c.ShouldBindJSON(&input); err != nil {
		pkg.PanicException(constants.InvalidRequest)
		return
	}

	err := a.loadData()
	if err != nil {
		pkg.PanicException(constants.InternalError)
		return
	}

	allow := false
	query := func(txn storage.Transaction) error {
		r := rego.New(
			rego.Query("data.rbac.allow"),
			rego.Input(input),
			rego.Compiler(a.compiler),
			rego.Store(a.pluginManager.Store),
			rego.Transaction(txn))

		result, err := r.Eval(context.Background())
		if err != nil {
			return err
		} else if len(result) == 0 {
			return errors.New("Undefined query.")
		} else if len(result) > 1 {
			return errors.New("Attempt to evaluate non-boolean decision.")
		} else if value, ok := result[0].Expressions[0].Value.(bool); !ok {
			return errors.New("Attempt to evaluate non-boolean decision.")
		} else {
			allow = value
		}
		return nil
	}

	// Execute the query
	if err := storage.Txn(context.Background(), a.pluginManager.Store, storage.TransactionParams{}, query); err != nil {
		pkg.PanicException(constants.InternalError)
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constants.Success, allow))
}

func (a ApiServiceImpl) loadData() error {

	str, err := a.dataRepo.Select()
	if err != nil {
		return err
	}

	str = "{\n  \"roles\": {\n    \"id:customer_1:project:project_1:user:user_1\": {\n      \"id:customer_1:project:project_1\": [\n        \"administrator\"\n      ]\n    },\n    \"id:customer_1:project:project_2:user:user_2\": {\n      \"id:customer_1:project:project_2\": [\n        \"administrator\"\n      ]\n    },\n    \"id:customer_1:project:project_1:user:user_3\": {\n      \"id:customer_1:project:project_1\": [\n        \"viewer\"\n      ]\n    },\n    \"id:customer_1:project:project_2:user:user_4\": {\n      \"id:customer_1:project:project_2\": [\n        \"manager\"\n      ]\n    }\n  },\n  \"permissions\": {\n    \"administrator\": [\n      \"view:resource\",\n      \"update:resource\",\n      \"create:resource\",\n      \"delete:resource\"\n    ],\n    \"viewer\": [\n      \"view:resource\"\n    ],\n    \"manager\": [\n      \"view:resource\",\n      \"update:resource\"\n    ]\n  }\n}"
	data := make(map[string]interface{})
	if err = json.NewDecoder(bytes.NewReader([]byte(str))).Decode(&data); err != nil {
		return err
	}

	// Store the RBAC data in the configured Rego.Store.
	path := make([]string, 0)
	store := a.pluginManager.Store
	txn := storage.NewTransactionOrDie(context.Background(), store, storage.WriteParams)
	err = store.Write(context.Background(), txn, storage.AddOp, path, data)
	if err != nil {
		store.Abort(context.Background(), txn)
		return err
	}
	err = store.Commit(context.Background(), txn)
	if err != nil {
		return err
	}

	return nil

}
