package bpfloader

import (
	"log"
	"github.com/cilium/ebpf"
)

func LoadBPFProgram(path string) *ebpf.Program {
	spec, err := ebpf.LoadCollectionSpec(path)
	if err != nil {
		log.Fatalf("Failed to load BPF spec: %v", err)
	}

	coll, err := ebpf.NewCollection(spec)
	if err != nil {
		log.Fatalf("Failed to load BPF program: %v", err)
	}

	return coll.Programs["trace_packet"]
}