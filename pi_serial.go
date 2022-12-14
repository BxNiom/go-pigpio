package pigpio

func (p *Pi) OpenSerial(tty string, baud int) (*Serial, error) {
	r, e := p.socket.SendCommand(cmdSerialOpen, baud, 0, []byte(tty))
	if e != nil || r < 0 {
		return nil, newPiError(r, e, "OpenSerial(tty: %s, baud: %d)", tty, baud)
	}

	return &Serial{tty: tty, handle: r, baud: baud, pi: p}, nil
}
