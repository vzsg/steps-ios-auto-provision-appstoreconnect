package autoprovision

import "github.com/bitrise-steplib/steps-ios-auto-provision-appstoreconnect/appstoreconnect"

// ListDevices returns the registered devices on the Apple Developer portal
func ListDevices(client *appstoreconnect.Client, udid string, platform appstoreconnect.DevicePlatform) ([]appstoreconnect.Device, error) {
	var nextPageURL string
	var devices []appstoreconnect.Device
	for {
		response, err := client.Provisioning.ListDevices(&appstoreconnect.ListDevicesOptions{
			PagingOptions: appstoreconnect.PagingOptions{
				Limit: 20,
				Next:  nextPageURL,
			},
			FilterUDID:     udid,
			FilterPlatform: platform,
		})
		if err != nil {
			return nil, err
		}

		devices = append(devices, response.Data...)

		nextPageURL = response.Links.Next
		if nextPageURL == "" {
			return devices, nil
		}
	}
}
