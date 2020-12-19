.PHONY: daemon_start
daemon_start:
	vw -i train-model/vw.model --daemon --port 58000 --num_children 1 --oaa 3 --loss_function=logistic

.PHONY: daemon_stop
daemon_stop:
	pkill -9 -f `vw.*--port 58000`