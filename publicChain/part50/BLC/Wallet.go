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
	PublicKey  []byte
}

//创建一个钱包
func NewWallet() *Wallet {
	private, public := newKeyPair()
	wallet := Wallet{private, public}

	return &wallet
}

//返回钱包地址
func (w *Wallet) GetAddress() []byte {
	//先将publickKey 256Hash  160Hash
	pubKeyHash := HashPubKey(w.PublicKey)

	versionedPayload := append([]byte{version}, pubKeyHash...)
	checksum := checksum(versionedPayload)

	fullPayload := append(versionedPayload, checksum...)
	address := Base58Encode(fullPayload)

	return address
}

//对公钥进行Hash生成人肉眼睛能看懂的内容
func HashPubKey(pubKey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubKey)
	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		log.Panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)

	return publicRIPEMD160
}

//验证地址的有效性
func ValidateAddress(address string) bool {
	publicKeyHash := Base58Decode([]byte(address))

	actualChecksum := publicKeyHash[len(publicKeyHash)-addressChecksumLen:]

	version := publicKeyHash[0]

	publicKeyHash = publicKeyHash[1 : len(publicKeyHash)-addressChecksumLen]

	targetChecksum := checksum(append([]byte{version}, publicKeyHash...))

	return bytes.Compare(actualChecksum, targetChecksum) == 0
}

//Checksum 为一个公钥生成 checksum
func checksum(payload []byte) []byte {
	//两次进行Sum256 Hash
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])

	return secondSHA[:addressChecksumLen]
}

//创建私钥 、 公钥
func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()

	private, err := ecdsa.GenerateKey(curve,rand.Reader)

	if err != nil{
		log.Panic(err)
	}

	pubKey := append(private.PublicKey.X.Bytes(),private.PublicKey.Y.Bytes()...)

	return *private,pubKey
}
