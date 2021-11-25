package model

import "innogrid.com/hcloud-classic/hcc_errors"

// Quota : Contain infos of the quota
type Quota struct {
	GroupID            int64                    `json:"group_id"`
	GroupName          string                   `json:"group_name"`
	TotalCPUCores      int                      `json:"total_cpu_cores"`
	TotalMemoryGB      int                      `json:"total_memory_gb"`
	LimitSubnetCnt     int                      `json:"limit_subnet_cnt"`
	LimitAdaptiveIPCnt int                      `json:"limit_adaptive_ip_cnt"`
	LimitNodeCnt       int                      `json:"limit_node_cnt"`
	PoolName           string                   `json:"pool_name"`
	LimitSSDGB         int                      `json:"limit_ssd_gb"`
	LimitHDDGB         int                      `json:"limit_hdd_gb"`
	Errors             hcc_errors.HccErrorStack `json:"errors"`
}
