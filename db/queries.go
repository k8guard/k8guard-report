package db

import (
	"fmt"
	"time"

	"github.com/k8guard/k8guard-report/db/stmts"

	libs "github.com/k8guard/k8guardlibs"
)

type VActionResponseModel struct {
	Namespace       string
	Etype           string
	Created_at      time.Time
	Action          string
	Cluster         string
	Source          string
	ViolationSource string
	ViolationType   string
}

type Context struct {
	Namespace string
	Results   map[string][]VActionResponseModel
}

func (m VActionResponseModel) GetAllByNameSpace(namespace string) Context {
	result := make(map[string][]VActionResponseModel, 0)

	listOfActions := []string{"ActionNamespace", "ActionDeployment", "ActionDaemonSet", "ActionIngress", "ActionPod", "ActionJob", "ActionCronJob"}

	for _, action := range listOfActions {

		iter := Sess.Query(fmt.Sprintf(stmts.SELECT_ACTIONS_BY_NAMESPACE, libs.Cfg.CassandraKeyspace), namespace, action, 20).Iter()

		for iter.Scan(&m.Namespace, &m.Etype, &m.Created_at, &m.Action, &m.Cluster, &m.Source, &m.ViolationSource, &m.ViolationType) {
			result[action] = append(result[action], m)
		}
		if err := iter.Close(); err != nil {
			libs.Log.Fatal("Error closing iter ", err)
		}
	}

	libs.Log.Debug("The Results for", namespace, "are ", result)
	return Context{Namespace: namespace, Results: result}
}

func (m VActionResponseModel) GetLastActions(numberOfRecent int) Context {
	libs.Log.Debug("Trying to get last action ", numberOfRecent)
	result := make(map[string][]VActionResponseModel, 0)

	iter := Sess.Query(fmt.Sprintf(stmts.SELECT_ACTIONS, libs.Cfg.CassandraKeyspace), numberOfRecent).Iter()

	for iter.Scan(&m.Namespace, &m.Etype, &m.Created_at, &m.Action, &m.Cluster, &m.Source, &m.ViolationSource, &m.ViolationType) {
		result["recent"] = append(result["recent"], m)
	}
	if err := iter.Close(); err != nil {
		libs.Log.Fatal("Error closing iter ", err)
	}
	libs.Log.Debug("The Results for", result)
	return Context{Results: result}
}

func (m VActionResponseModel) GetLastAction() (*VActionResponseModel, error) {
	if err := Sess.Query(fmt.Sprintf(stmts.SELECT_ACTIONS, libs.Cfg.CassandraKeyspace), 10).Scan(&m.Namespace, &m.Etype, &m.Created_at, &m.Action, &m.Cluster, &m.Source, &m.ViolationSource, &m.ViolationType); err != nil {
		libs.Log.Error(err)
		return &m, err
	}
	return &m, nil
}

func (m VActionResponseModel) Ping() error {
	if err := Sess.Query(stmts.SELECT_CURRENT_TIMEUUID).Exec(); err != nil {
		libs.Log.Error(err)
		return err
	}
	return nil
}
