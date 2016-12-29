package wiringPi

type Moter struct {
	Out1 *Gpio
	Out2 *Gpio
}

type MoterBehavior interface {
	Foward() error
	Back() error
	Stop() error
}

func NewMoter (pin1 uint, pin2 uint) (*Moter, error) {
	moter := new(Moter)
	var err error
	moter.Out1, err = NewGpio(pin1, OUTPUT)
	if err != nil {
		return nil, err
	}
	moter.Out2, err = NewGpio(pin2, OUTPUT)
	if err != nil {
		return nil, err
	}

	return moter, nil
}

func (m *Moter) Foward() error {
	if err := m.Out1.DigitalWrite(1); err != nil {
		return err
	}
	if err := m.Out2.DigitalWrite(0); err != nil {
		return err
	}

	return nil
}

func (m *Moter) Back() error {
	if err := m.Out1.DigitalWrite(0); err != nil {
		return err
	}
	if err := m.Out2.DigitalWrite(1); err != nil {
		return err
	}

	return nil
}

func (m *Moter) Stop() error {
	if err := m.Out1.DigitalWrite(0); err != nil {
		return err
	}
	if err := m.Out2.DigitalWrite(0); err != nil {
		return err
	}

	return nil
}
