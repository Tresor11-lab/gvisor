// automatically generated by stateify.

// +build arm64

package arch

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (r *Registers) StateTypeName() string {
	return "pkg/sentry/arch.Registers"
}

func (r *Registers) StateFields() []string {
	return []string{
		"PtraceRegs",
		"TPIDR_EL0",
	}
}

func (r *Registers) beforeSave() {}

// +checklocksignore
func (r *Registers) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.PtraceRegs)
	stateSinkObject.Save(1, &r.TPIDR_EL0)
}

func (r *Registers) afterLoad() {}

// +checklocksignore
func (r *Registers) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.PtraceRegs)
	stateSourceObject.Load(1, &r.TPIDR_EL0)
}

func (s *State) StateTypeName() string {
	return "pkg/sentry/arch.State"
}

func (s *State) StateFields() []string {
	return []string{
		"Regs",
		"aarch64FPState",
		"FeatureSet",
		"OrigR0",
	}
}

func (s *State) beforeSave() {}

// +checklocksignore
func (s *State) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.Regs)
	stateSinkObject.Save(1, &s.aarch64FPState)
	stateSinkObject.Save(2, &s.FeatureSet)
	stateSinkObject.Save(3, &s.OrigR0)
}

func (s *State) afterLoad() {}

// +checklocksignore
func (s *State) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.Regs)
	stateSourceObject.LoadWait(1, &s.aarch64FPState)
	stateSourceObject.Load(2, &s.FeatureSet)
	stateSourceObject.Load(3, &s.OrigR0)
}

func init() {
	state.Register((*Registers)(nil))
	state.Register((*State)(nil))
}
