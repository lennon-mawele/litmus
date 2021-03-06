// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type ActionPayload struct {
	RequestType  *string `json:"request_type"`
	K8sManifest  *string `json:"k8s_manifest"`
	Namespace    *string `json:"namespace"`
	ExternalData *string `json:"external_data"`
}

type ChaosWorkFlowInput struct {
	WorkflowManifest    string             `json:"workflow_manifest"`
	CronSyntax          string             `json:"cronSyntax"`
	WorkflowName        string             `json:"workflow_name"`
	WorkflowDescription string             `json:"workflow_description"`
	Weightages          []*WeightagesInput `json:"weightages"`
	IsCustomWorkflow    bool               `json:"isCustomWorkflow"`
	ProjectID           string             `json:"project_id"`
	ClusterID           string             `json:"cluster_id"`
}

type ChaosWorkFlowResponse struct {
	WorkflowID          string `json:"workflow_id"`
	CronSyntax          string `json:"cronSyntax"`
	WorkflowName        string `json:"workflow_name"`
	WorkflowDescription string `json:"workflow_description"`
	IsCustomWorkflow    bool   `json:"isCustomWorkflow"`
}

type Cluster struct {
	ClusterID          string  `json:"cluster_id"`
	ProjectID          string  `json:"project_id"`
	ClusterName        string  `json:"cluster_name"`
	Description        *string `json:"description"`
	PlatformName       string  `json:"platform_name"`
	AccessKey          string  `json:"access_key"`
	IsRegistered       bool    `json:"is_registered"`
	IsClusterConfirmed bool    `json:"is_cluster_confirmed"`
	IsActive           bool    `json:"is_active"`
	UpdatedAt          string  `json:"updated_at"`
	CreatedAt          string  `json:"created_at"`
	ClusterType        string  `json:"cluster_type"`
}

type ClusterAction struct {
	ProjectID string         `json:"project_id"`
	Action    *ActionPayload `json:"action"`
}

type ClusterActionInput struct {
	ClusterID string `json:"cluster_id"`
	Action    string `json:"action"`
}

type ClusterConfirmResponse struct {
	IsClusterConfirmed bool    `json:"isClusterConfirmed"`
	NewClusterKey      *string `json:"newClusterKey"`
	ClusterID          *string `json:"cluster_id"`
}

type ClusterEvent struct {
	EventID     string   `json:"event_id"`
	EventType   string   `json:"event_type"`
	EventName   string   `json:"event_name"`
	Description string   `json:"description"`
	Cluster     *Cluster `json:"cluster"`
}

type ClusterEventInput struct {
	EventName   string `json:"event_name"`
	Description string `json:"description"`
	ClusterID   string `json:"cluster_id"`
	AccessKey   string `json:"access_key"`
}

type ClusterIdentity struct {
	ClusterID string `json:"cluster_id"`
	AccessKey string `json:"access_key"`
}

type ClusterInput struct {
	ClusterName  string  `json:"cluster_name"`
	Description  *string `json:"description"`
	PlatformName string  `json:"platform_name"`
	ProjectID    string  `json:"project_id"`
	ClusterType  string  `json:"cluster_type"`
}

type Member struct {
	UserID     string `json:"user_id"`
	UserName   string `json:"user_name"`
	Role       string `json:"role"`
	Invitation string `json:"invitation"`
}

type PodLog struct {
	ClusterID     *ClusterIdentity `json:"cluster_id"`
	RequestID     string           `json:"request_id"`
	WorkflowRunID string           `json:"workflow_run_id"`
	PodName       string           `json:"pod_name"`
	PodType       string           `json:"pod_type"`
	Log           string           `json:"log"`
}

type PodLogRequest struct {
	ClusterID      string  `json:"cluster_id"`
	WorkflowRunID  string  `json:"workflow_run_id"`
	PodName        string  `json:"pod_name"`
	PodNamespace   string  `json:"pod_namespace"`
	PodType        string  `json:"pod_type"`
	ExpPod         *string `json:"exp_pod"`
	RunnerPod      *string `json:"runner_pod"`
	ChaosNamespace *string `json:"chaos_namespace"`
}

type PodLogResponse struct {
	WorkflowRunID string `json:"workflow_run_id"`
	PodName       string `json:"pod_name"`
	PodType       string `json:"pod_type"`
	Log           string `json:"log"`
}

type Project struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Members   []*Member `json:"members"`
	State     *string   `json:"state"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	RemovedAt string    `json:"removed_at"`
}

type User struct {
	ID              string     `json:"id"`
	Username        string     `json:"username"`
	Email           *string    `json:"email"`
	IsEmailVerified *bool      `json:"is_email_verified"`
	CompanyName     *string    `json:"company_name"`
	Name            *string    `json:"name"`
	Project         []*Project `json:"project"`
	Role            *string    `json:"role"`
	State           *string    `json:"state"`
	CreatedAt       string     `json:"created_at"`
	UpdatedAt       string     `json:"updated_at"`
	RemovedAt       string     `json:"removed_at"`
}

type UserInput struct {
	Username    string  `json:"username"`
	Email       *string `json:"email"`
	CompanyName *string `json:"company_name"`
	Name        *string `json:"name"`
	ProjectName string  `json:"project_name"`
}

type WeightagesInput struct {
	ExperimentName string `json:"experiment_name"`
	Weightage      int    `json:"weightage"`
}

type WorkflowRun struct {
	WorkflowRunID string `json:"workflow_run_id"`
	WorkflowID    string `json:"workflow_id"`
	ClusterName   string `json:"cluster_name"`
	LastUpdated   string `json:"last_updated"`
	ProjectID     string `json:"project_id"`
	ClusterID     string `json:"cluster_id"`
	WorkflowName  string `json:"workflow_name"`
	ExecutionData string `json:"execution_data"`
}

type WorkflowRunInput struct {
	WorkflowRunID string           `json:"workflow_run_id"`
	WorkflowName  string           `json:"workflow_name"`
	ExecutionData string           `json:"execution_data"`
	ClusterID     *ClusterIdentity `json:"cluster_id"`
}

type Weightages struct {
	ExperimentName string `json:"experiment_name"`
	Weightage      int    `json:"weightage"`
}
