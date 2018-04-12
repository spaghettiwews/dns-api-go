//
// This file was generated via github.com/skx/implant/
//
// Local edits will be lost.
//
package main

import (
	"bytes"
	"compress/gzip"
	"encoding/hex"
	"errors"
	"io/ioutil"
)

//
// EmbeddedResource is the structure which is used to record details of
// each embedded resource in your binary.
//
// The resource contains the (original) filename, relative to the input
// directory `implant` was generated with, along with the original size
// and the compressed/encoded data.
//
type EmbeddedResource struct {
	Filename string
	Contents string
	Length   int
}

//
// RESOURCES is a simple array containing one entry for each embedded
// resource.
//
// It is exposed to callers via the `getResources()` function.
//
var RESOURCES []EmbeddedResource

//
// Populate our resources
//
func init() {

	var tmp EmbeddedResource

	tmp.Filename = "data/favicon.ico"
	tmp.Contents = "1f8b08000000000004ff6ccab10d85300c45d1ebbf405cfd9a928e15d88181602446cb043c640c518a1ce9baf103c370ff114ee00fcc800313f90f6b9c57aed3b6c4325d3bad5a18d66f3e124f76d02a9561fd460209ee000000fffffdf4ae2dc6000000"
	tmp.Length = 198
	RESOURCES = append(RESOURCES, tmp)

	tmp.Filename = "data/humans.txt"
	tmp.Contents = "1f8b08000000000004ff04c0b1cac2301007f039f71437978f76cff449b108e264371109fa5743d33bc95de3ebfb1b3a9ef7bb13770385f18d673c3b1af888f5436154f174f7c8e668e04bf22b9ba3a1d7faeab785c2fccdeea891ffcdb5406ea65b310a53d5351e502ccb92ff78ca52923ce8170000ffffab0427536e000000"
	tmp.Length = 110
	RESOURCES = append(RESOURCES, tmp)

	tmp.Filename = "data/index.html"
	tmp.Contents = "1f8b08000000000004ffe458db72db38d2becf53f4cfa4e6e637491f3293c4a158a5b16762e7e0782d4d26d99b2d106889b0403403803aaccbefbe05909229d9d971b66a77ab66a10b0124bad1fda1fb4383d9ff9d7e3c197fb9fc054a57a9fc49e6ff40313d1d44a8a3fc0940562213be039055e818f092198b6e10356e12bf8cd2ee9d927a0606d520628d2bc944501a9c0ca2b46c2aa66de2962e82f5e4a048b30a079140cb8dac9d241d0127ed50bb4134bc3c87091938bd1881229a35b57d507a86ab0519617ba242db3db8b6a4f780d5f24eca49a730bfb949cec83abff6ed2dc430d4405a498dbda5c0a2994b8e59daca3c818df35e6e10cd252e6a32aeb7ea420a570e047ab9380cf6406ae92453b1e54ce1e020c009d022d582533a57dbe334add8920b9d1444ce3ac36a3fe054a59b07e9517294bc48b9b577cf924aea845b1bb5b05bb752684b441781d40ea746bad520b2253b7af93cfef9d3172947e7bfe2bb03f1a67a7b359cad7873363cbb9a1e1d7eac7ee38bc50bd247575fc4f4f927f6ff97d5686cff9ebefbe9e5bc10bf5c97cf9b08b8216bc9c8a9d4838869d2ab8a1abb76aadd44b0860fa234e52430b9fedaa059053fda6e7c901c1c26cf83ddd736cab3b495ea006e07ad8aef05e67a17976bfb100a63fee3f95f64b17ff8e2eb7c753dfa3039bbfef881bd9f4d9adf3f2dffbafced529fbc1dbe5087d5c9ef17e7f59b57d59b93d3978b3717e7fcf2f4c578c9a26fa2b0eb8bdf0c70ab1a0791c3a5f33bd74115520d6e42500140c1f86c6aa8d122e6a4c81c3f9dbcf2bfd7ed84dbf62fb10ee7184f881c9a3be19a0921f5f4189eefd74bd8ef64005a4df0f4d5abb59e071682ed95000a32024deca83e86c37a09969414f0145ff95fa7fb417b12cde6b1e789be6d9d0d85627cd609034c48bb7881725aba63284889ee4dab364b036e3e1eb2d4eb0bbd82c4ca770032cde6c015b316061069362f9881f62f967a8ec6e27a2870c21ae52230a4b09b2da7aca599b0ca20ead08b154edcf18ffbf5f2f51ad0d8040bc3b3f5b6412664b7faa05bbbf37933c3b3e5419e3168b33b8db6e9264b590e3164b6624ae5ff8c79922c6d27656979d0fa0e90a542cedb41966ab6eef6acf2f4c9a4c63b8bb2f2203fbd18c1fbc0a1305a135b4feb43f2f1443552f4fdeacd32b4e8bdd9c685938a6d151ff8dd8f2b11af49af6d3d0fdad657bb16bd93dd1606c8ea7c5c4abba667604ad1c2c28a1a7004359a0999caf378ac82bb16e692819555ad10cec6e34b30f8b541ebec1e18e428e7524fbdb80183b651ce866c0466e1ede8e3458cdaf398002aae913b9b6469ddc2bf6ec120848aac838259c9a1b16c8a804b16d65c50a34438bf40c919822ba53dded1b283c9eef05fc4fd707b87ee22a76d7db5f771dfdf120eb8671e896de77d7b06bc310ad66cbd75b8a6cf3cfda5cf4ab2ae25b1ae65e178c8ff23386c87d077e2b02d1c701813086939cdd1c0a294bc04ef9e859269a110b0625285c2c595082d6853ff2ce154757e83a08a49ede30e7808908a85d84098908f68a9a7eb40fd3306cb861e1f8e9a0f9fd30d6251fec7733ca9fe69c2ebfc72fe53e08ba6b6b02033f354c29d5a794ef23cd7320a2e6be46ecf3f148436045b4b794013187f1e6ff8cc11fd0fc6d070381ca642db98d5322133fd561cedce7b642c052668937b2bf27ceb18f10f827cfc799c868a2e99c86f19d79ff348c33a7eeb1975732327905ca190f6b6adb1da969547f9157318bf9795749e717e6055fd1a86456331be34e490fb5a294bcba3be93fd8363b70ef89e4a60ab9a59b7bef2fba7d22e1b875d385112b5b3c00c82f2aea0004770b8bfdf2584f5750194d4983d605a00ef0416250521563456ce11c800b912cd425a84a941142b5848a5a0402814f1190a5848578226d0e424f735da4e39102c3a9f8444cd8adca05a6569914353930e67ffa67271256b0f005b86138093b6d257cff7e2662a5dd914fef848ed6cb90eea784a512eb5754c29bf792b6a0cd0424378a6fdd595e5811f7cc901d27a7b34b9600e034e5585864ba636162dbc6f4caf60da30c3b443f454c2e64c2a564825ddcad74c4a6e06648071de18c657f70ba38044a8d582414a85c2dbdb611aadbdc56dfa74391a4f096e6e924f68ac247d7bdb9d94f7f57601be0e987b517473835adcdede65c096406fd0efd6f90fbab0f56b5f903c72102cc8badb981483a8ed46ede5641085dcee6e6b9b626a3bbcf54ea57e3f7b2a111f6d84efb9baad2d14dbfdd90059a3d6e66836879e26df32257b6f3717992b54cce7507751b059aa64ebedba654adedd70caeefb8528799d484aa3fc57665d48b4ab102c0ae174a55925797c7a317a80c61ea7bea372ff4924cadf481717cca208a5beff9a23f543d4bdad3a4b1b75e7496fff7dfbf780390c1fc2e0bdd4b3c70259289ab677fd6422d328ff595170ed91fbe016d2393401a8912385fa6f236a948df271fbc6efc0aeaeff02322754afc2157bd796f079ec7e7c857c4a6658d5e1c01cf98f21f00eabfa3bfde96d7b96b639eb63224b0b12abfc499696ae52f9937f040000ffff0784c8b317150000"
	tmp.Length = 5399
	RESOURCES = append(RESOURCES, tmp)

	tmp.Filename = "data/robots.txt"
	tmp.Contents = "1f8b08000000000004ff04c0510ac0200800d07f4fe177206ebffeae2bec00d6220429a881ecf67bf76e8bb4b7f10a26b89686d3d35c3fc1f3806c5bdd670872ed46c50627c8b6d57d8620d76e546c30fc010000ffffab8f538647000000"
	tmp.Length = 71
	RESOURCES = append(RESOURCES, tmp)

}

//
// Return the contents of a resource.
//
func getResource(path string) ([]byte, error) {
	for _, entry := range RESOURCES {
		//
		// We found the file contents.
		//
		if entry.Filename == path {
			var raw bytes.Buffer

			// Decode the data.
			in, err := hex.DecodeString(entry.Contents)
			if err != nil {
				return nil, err
			}

			// Gunzip the data to the client
			gr, err := gzip.NewReader(bytes.NewBuffer(in))
			defer gr.Close()
			data, err := ioutil.ReadAll(gr)
			if err != nil {
				return nil, err
			}
			_, err = raw.Write(data)
			if err != nil {
				return nil, err
			}

			// Return it.
			return raw.Bytes(), nil
		}
	}
	return nil, errors.New("Failed to find resource")
}

//
// Return the available resources.
//
func getResources() []EmbeddedResource {
	return RESOURCES
}
