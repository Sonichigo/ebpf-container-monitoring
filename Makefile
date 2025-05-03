# Define architecture and kernel versions
ARCH := $(shell uname -m)
KERNEL_VERSION := $(shell uname -r)

# Define include paths
KERNEL_INCLUDE := /usr/src/linux-headers-$(KERNEL_VERSION)/include
ARCH_INCLUDE := /usr/src/linux-headers-$(KERNEL_VERSION)/arch/$(subst x86_64,x86,$(ARCH))/include
LIBC_INCLUDE := /usr/include/$(ARCH)-linux-gnu
INC_FLAGS := -I$(KERNEL_INCLUDE) -I$(ARCH_INCLUDE) -I$(LIBC_INCLUDE)

# Define preprocessor flags
CFLAGS := -O2 -g -Wall
BPF_CFLAGS := $(CFLAGS) -target bpf $(INC_FLAGS) -D__KERNEL__ -D__BPF_TRACING__

# Define targets
BPF_TARGETS = bpf/net_monitor.bpf.o bpf/tcp_rtt.bpf.o

# Default target
all: $(BPF_TARGETS)

# Rule to build BPF objects
bpf/%.bpf.o: bpf/%.bpf.c
	clang $(BPF_CFLAGS) -c $< -o $@

# Clean target
clean:
	rm -f $(BPF_TARGETS)

.PHONY: all clean