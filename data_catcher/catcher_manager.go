package data_catcher

import(
	"sync"
	"fmt"
)

var (
	_g_DataCatcherMng *FinDataCatcherMng
)

func init() {
	_g_DataCatcherMng = new(FinDataCatcherMng)
}	

func GetGlobalFinDataCatcherMng() *FinDataCatcherMng {
	return _g_DataCatcherMng
}

/////
type FinDataCatcherMng struct {
	_catcher_queue []*FinDataHttpCatcher
	_catcher_queue_lock sync.Mutex

	_jobs_ctx_queue []*FinDataCatcherJobContext
	_jobs_ctx_queue_lock     sync.Mutex
}

func (d *FinDataCatcherMng) Init(num int) {
	d._catcher_queue = make([]*FinDataHttpCatcher, 0)
	d._jobs_ctx_queue = make([]*FinDataCatcherJobContext, 0)

	// 
	catcher := CreateFinDataHttpCatcher()
	d.AddHttpCatcher(catcher)
	go catcher.Run()
}

func (d *FinDataCatcherMng) AddJobContext(ctx *FinDataCatcherJobContext) {
	defer d._jobs_ctx_queue_lock.Unlock()

	d._jobs_ctx_queue_lock.Lock()
	d._jobs_ctx_queue = append(d._jobs_ctx_queue, ctx)	
}

func (d *FinDataCatcherMng) AddHttpCatcher(c *FinDataHttpCatcher) {
	defer d._catcher_queue_lock.Unlock()

	d._catcher_queue_lock.Lock()
	d._catcher_queue = append(d._catcher_queue, c)
}

func (d *FinDataCatcherMng) Dispatch(ctx *FinDataCatcherJobContext) bool {
	c := d.PopCatcher()
	if nil == c {
		fmt.Printf("Catcher is nil\n")
		return false
	}

	c.AssignJobContext(ctx)
	return true
}

func (d *FinDataCatcherMng) PopCatcher() *FinDataHttpCatcher {
	var c *FinDataHttpCatcher
	defer d._catcher_queue_lock.Unlock()
	
	d._catcher_queue_lock.Lock()
	if len(d._catcher_queue) != 0 {
		c = d._catcher_queue[0]
		d._catcher_queue = d._catcher_queue[1:]
	}

	return c
}

func (d *FinDataCatcherMng) GetJobContextNum() int {
	defer d._jobs_ctx_queue_lock.Unlock()

	d._jobs_ctx_queue_lock.Lock()
	return len(d._jobs_ctx_queue)
}

func (d *FinDataCatcherMng) GetHttpCatcherNum() int {
	defer d._catcher_queue_lock.Unlock()

	d._catcher_queue_lock.Lock()
	return len(d._catcher_queue)
}

func (d *FinDataCatcherMng) PopJobContext(num int) []*FinDataCatcherJobContext {
	defer d._jobs_ctx_queue_lock.Unlock()
		
	ret := make([]*FinDataCatcherJobContext, 0)
	d._jobs_ctx_queue_lock.Lock()
	for i:=0; i < num && i < len(d._jobs_ctx_queue); i++ {
		ctx := d._jobs_ctx_queue[0]
		ret = append(ret, ctx)
		d._jobs_ctx_queue = d._jobs_ctx_queue[1:]
	}

	return ret
}
