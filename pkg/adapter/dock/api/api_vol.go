// Copyright (c) 2016 Huawei Technologies Co., Ltd. All Rights Reserved.
//
//    Licensed under the Apache License, Version 2.0 (the "License"); you may
//    not use this file except in compliance with the License. You may obtain
//    a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//    License for the specific language governing permissions and limitations
//    under the License.

/*
This module implements the entry into operations of storageDock module.

*/

package api

import (
	"log"

	"github.com/opensds/opensds/pkg/adapter/dock"
)

func CreateVolume(resourceType string, name string, size int) (string, error) {
	result, err := dock.CreateVolume(resourceType, name, size)

	if err != nil {
		log.Println("Error occured in adapter module when create volume!")
		return "", err
	} else {
		return result, nil
	}
}

func GetVolume(resourceType string, volID string) (string, error) {
	result, err := dock.GetVolume(resourceType, volID)

	if err != nil {
		log.Println("Error occured in adapter module when get volume!")
		return "", err
	} else {
		return result, nil
	}
}

func GetAllVolumes(resourceType string, allowDetails bool) (string, error) {
	result, err := dock.GetAllVolumes(resourceType, allowDetails)

	if err != nil {
		log.Println("Error occured in adapter module when get all volumes!")
		return "", err
	} else {
		return result, nil
	}
}

func UpdateVolume(resourceType string, volID string, name string) (string, error) {
	result, err := dock.UpdateVolume(resourceType, volID, name)

	if err != nil {
		log.Println("Error occured in adapter module when update volume!")
		return "", err
	} else {
		return result, nil
	}
}

func DeleteVolume(resourceType string, volID string) (string, error) {
	result, err := dock.DeleteVolume(resourceType, volID)

	if err != nil {
		log.Println("Error occured in adapter module when delete volume!")
		return "", err
	} else {
		return result, nil
	}
}

func AttachVolume(resourceType, volID, host, device string) (string, error) {
	result, err := dock.AttachVolume(resourceType, volID, host, device)

	if err != nil {
		log.Println("Error occured in adapter module when attach volume!")
		return "", err
	} else {
		return result, nil
	}
}

func DetachVolume(resourceType, volID, attachment string) (string, error) {
	result, err := dock.DetachVolume(resourceType, volID, attachment)

	if err != nil {
		log.Println("Error occured in adapter module when detach volume!")
		return "", err
	} else {
		return result, nil
	}
}

func MountVolume(mountDir, device, fsType string) (string, error) {
	result, err := dock.MountVolume(mountDir, device, fsType)

	if err != nil {
		log.Println("Error occured in adapter module when mount volume!")
		return "", err
	} else {
		return result, nil
	}
}

func UnmountVolume(mountDir string) (string, error) {
	result, err := dock.UnmountVolume(mountDir)

	if err != nil {
		log.Println("Error occured in adapter module when unmount volume!")
		return "", err
	} else {
		return result, nil
	}
}