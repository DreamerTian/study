package BLC

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"log"
)

const version = byte(0x00)
const addressChecksumLen = 4

//Wallet存储 private 和 public keys
type Wallet struct {
	//1.私钥 使用椭圆曲线加密
	PrivateKey ecdsa.PrivateKey
	//2.通过私钥生成的公钥
	PublicKey []byte
}

//创建一个钱包
func NewWallet() *Wallet {
	private, public := newKeyPair()
	wallet := Wallet{private, public}

	return &wallet
}

//返回钱包地址
func (w *Wallet) GetAddress() []byte {
	//1.hash 160
	ripemd160Hash := w.Ripemd160Hash(w.PublicKey)

	version_ripemd160Hash := append([]byte{version}, ripemd160Hash...)

	checkSumBytes := CheckSum(version_ripemd160Hash)

	bytes := append(version_ripemd160Hash, checkSumBytes...)

	return Base58Encode(bytes)

}

func CheckSum(payload []byte) []byte {

	hash1 := sha256.Sum256(payload)
	hash2 := sha256.Sum256(hash1[:])

	return hash2[:addressChecksumLen]
}

func (w *Wallet) Ripemd160Hash(publicKey []byte) []byte {
	//1.256
	hash256 := sha256.New()
	hash256.Write(publicKey)
	hash := hash256.Sum(nil)

	//2.160
	ripemd160 := ripemd160.New()
	ripemd160.Write(hash)
	return ripemd160.Sum(nil)
}

func IsValidForAddress(address []byte) bool {

	version_public_checksumBytes := Base58Decode(address)

	checkSumBytes := version_public_checksumBytes[len(version_public_checksumBytes)-addressChecksumLen:]

	version_ripemd160 := version_public_checksumBytes[:len(version_public_checksumBytes)-addressChecksumLen]

	checkBytes := CheckSum(version_ripemd160)

	if bytes.Compare(checkSumBytes,checkBytes) == 0{
		return true
	}
	return false
}

//创建私钥 、 公钥
func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()

	private, err := ecdsa.GenerateKey(curve, rand.Reader)

	if err != nil {
		log.Panic(err)
	}

	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pubKey
}
