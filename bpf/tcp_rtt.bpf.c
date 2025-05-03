#include <linux/bpf.h>
#include <linux/tcp.h>
#include <linux/ip.h>
#include <linux/ptrace.h>
#include <bpf/bpf_helpers.h>

struct rtt_info {
    __u64 srtt_us;
};

struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __uint(max_entries, 1024);
    __type(key, __u64);
    __type(value, struct rtt_info);
} rtt_stats SEC(".maps");

SEC("kprobe/tcp_ack")
int trace_tcp_ack(struct pt_regs *ctx) {
    struct sock *sk = (struct sock *)PT_REGS_PARM1(ctx);
    if (!sk) return 0;

    struct tcp_sock *tp = (struct tcp_sock *)sk;
    __u64 pid = bpf_get_current_pid_tgid();
    struct rtt_info info = {};

    info.srtt_us = tp->srtt_us >> 3;
    bpf_map_update_elem(&rtt_stats, &pid, &info, BPF_ANY);

    return 0;
}

char LICENSE[] SEC("license") = "GPL";