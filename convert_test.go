package apple_price

import "testing"

func TestConvert(t *testing.T) {
	for _, storefront := range storefronts {
		countryCode, ok := StorefrontToCountryCodeMap[storefront]
		if !ok {
			t.Errorf("storefront %s not found in storefrontToCountryCodeMap", storefront)
			return
		}
		territory, ok := CountryCodeToTerritoryMap[countryCode]
		if !ok {
			t.Errorf("country code %s not found in countryCodeToTerritoryMap", countryCode)
			return
		}
		if territory != storefront {
			t.Errorf("storefront %s not equal to territory %s", storefront, territory)
			return
		}
		currency, ok := TerritoryToCurrencyMap[territory]
		if !ok {
			t.Errorf("territory %s not found in territoryToCurrencyMap", territory)
			return
		}
		t.Logf("storefront %s, country code %s, territory %s, currency %s", storefront, countryCode, territory, currency)
	}
}
