package characteristic

import "github.com/brutella/hap/characteristic"

type RadonShortTermAverage struct {
	*characteristic.Float
}

const TypeRadonShortTermAverage = "03147F95-F731-4D19-A409-AB7858A835E6"

func NewRadonShortTermAverage() *RadonShortTermAverage {
	c := &RadonShortTermAverage{}
	c.Float = characteristic.NewFloat(TypeRadonShortTermAverage)
	c.Float.Description = "Radon (24h avg)"
	c.Float.Unit = "Bq/m³"
	c.Float.Format = characteristic.FormatFloat
	c.Float.Permissions = []string{characteristic.PermissionRead, characteristic.PermissionEvents}
	c.Float.SetMinValue(0)
	c.Float.SetMaxValue(16383)
	c.Float.SetStepValue(1)
	return c
}
