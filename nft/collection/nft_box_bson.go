package collection

import (
	bsonenc "github.com/spikeekips/mitum-currency/digest/util/bson"
	"github.com/spikeekips/mitum/util"
	"github.com/spikeekips/mitum/util/hint"
	"go.mongodb.org/mongo-driver/bson"
)

func (nbx NFTBox) MarshalBSON() ([]byte, error) {
	return bsonenc.Marshal(
		bson.M{
			"_hint": nbx.Hint().String(),
			"nfts":  nbx.nfts,
		},
	)
}

type NFTBoxBSONUnmarshaler struct {
	Hint string   `bson:"_hint"`
	NFTs bson.Raw `bson:"nfts"`
}

func (nbx *NFTBox) DecodeBSON(b []byte, enc *bsonenc.Encoder) error {
	e := util.StringErrorFunc("failed to decode bson of NFTBox")

	var u NFTBoxBSONUnmarshaler
	if err := bsonenc.Unmarshal(b, &u); err != nil {
		return e(err, "")
	}

	ht, err := hint.ParseHint(u.Hint)
	if err != nil {
		return e(err, "")
	}

	return nbx.unmarshal(enc, ht, u.NFTs)
}
