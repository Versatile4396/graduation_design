package snowflake

import (
	"fmt"
	"time"

	"github.com/sony/sonyflake"
)

var (
	sonyFlake     *sonyflake.Sonyflake // 实例
	sonyMachineID uint16
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

func Init(machineID uint16) (err error) {
	sonyMachineID = machineID
	t, _ := time.Parse("2006-01-02", "2025-01-12")
	settings := sonyflake.Settings{
		StartTime: t,
		MachineID: getMachineID, // 指定机器ID
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	return
}

// GetId 返回生成的Id值

func GetID() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("sony flake not inited")
	}
	id, err = sonyFlake.NextID()
	return
}
