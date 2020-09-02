package xcassets

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestDevices_idioms(t *testing.T) {
	d := Devices{}
	d.Universal()
	d.IPhone()
	d.IPad()
	d.Catalyst()
	d.CarPlay()
	d.AppleWatch()
	d.AppleTV()
	d.Mac()

	idioms := d.idioms()
	expected := []string{"universal", "iphone", "ipad", "car", "watch", "tv", "mac"}

	assert.Equal(t, expected, idioms, "Idioms should equal specified idioms")
}

func TestDevices_subtypes(t *testing.T) {
	d := Devices{}
	d.Catalyst()

	assert.True(t, d.iPad)
	assert.True(t, d.catalyst)

	expected := []string{"mac-catalyst"}
	subtypes := d.subtypes()

	assert.Equal(t, expected, subtypes)
}