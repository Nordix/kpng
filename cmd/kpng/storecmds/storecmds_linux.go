package storecmds

import (
	_ "sigs.k8s.io/kpng/backends/ipvs-as-sink"
	_ "sigs.k8s.io/kpng/backends/ipvs-as-sink2"
	_ "sigs.k8s.io/kpng/backends/nft"
	_ "sigs.k8s.io/kpng/backends/iptables"
)
