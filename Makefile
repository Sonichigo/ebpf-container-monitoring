BPF_TARGETS = bpf/net_monitor.bpf.o bpf/tcp_rtt.bpf.o

all: $(BPF_TARGETS)

bpf/%.bpf.o: bpf/%.bpf.c
	clang -O2 -g -Wall -target bpf -c $< -o $@