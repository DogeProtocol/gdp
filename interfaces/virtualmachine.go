package interfaces

type VirtualMachineInput interface {
}

type VirtualMachineOutput interface {
}

type VirtualMachineCreateContext interface {
}

type VirtualMachine interface {
	New(context VirtualMachineCreateContext)
	Call(caller Contract, contractAddress Address, input VirtualMachineInput) (output VirtualMachineOutput, err error)
}
