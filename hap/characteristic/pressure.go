package characteristic

import "github.com/brutella/hap/characteristic"

type PressureCharacteristic struct {
	*characteristic.Float
}

const TypePressure = "7BA02876-A60F-487A-B9EA-DC9A096B447A"

func NewPressure() *PressureCharacteristic {
	c := &PressureCharacteristic{}
	c.Float = characteristic.NewFloat(TypePressure)
	c.Float.Description = "Air Pressure"
	c.Float.Unit = "hPa"
	c.Float.Format = characteristic.FormatFloat
	c.Float.Permissions = []string{characteristic.PermissionRead, characteristic.PermissionEvents}
	c.Float.SetMinValue(850)
	c.Float.SetMaxValue(1100)
	c.Float.SetStepValue(1)
	return c
}
