// This file is auto-generated by jwt/internal/cmd/gentoken/main.go. DO NOT EDIT
package jwt

import (
	"bytes"
	"context"
	"github.com/lestrrat-go/jwx/internal/json"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/lestrrat-go/iter/mapiter"
	"github.com/lestrrat-go/jwx/internal/iter"
	"github.com/lestrrat-go/jwx/jwt/internal/types"
	"github.com/pkg/errors"
)

const (
	AudienceKey   = "aud"
	ExpirationKey = "exp"
	IssuedAtKey   = "iat"
	IssuerKey     = "iss"
	JwtIDKey      = "jti"
	NotBeforeKey  = "nbf"
	SubjectKey    = "sub"
)

// Token represents a generic JWT token.
// which are type-aware (to an extent). Other claims may be accessed via the `Get`/`Set`
// methods but their types are not taken into consideration at all. If you have non-standard
// claims that you must frequently access, consider creating accessors functions
// like the following
//
// func SetFoo(tok jwt.Token) error
// func GetFoo(tok jwt.Token) (*Customtyp, error)
//
// Embedding jwt.Token into another struct is not recommended, becase
// jwt.Token needs to handle private claims, and this really does not
// work well when it is embedded in other structure
type Token interface {
	Audience() []string
	Expiration() time.Time
	IssuedAt() time.Time
	Issuer() string
	JwtID() string
	NotBefore() time.Time
	Subject() string
	PrivateClaims() map[string]interface{}
	Get(string) (interface{}, bool)
	Set(string, interface{}) error
	Iterate(context.Context) Iterator
	Walk(context.Context, Visitor) error
	AsMap(context.Context) (map[string]interface{}, error)
}
type stdToken struct {
	audience      types.StringList       // https://tools.ietf.org/html/rfc7519#section-4.1.3
	expiration    *types.NumericDate     // https://tools.ietf.org/html/rfc7519#section-4.1.4
	issuedAt      *types.NumericDate     // https://tools.ietf.org/html/rfc7519#section-4.1.6
	issuer        *string                // https://tools.ietf.org/html/rfc7519#section-4.1.1
	jwtID         *string                // https://tools.ietf.org/html/rfc7519#section-4.1.7
	notBefore     *types.NumericDate     // https://tools.ietf.org/html/rfc7519#section-4.1.5
	subject       *string                // https://tools.ietf.org/html/rfc7519#section-4.1.2
	privateClaims map[string]interface{} `json:"-"`
}

type stdTokenMarshalProxy struct {
	Xaudience   types.StringList   `json:"aud,omitempty"`
	Xexpiration *types.NumericDate `json:"exp,omitempty"`
	XissuedAt   *types.NumericDate `json:"iat,omitempty"`
	Xissuer     *string            `json:"iss,omitempty"`
	XjwtID      *string            `json:"jti,omitempty"`
	XnotBefore  *types.NumericDate `json:"nbf,omitempty"`
	Xsubject    *string            `json:"sub,omitempty"`
}

// New creates a standard token, with minimal knowledge of
// possible claims. Standard claims include"aud", "exp", "iat", "iss", "jti", "nbf" and "sub".
// Convenience accessors are provided for these standard claims
func New() Token {
	return &stdToken{
		privateClaims: make(map[string]interface{}),
	}
}

// Size returns the number of valid claims stored in this token
func (t *stdToken) Size() int {
	var count int
	if len(t.audience) > 0 {
		count++
	}
	count += len(t.privateClaims)
	return count
}

func (t *stdToken) Get(name string) (interface{}, bool) {
	switch name {
	case AudienceKey:
		if t.audience == nil {
			return nil, false
		}
		v := t.audience.Get()
		return v, true
	case ExpirationKey:
		if t.expiration == nil {
			return nil, false
		}
		v := t.expiration.Get()
		return v, true
	case IssuedAtKey:
		if t.issuedAt == nil {
			return nil, false
		}
		v := t.issuedAt.Get()
		return v, true
	case IssuerKey:
		if t.issuer == nil {
			return nil, false
		}
		v := *(t.issuer)
		return v, true
	case JwtIDKey:
		if t.jwtID == nil {
			return nil, false
		}
		v := *(t.jwtID)
		return v, true
	case NotBeforeKey:
		if t.notBefore == nil {
			return nil, false
		}
		v := t.notBefore.Get()
		return v, true
	case SubjectKey:
		if t.subject == nil {
			return nil, false
		}
		v := *(t.subject)
		return v, true
	default:
		v, ok := t.privateClaims[name]
		return v, ok
	}
}

func (t *stdToken) Set(name string, value interface{}) error {
	switch name {
	case AudienceKey:
		var acceptor types.StringList
		if err := acceptor.Accept(value); err != nil {
			return errors.Wrapf(err, `invalid value for %s key`, AudienceKey)
		}
		t.audience = acceptor
		return nil
	case ExpirationKey:
		var acceptor types.NumericDate
		if err := acceptor.Accept(value); err != nil {
			return errors.Wrapf(err, `invalid value for %s key`, ExpirationKey)
		}
		t.expiration = &acceptor
		return nil
	case IssuedAtKey:
		var acceptor types.NumericDate
		if err := acceptor.Accept(value); err != nil {
			return errors.Wrapf(err, `invalid value for %s key`, IssuedAtKey)
		}
		t.issuedAt = &acceptor
		return nil
	case IssuerKey:
		if v, ok := value.(string); ok {
			t.issuer = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, IssuerKey, value)
	case JwtIDKey:
		if v, ok := value.(string); ok {
			t.jwtID = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, JwtIDKey, value)
	case NotBeforeKey:
		var acceptor types.NumericDate
		if err := acceptor.Accept(value); err != nil {
			return errors.Wrapf(err, `invalid value for %s key`, NotBeforeKey)
		}
		t.notBefore = &acceptor
		return nil
	case SubjectKey:
		if v, ok := value.(string); ok {
			t.subject = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, SubjectKey, value)
	default:
		if t.privateClaims == nil {
			t.privateClaims = map[string]interface{}{}
		}
		t.privateClaims[name] = value
	}
	return nil
}

func (t *stdToken) Audience() []string {
	if t.audience != nil {
		return t.audience.Get()
	}
	return nil
}

func (t *stdToken) Expiration() time.Time {
	if t.expiration != nil {
		return t.expiration.Get()
	}
	return time.Time{}
}

func (t *stdToken) IssuedAt() time.Time {
	if t.issuedAt != nil {
		return t.issuedAt.Get()
	}
	return time.Time{}
}

func (t *stdToken) Issuer() string {
	if t.issuer != nil {
		return *(t.issuer)
	}
	return ""
}

func (t *stdToken) JwtID() string {
	if t.jwtID != nil {
		return *(t.jwtID)
	}
	return ""
}

func (t *stdToken) NotBefore() time.Time {
	if t.notBefore != nil {
		return t.notBefore.Get()
	}
	return time.Time{}
}

func (t *stdToken) Subject() string {
	if t.subject != nil {
		return *(t.subject)
	}
	return ""
}

func (t *stdToken) PrivateClaims() map[string]interface{} {
	return t.privateClaims
}

func (t *stdToken) iterate(ctx context.Context, ch chan *ClaimPair) {
	defer close(ch)

	var pairs []*ClaimPair
	if t.audience != nil {
		v := t.audience.Get()
		pairs = append(pairs, &ClaimPair{Key: AudienceKey, Value: v})
	}
	if t.expiration != nil {
		v := t.expiration.Get()
		pairs = append(pairs, &ClaimPair{Key: ExpirationKey, Value: v})
	}
	if t.issuedAt != nil {
		v := t.issuedAt.Get()
		pairs = append(pairs, &ClaimPair{Key: IssuedAtKey, Value: v})
	}
	if t.issuer != nil {
		v := *(t.issuer)
		pairs = append(pairs, &ClaimPair{Key: IssuerKey, Value: v})
	}
	if t.jwtID != nil {
		v := *(t.jwtID)
		pairs = append(pairs, &ClaimPair{Key: JwtIDKey, Value: v})
	}
	if t.notBefore != nil {
		v := t.notBefore.Get()
		pairs = append(pairs, &ClaimPair{Key: NotBeforeKey, Value: v})
	}
	if t.subject != nil {
		v := *(t.subject)
		pairs = append(pairs, &ClaimPair{Key: SubjectKey, Value: v})
	}
	for k, v := range t.privateClaims {
		pairs = append(pairs, &ClaimPair{Key: k, Value: v})
	}
	for _, pair := range pairs {
		select {
		case <-ctx.Done():
			return
		case ch <- pair:
		}
	}
}

// this is almost identical to json.Encoder.Encode(), but we use Marshal
// to avoid having to remove the trailing newline for each successive
// call to Encode()
func writeJSON(buf *bytes.Buffer, v interface{}, keyName string) error {
	enc, err := json.Marshal(v)
	if err != nil {
		return errors.Wrapf(err, `failed to encode '%s'`, keyName)
	}
	buf.Write(enc)

	return nil
}

func (t *stdToken) UnmarshalJSON(buf []byte) error {
	var proxy stdTokenMarshalProxy
	if err := json.Unmarshal(buf, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal stdToken`)
	}
	t.audience = proxy.Xaudience
	t.expiration = proxy.Xexpiration
	t.issuedAt = proxy.XissuedAt
	t.issuer = proxy.Xissuer
	t.jwtID = proxy.XjwtID
	t.notBefore = proxy.XnotBefore
	t.subject = proxy.Xsubject
	var m map[string]interface{}
	if err := json.Unmarshal(buf, &m); err != nil {
		return errors.Wrap(err, `failed to parse privsate parameters`)
	}
	delete(m, AudienceKey)
	delete(m, ExpirationKey)
	delete(m, IssuedAtKey)
	delete(m, IssuerKey)
	delete(m, JwtIDKey)
	delete(m, NotBeforeKey)
	delete(m, SubjectKey)
	t.privateClaims = m
	return nil
}

func (t stdToken) MarshalJSON() ([]byte, error) {
	var proxy stdTokenMarshalProxy
	proxy.Xaudience = t.audience
	proxy.Xexpiration = t.expiration
	proxy.XissuedAt = t.issuedAt
	proxy.Xissuer = t.issuer
	proxy.XjwtID = t.jwtID
	proxy.XnotBefore = t.notBefore
	proxy.Xsubject = t.subject
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(proxy); err != nil {
		return nil, errors.Wrap(err, `failed to encode proxy to JSON`)
	}
	hasContent := buf.Len() > 3 // encoding/json always adds a newline, so "{}\n" is the empty hash
	if l := len(t.privateClaims); l > 0 {
		buf.Truncate(buf.Len() - 2)
		keys := make([]string, 0, l)
		for k := range t.privateClaims {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for i, k := range keys {
			if hasContent || i > 0 {
				fmt.Fprintf(&buf, `,`)
			}
			fmt.Fprintf(&buf, `%s:`, strconv.Quote(k))
			if err := enc.Encode(t.privateClaims[k]); err != nil {
				return nil, errors.Wrapf(err, `failed to encode private param %s`, k)
			}
		}
		fmt.Fprintf(&buf, `}`)
	}
	return buf.Bytes(), nil
}

func (t *stdToken) Iterate(ctx context.Context) Iterator {
	ch := make(chan *ClaimPair)
	go t.iterate(ctx, ch)
	return mapiter.New(ch)
}

func (t *stdToken) Walk(ctx context.Context, visitor Visitor) error {
	return iter.WalkMap(ctx, t, visitor)
}

func (t *stdToken) AsMap(ctx context.Context) (map[string]interface{}, error) {
	return iter.AsMap(ctx, t)
}
