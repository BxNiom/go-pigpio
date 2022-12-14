package pigpio

func (p *Pi) StoreScript(code string) (*Script, error) {
	r, e := p.socket.SendCommand(cmdScript, 0, 0, []byte(code))
	if e != nil || r < 0 {
		return nil, newPiError(r, e, "StoreScript(code: %s)", code)
	}

	return &Script{pi: p, handle: r, code: code}, nil
}

func (p *Pi) AttachToScript(handle int) *Script {
	return &Script{pi: p, handle: handle, code: ""}
}
