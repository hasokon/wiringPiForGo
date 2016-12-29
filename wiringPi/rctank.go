package wiringPi

type RcTank struct {
	RightMoter *Moter
	LeftMoter *Moter
}

type RcTankBehavior interface {
	Foward() error
	Back() error
	TurnRight() error
	TurnLeft() error
	Stop() error
}

func NewRcTank() (*RcTank, error) {
	rcTank := new(RcTank)
	var err error
	rcTank.RightMoter, err = NewMoter(27, 22)
	if err != nil {
		return nil,err
	}
	rcTank.LeftMoter, err = NewMoter(17, 4)
	if err != nil {
		return nil, err
	}

	return rcTank, nil
}

func (rt *RcTank) Foward() error {
	if err := rt.RightMoter.Foward(); err != nil {
		return err
	}
	if err := rt.LeftMoter.Foward(); err != nil {
		return err
	}
	return nil
}

func (rt *RcTank) Back() error {
	if err := rt.RightMoter.Back(); err != nil {
		return err
	}
	if err := rt.LeftMoter.Back(); err != nil {
		return err
	}
	return nil
}

func (rt *RcTank) TurnRight() error {
	if err := rt.RightMoter.Back(); err != nil {
		return err
	}
	if err := rt.LeftMoter.Foward(); err != nil {
		return err
	}
	return nil
}

func (rt *RcTank) TurnLeft() error {
	if err := rt.RightMoter.Foward(); err != nil {
		return err
	}
	if err := rt.LeftMoter.Back(); err != nil {
		return err
	}
	return nil
}

func (rt *RcTank) Stop() error {
	if err := rt.RightMoter.Stop(); err != nil {
		return err
	}
	if err := rt.LeftMoter.Stop(); err != nil {
		return err
	}
	return nil
}
