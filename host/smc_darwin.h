#ifndef __SMC_H__
#define __SMC_H__ 1

#include <IOKit/IOKitLib.h>

//Intel keys (Stats App)
#define CPU_DIODE_FILTERED			"TC0F"
#define CPU_PROXIMITY				"TC0P"
#define CPU_CORE_1 					"TC1C" //or TC1c
#define CPU_CORE_2 					"TC2C" //or TC2c
#define CPU_CORE_3 					"TC3C" //or TC3c
#define CPU_CORE_4 					"TC4C" //or TC4c
#define CPU_CORE_5 					"TC5C" //or TC5c
#define CPU_CORE_6 					"TC6C" //or TC6c
#define GPU_INTEL_GRAPHICS			"TCGC"
#define AIRPORT						"TW0P"
//(Mac-exporter)
#define MEMORY_SLOTS_PROXIMITY 		"TM0P"


//M1 keys (Stats App)
#define M1_CPU_EFFICIENCY_CORE_1	"Tp09"
#define M1_CPU_EFFICIENCY_CORE_2	"Tp0T"
#define M1_CPU_PERFORMANCE_CORE_1	"Tp01"
#define M1_CPU_PERFORMANCE_CORE_2	"Tp05"
#define M1_CPU_PERFORMANCE_CORE_3	"Tp0D"
#define M1_CPU_PERFORMANCE_CORE_4	"Tp0H"
#define M1_CPU_PERFORMANCE_CORE_5	"Tp0L"
#define M1_CPU_PERFORMANCE_CORE_6	"Tp0P"
#define M1_CPU_PERFORMANCE_CORE_7	"Tp0X"
#define M1_CPU_PERFORMANCE_CORE_8	"Tp0b"
#define M1_GPU_1					"Tg05"
#define M1_GPU_2					"Tg0D"
#define M1_GPU_3					"Tg0L"
#define M1_GPU_4					"Tg0T"
#define M1_NAND						"TH0x"
#define M1_AIRPORT					"TW0P"

//M2 keys (Stats App)
#define M2_AIRPORT					"TW0P"
#define M2_NAND						"TH0x"
#define M2_DISK_1_B					"TH1B"
#define M2_CPU_EFFICIENCY_CORE_1	"Tp05"
#define M2_CPU_EFFICIENCY_CORE_2	"Tp0D"
#define M2_CPU_EFFICIENCY_CORE_3	"Tp0j"
#define M2_CPU_EFFICIENCY_CORE_4	"Tp0r"
#define M2_CPU_PERFORMANCE_CORE_1	"Tp01"
#define M2_CPU_PERFORMANCE_CORE_2	"Tp09"
#define M2_CPU_PERFORMANCE_CORE_3	"Tp0f"
#define M2_CPU_PERFORMANCE_CORE_4	"Tp0n"
#define M2_GPU_1					"Tg0f"
#define M2_GPU_2					"Tg0n"


kern_return_t gopsutil_v3_open_smc(void);
kern_return_t gopsutil_v3_close_smc(void);
double gopsutil_v3_get_temperature(char *);

#endif // __SMC_H__
