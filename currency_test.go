package numeric

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestUSDToCents(t *testing.T) {
	gunit.Run(new(USDToCentsFixture), t)
}

type USDToCentsFixture struct {
	*gunit.Fixture
}

func (this *USDToCentsFixture) assertParsed(input string, expectedCents uint64) {
	cents, err := TryUSDToCents(input)
	this.So(err, should.BeNil)
	this.So(cents, should.Equal, expectedCents)
}
func (this *USDToCentsFixture) assertInvalid(input string) {
	cents, err := TryUSDToCents(input)
	this.So(err, should.NotBeNil)
	this.So(cents, should.Equal, 0)
}

func (this *USDToCentsFixture) TestTryUSDToCents_Valid() {
	this.assertParsed("$0.00 ", 0)
	this.assertParsed(" $0.00", 0)
	this.assertParsed("$0.00", 0)
	this.assertParsed("$0.01", 1)
	this.assertParsed("$0.10", 10)
	this.assertParsed("$1.00", 100)
	this.assertParsed("$10.00", 1000)
	this.assertParsed("$19.95", 1995)
	this.assertParsed("$88.99", 8899)
	this.assertParsed("$888.99", 88899)
	this.assertParsed("$8888.99", 888899)
	this.assertParsed("$7,888.99", 788899)
	this.assertParsed("$66,777,888.99", 6677788899)
}
func (this *USDToCentsFixture) TestTryUSDToCents_Invalid_MissingCharacters() {
	this.assertInvalid("0")
	this.assertInvalid("00")
	this.assertInvalid("000")
	this.assertInvalid("0.00")
	this.assertInvalid("$.")
	this.assertInvalid("$0.")
	this.assertInvalid("$.0")
	this.assertInvalid("$0.0")
	this.assertInvalid("$000")
}
func (this *USDToCentsFixture) TestTryUSDToCents_Invalid_BadSyntax() {
	this.assertInvalid("$ 0.00")
	this.assertInvalid("$0 .00")
	this.assertInvalid("$0. 00")
	this.assertInvalid("$0.0 0")
	this.assertInvalid("0.0$0")
}
func (this *USDToCentsFixture) TestTryUSDToCents_Invalid_IllegalCharacters() {
	this.assertInvalid("$a.00")
	this.assertInvalid("$0.a0")
	this.assertInvalid("$0.0a")
	this.assertInvalid("$0a.00")
	this.assertInvalid("$a0.00")
	this.assertInvalid("$0.00a")
}
func (this *USDToCentsFixture) TestTryUSDToCents_Invalid_IncorrectCommaPlacement() {
	this.assertInvalid("$77,88.99")
	this.assertInvalid("$6,77,888.99")
}

func (this *USDToCentsFixture) TestUSDToCentsPanicsForUnderlyingError() {
	this.So(USDToCents("$1.23"), should.Equal, 123)
	this.So(func() { USDToCents("invalid") }, should.Panic)
}
