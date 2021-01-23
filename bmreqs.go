package bmreqs

type bmRequirements struct {
	Config interface{} // Global compiler config
	Procr  map[int]*ProcRequirements
	IOr    map[int]*IORequirements
	Chanr  map[int]*ChanRequirements
	Shrdr  map[int]*SharedMemRequirements
}

func (reqmnt *BondgoRequirements) requirementsMonitor(useditem chan UsageNotify, usagedone chan bool) {
	//debug := reqmnt.Config.Debug
UB:
	for {
		notif := <-useditem

		targettype := notif.TargetType
		targetid := notif.TargetId
		componenttype := notif.ComponentType
		components := notif.Components
		componenti := notif.Componenti

		switch targettype {
		case TR_PROC:
			var proc *ProcRequirements
			if exists, ok := reqmnt.Procr[targetid]; ok {
				proc = exists
			} else {
				proc = new(ProcRequirements)
				reqmnt.Procr[targetid] = proc
			}

			switch componenttype {
			case C_OPCODE:
				present := false
				for _, op := range proc.Opcodes {
					if op == components {
						present = true
						break
					}
				}
				if !present {
					proc.Opcodes = append(proc.Opcodes, components)
				}
			case C_REGSIZE:
				if componenti > proc.Registersize {
					proc.Registersize = componenti
				}
			case C_INPUT:

				var ior *IORequirements
				if exists, ok := reqmnt.IOr[targetid]; ok {
					ior = exists
				} else {
					ior = new(IORequirements)
					ior.Inputs_ids = make([]int, 0)
					ior.Outputs_ids = make([]int, 0)
					reqmnt.IOr[targetid] = ior
				}

				present := false
				for _, inp := range ior.Inputs_ids {
					if inp == componenti {
						present = true
						break
					}
				}

				if !present {
					ior.Inputs_ids = append(ior.Inputs_ids, componenti)
					proc.Inputs = len(ior.Inputs_ids)
				}

			case C_OUTPUT:

				var ior *IORequirements
				if exists, ok := reqmnt.IOr[targetid]; ok {
					ior = exists
				} else {
					ior = new(IORequirements)
					ior.Inputs_ids = make([]int, 0)
					ior.Outputs_ids = make([]int, 0)
					reqmnt.IOr[targetid] = ior
				}

				present := false
				for _, outp := range ior.Outputs_ids {
					if outp == componenti {
						present = true
						break
					}
				}

				if !present {
					ior.Outputs_ids = append(ior.Outputs_ids, componenti)
					proc.Outputs = len(ior.Outputs_ids)
				}

			case C_ROMSIZE:
				if componenti > proc.Romsize {
					proc.Romsize = componenti
				}
			case C_RAMSIZE:
				if componenti > proc.Ramsize {
					proc.Ramsize = componenti
				}
			case C_SHAREDOBJECT:
				proc.SharedObjects = append(proc.SharedObjects, components)
			case C_DEVICE:
				proc.Device = components
			}

			// TODO Other cases
		case TR_CHAN:
			var cchan *ChanRequirements
			if exists, ok := reqmnt.Chanr[targetid]; ok {
				cchan = exists
			} else {
				cchan = new(ChanRequirements)
				reqmnt.Chanr[targetid] = cchan
			}
			switch componenttype {
			case C_CONNECTED:
				present := false
				for _, op := range cchan.Connected {
					if op == componenti {
						present = true
						break
					}
				}
				if !present {
					cchan.Connected = append(cchan.Connected, componenti)
				}
			}

		case TR_EXIT:
			break UB
		}
	}
	usagedone <- true
}
