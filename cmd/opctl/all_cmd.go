// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package main

import (
	"openpitrix.io/openpitrix/test"
	"openpitrix.io/openpitrix/test/client/app_manager"
	"openpitrix.io/openpitrix/test/client/category_manager"
	"openpitrix.io/openpitrix/test/client/cluster_manager"
	"openpitrix.io/openpitrix/test/client/job_manager"
	"openpitrix.io/openpitrix/test/client/repo_indexer"
	"openpitrix.io/openpitrix/test/client/repo_manager"
	"openpitrix.io/openpitrix/test/client/runtime_manager"
	"openpitrix.io/openpitrix/test/client/task_manager"
	"openpitrix.io/openpitrix/test/models"
)

var AllCmd = []Cmd{
	NewCreateAppCmd(),
	NewCreateAppVersionCmd(),
	NewDeleteAppVersionsCmd(),
	NewDeleteAppsCmd(),
	NewDescribeAppVersionsCmd(),
	NewDescribeAppsCmd(),
	NewGetAppStatisticsCmd(),
	NewGetAppVersionPackageCmd(),
	NewGetAppVersionPackageFilesCmd(),
	NewModifyAppCmd(),
	NewModifyAppVersionCmd(),
	NewCreateCategoryCmd(),
	NewDeleteCategoriesCmd(),
	NewDescribeCategoriesCmd(),
	NewModifyCategoryCmd(),
	NewAddClusterNodesCmd(),
	NewAttachKeyPairsCmd(),
	NewCeaseClustersCmd(),
	NewCreateClusterCmd(),
	NewCreateKeyPairCmd(),
	NewDeleteClusterNodesCmd(),
	NewDeleteClustersCmd(),
	NewDeleteKeyPairsCmd(),
	NewDescribeClusterNodesCmd(),
	NewDescribeClustersCmd(),
	NewDescribeKeyPairsCmd(),
	NewDescribeSubnetsCmd(),
	NewDetachKeyPairsCmd(),
	NewGetClusterStatisticsCmd(),
	NewRecoverClustersCmd(),
	NewResizeClusterCmd(),
	NewRollbackClusterCmd(),
	NewStartClustersCmd(),
	NewStopClustersCmd(),
	NewUpdateClusterEnvCmd(),
	NewUpgradeClusterCmd(),
	NewDescribeJobsCmd(),
	NewDescribeRepoEventsCmd(),
	NewIndexRepoCmd(),
	NewCreateRepoCmd(),
	NewDeleteReposCmd(),
	NewDescribeReposCmd(),
	NewModifyRepoCmd(),
	NewValidateRepoCmd(),
	NewCreateRuntimeCmd(),
	NewDeleteRuntimesCmd(),
	NewDescribeRuntimeProviderZonesCmd(),
	NewDescribeRuntimesCmd(),
	NewGetRuntimeStatisticsCmd(),
	NewModifyRuntimeCmd(),
	NewDescribeTasksCmd(),
	NewRetryTasksCmd(),
}

type CreateAppCmd struct {
	*models.OpenpitrixCreateAppRequest
}

func NewCreateAppCmd() Cmd {
	return &CreateAppCmd{
		&models.OpenpitrixCreateAppRequest{},
	}
}

func (*CreateAppCmd) GetActionName() string {
	return "CreateApp"
}

func (c *CreateAppCmd) ParseFlag(f Flag) {
	f.StringVarP(&c.CategoryID, "category_id", "", "", "")
	f.StringVarP(&c.ChartName, "chart_name", "", "", "")
	f.StringVarP(&c.Description, "description", "", "", "")
	f.StringVarP(&c.Home, "home", "", "", "")
	f.StringVarP(&c.Icon, "icon", "", "", "")
	f.StringVarP(&c.Keywords, "keywords", "", "", "")
	f.StringVarP(&c.Maintainers, "maintainers", "", "", "")
	f.StringVarP(&c.Name, "name", "", "", "")
	f.StringVarP(&c.Owner, "owner", "", "", "")
	f.StringVarP(&c.Readme, "readme", "", "", "")
	f.StringVarP(&c.RepoID, "repo_id", "", "", "")
	f.StringVarP(&c.Screenshots, "screenshots", "", "", "")
	f.StringVarP(&c.Sources, "sources", "", "", "")
}

func (c *CreateAppCmd) Run(out Out) error {
	params := app_manager.NewCreateAppParams()
	params.WithBody(c.OpenpitrixCreateAppRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.AppManager.CreateApp(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type CreateAppVersionCmd struct {
	*models.OpenpitrixCreateAppVersionRequest
}

func NewCreateAppVersionCmd() Cmd {
	return &CreateAppVersionCmd{
		&models.OpenpitrixCreateAppVersionRequest{},
	}
}

func (*CreateAppVersionCmd) GetActionName() string {
	return "CreateAppVersion"
}

func (c *CreateAppVersionCmd) ParseFlag(f Flag) {
	f.StringVarP(&c.AppID, "app_id", "", "", "")
	f.StringVarP(&c.Description, "description", "", "", "")
	f.StringVarP(&c.Name, "name", "", "", "")
	f.StringVarP(&c.Owner, "owner", "", "", "")
	f.StringVarP(&c.PackageName, "package_name", "", "", "")
}

func (c *CreateAppVersionCmd) Run(out Out) error {
	params := app_manager.NewCreateAppVersionParams()
	params.WithBody(c.OpenpitrixCreateAppVersionRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.AppManager.CreateAppVersion(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type DeleteAppVersionsCmd struct {
	*models.OpenpitrixDeleteAppVersionsRequest
}

func NewDeleteAppVersionsCmd() Cmd {
	return &DeleteAppVersionsCmd{
		&models.OpenpitrixDeleteAppVersionsRequest{},
	}
}

func (*DeleteAppVersionsCmd) GetActionName() string {
	return "DeleteAppVersions"
}

func (c *DeleteAppVersionsCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.VersionID, "version_id", "", []string{}, "")
}

func (c *DeleteAppVersionsCmd) Run(out Out) error {
	params := app_manager.NewDeleteAppVersionsParams()
	params.WithBody(c.OpenpitrixDeleteAppVersionsRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.AppManager.DeleteAppVersions(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type DeleteAppsCmd struct {
	*models.OpenpitrixDeleteAppsRequest
}

func NewDeleteAppsCmd() Cmd {
	return &DeleteAppsCmd{
		&models.OpenpitrixDeleteAppsRequest{},
	}
}

func (*DeleteAppsCmd) GetActionName() string {
	return "DeleteApps"
}

func (c *DeleteAppsCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.AppID, "app_id", "", []string{}, "")
}

func (c *DeleteAppsCmd) Run(out Out) error {
	params := app_manager.NewDeleteAppsParams()
	params.WithBody(c.OpenpitrixDeleteAppsRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.AppManager.DeleteApps(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type DescribeAppVersionsCmd struct {
	*app_manager.DescribeAppVersionsParams
}

func NewDescribeAppVersionsCmd() Cmd {
	return &DescribeAppVersionsCmd{
		app_manager.NewDescribeAppVersionsParams(),
	}
}

func (*DescribeAppVersionsCmd) GetActionName() string {
	return "DescribeAppVersions"
}

func (c *DescribeAppVersionsCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.AppID, "app_id", "", []string{}, "")
	f.StringSliceVarP(&c.Description, "description", "", []string{}, "")
	c.Limit = new(int64)
	f.Int64VarP(c.Limit, "limit", "", 20, "")
	f.StringSliceVarP(&c.Name, "name", "", []string{}, "")
	c.Offset = new(int64)
	f.Int64VarP(c.Offset, "offset", "", 0, "")
	f.StringSliceVarP(&c.Owner, "owner", "", []string{}, "")
	f.StringSliceVarP(&c.PackageName, "package_name", "", []string{}, "")
	c.Reverse = new(bool)
	f.BoolVarP(c.Reverse, "reverse", "", false, "")
	c.SearchWord = new(string)
	f.StringVarP(c.SearchWord, "search_word", "", "", "")
	c.SortKey = new(string)
	f.StringVarP(c.SortKey, "sort_key", "", "", "")
	f.StringSliceVarP(&c.Status, "status", "", []string{}, "")
	f.StringSliceVarP(&c.VersionID, "version_id", "", []string{}, "")
}

func (c *DescribeAppVersionsCmd) Run(out Out) error {

	out.WriteRequest(c.DescribeAppVersionsParams)

	client := test.GetClient(clientConfig)
	res, err := client.AppManager.DescribeAppVersions(c.DescribeAppVersionsParams)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type DescribeAppsCmd struct {
	*app_manager.DescribeAppsParams
}

func NewDescribeAppsCmd() Cmd {
	return &DescribeAppsCmd{
		app_manager.NewDescribeAppsParams(),
	}
}

func (*DescribeAppsCmd) GetActionName() string {
	return "DescribeApps"
}

func (c *DescribeAppsCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.AppID, "app_id", "", []string{}, "")
	f.StringSliceVarP(&c.CategoryID, "category_id", "", []string{}, "")
	f.StringSliceVarP(&c.ChartName, "chart_name", "", []string{}, "")
	c.Limit = new(int64)
	f.Int64VarP(c.Limit, "limit", "", 20, "")
	f.StringSliceVarP(&c.Name, "name", "", []string{}, "")
	c.Offset = new(int64)
	f.Int64VarP(c.Offset, "offset", "", 0, "")
	f.StringSliceVarP(&c.Owner, "owner", "", []string{}, "")
	f.StringSliceVarP(&c.RepoID, "repo_id", "", []string{}, "")
	c.Reverse = new(bool)
	f.BoolVarP(c.Reverse, "reverse", "", false, "")
	c.SearchWord = new(string)
	f.StringVarP(c.SearchWord, "search_word", "", "", "")
	c.SortKey = new(string)
	f.StringVarP(c.SortKey, "sort_key", "", "", "")
	f.StringSliceVarP(&c.Status, "status", "", []string{}, "")
}

func (c *DescribeAppsCmd) Run(out Out) error {

	out.WriteRequest(c.DescribeAppsParams)

	client := test.GetClient(clientConfig)
	res, err := client.AppManager.DescribeApps(c.DescribeAppsParams)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type GetAppStatisticsCmd struct {
	*app_manager.GetAppStatisticsParams
}

func NewGetAppStatisticsCmd() Cmd {
	return &GetAppStatisticsCmd{
		app_manager.NewGetAppStatisticsParams(),
	}
}

func (*GetAppStatisticsCmd) GetActionName() string {
	return "GetAppStatistics"
}

func (c *GetAppStatisticsCmd) ParseFlag(f Flag) {
}

func (c *GetAppStatisticsCmd) Run(out Out) error {

	out.WriteRequest(c.GetAppStatisticsParams)

	client := test.GetClient(clientConfig)
	res, err := client.AppManager.GetAppStatistics(c.GetAppStatisticsParams)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type GetAppVersionPackageCmd struct {
	*app_manager.GetAppVersionPackageParams
}

func NewGetAppVersionPackageCmd() Cmd {
	return &GetAppVersionPackageCmd{
		app_manager.NewGetAppVersionPackageParams(),
	}
}

func (*GetAppVersionPackageCmd) GetActionName() string {
	return "GetAppVersionPackage"
}

func (c *GetAppVersionPackageCmd) ParseFlag(f Flag) {
	c.VersionID = new(string)
	f.StringVarP(c.VersionID, "version_id", "", "", "")
}

func (c *GetAppVersionPackageCmd) Run(out Out) error {

	out.WriteRequest(c.GetAppVersionPackageParams)

	client := test.GetClient(clientConfig)
	res, err := client.AppManager.GetAppVersionPackage(c.GetAppVersionPackageParams)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type GetAppVersionPackageFilesCmd struct {
	*app_manager.GetAppVersionPackageFilesParams
}

func NewGetAppVersionPackageFilesCmd() Cmd {
	return &GetAppVersionPackageFilesCmd{
		app_manager.NewGetAppVersionPackageFilesParams(),
	}
}

func (*GetAppVersionPackageFilesCmd) GetActionName() string {
	return "GetAppVersionPackageFiles"
}

func (c *GetAppVersionPackageFilesCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.Files, "files", "", []string{}, "")
	c.VersionID = new(string)
	f.StringVarP(c.VersionID, "version_id", "", "", "")
}

func (c *GetAppVersionPackageFilesCmd) Run(out Out) error {

	out.WriteRequest(c.GetAppVersionPackageFilesParams)

	client := test.GetClient(clientConfig)
	res, err := client.AppManager.GetAppVersionPackageFiles(c.GetAppVersionPackageFilesParams)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type ModifyAppCmd struct {
	*models.OpenpitrixModifyAppRequest
}

func NewModifyAppCmd() Cmd {
	return &ModifyAppCmd{
		&models.OpenpitrixModifyAppRequest{},
	}
}

func (*ModifyAppCmd) GetActionName() string {
	return "ModifyApp"
}

func (c *ModifyAppCmd) ParseFlag(f Flag) {
	f.StringVarP(&c.AppID, "app_id", "", "", "")
	f.StringVarP(&c.CategoryID, "category_id", "", "", "")
	f.StringVarP(&c.ChartName, "chart_name", "", "", "")
	f.StringVarP(&c.Description, "description", "", "", "")
	f.StringVarP(&c.Home, "home", "", "", "")
	f.StringVarP(&c.Icon, "icon", "", "", "")
	f.StringVarP(&c.Keywords, "keywords", "", "", "")
	f.StringVarP(&c.Maintainers, "maintainers", "", "", "")
	f.StringVarP(&c.Name, "name", "", "", "")
	f.StringVarP(&c.Owner, "owner", "", "", "")
	f.StringVarP(&c.Readme, "readme", "", "", "")
	f.StringVarP(&c.RepoID, "repo_id", "", "", "")
	f.StringVarP(&c.Screenshots, "screenshots", "", "", "")
	f.StringVarP(&c.Sources, "sources", "", "", "")
}

func (c *ModifyAppCmd) Run(out Out) error {
	params := app_manager.NewModifyAppParams()
	params.WithBody(c.OpenpitrixModifyAppRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.AppManager.ModifyApp(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type ModifyAppVersionCmd struct {
	*models.OpenpitrixModifyAppVersionRequest
}

func NewModifyAppVersionCmd() Cmd {
	return &ModifyAppVersionCmd{
		&models.OpenpitrixModifyAppVersionRequest{},
	}
}

func (*ModifyAppVersionCmd) GetActionName() string {
	return "ModifyAppVersion"
}

func (c *ModifyAppVersionCmd) ParseFlag(f Flag) {
	f.StringVarP(&c.Description, "description", "", "", "")
	f.StringVarP(&c.Name, "name", "", "", "")
	f.StringVarP(&c.Owner, "owner", "", "", "")
	f.StringVarP(&c.PackageName, "package_name", "", "", "")
	f.StringVarP(&c.VersionID, "version_id", "", "", "")
}

func (c *ModifyAppVersionCmd) Run(out Out) error {
	params := app_manager.NewModifyAppVersionParams()
	params.WithBody(c.OpenpitrixModifyAppVersionRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.AppManager.ModifyAppVersion(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type CreateCategoryCmd struct {
	*models.OpenpitrixCreateCategoryRequest
}

func NewCreateCategoryCmd() Cmd {
	return &CreateCategoryCmd{
		&models.OpenpitrixCreateCategoryRequest{},
	}
}

func (*CreateCategoryCmd) GetActionName() string {
	return "CreateCategory"
}

func (c *CreateCategoryCmd) ParseFlag(f Flag) {
	f.StringVarP(&c.Description, "description", "", "", "")
	f.StringVarP(&c.Locale, "locale", "", "", "")
	f.StringVarP(&c.Name, "name", "", "", "")
}

func (c *CreateCategoryCmd) Run(out Out) error {
	params := category_manager.NewCreateCategoryParams()
	params.WithBody(c.OpenpitrixCreateCategoryRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.CategoryManager.CreateCategory(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type DeleteCategoriesCmd struct {
	*models.OpenpitrixDeleteCategoriesRequest
}

func NewDeleteCategoriesCmd() Cmd {
	return &DeleteCategoriesCmd{
		&models.OpenpitrixDeleteCategoriesRequest{},
	}
}

func (*DeleteCategoriesCmd) GetActionName() string {
	return "DeleteCategories"
}

func (c *DeleteCategoriesCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.CategoryID, "category_id", "", []string{}, "")
}

func (c *DeleteCategoriesCmd) Run(out Out) error {
	params := category_manager.NewDeleteCategoriesParams()
	params.WithBody(c.OpenpitrixDeleteCategoriesRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.CategoryManager.DeleteCategories(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type DescribeCategoriesCmd struct {
	*category_manager.DescribeCategoriesParams
}

func NewDescribeCategoriesCmd() Cmd {
	return &DescribeCategoriesCmd{
		category_manager.NewDescribeCategoriesParams(),
	}
}

func (*DescribeCategoriesCmd) GetActionName() string {
	return "DescribeCategories"
}

func (c *DescribeCategoriesCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.CategoryID, "category_id", "", []string{}, "")
	c.Limit = new(int64)
	f.Int64VarP(c.Limit, "limit", "", 20, "")
	f.StringSliceVarP(&c.Name, "name", "", []string{}, "")
	c.Offset = new(int64)
	f.Int64VarP(c.Offset, "offset", "", 0, "")
	f.StringSliceVarP(&c.Owner, "owner", "", []string{}, "")
	c.Reverse = new(bool)
	f.BoolVarP(c.Reverse, "reverse", "", false, "")
	c.SearchWord = new(string)
	f.StringVarP(c.SearchWord, "search_word", "", "", "")
	c.SortKey = new(string)
	f.StringVarP(c.SortKey, "sort_key", "", "", "")
}

func (c *DescribeCategoriesCmd) Run(out Out) error {

	out.WriteRequest(c.DescribeCategoriesParams)

	client := test.GetClient(clientConfig)
	res, err := client.CategoryManager.DescribeCategories(c.DescribeCategoriesParams)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type ModifyCategoryCmd struct {
	*models.OpenpitrixModifyCategoryRequest
}

func NewModifyCategoryCmd() Cmd {
	return &ModifyCategoryCmd{
		&models.OpenpitrixModifyCategoryRequest{},
	}
}

func (*ModifyCategoryCmd) GetActionName() string {
	return "ModifyCategory"
}

func (c *ModifyCategoryCmd) ParseFlag(f Flag) {
	f.StringVarP(&c.CategoryID, "category_id", "", "", "")
	f.StringVarP(&c.Description, "description", "", "", "")
	f.StringVarP(&c.Locale, "locale", "", "", "")
	f.StringVarP(&c.Name, "name", "", "", "")
}

func (c *ModifyCategoryCmd) Run(out Out) error {
	params := category_manager.NewModifyCategoryParams()
	params.WithBody(c.OpenpitrixModifyCategoryRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.CategoryManager.ModifyCategory(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type AddClusterNodesCmd struct {
	*models.OpenpitrixAddClusterNodesRequest
}

func NewAddClusterNodesCmd() Cmd {
	return &AddClusterNodesCmd{
		&models.OpenpitrixAddClusterNodesRequest{},
	}
}

func (*AddClusterNodesCmd) GetActionName() string {
	return "AddClusterNodes"
}

func (c *AddClusterNodesCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.AdvancedParam, "advanced_param", "", []string{}, "")
	f.StringVarP(&c.ClusterID, "cluster_id", "", "", "")
	f.StringVarP(&c.Role, "role", "", "", "")
}

func (c *AddClusterNodesCmd) Run(out Out) error {
	params := cluster_manager.NewAddClusterNodesParams()
	params.WithBody(c.OpenpitrixAddClusterNodesRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.ClusterManager.AddClusterNodes(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type AttachKeyPairsCmd struct {
	*models.OpenpitrixAttachKeyPairsRequest
}

func NewAttachKeyPairsCmd() Cmd {
	return &AttachKeyPairsCmd{
		&models.OpenpitrixAttachKeyPairsRequest{},
	}
}

func (*AttachKeyPairsCmd) GetActionName() string {
	return "AttachKeyPairs"
}

func (c *AttachKeyPairsCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.KeyPairID, "key_pair_id", "", []string{}, "")
	f.StringSliceVarP(&c.NodeID, "node_id", "", []string{}, "")
}

func (c *AttachKeyPairsCmd) Run(out Out) error {
	params := cluster_manager.NewAttachKeyPairsParams()
	params.WithBody(c.OpenpitrixAttachKeyPairsRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.ClusterManager.AttachKeyPairs(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type CeaseClustersCmd struct {
	*models.OpenpitrixCeaseClustersRequest
}

func NewCeaseClustersCmd() Cmd {
	return &CeaseClustersCmd{
		&models.OpenpitrixCeaseClustersRequest{},
	}
}

func (*CeaseClustersCmd) GetActionName() string {
	return "CeaseClusters"
}

func (c *CeaseClustersCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.AdvancedParam, "advanced_param", "", []string{}, "")
	f.StringSliceVarP(&c.ClusterID, "cluster_id", "", []string{}, "")
}

func (c *CeaseClustersCmd) Run(out Out) error {
	params := cluster_manager.NewCeaseClustersParams()
	params.WithBody(c.OpenpitrixCeaseClustersRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.ClusterManager.CeaseClusters(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type CreateClusterCmd struct {
	*models.OpenpitrixCreateClusterRequest
}

func NewCreateClusterCmd() Cmd {
	return &CreateClusterCmd{
		&models.OpenpitrixCreateClusterRequest{},
	}
}

func (*CreateClusterCmd) GetActionName() string {
	return "CreateCluster"
}

func (c *CreateClusterCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.AdvancedParam, "advanced_param", "", []string{}, "")
	f.StringVarP(&c.AppID, "app_id", "", "", "")
	f.StringVarP(&c.Conf, "conf", "", "", "")
	f.StringVarP(&c.RuntimeID, "runtime_id", "", "", "")
	f.StringVarP(&c.VersionID, "version_id", "", "", "")
}

func (c *CreateClusterCmd) Run(out Out) error {
	params := cluster_manager.NewCreateClusterParams()
	params.WithBody(c.OpenpitrixCreateClusterRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.ClusterManager.CreateCluster(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type CreateKeyPairCmd struct {
	*models.OpenpitrixCreateKeyPairRequest
}

func NewCreateKeyPairCmd() Cmd {
	return &CreateKeyPairCmd{
		&models.OpenpitrixCreateKeyPairRequest{},
	}
}

func (*CreateKeyPairCmd) GetActionName() string {
	return "CreateKeyPair"
}

func (c *CreateKeyPairCmd) ParseFlag(f Flag) {
	f.StringVarP(&c.Description, "description", "", "", "")
	f.StringVarP(&c.Name, "name", "", "", "")
	f.StringVarP(&c.PubKey, "pub_key", "", "", "")
}

func (c *CreateKeyPairCmd) Run(out Out) error {
	params := cluster_manager.NewCreateKeyPairParams()
	params.WithBody(c.OpenpitrixCreateKeyPairRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.ClusterManager.CreateKeyPair(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type DeleteClusterNodesCmd struct {
	*models.OpenpitrixDeleteClusterNodesRequest
}

func NewDeleteClusterNodesCmd() Cmd {
	return &DeleteClusterNodesCmd{
		&models.OpenpitrixDeleteClusterNodesRequest{},
	}
}

func (*DeleteClusterNodesCmd) GetActionName() string {
	return "DeleteClusterNodes"
}

func (c *DeleteClusterNodesCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.AdvancedParam, "advanced_param", "", []string{}, "")
	f.StringVarP(&c.ClusterID, "cluster_id", "", "", "")
	f.StringSliceVarP(&c.NodeID, "node_id", "", []string{}, "")
	f.StringVarP(&c.Role, "role", "", "", "")
}

func (c *DeleteClusterNodesCmd) Run(out Out) error {
	params := cluster_manager.NewDeleteClusterNodesParams()
	params.WithBody(c.OpenpitrixDeleteClusterNodesRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.ClusterManager.DeleteClusterNodes(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type DeleteClustersCmd struct {
	*models.OpenpitrixDeleteClustersRequest
}

func NewDeleteClustersCmd() Cmd {
	return &DeleteClustersCmd{
		&models.OpenpitrixDeleteClustersRequest{},
	}
}

func (*DeleteClustersCmd) GetActionName() string {
	return "DeleteClusters"
}

func (c *DeleteClustersCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.AdvancedParam, "advanced_param", "", []string{}, "")
	f.StringSliceVarP(&c.ClusterID, "cluster_id", "", []string{}, "")
}

func (c *DeleteClustersCmd) Run(out Out) error {
	params := cluster_manager.NewDeleteClustersParams()
	params.WithBody(c.OpenpitrixDeleteClustersRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.ClusterManager.DeleteClusters(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type DeleteKeyPairsCmd struct {
	*models.OpenpitrixDeleteKeyPairsRequest
}

func NewDeleteKeyPairsCmd() Cmd {
	return &DeleteKeyPairsCmd{
		&models.OpenpitrixDeleteKeyPairsRequest{},
	}
}

func (*DeleteKeyPairsCmd) GetActionName() string {
	return "DeleteKeyPairs"
}

func (c *DeleteKeyPairsCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.KeyPairID, "key_pair_id", "", []string{}, "")
}

func (c *DeleteKeyPairsCmd) Run(out Out) error {
	params := cluster_manager.NewDeleteKeyPairsParams()
	params.WithBody(c.OpenpitrixDeleteKeyPairsRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.ClusterManager.DeleteKeyPairs(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type DescribeClusterNodesCmd struct {
	*cluster_manager.DescribeClusterNodesParams
}

func NewDescribeClusterNodesCmd() Cmd {
	return &DescribeClusterNodesCmd{
		cluster_manager.NewDescribeClusterNodesParams(),
	}
}

func (*DescribeClusterNodesCmd) GetActionName() string {
	return "DescribeClusterNodes"
}

func (c *DescribeClusterNodesCmd) ParseFlag(f Flag) {
	c.ClusterID = new(string)
	f.StringVarP(c.ClusterID, "cluster_id", "", "", "")
	c.Limit = new(int64)
	f.Int64VarP(c.Limit, "limit", "", 20, "default is 20, max value is 200.")
	f.StringSliceVarP(&c.NodeID, "node_id", "", []string{}, "")
	c.Offset = new(int64)
	f.Int64VarP(c.Offset, "offset", "", 0, "default is 0.")
	c.SearchWord = new(string)
	f.StringVarP(c.SearchWord, "search_word", "", "", "")
	f.StringSliceVarP(&c.Status, "status", "", []string{}, "")
}

func (c *DescribeClusterNodesCmd) Run(out Out) error {

	out.WriteRequest(c.DescribeClusterNodesParams)

	client := test.GetClient(clientConfig)
	res, err := client.ClusterManager.DescribeClusterNodes(c.DescribeClusterNodesParams)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type DescribeClustersCmd struct {
	*cluster_manager.DescribeClustersParams
}

func NewDescribeClustersCmd() Cmd {
	return &DescribeClustersCmd{
		cluster_manager.NewDescribeClustersParams(),
	}
}

func (*DescribeClustersCmd) GetActionName() string {
	return "DescribeClusters"
}

func (c *DescribeClustersCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.AppID, "app_id", "", []string{}, "")
	f.StringSliceVarP(&c.ClusterID, "cluster_id", "", []string{}, "")
	c.ExternalClusterID = new(string)
	f.StringVarP(c.ExternalClusterID, "external_cluster_id", "", "", "")
	f.StringSliceVarP(&c.FrontgateID, "frontgate_id", "", []string{}, "")
	c.Limit = new(int64)
	f.Int64VarP(c.Limit, "limit", "", 20, "default is 20, max value is 200.")
	c.Offset = new(int64)
	f.Int64VarP(c.Offset, "offset", "", 0, "default is 0.")
	f.StringSliceVarP(&c.RuntimeID, "runtime_id", "", []string{}, "")
	c.SearchWord = new(string)
	f.StringVarP(c.SearchWord, "search_word", "", "", "")
	f.StringSliceVarP(&c.Status, "status", "", []string{}, "")
	f.StringSliceVarP(&c.VersionID, "version_id", "", []string{}, "")
}

func (c *DescribeClustersCmd) Run(out Out) error {

	out.WriteRequest(c.DescribeClustersParams)

	client := test.GetClient(clientConfig)
	res, err := client.ClusterManager.DescribeClusters(c.DescribeClustersParams)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type DescribeKeyPairsCmd struct {
	*cluster_manager.DescribeKeyPairsParams
}

func NewDescribeKeyPairsCmd() Cmd {
	return &DescribeKeyPairsCmd{
		cluster_manager.NewDescribeKeyPairsParams(),
	}
}

func (*DescribeKeyPairsCmd) GetActionName() string {
	return "DescribeKeyPairs"
}

func (c *DescribeKeyPairsCmd) ParseFlag(f Flag) {
	c.Description = new(string)
	f.StringVarP(c.Description, "description", "", "", "")
	c.KeyPairID = new(string)
	f.StringVarP(c.KeyPairID, "key_pair_id", "", "", "")
	c.Limit = new(int64)
	f.Int64VarP(c.Limit, "limit", "", 20, "")
	c.Name = new(string)
	f.StringVarP(c.Name, "name", "", "", "")
	c.Offset = new(int64)
	f.Int64VarP(c.Offset, "offset", "", 0, "")
	c.Owner = new(string)
	f.StringVarP(c.Owner, "owner", "", "", "")
	c.PubKey = new(string)
	f.StringVarP(c.PubKey, "pub_key", "", "", "")
	c.SearchWord = new(string)
	f.StringVarP(c.SearchWord, "search_word", "", "", "")
}

func (c *DescribeKeyPairsCmd) Run(out Out) error {

	out.WriteRequest(c.DescribeKeyPairsParams)

	client := test.GetClient(clientConfig)
	res, err := client.ClusterManager.DescribeKeyPairs(c.DescribeKeyPairsParams)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type DescribeSubnetsCmd struct {
	*cluster_manager.DescribeSubnetsParams
}

func NewDescribeSubnetsCmd() Cmd {
	return &DescribeSubnetsCmd{
		cluster_manager.NewDescribeSubnetsParams(),
	}
}

func (*DescribeSubnetsCmd) GetActionName() string {
	return "DescribeSubnets"
}

func (c *DescribeSubnetsCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.AdvancedParam, "advanced_param", "", []string{}, "")
	c.Limit = new(int64)
	f.Int64VarP(c.Limit, "limit", "", 20, "")
	c.Offset = new(int64)
	f.Int64VarP(c.Offset, "offset", "", 0, "")
	c.RuntimeID = new(string)
	f.StringVarP(c.RuntimeID, "runtime_id", "", "", "")
	f.StringSliceVarP(&c.SubnetID, "subnet_id", "", []string{}, "")
	c.SubnetTypeValue = new(int64)
	f.Int64VarP(c.SubnetTypeValue, "subnet_type_value", "", 0, "The uint32 value.")
	f.StringSliceVarP(&c.Zone, "zone", "", []string{}, "")
}

func (c *DescribeSubnetsCmd) Run(out Out) error {

	out.WriteRequest(c.DescribeSubnetsParams)

	client := test.GetClient(clientConfig)
	res, err := client.ClusterManager.DescribeSubnets(c.DescribeSubnetsParams)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type DetachKeyPairsCmd struct {
	*models.OpenpitrixDetachKeyPairsRequest
}

func NewDetachKeyPairsCmd() Cmd {
	return &DetachKeyPairsCmd{
		&models.OpenpitrixDetachKeyPairsRequest{},
	}
}

func (*DetachKeyPairsCmd) GetActionName() string {
	return "DetachKeyPairs"
}

func (c *DetachKeyPairsCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.KeyPairID, "key_pair_id", "", []string{}, "")
	f.StringSliceVarP(&c.NodeID, "node_id", "", []string{}, "")
}

func (c *DetachKeyPairsCmd) Run(out Out) error {
	params := cluster_manager.NewDetachKeyPairsParams()
	params.WithBody(c.OpenpitrixDetachKeyPairsRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.ClusterManager.DetachKeyPairs(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type GetClusterStatisticsCmd struct {
	*cluster_manager.GetClusterStatisticsParams
}

func NewGetClusterStatisticsCmd() Cmd {
	return &GetClusterStatisticsCmd{
		cluster_manager.NewGetClusterStatisticsParams(),
	}
}

func (*GetClusterStatisticsCmd) GetActionName() string {
	return "GetClusterStatistics"
}

func (c *GetClusterStatisticsCmd) ParseFlag(f Flag) {
}

func (c *GetClusterStatisticsCmd) Run(out Out) error {

	out.WriteRequest(c.GetClusterStatisticsParams)

	client := test.GetClient(clientConfig)
	res, err := client.ClusterManager.GetClusterStatistics(c.GetClusterStatisticsParams)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type RecoverClustersCmd struct {
	*models.OpenpitrixRecoverClustersRequest
}

func NewRecoverClustersCmd() Cmd {
	return &RecoverClustersCmd{
		&models.OpenpitrixRecoverClustersRequest{},
	}
}

func (*RecoverClustersCmd) GetActionName() string {
	return "RecoverClusters"
}

func (c *RecoverClustersCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.AdvancedParam, "advanced_param", "", []string{}, "")
	f.StringSliceVarP(&c.ClusterID, "cluster_id", "", []string{}, "")
}

func (c *RecoverClustersCmd) Run(out Out) error {
	params := cluster_manager.NewRecoverClustersParams()
	params.WithBody(c.OpenpitrixRecoverClustersRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.ClusterManager.RecoverClusters(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type ResizeClusterCmd struct {
	*models.OpenpitrixResizeClusterRequest
}

func NewResizeClusterCmd() Cmd {
	return &ResizeClusterCmd{
		&models.OpenpitrixResizeClusterRequest{},
	}
}

func (*ResizeClusterCmd) GetActionName() string {
	return "ResizeCluster"
}

func (c *ResizeClusterCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.AdvancedParam, "advanced_param", "", []string{}, "")
	f.StringVarP(&c.ClusterID, "cluster_id", "", "", "")
	f.StringVarP(&c.Role, "role", "", "", "")
}

func (c *ResizeClusterCmd) Run(out Out) error {
	params := cluster_manager.NewResizeClusterParams()
	params.WithBody(c.OpenpitrixResizeClusterRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.ClusterManager.ResizeCluster(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type RollbackClusterCmd struct {
	*models.OpenpitrixRollbackClusterRequest
}

func NewRollbackClusterCmd() Cmd {
	return &RollbackClusterCmd{
		&models.OpenpitrixRollbackClusterRequest{},
	}
}

func (*RollbackClusterCmd) GetActionName() string {
	return "RollbackCluster"
}

func (c *RollbackClusterCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.AdvancedParam, "advanced_param", "", []string{}, "")
	f.StringVarP(&c.ClusterID, "cluster_id", "", "", "")
}

func (c *RollbackClusterCmd) Run(out Out) error {
	params := cluster_manager.NewRollbackClusterParams()
	params.WithBody(c.OpenpitrixRollbackClusterRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.ClusterManager.RollbackCluster(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type StartClustersCmd struct {
	*models.OpenpitrixStartClustersRequest
}

func NewStartClustersCmd() Cmd {
	return &StartClustersCmd{
		&models.OpenpitrixStartClustersRequest{},
	}
}

func (*StartClustersCmd) GetActionName() string {
	return "StartClusters"
}

func (c *StartClustersCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.AdvancedParam, "advanced_param", "", []string{}, "")
	f.StringSliceVarP(&c.ClusterID, "cluster_id", "", []string{}, "")
}

func (c *StartClustersCmd) Run(out Out) error {
	params := cluster_manager.NewStartClustersParams()
	params.WithBody(c.OpenpitrixStartClustersRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.ClusterManager.StartClusters(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type StopClustersCmd struct {
	*models.OpenpitrixStopClustersRequest
}

func NewStopClustersCmd() Cmd {
	return &StopClustersCmd{
		&models.OpenpitrixStopClustersRequest{},
	}
}

func (*StopClustersCmd) GetActionName() string {
	return "StopClusters"
}

func (c *StopClustersCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.AdvancedParam, "advanced_param", "", []string{}, "")
	f.StringSliceVarP(&c.ClusterID, "cluster_id", "", []string{}, "")
}

func (c *StopClustersCmd) Run(out Out) error {
	params := cluster_manager.NewStopClustersParams()
	params.WithBody(c.OpenpitrixStopClustersRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.ClusterManager.StopClusters(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type UpdateClusterEnvCmd struct {
	*models.OpenpitrixUpdateClusterEnvRequest
}

func NewUpdateClusterEnvCmd() Cmd {
	return &UpdateClusterEnvCmd{
		&models.OpenpitrixUpdateClusterEnvRequest{},
	}
}

func (*UpdateClusterEnvCmd) GetActionName() string {
	return "UpdateClusterEnv"
}

func (c *UpdateClusterEnvCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.AdvancedParam, "advanced_param", "", []string{}, "")
	f.StringVarP(&c.ClusterID, "cluster_id", "", "", "")
	f.StringVarP(&c.Env, "env", "", "", "")
}

func (c *UpdateClusterEnvCmd) Run(out Out) error {
	params := cluster_manager.NewUpdateClusterEnvParams()
	params.WithBody(c.OpenpitrixUpdateClusterEnvRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.ClusterManager.UpdateClusterEnv(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type UpgradeClusterCmd struct {
	*models.OpenpitrixUpgradeClusterRequest
}

func NewUpgradeClusterCmd() Cmd {
	return &UpgradeClusterCmd{
		&models.OpenpitrixUpgradeClusterRequest{},
	}
}

func (*UpgradeClusterCmd) GetActionName() string {
	return "UpgradeCluster"
}

func (c *UpgradeClusterCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.AdvancedParam, "advanced_param", "", []string{}, "")
	f.StringVarP(&c.ClusterID, "cluster_id", "", "", "")
	f.StringVarP(&c.VersionID, "version_id", "", "", "")
}

func (c *UpgradeClusterCmd) Run(out Out) error {
	params := cluster_manager.NewUpgradeClusterParams()
	params.WithBody(c.OpenpitrixUpgradeClusterRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.ClusterManager.UpgradeCluster(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type DescribeJobsCmd struct {
	*job_manager.DescribeJobsParams
}

func NewDescribeJobsCmd() Cmd {
	return &DescribeJobsCmd{
		job_manager.NewDescribeJobsParams(),
	}
}

func (*DescribeJobsCmd) GetActionName() string {
	return "DescribeJobs"
}

func (c *DescribeJobsCmd) ParseFlag(f Flag) {
	c.AppID = new(string)
	f.StringVarP(c.AppID, "app_id", "", "", "")
	c.ClusterID = new(string)
	f.StringVarP(c.ClusterID, "cluster_id", "", "", "")
	c.Executor = new(string)
	f.StringVarP(c.Executor, "executor", "", "", "")
	f.StringSliceVarP(&c.JobID, "job_id", "", []string{}, "")
	c.Limit = new(int64)
	f.Int64VarP(c.Limit, "limit", "", 20, "default is 20, max value is 200.")
	c.Offset = new(int64)
	f.Int64VarP(c.Offset, "offset", "", 0, "default is 0.")
	c.Provider = new(string)
	f.StringVarP(c.Provider, "provider", "", "", "")
	c.SearchWord = new(string)
	f.StringVarP(c.SearchWord, "search_word", "", "", "")
	f.StringSliceVarP(&c.Status, "status", "", []string{}, "")
	c.VersionID = new(string)
	f.StringVarP(c.VersionID, "version_id", "", "", "")
}

func (c *DescribeJobsCmd) Run(out Out) error {

	out.WriteRequest(c.DescribeJobsParams)

	client := test.GetClient(clientConfig)
	res, err := client.JobManager.DescribeJobs(c.DescribeJobsParams)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type DescribeRepoEventsCmd struct {
	*repo_indexer.DescribeRepoEventsParams
}

func NewDescribeRepoEventsCmd() Cmd {
	return &DescribeRepoEventsCmd{
		repo_indexer.NewDescribeRepoEventsParams(),
	}
}

func (*DescribeRepoEventsCmd) GetActionName() string {
	return "DescribeRepoEvents"
}

func (c *DescribeRepoEventsCmd) ParseFlag(f Flag) {
	c.Limit = new(int64)
	f.Int64VarP(c.Limit, "limit", "", 20, "")
	c.Offset = new(int64)
	f.Int64VarP(c.Offset, "offset", "", 0, "")
	f.StringSliceVarP(&c.Owner, "owner", "", []string{}, "")
	f.StringSliceVarP(&c.RepoEventID, "repo_event_id", "", []string{}, "")
	f.StringSliceVarP(&c.RepoID, "repo_id", "", []string{}, "")
	f.StringSliceVarP(&c.Status, "status", "", []string{}, "")
}

func (c *DescribeRepoEventsCmd) Run(out Out) error {

	out.WriteRequest(c.DescribeRepoEventsParams)

	client := test.GetClient(clientConfig)
	res, err := client.RepoIndexer.DescribeRepoEvents(c.DescribeRepoEventsParams)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type IndexRepoCmd struct {
	*models.OpenpitrixIndexRepoRequest
}

func NewIndexRepoCmd() Cmd {
	return &IndexRepoCmd{
		&models.OpenpitrixIndexRepoRequest{},
	}
}

func (*IndexRepoCmd) GetActionName() string {
	return "IndexRepo"
}

func (c *IndexRepoCmd) ParseFlag(f Flag) {
	f.StringVarP(&c.RepoID, "repo_id", "", "", "")
}

func (c *IndexRepoCmd) Run(out Out) error {
	params := repo_indexer.NewIndexRepoParams()
	params.WithBody(c.OpenpitrixIndexRepoRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.RepoIndexer.IndexRepo(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type CreateRepoCmd struct {
	*models.OpenpitrixCreateRepoRequest
}

func NewCreateRepoCmd() Cmd {
	return &CreateRepoCmd{
		&models.OpenpitrixCreateRepoRequest{},
	}
}

func (*CreateRepoCmd) GetActionName() string {
	return "CreateRepo"
}

func (c *CreateRepoCmd) ParseFlag(f Flag) {
	f.StringVarP(&c.CategoryID, "category_id", "", "", "")
	f.StringVarP(&c.Credential, "credential", "", "", "")
	f.StringVarP(&c.Description, "description", "", "", "")
	f.StringVarP(&c.Name, "name", "", "", "")
	f.StringSliceVarP(&c.Providers, "providers", "", []string{}, "")
	f.StringVarP(&c.Type, "type", "", "", "")
	f.StringVarP(&c.URL, "url", "", "", "")
	f.StringVarP(&c.Visibility, "visibility", "", "", "")
}

func (c *CreateRepoCmd) Run(out Out) error {
	params := repo_manager.NewCreateRepoParams()
	params.WithBody(c.OpenpitrixCreateRepoRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.RepoManager.CreateRepo(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type DeleteReposCmd struct {
	*models.OpenpitrixDeleteReposRequest
}

func NewDeleteReposCmd() Cmd {
	return &DeleteReposCmd{
		&models.OpenpitrixDeleteReposRequest{},
	}
}

func (*DeleteReposCmd) GetActionName() string {
	return "DeleteRepos"
}

func (c *DeleteReposCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.RepoID, "repo_id", "", []string{}, "")
}

func (c *DeleteReposCmd) Run(out Out) error {
	params := repo_manager.NewDeleteReposParams()
	params.WithBody(c.OpenpitrixDeleteReposRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.RepoManager.DeleteRepos(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type DescribeReposCmd struct {
	*repo_manager.DescribeReposParams
}

func NewDescribeReposCmd() Cmd {
	return &DescribeReposCmd{
		repo_manager.NewDescribeReposParams(),
	}
}

func (*DescribeReposCmd) GetActionName() string {
	return "DescribeRepos"
}

func (c *DescribeReposCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.CategoryID, "category_id", "", []string{}, "")
	c.Label = new(string)
	f.StringVarP(c.Label, "label", "", "", "")
	c.Limit = new(int64)
	f.Int64VarP(c.Limit, "limit", "", 20, "")
	f.StringSliceVarP(&c.Name, "name", "", []string{}, "")
	c.Offset = new(int64)
	f.Int64VarP(c.Offset, "offset", "", 0, "")
	f.StringSliceVarP(&c.Provider, "provider", "", []string{}, "")
	f.StringSliceVarP(&c.RepoID, "repo_id", "", []string{}, "")
	c.Reverse = new(bool)
	f.BoolVarP(c.Reverse, "reverse", "", false, "")
	c.SearchWord = new(string)
	f.StringVarP(c.SearchWord, "search_word", "", "", "")
	c.Selector = new(string)
	f.StringVarP(c.Selector, "selector", "", "", "")
	c.SortKey = new(string)
	f.StringVarP(c.SortKey, "sort_key", "", "", "")
	f.StringSliceVarP(&c.Status, "status", "", []string{}, "")
	f.StringSliceVarP(&c.Type, "type", "", []string{}, "")
	f.StringSliceVarP(&c.Visibility, "visibility", "", []string{}, "")
}

func (c *DescribeReposCmd) Run(out Out) error {

	out.WriteRequest(c.DescribeReposParams)

	client := test.GetClient(clientConfig)
	res, err := client.RepoManager.DescribeRepos(c.DescribeReposParams)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type ModifyRepoCmd struct {
	*models.OpenpitrixModifyRepoRequest
}

func NewModifyRepoCmd() Cmd {
	return &ModifyRepoCmd{
		&models.OpenpitrixModifyRepoRequest{},
	}
}

func (*ModifyRepoCmd) GetActionName() string {
	return "ModifyRepo"
}

func (c *ModifyRepoCmd) ParseFlag(f Flag) {
	f.StringVarP(&c.CategoryID, "category_id", "", "", "")
	f.StringVarP(&c.Credential, "credential", "", "", "")
	f.StringVarP(&c.Description, "description", "", "", "")
	f.StringVarP(&c.Name, "name", "", "", "")
	f.StringSliceVarP(&c.Providers, "providers", "", []string{}, "")
	f.StringVarP(&c.RepoID, "repo_id", "", "", "")
	f.StringVarP(&c.Type, "type", "", "", "")
	f.StringVarP(&c.URL, "url", "", "", "")
	f.StringVarP(&c.Visibility, "visibility", "", "", "")
}

func (c *ModifyRepoCmd) Run(out Out) error {
	params := repo_manager.NewModifyRepoParams()
	params.WithBody(c.OpenpitrixModifyRepoRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.RepoManager.ModifyRepo(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type ValidateRepoCmd struct {
	*repo_manager.ValidateRepoParams
}

func NewValidateRepoCmd() Cmd {
	return &ValidateRepoCmd{
		repo_manager.NewValidateRepoParams(),
	}
}

func (*ValidateRepoCmd) GetActionName() string {
	return "ValidateRepo"
}

func (c *ValidateRepoCmd) ParseFlag(f Flag) {
	c.Credential = new(string)
	f.StringVarP(c.Credential, "credential", "", "", "")
	c.Type = new(string)
	f.StringVarP(c.Type, "type", "", "", "")
	c.URL = new(string)
	f.StringVarP(c.URL, "url", "", "", "")
}

func (c *ValidateRepoCmd) Run(out Out) error {

	out.WriteRequest(c.ValidateRepoParams)

	client := test.GetClient(clientConfig)
	res, err := client.RepoManager.ValidateRepo(c.ValidateRepoParams)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type CreateRuntimeCmd struct {
	*models.OpenpitrixCreateRuntimeRequest
}

func NewCreateRuntimeCmd() Cmd {
	return &CreateRuntimeCmd{
		&models.OpenpitrixCreateRuntimeRequest{},
	}
}

func (*CreateRuntimeCmd) GetActionName() string {
	return "CreateRuntime"
}

func (c *CreateRuntimeCmd) ParseFlag(f Flag) {
	f.StringVarP(&c.Description, "description", "", "", "")
	f.StringVarP(&c.Labels, "labels", "", "", "")
	f.StringVarP(&c.Name, "name", "", "", "")
	f.StringVarP(&c.Provider, "provider", "", "", "")
	f.StringVarP(&c.RuntimeCredential, "runtime_credential", "", "", "")
	f.StringVarP(&c.RuntimeURL, "runtime_url", "", "", "")
	f.StringVarP(&c.Zone, "zone", "", "", "")
}

func (c *CreateRuntimeCmd) Run(out Out) error {
	params := runtime_manager.NewCreateRuntimeParams()
	params.WithBody(c.OpenpitrixCreateRuntimeRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.RuntimeManager.CreateRuntime(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type DeleteRuntimesCmd struct {
	*models.OpenpitrixDeleteRuntimesRequest
}

func NewDeleteRuntimesCmd() Cmd {
	return &DeleteRuntimesCmd{
		&models.OpenpitrixDeleteRuntimesRequest{},
	}
}

func (*DeleteRuntimesCmd) GetActionName() string {
	return "DeleteRuntimes"
}

func (c *DeleteRuntimesCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.RuntimeID, "runtime_id", "", []string{}, "")
}

func (c *DeleteRuntimesCmd) Run(out Out) error {
	params := runtime_manager.NewDeleteRuntimesParams()
	params.WithBody(c.OpenpitrixDeleteRuntimesRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.RuntimeManager.DeleteRuntimes(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type DescribeRuntimeProviderZonesCmd struct {
	*runtime_manager.DescribeRuntimeProviderZonesParams
}

func NewDescribeRuntimeProviderZonesCmd() Cmd {
	return &DescribeRuntimeProviderZonesCmd{
		runtime_manager.NewDescribeRuntimeProviderZonesParams(),
	}
}

func (*DescribeRuntimeProviderZonesCmd) GetActionName() string {
	return "DescribeRuntimeProviderZones"
}

func (c *DescribeRuntimeProviderZonesCmd) ParseFlag(f Flag) {
	c.Provider = new(string)
	f.StringVarP(c.Provider, "provider", "", "", "")
	c.RuntimeCredential = new(string)
	f.StringVarP(c.RuntimeCredential, "runtime_credential", "", "", "")
	c.RuntimeURL = new(string)
	f.StringVarP(c.RuntimeURL, "runtime_url", "", "", "")
}

func (c *DescribeRuntimeProviderZonesCmd) Run(out Out) error {

	out.WriteRequest(c.DescribeRuntimeProviderZonesParams)

	client := test.GetClient(clientConfig)
	res, err := client.RuntimeManager.DescribeRuntimeProviderZones(c.DescribeRuntimeProviderZonesParams)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type DescribeRuntimesCmd struct {
	*runtime_manager.DescribeRuntimesParams
}

func NewDescribeRuntimesCmd() Cmd {
	return &DescribeRuntimesCmd{
		runtime_manager.NewDescribeRuntimesParams(),
	}
}

func (*DescribeRuntimesCmd) GetActionName() string {
	return "DescribeRuntimes"
}

func (c *DescribeRuntimesCmd) ParseFlag(f Flag) {
	c.Label = new(string)
	f.StringVarP(c.Label, "label", "", "", "")
	c.Limit = new(int64)
	f.Int64VarP(c.Limit, "limit", "", 20, "")
	c.Offset = new(int64)
	f.Int64VarP(c.Offset, "offset", "", 0, "")
	f.StringSliceVarP(&c.Owner, "owner", "", []string{}, "")
	f.StringSliceVarP(&c.Provider, "provider", "", []string{}, "")
	f.StringSliceVarP(&c.RuntimeID, "runtime_id", "", []string{}, "")
	c.SearchWord = new(string)
	f.StringVarP(c.SearchWord, "search_word", "", "", "")
	f.StringSliceVarP(&c.Status, "status", "", []string{}, "")
}

func (c *DescribeRuntimesCmd) Run(out Out) error {

	out.WriteRequest(c.DescribeRuntimesParams)

	client := test.GetClient(clientConfig)
	res, err := client.RuntimeManager.DescribeRuntimes(c.DescribeRuntimesParams)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type GetRuntimeStatisticsCmd struct {
	*runtime_manager.GetRuntimeStatisticsParams
}

func NewGetRuntimeStatisticsCmd() Cmd {
	return &GetRuntimeStatisticsCmd{
		runtime_manager.NewGetRuntimeStatisticsParams(),
	}
}

func (*GetRuntimeStatisticsCmd) GetActionName() string {
	return "GetRuntimeStatistics"
}

func (c *GetRuntimeStatisticsCmd) ParseFlag(f Flag) {
}

func (c *GetRuntimeStatisticsCmd) Run(out Out) error {

	out.WriteRequest(c.GetRuntimeStatisticsParams)

	client := test.GetClient(clientConfig)
	res, err := client.RuntimeManager.GetRuntimeStatistics(c.GetRuntimeStatisticsParams)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type ModifyRuntimeCmd struct {
	*models.OpenpitrixModifyRuntimeRequest
}

func NewModifyRuntimeCmd() Cmd {
	return &ModifyRuntimeCmd{
		&models.OpenpitrixModifyRuntimeRequest{},
	}
}

func (*ModifyRuntimeCmd) GetActionName() string {
	return "ModifyRuntime"
}

func (c *ModifyRuntimeCmd) ParseFlag(f Flag) {
	f.StringVarP(&c.Description, "description", "", "", "")
	f.StringVarP(&c.Labels, "labels", "", "", "")
	f.StringVarP(&c.Name, "name", "", "", "")
	f.StringVarP(&c.RuntimeID, "runtime_id", "", "", "")
}

func (c *ModifyRuntimeCmd) Run(out Out) error {
	params := runtime_manager.NewModifyRuntimeParams()
	params.WithBody(c.OpenpitrixModifyRuntimeRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.RuntimeManager.ModifyRuntime(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type DescribeTasksCmd struct {
	*task_manager.DescribeTasksParams
}

func NewDescribeTasksCmd() Cmd {
	return &DescribeTasksCmd{
		task_manager.NewDescribeTasksParams(),
	}
}

func (*DescribeTasksCmd) GetActionName() string {
	return "DescribeTasks"
}

func (c *DescribeTasksCmd) ParseFlag(f Flag) {
	c.Executor = new(string)
	f.StringVarP(c.Executor, "executor", "", "", "")
	f.StringSliceVarP(&c.JobID, "job_id", "", []string{}, "")
	c.Limit = new(int64)
	f.Int64VarP(c.Limit, "limit", "", 20, "default is 20, max value is 200.")
	c.Offset = new(int64)
	f.Int64VarP(c.Offset, "offset", "", 0, "default is 0.")
	c.SearchWord = new(string)
	f.StringVarP(c.SearchWord, "search_word", "", "", "")
	f.StringSliceVarP(&c.Status, "status", "", []string{}, "")
	c.Target = new(string)
	f.StringVarP(c.Target, "target", "", "", "")
	f.StringSliceVarP(&c.TaskID, "task_id", "", []string{}, "")
}

func (c *DescribeTasksCmd) Run(out Out) error {

	out.WriteRequest(c.DescribeTasksParams)

	client := test.GetClient(clientConfig)
	res, err := client.TaskManager.DescribeTasks(c.DescribeTasksParams)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}

type RetryTasksCmd struct {
	*models.OpenpitrixRetryTasksRequest
}

func NewRetryTasksCmd() Cmd {
	return &RetryTasksCmd{
		&models.OpenpitrixRetryTasksRequest{},
	}
}

func (*RetryTasksCmd) GetActionName() string {
	return "RetryTasks"
}

func (c *RetryTasksCmd) ParseFlag(f Flag) {
	f.StringSliceVarP(&c.TaskID, "task_id", "", []string{}, "")
}

func (c *RetryTasksCmd) Run(out Out) error {
	params := task_manager.NewRetryTasksParams()
	params.WithBody(c.OpenpitrixRetryTasksRequest)

	out.WriteRequest(params)

	client := test.GetClient(clientConfig)
	res, err := client.TaskManager.RetryTasks(params)
	if err != nil {
		return err
	}

	out.WriteResponse(res.Payload)

	return nil
}
