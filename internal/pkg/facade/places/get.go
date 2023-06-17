package places

import "context"

// GetAppleAddressByPlaceID ...
func (f *Facade) GetAppleAddressByPlaceID(ctx context.Context, placeID int32) (string, error) {
	place, err := f.placeStorage.GetPlaceByExternalPlaceID(ctx, int(placeID))
	if err != nil {
		return "", err
	}

	return place.AppleAddress, nil
}
