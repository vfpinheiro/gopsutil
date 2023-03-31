//go:build darwin && cgo
// +build darwin,cgo

package host

// #cgo LDFLAGS: -framework IOKit
// #include "smc_darwin.h"
import "C"
import "context"

func SensorsTemperaturesWithContext(ctx context.Context) ([]TemperatureStat, error) {
	temperatureKeys := []string{
		//Intel keys (Stats App)
		C.CPU_DIODE_FILTERED,
		C.CPU_PROXIMITY,
		C.CPU_CORE_1,
		C.CPU_CORE_2,
		C.CPU_CORE_3,
		C.CPU_CORE_4,
		C.CPU_CORE_5,
		C.CPU_CORE_6,
		C.GPU_INTEL_GRAPHICS,
		C.AIRPORT,
		C.MEMORY_SLOTS_PROXIMITY,

		//M1 Keys
		C.M1_CPU_EFFICIENCY_CORE_1,
		C.M1_CPU_EFFICIENCY_CORE_2,
		C.M1_CPU_PERFORMANCE_CORE_1,
		C.M1_CPU_PERFORMANCE_CORE_2,
		C.M1_CPU_PERFORMANCE_CORE_3,
		C.M1_CPU_PERFORMANCE_CORE_4,
		C.M1_CPU_PERFORMANCE_CORE_5,
		C.M1_CPU_PERFORMANCE_CORE_6,
		C.M1_CPU_PERFORMANCE_CORE_7,
		C.M1_CPU_PERFORMANCE_CORE_8,
		C.M1_GPU_1,
		C.M1_GPU_2,
		C.M1_GPU_3,
		C.M1_GPU_4,
		C.M1_NAND,
		C.M1_AIRPORT,

		//M2 Keys
		C.M2_AIRPORT,
		C.M2_NAND,
		C.M2_DISK_1_B,
		C.M2_CPU_EFFICIENCY_CORE_1,
		C.M2_CPU_EFFICIENCY_CORE_2,
		C.M2_CPU_EFFICIENCY_CORE_3,
		C.M2_CPU_EFFICIENCY_CORE_4,
		C.M2_CPU_PERFORMANCE_CORE_1,
		C.M2_CPU_PERFORMANCE_CORE_2,
		C.M2_CPU_PERFORMANCE_CORE_3,
		C.M2_CPU_PERFORMANCE_CORE_4,
		C.M2_GPU_1,
		C.M2_GPU_2,
	}
	var temperatures []TemperatureStat

	C.gopsutil_v3_open_smc()
	defer C.gopsutil_v3_close_smc()

	for _, key := range temperatureKeys {
		temperatures = append(temperatures, TemperatureStat{
			SensorKey:   key,
			Temperature: float64(C.gopsutil_v3_get_temperature(C.CString(key))),
		})
	}
	return temperatures, nil
}
