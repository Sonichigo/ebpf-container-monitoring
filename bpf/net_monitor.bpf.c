#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>

struct packet_data {
    __u64 packet_count;
    __u64 total_bytes;
};

struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __uint(max_entries, 1024);
    __type(key, __u64);
    __type(value, struct packet_data);
} packet_stats SEC(".maps");

SEC("tracepoint/net/net_dev_queue")
int trace_packet(struct trace_event_raw_net_dev_template *ctx) {
    __u64 pid = bpf_get_current_pid_tgid();
    struct packet_data *data = bpf_map_lookup_elem(&packet_stats, &pid);

    if (data) {
        __sync_fetch_and_add(&data->packet_count, 1);
        __sync_fetch_and_add(&data->total_bytes, ctx->len);
    } else {
        struct packet_data init = { .packet_count = 1, .total_bytes = ctx->len };
        bpf_map_update_elem(&packet_stats, &pid, &init, BPF_ANY);
    }
    return 0;
}

char LICENSE[] SEC("license") = "GPL";