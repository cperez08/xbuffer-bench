package allocation

import (
	"encoding/json"
	"fmt"

	model "github.com/cperez08/xbuffer-test/model/native"
)

func GetNativeLevel1() []byte {
	// GetBaseNative(1, 0, 0)
	// 1 -> address for customer
	// 0 -> nested addresses for customer
	// 0 -> friends
	bt, _ := json.Marshal(GetBaseNative(1, 0, 0))
	return bt
}

func GetNativeLevel2() []byte {
	// GetBaseNative(20, 10, 10)
	// 20 -> address for customer
	// 10 -> nested addresses for customer
	// 10 -> friends
	bt, _ := json.Marshal(GetBaseNative(20, 10, 10))
	return bt
}

func GetNativeLevel3() []byte {
	// GetBaseNative(100, 100, 200)
	// 100 -> address for customer
	// 100 -> nested addresses for customer
	// 200 -> friends
	bt, _ := json.Marshal(GetBaseNative(100, 100, 200))
	return bt
}

func UnmarshalNative(bts []byte) {
	customer := new(model.Customer)
	json.Unmarshal(bts, customer)
	if customer == nil {
		panic("no native customer found")
	}
}

func GetBaseNative(addrCount, nestedAddrCount, friendCount int) *model.Customer {
	return &model.Customer{
		FirstName: "Native First Name",
		LastName:  "Native Last Name",
		Age:       999,
		Balance:   99994859059012502541295.993269264,
		Debt:      450505050.9139519591259,
		Addresses: generateAddressNative(addrCount),
		Friends:   generateFriendsNative(nestedAddrCount, friendCount),
		Preferences: &model.Preferences{
			DarkMode: true,
			Language: "en_us",
			Notification: &model.Notification{
				Sms:   true,
				Push:  true,
				Email: true,
			},
		},
	}
}

func generateAddressNative(addrCount int) map[string]*model.Location {
	addrs := make(map[string]*model.Location)
	for i := 1; i <= addrCount; i++ {
		addrs[fmt.Sprintf("my Native address # %d", i)] = &model.Location{
			Address:   fmt.Sprintf("street # %d NY city", i),
			Latitude:  40.6895541,
			Longitude: -74.0446034,
			IsActive:  true,
		}
	}

	return addrs
}

func generateFriendsNative(nestedAddrCount, friendsCount int) []*model.Customer {
	var friends []*model.Customer
	for i := 1; i <= friendsCount; i++ {
		friends = append(friends, &model.Customer{
			FirstName: "Native Nested First Name",
			LastName:  "Native Nested Last Name",
			Age:       888,
			Balance:   77774859059012502541295.883269264,
			Debt:      77774859059012502541295.883269264,
			Addresses: generateAddressNative(nestedAddrCount),
			Preferences: &model.Preferences{
				DarkMode: true,
				Language: "en_gb",
				Notification: &model.Notification{
					Sms:   true,
					Push:  true,
					Email: true,
				},
			},
		})
	}

	return friends
}
