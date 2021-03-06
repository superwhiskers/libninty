/*

fennel - nintendo network utility library for golang
Copyright (C) 2018-2019 superwhiskers <whiskerdev@protonmail.com>

this source code form is subject to the terms of the mozilla public
license, v. 2.0. if a copy of the mpl was not distributed with this
file, you can obtain one at http://mozilla.org/MPL/2.0/.

*/

package types

import (
	crunch "github.com/superwhiskers/crunch/v3"
	"github.com/superwhiskers/fennel/utils"
)

// swaps the endianness of a mii binary format to little-endian
func swapMiiEndiannessToLittle(data []byte) []byte {

	data = utils.Swapu32Little(data, 0x00)

	for i := 0x18; i <= 0x2E; i += 2 {

		data = utils.Swapu16Little(data, i)

	}

	for i := 0x30; i <= 0x48; i += 2 {

		data = utils.Swapu16Little(data, i)

	}

	for i := 0x48; i <= 0x5C; i += 2 {

		data = utils.Swapu16Little(data, i)

	}

	data = utils.Swapu16Little(data, 0x5C)

	return data

}

// swaps the endianness of a mii binary format to big-endian
func swapMiiEndiannessToBig(data []byte) []byte {

	data = utils.Swapu32Big(data, 0x00)

	for i := 0x18; i <= 0x2E; i += 2 {

		data = utils.Swapu16Big(data, i)

	}

	for i := 0x30; i <= 0x48; i += 2 {

		data = utils.Swapu16Big(data, i)

	}

	for i := 0x48; i <= 0x5C; i += 2 {

		data = utils.Swapu16Big(data, i)

	}

	data = utils.Swapu16Big(data, 0x5C)

	return data

}

// converts a bit to a bool
func bitToBool(data byte) bool {

	return data != 0x00

}

// converts a bool to a bit
func boolToBit(data bool) byte {

	if !data {

		return 0x00

	}
	return 0x01

}

// Mii contains all of the data that a mii can have
type Mii struct {
	// unknown fields
	Unknown1  uint64
	Unknown2  uint64
	Unknown3  uint64
	Unknown4  byte
	Unknown5  []byte
	Unknown6  byte
	Unknown7  byte
	Unknown8  uint8
	Unknown9  byte
	Unknown10 []byte

	// attributes
	BirthPlatform uint64
	FontRegion    uint64
	RegionMove    uint64
	Copyable      bool
	MiiVersion    uint8
	AuthorID      []uint8
	MiiID         []uint8
	LocalOnly     bool
	Color         uint64
	BirthDay      uint64
	BirthMonth    uint64
	Gender        byte
	MiiName       string
	Size          uint8
	Fatness       uint8
	AuthorName    string
	Checksum      uint16

	// face
	BlushType uint64
	FaceStyle uint64
	FaceColor uint64
	FaceType  uint64

	// hair
	HairMirrored uint64
	HairColor    uint64
	HairType     uint8

	// eyes
	EyeThickness uint64
	EyeScale     uint64
	EyeColor     uint64
	EyeType      uint64
	EyeHeight    uint64
	EyeDistance  uint64
	EyeRotation  uint64

	// eyebrow
	EyebrowThickness uint64
	EyebrowScale     uint64
	EyebrowColor     uint64
	EyebrowType      uint64
	EyebrowHeight    uint64
	EyebrowDistance  uint64
	EyebrowRotation  uint64

	// nose
	NoseHeight uint64
	NoseScale  uint64
	NoseType   uint64

	// mouth
	MouthThickness uint64
	MouthScale     uint64
	MouthColor     uint64
	MouthType      uint64
	MouthHeight    uint64

	// mustache
	MustacheType   uint64
	MustacheHeight uint64
	MustacheScale  uint64

	// beard
	BeardColor uint64
	BeardType  uint64

	// glasses
	GlassesHeight uint64
	GlassesScale  uint64
	GlassesColor  uint64
	GlassesType   uint64

	// mole
	MoleY       uint64
	MoleX       uint64
	MoleScale   uint64
	MoleEnabled bool
}

// NilMii is a Mii with no data
var NilMii = Mii{}

// ParseMii takes a mii as a byte array and returns a parsed Mii
func ParseMii(miiByte []byte) *Mii {

	mii := &Mii{}
	mii.Parse(miiByte)
	return mii

}

// Parse takes a mii as a byte array and parses it into a Mii
// TODO: potentially hardcode offsets for `Seek` calls
func (mii *Mii) Parse(miiByte []byte) {

	var (
		tmp1 byte
		tmp2 []byte
		tmp3 []uint16
	)

	buf := &crunch.MiniBuffer{}
	crunch.NewMiniBuffer(&buf, swapMiiEndiannessToLittle(miiByte))

	buf.ReadBitsNext(&mii.BirthPlatform, 4)
	buf.ReadBitsNext(&mii.Unknown1, 4)
	buf.ReadBitsNext(&mii.Unknown2, 4)
	buf.ReadBitsNext(&mii.Unknown3, 4)
	buf.ReadBitsNext(&mii.FontRegion, 4)
	buf.ReadBitsNext(&mii.RegionMove, 2)
	buf.ReadBitNext(&mii.Unknown4)

	buf.ReadBitNext(&tmp1)
	mii.Copyable = bitToBool(tmp1)

	buf.AlignByte()

	buf.ReadBytesNext(&tmp2, 1)
	mii.MiiVersion = tmp2[0]

	buf.ReadBytesNext(&mii.AuthorID, 8)
	buf.ReadBytesNext(&mii.MiiID, 10)
	buf.ReadBytesNext(&mii.Unknown5, 2)

	buf.AlignBit()

	buf.ReadBitNext(&mii.Unknown6)
	buf.ReadBitNext(&mii.Unknown7)
	buf.ReadBitsNext(&mii.Color, 4)
	buf.ReadBitsNext(&mii.BirthDay, 5)
	buf.ReadBitsNext(&mii.BirthMonth, 4)
	buf.ReadBitNext(&mii.Gender)

	buf.AlignByte()

	buf.ReadBytesNext(&tmp2, 20)
	mii.MiiName = utils.DecodeUTF8StringFromBytes(tmp2)

	buf.ReadBytesNext(&tmp2, 1)
	mii.Fatness = tmp2[0]

	buf.ReadBytesNext(&tmp2, 1)
	mii.Size = tmp2[0]

	buf.AlignBit()

	buf.ReadBitsNext(&mii.BlushType, 4)
	buf.ReadBitsNext(&mii.FaceStyle, 4)
	buf.ReadBitsNext(&mii.FaceColor, 3)
	buf.ReadBitsNext(&mii.FaceType, 4)

	buf.ReadBitNext(&tmp1)
	mii.LocalOnly = bitToBool(tmp1)

	buf.ReadBitsNext(&mii.HairMirrored, 5)
	buf.ReadBitsNext(&mii.HairColor, 3)

	buf.AlignByte()

	buf.ReadBytesNext(&tmp2, 1)
	mii.HairType = tmp2[0]

	buf.AlignBit()

	buf.ReadBitsNext(&mii.EyeThickness, 3)
	buf.ReadBitsNext(&mii.EyeScale, 4)
	buf.ReadBitsNext(&mii.EyeColor, 3)
	buf.ReadBitsNext(&mii.EyeType, 6)
	buf.ReadBitsNext(&mii.EyeHeight, 7)
	buf.ReadBitsNext(&mii.EyeDistance, 4)
	buf.ReadBitsNext(&mii.EyeRotation, 5)

	buf.ReadBitsNext(&mii.EyebrowThickness, 4)
	buf.ReadBitsNext(&mii.EyebrowScale, 4)
	buf.ReadBitsNext(&mii.EyebrowColor, 3)
	buf.ReadBitsNext(&mii.EyebrowType, 5)
	buf.ReadBitsNext(&mii.EyebrowHeight, 7)
	buf.ReadBitsNext(&mii.EyebrowDistance, 4)
	buf.ReadBitsNext(&mii.EyebrowRotation, 5)

	buf.ReadBitsNext(&mii.NoseHeight, 7)
	buf.ReadBitsNext(&mii.NoseScale, 4)
	buf.ReadBitsNext(&mii.NoseType, 5)

	buf.ReadBitsNext(&mii.MouthThickness, 3)
	buf.ReadBitsNext(&mii.MouthScale, 4)
	buf.ReadBitsNext(&mii.MouthColor, 3)
	buf.ReadBitsNext(&mii.MouthType, 6)

	buf.AlignByte()

	buf.ReadBytesNext(&tmp2, 1)
	mii.Unknown8 = tmp2[0]

	buf.AlignBit()

	buf.ReadBitsNext(&mii.MustacheType, 3)
	buf.ReadBitsNext(&mii.MouthHeight, 5)
	buf.ReadBitsNext(&mii.MustacheHeight, 6)
	buf.ReadBitsNext(&mii.MustacheScale, 4)
	buf.ReadBitsNext(&mii.BeardColor, 3)
	buf.ReadBitsNext(&mii.BeardType, 3)

	buf.ReadBitsNext(&mii.GlassesHeight, 5)
	buf.ReadBitsNext(&mii.GlassesScale, 4)
	buf.ReadBitsNext(&mii.GlassesColor, 3)
	buf.ReadBitsNext(&mii.GlassesType, 4)
	buf.ReadBitNext(&mii.Unknown9)

	buf.ReadBitsNext(&mii.MoleY, 5)
	buf.ReadBitsNext(&mii.MoleX, 5)
	buf.ReadBitsNext(&mii.MoleScale, 4)

	buf.ReadBitNext(&tmp1)
	mii.MoleEnabled = bitToBool(tmp1)

	buf.AlignByte()

	buf.ReadBytesNext(&tmp2, 20)
	mii.AuthorName = utils.DecodeUTF8StringFromBytes(tmp2)

	buf.ReadBytesNext(&mii.Unknown10, 2)

	tmp3 = []uint16{0x00}
	buf.ReadU16LENext(&tmp3, 1)
	mii.Checksum = tmp3[0]

	// TODO: add proper checksum validation

}

// Encode takes a Mii and encodes it as a byte array
// TODO: potentially hardcode offsets for `Seek` calls
// TODO: switch to `MiniBuffer` instead of using `Buffer`
func (mii *Mii) Encode() []byte {

	buf := crunch.NewBuffer(make([]byte, 0x60))

	/*buf.SetBitsNext(mii.BirthPlatform, 4)
	buf.SetBitsNext(mii.Unknown1, 4)
	buf.SetBitsNext(mii.Unknown2, 4)
	buf.SetBitsNext(mii.Unknown3, 4)
	buf.SetBitsNext(mii.FontRegion, 4)
	buf.SetBitsNext(mii.RegionMove, 2)
	buf.SetBitNext(mii.Unknown4)
	buf.SetBitNext(boolToBit(mii.Copyable))

	buf.AlignByte()

	buf.WriteByteNext(mii.MiiVersion)
	buf.WriteBytesNext(mii.AuthorID)
	buf.WriteBytesNext(mii.MiiID)
	buf.WriteBytesNext(mii.Unknown5)

	buf.AlignBit()

	buf.SetBitNext(mii.Unknown6)
	buf.SetBitNext(mii.Unknown7)
	buf.SetBitsNext(mii.Color, 4)
	buf.SetBitsNext(mii.BirthDay, 5)
	buf.SetBitsNext(mii.BirthMonth, 4)
	buf.SetBitNext(mii.Gender)

	buf.AlignByte()

	buf.WriteBytesNext(utils.EncodeBytesFromUTF8String(mii.MiiName))
	buf.WriteByteNext(mii.Size)
	buf.WriteByteNext(mii.Fatness)

	buf.AlignBit()

	buf.SetBitsNext(mii.BlushType, 4)
	buf.SetBitsNext(mii.FaceStyle, 4)
	buf.SetBitsNext(mii.FaceColor, 3)
	buf.SetBitsNext(mii.FaceType, 4)
	buf.SetBitNext(boolToBit(mii.LocalOnly))
	buf.SetBitsNext(mii.HairMirrored, 5)
	buf.SetBitsNext(mii.HairColor, 3)

	buf.AlignByte()

	buf.WriteByteNext(mii.HairType)

	buf.AlignBit()

	buf.SetBitsNext(mii.EyeThickness, 3)
	buf.SetBitsNext(mii.EyeScale, 4)
	buf.SetBitsNext(mii.EyeColor, 3)
	buf.SetBitsNext(mii.EyeType, 6)
	buf.SetBitsNext(mii.EyeHeight, 7)
	buf.SetBitsNext(mii.EyeDistance, 4)
	buf.SetBitsNext(mii.EyeRotation, 5)

	buf.SetBitsNext(mii.EyebrowThickness, 4)
	buf.SetBitsNext(mii.EyebrowScale, 4)
	buf.SetBitsNext(mii.EyebrowColor, 3)
	buf.SetBitsNext(mii.EyebrowType, 5)
	buf.SetBitsNext(mii.EyebrowHeight, 7)
	buf.SetBitsNext(mii.EyebrowDistance, 4)
	buf.SetBitsNext(mii.EyebrowRotation, 5)

	buf.SetBitsNext(mii.NoseHeight, 7)
	buf.SetBitsNext(mii.NoseScale, 4)
	buf.SetBitsNext(mii.NoseType, 5)

	buf.SetBitsNext(mii.MouthThickness, 3)
	buf.SetBitsNext(mii.MouthScale, 4)
	buf.SetBitsNext(mii.MouthColor, 3)
	buf.SetBitsNext(mii.MouthType, 6)

	buf.AlignByte()

	buf.WriteByteNext(mii.Unknown8)

	buf.AlignBit()

	buf.SetBitsNext(mii.MustacheType, 3)
	buf.SetBitsNext(mii.MouthHeight, 5)
	buf.SetBitsNext(mii.MustacheHeight, 6)
	buf.SetBitsNext(mii.MustacheScale, 4)
	buf.SetBitsNext(mii.BeardColor, 3)
	buf.SetBitsNext(mii.BeardType, 3)

	buf.SetBitsNext(mii.GlassesHeight, 5)
	buf.SetBitsNext(mii.GlassesScale, 4)
	buf.SetBitsNext(mii.GlassesColor, 3)
	buf.SetBitsNext(mii.GlassesType, 4)
	buf.SetBitNext(mii.Unknown9)
	buf.SetBitsNext(mii.MoleY, 5)
	buf.SetBitsNext(mii.MoleX, 5)
	buf.SetBitsNext(mii.MoleScale, 4)
	buf.SetBitNext(boolToBit(mii.MoleEnabled))

	buf.AlignByte()

	buf.WriteBytesNext(utils.EncodeBytesFromUTF8String(mii.AuthorName))
	buf.WriteBytesNext(mii.Unknown10)
	buf.WriteBytes(0x00, swapMiiEndiannessToBig(buf.Bytes()))
	buf.WriteU16LENext([]uint16{utils.CRC16(buf.Bytes())})*/

	return buf.Bytes()

}
