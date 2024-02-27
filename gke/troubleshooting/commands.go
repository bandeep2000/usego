package troubleshooting

// Add this also to check for cert
/*

cat /var/lib/kubelet/config.yaml
apiVersion: kubelet.config.k8s.io/v1beta1
authentication:
  anonymous:
    enabled: false
  webhook:
*/

func GetNodeFailureCmds() []string {

	str1 := []string{
		"kubelet get nodes",
		"k describe node <nodename>",
		"top",
		"df -h", // Disk issues
		"service kubelet status",
		"journalctl -u kubelet",
		"openssl x509 -in /var/lib/kbulente/worker-1.crt -text",cd j
		"cat /etc/kubernetes/kubelet.conf",
		"cat /var/lib/kubelet/config.yaml", //check for ports also
	}

	return str1
}

func CoreDNSFailures() []string{

	/*

	k describe configmap coredns -n kube-system
Name:         coredns
Namespace:    kube-system
Labels:       <none>
Annotations:  <none>

Data
====
Corefile:
----
.:53 {
    errors
    health {
       lameduck 5s
    }
    ready
    kubernetes cluster.local in-addr.arpa ip6.arpa {
       pods insecure
       fallthrough in-addr.arpa ip6.arpa
       ttl 30
    }
    prometheus :9153
    forward . /etc/resolv.conf {
       max_concurrent 1000
    }
    cache 30


	*/

	// https://kubernetes.io/docs/tasks/administer-cluster/dns-debugging-resolution/
	comm_notes := []string {
		"k get clusterroles | grep dns"
		"k get clusterrolebinding | grep dns"
		"k describe configmap coredns -n kube-system"
		"kubectl get pods --namespace=kube-system -l k8s-app=kube-dns"
		"kubectl logs --namespace=kube-system -l k8s-app=kube-dns"
		"kubectl get svc --namespace=kube-system"
		"kubectl get endpoints kube-dns --namespace=kube-system" //Check end points
	}

	return comm_notes
}
