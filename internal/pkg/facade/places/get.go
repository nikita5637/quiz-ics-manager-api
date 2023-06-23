package places

import "context"

// GetAppleAddressByPlaceID ...
func (f *Facade) GetAppleAddressByPlaceID(ctx context.Context, placeID int32) (string, error) {
	var appleAddress string
	err := f.db.RunTX(ctx, "GetAppleAddressByPlaceID", func(ctx context.Context) error {
		place, err := f.placeStorage.GetPlaceByExternalPlaceID(ctx, int(placeID))
		if err != nil {
			return err
		}

		appleAddress = place.AppleAddress
		return nil
	})
	if err != nil {
		return "", err
	}

	return appleAddress, nil
}
