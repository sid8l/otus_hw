package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func checkingDoneStage(in In, done In) Out {
	out := make(Bi)
	go func() {
		defer close(out)
		for {
			select {
			case v, ok := <-in:
				if !ok {
					return
				}
				out <- v
			case <-done:
				return
			}
		}
	}()
	return out
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := make(Out)

	if len(stages) == 0 {
		out = checkingDoneStage(in, done)
	}

	for _, s := range stages {
		out = s(checkingDoneStage(in, done))
		in = out
	}
	return out
}
