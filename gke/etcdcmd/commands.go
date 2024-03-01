package etcdcmd

func GetCmds() []string {

	cmds := []string{
		"k get all --all-namespaces -o yaml > all-service.yaml # back up everything",
		"export ETCDCTL_API=3",
		"take back up of data directory # one way to back up",
		"ETCDCTL_API=3 etcdctl snapshot save snap.db",
		"ETCDCTL_API=3 etcdctl snapshot status snap.db",
		"stop api server to restore, as api server depends on it",
		"ETCDCTL_API=3 etcdctl snapshot restore snap.db --data-dir /var/example # etcd will shutdown cluster to prevent anyone joining",
		"Updated the etcdctl server start parameters with new data dir",
	}

	return cmds
}
