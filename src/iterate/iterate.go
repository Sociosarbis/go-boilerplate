package iterate

import (
	"context"
	"sync"
)

// reference from https://sergey.kamardin.org/articles/generic-concurrency-in-go/

func iterate[A, B any](
	ctx context.Context,
	prepare func(context.Context) (context.Context, context.CancelFunc),
	parallelism int,
	pull func() (A, bool),
	f func(context.Context, A) B,
	push func(A, B) bool) (err error) {
	subctx, cancel := context.WithCancel(ctx)
	defer cancel()

	/** 可以在函数内定义结构 */
	type result struct {
		a A
		b B
	}

	type loopInfo struct {
		dispatched int
		err        error
	}

	var (
		res  = make(chan result)
		term = make(chan loopInfo, 1)
	)

	var wg sync.WaitGroup

	go func() {
		wrk := make(chan A)
		var loop loopInfo
		defer func() {
			close(wrk)
			term <- loop
		}()
		var workersCount int
		sem := make(chan struct{})

		/** close可以让chan在每次select的时候可以触发 */
		close(sem)

		for {
			if err := subctx.Err(); err != nil {
				loop.err = err
				return
			}
			a, ok := pull()
			if !ok {
				return
			}
			if parallelism != 0 && workersCount == parallelism {
				/** 设为nil以后不会再触发select */
				sem = nil
			}
			select {
			case <-subctx.Done():
				loop.err = ctx.Err()
				return
			/** 如果wrk是buffer channel，在消息超过上限时发送消息会block住 */
			case wrk <- a:
				loop.dispatched++
				continue
			case <-sem:
				loop.dispatched++
			}
			workersCount++
			wg.Add(1)
			go (func(a A) {
				defer wg.Done()
				subctx := subctx
				if prepare != nil {
					var cancel context.CancelFunc
					subctx, cancel = prepare(subctx)
					defer cancel()
				}
				for {
					r := result{a: a}
					r.b = f(subctx, a)
					select {
					case res <- r:
					case <-subctx.Done():
						return
					}
					var ok bool
					a, ok = <-wrk
					if !ok {
						break
					}
				}
			})(a)
		}
	}()

collect:
	for i, num := 0, -1; num == -1 || i < num; {
		select {
		case <-ctx.Done():
			if err == nil {
				err = ctx.Err()
			}
			break collect
		case res := <-res:
			if !push(res.a, res.b) {
				cancel()
				break collect
			}
			i++
		case loop := <-term:
			num = loop.dispatched
			err = loop.err
		}
	}

	wg.Wait()
	return err
}
