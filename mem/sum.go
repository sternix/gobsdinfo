package mem

import "github.com/sternix/gobsdinfo/util"

type Summary struct {
	Version    string `json:"__version"`
	Statistics struct {
		ContextSwitches   int64 `json:"context-switches"`
		Interrupts        int64 `json:"interrupts"`
		SoftwareInterrups int64 `json:"software-interrupts"`
		Traps             int64 `json:"traps"`
		SystemCalls       int64 `json:"system-calls"`
		KernelThreads     int64 `json:"kernel-threads"`
		Forks             int64 `json:"forks"`
		Vforks            int64 `json:"vforks"`
		Rforks            int64 `json:"rforks"`
		SwapIns           int64 `json:"swap-ins"`
		SwapInpages       int64 `json:"swap-in-pages"`
		SwapOuts          int64 `json:"swap-outs"`
		SwapOutPages      int64 `json:"swap-out-pages"`
		VnodePageIns      int64 `json:"vnode-page-ins"`
		VnodePageInPages  int64 `json:"vnode-page-in-pages"`
		VnodePageOuts     int64 `json:"vnode-page-outs"`
		//TODO: dublicate key, might be vnode-page-out-pages, report this to the FreeBSD
		VnodePageOutPages          int64 `json:"vnode-page-outs"`
		PageDaemonWakeups          int64 `json:"page-daemon-wakeups"`
		PageDaemonPages            int64 `json:"page-daemon-pages"`
		PageReclamationShortfalls  int64 `json:"page-reclamation-shortfalls"`
		Reactivated                int64 `json:"reactivated"`
		CopyOnWriteFaults          int64 `json:"copy-on-write-faults"`
		CopyOnWriteOptimizedFaults int64 `json:"copy-on-write-optimized-faults"`
		ZeroFillPages              int64 `json:"zero-fill-pages"`
		ZeroFillPrezeroed          int64 `json:"zero-fill-prezeroed"`
		IntransitBlocking          int64 `json:"intransit-blocking"`
		TotalFaults                int64 `json:"total-faults"`
		FaultsRequiringIo          int64 `json:"faults-requiring-io"`
		FaultsFromThreadCreation   int64 `json:"faults-from-thread-creation"`
		FaultsFromFork             int64 `json:"faults-from-fork"`
		FaultsFromVfork            int64 `json:"faults-from-vfork"`
		PagesRfork                 int64 `json:"pages-rfork"`
		PagesFreed                 int64 `json:"pages-freed"`
		PagesFreedByDaemon         int64 `json:"pages-freed-by-daemon"`
		PagesFreedOnExit           int64 `json:"pages-freed-on-exit"`
		ActivePages                int64 `json:"active-pages"`
		InactivePages              int64 `json:"inactive-pages"`
		LaundryPages               int64 `json:"laundry-pages"`
		WiredPages                 int64 `json:"wired-pages"`
		FreePages                  int64 `json:"free-pages"`
		BytesPerPage               int64 `json:"bytes-per-page"`
		TotalNameLookups           int64 `json:"total-name-lookups"`
		PositiveCacheHits          int64 `json:"positive-cache-hits"`
		NegativeCacheHits          int64 `json:"negative-cache-hits"`
		CacheHitPercent            int64 `json:"cache-hit-percent"`
		Deletions                  int64 `json:"deletions"`
		FalseHits                  int64 `json:"false-hits"`
		TooLong                    int64 `json:"too-long"`
	} `json:"summary-statistics"`
}

func SumStats() (Summary, error) {
	var sum Summary
	if err := util.ExecAndParse(&sum, "/usr/bin/vmstat", "-s"); err != nil {
		return sum, err
	}

	return sum, nil
}
