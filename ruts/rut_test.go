package ruts

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestDV(t *testing.T) {
    assert.Equal(t, '3', GetDV(16089456))
    assert.Equal(t, '3', GetDV(16416888))
    assert.Equal(t, '9', GetDV(22652884))
    assert.Equal(t, '9', GetDV(24322384))
    assert.Equal(t, '2', GetDV(10858247))
    assert.Equal(t, '6', GetDV(10462669))
    assert.Equal(t, '9', GetDV(99999999))
    assert.Equal(t, 'K', GetDV(12341246))
    assert.Equal(t, 'K', GetDV(12341263))
}

func TestParse(t *testing.T) {
    rut, err := Parse("16.089.456-3")
    assert.Equal(t, 16089456, rut)
    assert.Nil(t, err)

    rut, err = Parse("160894563")
    assert.Equal(t, 16089456, rut)
    assert.Nil(t, err)

    rut, err = Parse("19")
    assert.Equal(t, 1, rut)
    assert.Nil(t, err)

    rut, err = Parse("99.999.999-9")
    assert.Equal(t, 99999999, rut)
    assert.Nil(t, err)

    rut, err = Parse("")
    assert.Equal(t, 0, rut)
    assert.NotNil(t, err)
}

func TestValidate(t *testing.T) {
    valid, err := Validate("16.089.456-3")
    assert.Nil(t, err)
    assert.True(t, valid)

    valid, err = Validate("16.416.888-3")
    assert.Nil(t, err)
    assert.True(t, valid)

    valid, err = Validate("226528849")
    assert.Nil(t, err)
    assert.True(t, valid)

    valid, err = Validate("243223849")
    assert.Nil(t, err)
    assert.True(t, valid)

    valid, err = Validate("12341246k")
    assert.Nil(t, err)
    assert.True(t, valid)

    valid, err = Validate("12341246K")
    assert.Nil(t, err)
    assert.True(t, valid)

    valid, err = Validate("11.111.111-1")
    assert.Nil(t, err)
    assert.True(t, valid)

    valid, err = Validate("104626691")
    assert.Nil(t, err)
    assert.False(t, valid)

    valid, err = Validate("11.111.111-k")
    assert.Nil(t, err)
    assert.False(t, valid)

    valid, err = Validate("")
    assert.NotNil(t, err)
    assert.False(t, valid)

    valid, err = Validate("19")
    assert.Nil(t, err)
    assert.True(t, valid)
}

func TestFormat(t *testing.T) {
    assert.Equal(t, "16.089.456-3", Format(16089456))
    assert.Equal(t, "16.416.888-3", Format(16416888))
    assert.Equal(t, "24.322.384-9", Format(24322384))
    assert.Equal(t, "22.652.884-9", Format(22652884))
    assert.Equal(t, "11.111.111-1", Format(11111111))
    assert.Equal(t, "1-9", Format(1))
}

