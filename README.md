# TOTP

**What it TOTP?**
Time-based one-time password

**Why do we need it?**
TOTP is commonly used for Multi-Factor Authentication(MFA).

**How does it work(simplified)?**
**Step 1:** pre-shared values * current time = new value
**Step 2:** Calculate Counter(Sliding Window/Offset)
**Step 3:** Generate HMAC SHA1 Hash
**Step 4:** Dynamic Truncation(6 or 8 digits)

**Why is it secure?**
- Offline - No internet required
- Replay Attack Resistance - rotates every 30 seconds
- Length is configurable(6-8 digits)


## References
- https://leapcell.io/blog/understanding-bitwise-operations-in-go
- https://www.loginradius.com/blog/engineering/what-is-totp-authentication