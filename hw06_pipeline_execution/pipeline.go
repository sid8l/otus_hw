package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func DoneChecker(done In, in In) Out {
	out := make(Bi)
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				out <- v
			}
		}
	}()
	return out
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	interim := in
	outStage := make(Out)
	for _, stage := range stages {
		outStage = stage(interim)
		interim = outStage
	}
	outPrime := DoneChecker(done, outStage)
	return outPrime
}
