module hcc/model

go 1.15

replace (
	innogrid.com/hcloud-classic/easygoorm => ../easygoorm
	innogrid.com/hcloud-classic/hcc_errors => ../hcc_errors
	innogrid.com/hcloud-classic/pb => ../pb
)

require innogrid.com/hcloud-classic/hcc_errors v0.0.0
