package addr

import (
	addresspb "go_protobuf/source/address"
)

//CreateAddress creates an instance of address
func CreateAddress() *addresspb.Address {
	a := addresspb.Address{
		CorrespondanceAddr: &addresspb.Location{
			Location: "loc 1",
			City: &addresspb.City{
				Name:    "Mumbai",
				ZipCode: "400005",
				Region:  addresspb.Division_WEST,
			},
		},

		AdditionalAddr: []*addresspb.Location{
			{
				Location: "loc 2",
				City: &addresspb.City{
					Name:    "Srinagar",
					ZipCode: "190001",
					Region:  addresspb.Division_NORTH,
				},
			},
			{
				Location: "loc 3",
				City: &addresspb.City{
					Name:    "Imphal",
					ZipCode: "795001",
					Region:  addresspb.Division_EAST,
				},
			},
			{
				Location: "loc 4",
				City: &addresspb.City{
					Name:    "Mysore",
					ZipCode: "570001",
					Region:  addresspb.Division_SOUTH,
				},
			},
		},
	}

	return &a
}
