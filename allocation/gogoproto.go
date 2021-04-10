package allocation

import (
	"fmt"

	model "github.com/cperez08/xbuffer-test/model/gogoproto"
	"github.com/gogo/protobuf/proto"
)

func GetGogoProtoLevel1() []byte {
	// GetBaseProto(1, 0, 0)
	// 1 -> address for customer
	// 0 -> nested addresses for customer
	// 0 -> friends
	bt, _ := proto.Marshal(GetBaseGogoProto(1, 0, 0))
	return bt
}

func GetGogoProtoLevel2() []byte {
	// GetBaseProto(20, 10, 10)
	// 20 -> address for customer
	// 10 -> nested addresses for customer
	// 10 -> friends
	bt, _ := proto.Marshal(GetBaseGogoProto(20, 10, 10))
	return bt
}

func GetGogoProtoLevel3() []byte {
	// GetBaseProto(100, 100, 200)
	// 100 -> address for customer
	// 100 -> nested addresses for customer
	// 200 -> friends
	bt, _ := proto.Marshal(GetBaseGogoProto(100, 100, 200))
	return bt
}

func UnmarshalGogoProto(bts []byte) {
	customer := new(model.Customer)
	proto.Unmarshal(bts, customer)
	if customer == nil {
		panic("no proto customer found")
	}
}

func GetBaseGogoProto(addrCount, nestedAddrCount, friendCount int) *model.Customer {
	return &model.Customer{
		FirstName: "ProtoF First Name",
		LastName:  "ProtoF Last Name",
		Age:       999,
		Balance:   99994859059012502541295.993269264,
		Debt:      450505050.9139519591259,
		Addresses: generateAddressGogoProto(addrCount),
		Friends:   generateFriendsGogoProto(nestedAddrCount, friendCount),
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

func generateAddressGogoProto(addrCount int) map[string]*model.Location {
	addrs := make(map[string]*model.Location)
	for i := 1; i <= addrCount; i++ {
		addrs[fmt.Sprintf("my Proto address # %d", i)] = &model.Location{
			Address:   fmt.Sprintf("street # %d NY city", i),
			Latitude:  40.6895541,
			Longitude: -74.0446034,
			IsActive:  true,
		}
	}

	return addrs
}

func generateFriendsGogoProto(nestedAddrCount, friendsCount int) []*model.Customer {
	var friends []*model.Customer
	for i := 1; i <= friendsCount; i++ {
		friends = append(friends, &model.Customer{
			FirstName: "Proto Nested First Name",
			LastName:  "Proto Nested Last Name",
			Age:       888,
			Balance:   77774859059012502541295.883269264,
			Debt:      77774859059012502541295.883269264,
			Addresses: generateAddressGogoProto(nestedAddrCount),
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
