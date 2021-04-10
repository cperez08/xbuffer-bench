package allocation

import (
	"fmt"

	model "github.com/cperez08/xbuffer-test/model/fbs"
	flatbuffers "github.com/google/flatbuffers/go"
)

func GetFBSLevel1() []byte {
	// GetBaseFBS(1, 0, 0)
	// 1 -> address for customer
	// 0 -> nested addresses for customer
	// 0 -> friends
	builder := GetBaseFBS(1, 0, 0)
	customer := model.CustomerEnd(builder)
	builder.Finish(customer)
	return builder.FinishedBytes()

}

func GetFBSLevel2() []byte {
	// GetBaseFBS(20, 10, 10)
	// 20 -> address for customer
	// 10 -> nested addresses for customer
	// 10 -> friends
	builder := GetBaseFBS(20, 10, 10)
	customer := model.CustomerEnd(builder)
	builder.Finish(customer)
	return builder.FinishedBytes()
}

func GetFBSLevel3() []byte {
	// GetBaseFBS(50, 20, 100)
	// 100 -> address for customer
	// 100 -> nested addresses for customer
	// 200 -> friends
	builder := GetBaseFBS(100, 100, 200)
	customer := model.CustomerEnd(builder)
	builder.Finish(customer)
	return builder.FinishedBytes()
}

func UnmarshalFBS(bts []byte) {
	customer := model.GetRootAsCustomer(bts, 0)
	if customer == nil {
		panic("no FBS customer found")
	}

	// you can access to the data like this
	// fmt.Println(string(customer.FirstName()))
	// fmt.Println(customer.AddressesLength())
	// fmt.Println(customer.FriendsLength())
	// friend := new(model.Customer)
	// for i := 0; i < customer.FriendsLength(); i++ {
	// 	if customer.Friends(friend, i) {
	// 		fmt.Println(string(friend.FirstName()))
	// 		fmt.Println(string(friend.LastName()))
	// 	}
	// }
}

func GetBaseFBS(addrCount, nestedAddrCount, friendsCount int) *flatbuffers.Builder {
	builder := flatbuffers.NewBuilder(256)

	firstName := builder.CreateString("FlatBuf First Name")
	lastName := builder.CreateString("FlatBuf Last Name")
	language := builder.CreateString("en_us")

	addresses := generateAddressFBS(builder, addrCount)
	friends := generateFriendsFBS(builder, nestedAddrCount, friendsCount)

	model.NotificationStart(builder)
	model.NotificationAddSms(builder, true)
	model.NotificationAddPush(builder, true)
	model.NotificationAddEmail(builder, true)
	notification := model.NotificationEnd(builder)

	model.PreferencesStart(builder)
	model.PreferencesAddDarkMode(builder, true)
	model.PreferencesAddLanguage(builder, language)
	model.PreferencesAddNotification(builder, notification)
	preferences := model.PreferencesEnd(builder)

	model.CustomerStart(builder)
	model.CustomerAddFirstName(builder, firstName)
	model.CustomerAddLastName(builder, lastName)
	model.CustomerAddAge(builder, 999)
	model.CustomerAddBalance(builder, 99994859059012502541295.993269264)
	model.CustomerAddDebt(builder, 99994859059012502541295.993269264)
	model.CustomerAddPreferences(builder, preferences)
	model.CustomerAddAddresses(builder, addresses)
	model.CustomerAddFriends(builder, friends)
	return builder
}

func generateAddressFBS(builder *flatbuffers.Builder, count int) flatbuffers.UOffsetT {
	var addrKV []flatbuffers.UOffsetT
	for i := 1; i <= count; i++ {
		addressAlias := builder.CreateString(fmt.Sprintf("my FlatBuf address # %d", i))
		address1 := builder.CreateString(fmt.Sprintf("street # %d NY city", i))

		model.LocationStart(builder)
		model.LocationAddAddress(builder, address1)
		model.LocationAddIsAtive(builder, true)
		model.LocationAddLatitude(builder, 40.6895541)
		model.LocationAddLongitude(builder, -74.0446034)
		location := model.LocationEnd(builder)

		model.AddressKVStart(builder)
		model.AddressKVAddAlias(builder, addressAlias)
		model.AddressKVAddLocation(builder, location)
		addrKV = append(addrKV, model.AddressKVEnd(builder))
	}

	model.CustomerStartAddressesVector(builder, count)
	for _, addr := range addrKV {
		builder.PrependUOffsetT(addr)
	}

	return builder.EndVector(count)
}

func generateFriendsFBS(builder *flatbuffers.Builder, addressCount, friendCount int) flatbuffers.UOffsetT {
	var friends []flatbuffers.UOffsetT
	for i := 1; i <= friendCount; i++ {
		firstName := builder.CreateString(fmt.Sprintf("FlatBuf Nested First Name # %d", i))
		lastName := builder.CreateString(fmt.Sprintf("FlatBuf Nested Last Name # %d", i))
		language := builder.CreateString("en_gb")

		addresses := generateAddressFBS(builder, addressCount)

		model.NotificationStart(builder)
		model.NotificationAddSms(builder, true)
		model.NotificationAddPush(builder, true)
		model.NotificationAddEmail(builder, true)
		notification := model.NotificationEnd(builder)

		model.PreferencesStart(builder)
		model.PreferencesAddDarkMode(builder, true)
		model.PreferencesAddLanguage(builder, language)
		model.PreferencesAddNotification(builder, notification)
		preferences := model.PreferencesEnd(builder)

		model.CustomerStart(builder)
		model.CustomerAddFirstName(builder, firstName)
		model.CustomerAddLastName(builder, lastName)
		model.CustomerAddAge(builder, 888)
		model.CustomerAddBalance(builder, 77774859059012502541295.883269264)
		model.CustomerAddDebt(builder, 77774859059012502541295.883269264)
		model.CustomerAddPreferences(builder, preferences)
		model.CustomerAddAddresses(builder, addresses)
		friends = append(friends, model.CustomerEnd(builder))
	}

	model.CustomerStartFriendsVector(builder, friendCount)
	for _, friend := range friends {
		builder.PrependUOffsetT(friend)
	}

	return builder.EndVector(friendCount)
}
