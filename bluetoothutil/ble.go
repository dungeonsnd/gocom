package bluetoothutil

// This example implements a NUS (Nordic UART Service) peripheral.
// I can't find much official documentation on the protocol, but this can be
// helpful:
// https://learn.adafruit.com/introducing-adafruit-ble-bluetooth-low-energy-friend/uart-service
//
// Code to interact with a raw terminal is in separate files with build tags.

import (
	"fmt"

	"github.com/dungeonsnd/bluetooth"
)

var (
	serviceUUID = bluetooth.NewUUID([16]byte{0x5e, 0x30, 0x00, 0x01, 0xb5, 0xa3, 0xf3, 0x93, 0xe0, 0xa9, 0xe5, 0x0e, 0x24, 0xdc, 0xca, 0x1e})
	rxUUID      = bluetooth.NewUUID([16]byte{0x5e, 0x30, 0x00, 0x02, 0xb5, 0xa3, 0xf3, 0x93, 0xe0, 0xa9, 0xe5, 0x0e, 0x24, 0xdc, 0xca, 0x1e})
	txUUID      = bluetooth.NewUUID([16]byte{0x5e, 0x30, 0x00, 0x03, 0xb5, 0xa3, 0xf3, 0x93, 0xe0, 0xa9, 0xe5, 0x0e, 0x24, 0xdc, 0xca, 0x1e})
)

var rxChar bluetooth.Characteristic
var txChar bluetooth.Characteristic

var adapter *bluetooth.Adapter

type OnRecvCallback func(client bluetooth.Connection, offset int, value []byte)

func StartByServiceUUID(serviceName, _serviceUUID string, recvCallBack OnRecvCallback) error {
	uuid, err := bluetooth.ParseUUID(_serviceUUID)
	if err != nil {
		return fmt.Errorf("Faled bluetooth.ParseUUID,  err:%v", err)
	}
	// fmt.Printf("uuid=%v \n", uuid)
	serviceUUID = uuid
	return Start(serviceName, recvCallBack)
}

func Start(serviceName string, recvCallBack OnRecvCallback) error {
	// uuid, err := bluetooth.ParseUUID("12342233-0000-1000-8000-A068189DFD22")
	// if err != nil {
	// 	return fmt.Errorf("Faled bluetooth.ParseUUID,  err:%v", err)
	// }
	// // fmt.Printf("uuid=%v \n", uuid)
	// serviceUUID = uuid

	// uuid, err = bluetooth.ParseUUID("12340001-0000-1000-8000-A068189DFD22")
	// if err != nil {
	// 	return fmt.Errorf("Faled bluetooth.ParseUUID,  err:%v", err)
	// }
	// rxUUID = uuid

	// uuid, err = bluetooth.ParseUUID("12340002-0000-1000-8000-A068189DFD22")
	// if err != nil {
	// 	return fmt.Errorf("Faled bluetooth.ParseUUID,  err:%v", err)
	// }
	// txUUID = uuid

	adapter = bluetooth.DefaultAdapter

	err := adapter.Enable()
	if err != nil {
		return fmt.Errorf("Faled adapter.Enable,  err:%v", err)
	}

	adv := adapter.DefaultAdvertisement()
	err = adv.Configure(bluetooth.AdvertisementOptions{
		LocalName:    serviceName, // Nordic UART Service
		ServiceUUIDs: []bluetooth.UUID{serviceUUID},
	})
	if err != nil {
		return fmt.Errorf("Faled adv.Configure, err:%v", err)
	}

	err = adv.Start()
	if err != nil {
		return fmt.Errorf("Faled adv.Start, err:%v", err)
	}

	err = adapter.AddService(&bluetooth.Service{
		UUID: serviceUUID,
		Characteristics: []bluetooth.CharacteristicConfig{
			{
				Handle: &rxChar,
				UUID:   rxUUID,
				Flags:  bluetooth.CharacteristicWritePermission,
				WriteEvent: func(client bluetooth.Connection, offset int, value []byte) {
					// txChar.Write(value)
					recvCallBack(client, offset, value)
				},
			},
			{
				Handle: &txChar,
				UUID:   txUUID,
				Flags:  bluetooth.CharacteristicNotifyPermission | bluetooth.CharacteristicReadPermission,
			},
		},
	})
	if err != nil {
		return fmt.Errorf("Faled adapter.AddService, err:%v", err)
	}

	return nil
}

func SendData(sendbuf []byte) (int, error) {
	if sendbuf == nil {
		return 0, fmt.Errorf("sendbuf == nil")
	}
	fmt.Printf("$$ sendbuf=%x\n", sendbuf)
	n, err := txChar.Write(sendbuf)
	if err != nil {
		return n, fmt.Errorf("Faled txChar.Write, err:%v", err)
	}
	return n, nil
}

func SendDataMultiTimes(sendbuf []byte) error {
	for len(sendbuf) != 0 {

		partlen := 15
		if len(sendbuf) < partlen {
			partlen = len(sendbuf)
		}
		part := sendbuf[:partlen]
		sendbuf = sendbuf[partlen:]
		// This also sends a notification.
		_, err := txChar.Write(part)
		if err != nil {
			return fmt.Errorf("Faled txChar.Write, err:%v", err)
		}
	}
	return nil
}
