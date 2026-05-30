# TOTP

### What it TOTP?
Time-based one-time password

### Why do we need it?
TOTP is commonly used for Multi-Factor Authentication(MFA).

### How does it work(simplified)?  
**Step 1:** pre-shared values * current time = new value  
**Step 2:** Calculate Counter(Sliding Window/Offset)  
**Step 3:** Generate HMAC SHA1 Hash  
**Step 4:** Dynamic Truncation(6 or 8 digits)  

### Why is it secure?  
- Offline - No internet required  
- Replay Attack Resistance - rotates every 30 seconds  
- Length is configurable(6-8 digits)  

### Attacks against TOTP
- Phishing/MITM: Proxy creds & TOTP to real server for session hijacking
- Code Interception: Malware/Spyware on device can intercept codes from TOTP apps
- Seed Exposure: Exposed Secret shared between server & client
- Time Traveler Attack: Allows an attacker w/ access to the TOTP device/hardware to generate a token for future time by changing hardware time.

## References  
- https://leapcell.io/blog/understanding-bitwise-operations-in-go
- https://www.loginradius.com/blog/engineering/what-is-totp-authentication
- https://www.beyondidentity.com/phishing-101/totp
- https://www.youtube.com/watch?v=C0pM6TIyvXI