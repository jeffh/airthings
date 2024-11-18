# Airthings API

This project provides HomeKit integration for Airthings air quality monitors,
allowing you to view your Airthings sensor data directly via HomeKit using HAP.
It supports various air quality measurements including temperature, humidity,
CO2, VOC, radon, and atmospheric pressure.

## Installation

```bash
go get github.com/jeffh/airthings
```

## Usage

Use `Authorize` to create a client that can access device metrics:

```go
// scopes can be nil for default
client, err := airthings.Authorize(ctx, airthingsClientId, airthingsClientSecret, nil)
```

Then you can list all available devices to your account:

```go
devices, err := client.ListDevices(airthings.ListDevicesOptions{})
```

You can then collect devices from each device:

```go
samples, err := client.GetLatestSamples(airthings.GetLatestSamplesOptions{
	SerialNumber: devices[0].SerialNumber,
})
```

### Throttling

The biggest limitation of the airthings API is the aggressive throttling. It's
preferrable to poll no greater than every 5 minutes for all API calls.

### HAP

The hap subpackage provides accessories for [HAP][hap].

```go
import ahap "github.com/jeffh/airthings/hap"

var accessories []*accessory.A
for _, device := range devices {
	group := ahap.New(client)

	accessories = append(accessories, group.A)
	// you have control on how often to update from the API since directly
	// proxying requests to the API would cause you to throttle quickly.
	go func(device airthings.Device){
		for {
			samples, err := client.GetLatestSamples(airthings)
			if err == nil {
				group.Update(samples)
			} else {
				// ...
			}
			select {
			case <-time.After(15 * time.Minute):
				continue
			case <-ctx.Done:
				return
			}
		}
	}(device)
}

store := hap.NewFsStore(".hap")
server := hap.NewServer(store, accessories...)
// ... normal hap setup
```

[hap]: https://github.com/brutella/hap
