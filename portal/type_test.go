package portal

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/ledgerwatch/erigon-lib/common/hexutil"
	"github.com/ledgerwatch/erigon-lib/common/hexutility"
	"github.com/ledgerwatch/erigon-lib/types/ssz"
	"github.com/stretchr/testify/assert"
)

func TestForkedLightClientUpdateRange(t *testing.T) {
	filePath := "testdata/light_client_updates_by_range.json"

	f, _ := os.Open(filePath)
	jsonStr, _ := io.ReadAll(f)

	var result map[string]interface{}
	_ = json.Unmarshal(jsonStr, &result)

	fmt.Println(result)

	for _, v := range result {
		decode, err := hexutil.Decode(v.(map[string]interface{})["content_value"].(string))
		if err != nil {
			return
		}
		var range1 ForkedLightClientUpdateRange
		range1, err = DecodeDynamicListForkedObject[*ForkedLightClientUpdate](decode, 0, uint32(len(decode)), 100)
		if err != nil {
			return
		}
		//fmt.Println(k, v)
		fmt.Println(range1)

		var buf []byte
		buf, err = ssz.EncodeDynamicList(buf, range1)
		hex := hexutility.Encode(buf)

		assert.Equal(t, v.(map[string]interface{})["content_value"], hex)
	}
}
