//
// Copyright 2020 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package authentication

const commonCaCert = `
-----BEGIN CERTIFICATE-----
MIIFpzCCA4+gAwIBAgIULS2/crhIlrfygVpLKzYHXCHX83YwDQYJKoZIhvcNAQEL
BQAwYzELMAkGA1UEBhMCVVMxETAPBgNVBAgMCE5ldyBZb3JrMQ8wDQYDVQQHDAZB
cm1vbmsxGjAYBgNVBAoMEUlCTSBDbG91ZCBQcml2YXRlMRQwEgYDVQQDDAt3d3cu
aWJtLmNvbTAeFw0yMDAxMDMwOTA2MzBaFw0yOTEyMzEwOTA2MzBaMGMxCzAJBgNV
BAYTAlVTMREwDwYDVQQIDAhOZXcgWW9yazEPMA0GA1UEBwwGQXJtb25rMRowGAYD
VQQKDBFJQk0gQ2xvdWQgUHJpdmF0ZTEUMBIGA1UEAwwLd3d3LmlibS5jb20wggIi
MA0GCSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQDWTqw35HaR3PUT7kC3uMDYnTWb
t+cdNOWYNb0jpZPDqNbbIgEh9s1RzypRmGaeDIC6MPC1fzqpGG76G4k5Ej7OvaEf
wo/JTyaki+ScMZIU9e8dtqxeEgGSdyWcwqQNjyCvK7HRNXTbpfs7xD4Oe9/HyexS
q1wxrSsJFpMCR0WVNpL1Wif5Tmrm1mYV7tqN64TqK+ZTCw5gFSlKSJzCDieuA82L
5HmJ107sY4CfiGmyGUlrhqA210HpXFtI8VEvhIKcUDdaHOm6vkO4W7lDcgSgq42K
AVfqxQ1levGKEMnrvrWkXmQFcnrP2oiUmBhGTheJhsXROuaYss8NekC9cVraxDaf
ZLhwwk0q9Tf8jCoFFSfGIATSQTACr0PuCRZTaNAVnrml0VXJcjLlP28SmAjYTQXF
8pu5tvh/vYCAwgKx4DADgopqkJ+BbXOfJ2FyDtqg7qstGlbaMAfvW6wi4tnouzTc
l9cCMFIoxeeoDH4NkEDQwqJCDug+Ie2w1hVs9Lel8BCDUQv4cfXf7BMqvPKTrRwW
WP27D2xkk2Leh8iDWbdZTiGj8ynYGdGDXr73f08QTLmCBVfcVeeihGBN3HdoR0F6
7AIIF7tbHl0FyyPIvSDRCZF66W4Q5/71Q5XuL2H89bO444vLATNrigqCEHZ9Duey
GJVgr6T3ghpXfaCYHQIDAQABo1MwUTAdBgNVHQ4EFgQUwthwBFAcoy91DauQMPcM
Kc0kJmcwHwYDVR0jBBgwFoAUwthwBFAcoy91DauQMPcMKc0kJmcwDwYDVR0TAQH/
BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAgEAQHCU0CuWCE4d+Wr4YKS7Q7BUZF3D
UBmDNesryt26hgiCI1hccjHma8tG85VjESjLch79CiLwA6TBuyraFLzYb+7SDJSx
cGuNRn6C2i688ibCYzqQpLUi0jiVOKBqHSKmSOD6o2DG3Ze8KyKrWtyiG03227xh
SGfIiM49dAQlU2kmOynlcRhcYeCLoDnmrDFSysGfgLUEJroc60A+VMSKlMkIa9Ts
NgHTZZ90ToGjoDPtXks0uxIMIAUqLkIylJZvinu1nw5S9Uui4ojgxVceq3SBQuS5
DHiUe3MJ25ImnvvcemdtxdfOkyfqC9aBEu+1Ozj7SoIRz09OAFxslfe8TCbh85H1
7X5HfS2Rz9GjK3YO2/5+15DaQZ7Zun7nJXDW5+lu7Sx/Jb9kpFuKbbRIqeDkpAdN
6vR5EW0J1PbaTLixXzu8IzuoXns9mOuYKxCsaDVGxTpcJY7U5qh6ZPBlQpxg9OQV
3jatCPAAIt7GXI5Qt7/P+aJwEBT7Zc9h0ZcUk0UjpcVo24g0962C3B5sqpZ9Cjxi
iiJo5LasgoR098WDzHSm007wHQT3ynztnoV9oowsIbfsp9+TJNZs5O23WNUcUacW
R23yjoF+pckrm/5vJUZTDv2RSvH0W6+26PQidT3ZqIby7EwOC756nPWzoxtVUi85
neyWMGYln+Yy1ws=
-----END CERTIFICATE-----
`

const papKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA4a77hSpSfr8ZRxpfpORosoQ/v0p2QcnxR3VNNYgfLUovTW7Q
axbMxIfO87WxXH4+krD8vm26KJu242qjPo/Iv0FhXOhb4Vfrxy82yPKCsuJvZewP
Vm+T9qvgoMh/OCwA6KNPToiyCOMfzey6zE1hA4DW8Mkoz3EDvyAVFQfFrX+26q8E
6FotQ87nP4ZJsZPRyMJJWW0fcH2XbxMooJ3E/aJJo2Nx9/hhGZTpK/OCfO8tfy4B
G6IpfkcrfW6K6R5CJvMVK9OSw3lj7zKLjYYxsmHvRs7VOV1K3W9McXwyTh1Zr8Ce
4ug9DxNK1HPPC1Vd3dsIte7FQAlWLLMDZKPXRwIDAQABAoIBADVcsowiHa9qcAen
7MWIXFkZfBk2+g9EOJC047SsovKKf/THJvCrS2+MzRpC89Ty5Mi6oIhGWaYQa4YK
SN8aEFVomCTnrBH0QAOw3jtVXCH/+GcZEedWKp5AtbROtAYMIYrgUng3SiFf0r1W
JDT6dt3lMO1uIwApsMEYTOPUGqGTB1i5Kb6qZKiucKu6yLq6JFuH8WnrIYcAacHN
eFYdA31d6Ljbbn7i2OofMdG6YMh+2ZEnNotQvyFSA1CoRuhAtK1bVdsXSDNWxip3
XcfoCT3zm+MZI1PaLeMu4+59VvIMNrx+Wwl8byLJTsceGzcPAEz869o8O2MJjMHj
sLpLnoECgYEA5SHlqHUSwOu2L/MrJSx2ZXAGUVmdTYiYdTrfRogQk0MGRYPZFYBN
fYXgg/A3GS8Qq7kKMb68EmFdLxPZjVMCR+Bt1do8kqRSwKallC2/SbbZNoUCBdvw
Vsoo7LRc7DHCNY89AgahF2LFb2iEh9rQNVbj0c9l+kMAeS9WwoCnB18CgYEA/CWO
gfzK3uD3Y3iS14KOcBXruui9hkX2kQ1cBoOD6pX6JdheOrpBbROjBFevavmOHh94
IpV58MBhhF07Qk1tC1tFGFtI+EffxgWVr1tzn9rPRDXjb5itBhfSWXAupazCFVAk
EQjTXKpyTuzYYBVCB69SQs8j3AyuM+VucuGrQRkCgYAnyJnqhOFLs+F/M9Zy7uRj
um6uY0PnuWbXO/CWe3t9Ri7plPn1PMC3oa4Y2nWGnuBjII6/kmFvQ459ZRHp8ta7
iiEn92t6/qMLpiOrtcG319KhH7j5MXXqa9FtP+e+bulMiWFLX8FKoTCsplYUjI6I
JF0MCdmn+Ug0wl5rCRF3OwKBgQClO+tWbZV4Vw5nRgVcavbyprrSwmAolMKOraWH
szmISf4iPNcLPzFOzJaAawHOZXlnbhHo//Fn/nopJnuF6H0z+vydiHyaD2eqOdCs
mI59zAMVXQcNA40nOAcliylWco7BLsJj01fHGv9Lj1QAKw4ZQR/0HhAOzNR/t7MY
X6A9SQKBgDKtCBOUYWxY5BgdQA9UUF+A283SfiGcaWhD38/AMu2Lesf2n8iN2NgZ
uW8/2Ek1IraCoOVqCF8hEA9FgpgIUhTkHhE04Ker2+I08uRn6PHqu7Z995C+Hu9c
kdokjntfXpgvl4Y16sjunAqZKSzQskjTh7mbwpdBKGtBUJrNDHsB
-----END RSA PRIVATE KEY-----
`

const papCert = `
-----BEGIN CERTIFICATE-----
MIIEbDCCAlSgAwIBAgIRAMjHRLPq9rZwFPOZuCFRosEwDQYJKoZIhvcNAQELBQAw
YzELMAkGA1UEBhMCVVMxETAPBgNVBAgMCE5ldyBZb3JrMQ8wDQYDVQQHDAZBcm1v
bmsxGjAYBgNVBAoMEUlCTSBDbG91ZCBQcml2YXRlMRQwEgYDVQQDDAt3d3cuaWJt
LmNvbTAeFw0yMDAyMTMxMTM0NDJaFw0yMDA1MTMxMTM0NDJaMCkxFTATBgNVBAoT
DGNlcnQtbWFuYWdlcjEQMA4GA1UEAxMHaWFtLXBhcDCCASIwDQYJKoZIhvcNAQEB
BQADggEPADCCAQoCggEBAOGu+4UqUn6/GUcaX6TkaLKEP79KdkHJ8Ud1TTWIHy1K
L01u0GsWzMSHzvO1sVx+PpKw/L5tuiibtuNqoz6PyL9BYVzoW+FX68cvNsjygrLi
b2XsD1Zvk/ar4KDIfzgsAOijT06IsgjjH83susxNYQOA1vDJKM9xA78gFRUHxa1/
tuqvBOhaLUPO5z+GSbGT0cjCSVltH3B9l28TKKCdxP2iSaNjcff4YRmU6Svzgnzv
LX8uARuiKX5HK31uiukeQibzFSvTksN5Y+8yi42GMbJh70bO1TldSt1vTHF8Mk4d
Wa/AnuLoPQ8TStRzzwtVXd3bCLXuxUAJViyzA2Sj10cCAwEAAaNVMFMwDgYDVR0P
AQH/BAQDAgWgMAwGA1UdEwEB/wQCMAAwHwYDVR0jBBgwFoAUwthwBFAcoy91DauQ
MPcMKc0kJmcwEgYDVR0RBAswCYIHaWFtLXBhcDANBgkqhkiG9w0BAQsFAAOCAgEA
MRDbbUbTw1r1XLxTS00FWWKMt0eVuweuTfEa6pWxgleNj0RQ4Dboev1FJq2gSGYy
BrUlp3dRpnSLP6tZg7BmovGKGp3K3sRL3FyKXxnx5AEIFrVeNGegqXmKPy1/oP6j
F7R1HHFKZ/nqiLoCzYO7lENvWNrMXa06RuULRHnOXvUcUvVivvQLDandNXWqRE2e
TREUuj6oksVdaI1WxP79Kl4p6lJLndXfZiD1yVMlfy96VnelyqZ4/wOx/VRh9az3
+TuACF20j/PKezcrXi01Da2gM8C7s8BYmO+khqYEK+UTNubydixQ5plqJat0hUrj
ZZE+/HUzBR/D5d+pgZtFHscC7o0pPtMQQc3FMnDUt3hABcsLkj/Lz3YJulnwnhG2
aZWjYGbxXzNH+vgpRxcxqCGKEqJ/ea0GY99WxOCk9rqnqBXjKzHgom1xOFl+A95+
/KlEBeRIVq0yL4SaPDN/dlWFHlp4oZxgYceyZSOGcN8GBv2zB8LobAGz8qhl4Tke
NuqMUU2sJl7LoYL+OaqRAF89xC0lvAmRc+nQzI9xnLIogDQw03iPKAIMg2KFJbtm
xXMl9HcB1mbOElgeNm+h2ObqFcdqIB9KT6spXiVpuTukC+OQ1KrFdVocg3NQeo9Y
veLyxtkzoLa3IXewcXZUDDMhNQXonBIqCVXx5bhQfxk=
-----END CERTIFICATE-----
`

const pdpCert = `
-----BEGIN CERTIFICATE-----
MIIEbDCCAlSgAwIBAgIRANO6/S0GHATBJGhwMxcEACcwDQYJKoZIhvcNAQELBQAw
YzELMAkGA1UEBhMCVVMxETAPBgNVBAgMCE5ldyBZb3JrMQ8wDQYDVQQHDAZBcm1v
bmsxGjAYBgNVBAoMEUlCTSBDbG91ZCBQcml2YXRlMRQwEgYDVQQDDAt3d3cuaWJt
LmNvbTAeFw0yMDAyMTMxMDM5MzZaFw0yMDA1MTMxMDM5MzZaMCkxFTATBgNVBAoT
DGNlcnQtbWFuYWdlcjEQMA4GA1UEAxMHaWFtLXBkcDCCASIwDQYJKoZIhvcNAQEB
BQADggEPADCCAQoCggEBALJUlVOwZcxvOZj4Fa19XXzSMD2dbpbwfZRRiNVSwhly
JRvk5Q5PIxfJjYapaKGo7e8cN9O+7oZeLsNqCRC8CE2JPrmqJg2/6kQEkZwrs73L
nIebv4pgZFRvDTgQxx8xh1RWNiMnfUw5nyQXUc7ZfdscDCG3uzWDbH2qKVTbLH2V
NPC1WgNku3jIov07MN5xUNGEA2bEelgmnD+ABZSEw93s947vDlAaDhZ+8Yv1xUqC
4VFfRQHoBiHcH3Ie0963CjbCUH1Afcgb5jAmU2eosFXsv68WTrcNQ8eYdojnTq6H
MVsQvc5gAfPaBpcoVthnpEATFQvERZ7MVUR8WZcUghUCAwEAAaNVMFMwDgYDVR0P
AQH/BAQDAgWgMAwGA1UdEwEB/wQCMAAwHwYDVR0jBBgwFoAUwthwBFAcoy91DauQ
MPcMKc0kJmcwEgYDVR0RBAswCYIHaWFtLXBkcDANBgkqhkiG9w0BAQsFAAOCAgEA
r/dINpVoPLRRJhIENaqw8qbWr4tzBbMgIW3mfeVG4C/t07dWtFMJCEcGJs5JeoH2
5IK6zXdrFPZ4c2VSRISutRyLhg9s6y2BAUvHkicp/h5WVxEc1W+Qpx53Z+ZsVvfp
/SbFtDCnZEwgaN1ZlVUAobJF1peX7ZrqEgEJ2mcnp/KFQxKOAV09SU6BJH5nuS1/
PH/XMNCxC77UCMO+PvVhAMnR4C350W0PYK1tNO9z9jrmcMbyqV0zti8XqmAapm+a
Q/nrjygH0FUzMXNnOgt5GYwkIarlYs/z2wfHm0toSm7TmbEjCts8kD7koxAdv/jS
r32GG/6I7ZlsFXTEN+YOt07JHr85GnesLYQkLArg8+PdT/qYruS6Fy3AhsAoLOFt
Th+UmMqssPi+eZdGs4McTpkwX6oZK+KcXoBUmXAssKQ4P9F/SE6lOjZa8Da0Mpmd
H/7swAWjKNcZycu4UMlqcYgPunoL5R66pcOiP9wN0szqkw67BZ2EWxFaOBqlaNS7
VZ2sT5Sj89g1gAsmaKn9SSUajS0G+snV9X7OWuIcvNeh7uTAFg2ROWyJVccV7TD2
qBngvi+0Em8PRxsH6Fw8m/Lb82K/OXOzW9gofWawU15MVc0qC6/cODSYF1MuyZpT
M7BIG+QxXIjoq5JOSknB0lMpR0fOuS6qzi5WdY5i2AU=
-----END CERTIFICATE-----
`

const pdpKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAslSVU7BlzG85mPgVrX1dfNIwPZ1ulvB9lFGI1VLCGXIlG+Tl
Dk8jF8mNhqlooajt7xw3077uhl4uw2oJELwITYk+uaomDb/qRASRnCuzvcuch5u/
imBkVG8NOBDHHzGHVFY2Iyd9TDmfJBdRztl92xwMIbe7NYNsfaopVNssfZU08LVa
A2S7eMii/Tsw3nFQ0YQDZsR6WCacP4AFlITD3ez3ju8OUBoOFn7xi/XFSoLhUV9F
AegGIdwfch7T3rcKNsJQfUB9yBvmMCZTZ6iwVey/rxZOtw1Dx5h2iOdOrocxWxC9
zmAB89oGlyhW2GekQBMVC8RFnsxVRHxZlxSCFQIDAQABAoIBAQCvufG8GiL2JNQZ
DL4vy02IZobq8Gu9iRX2RbIUpAxojrZKmm+sfgcStaYkgGN+ibyM1r2chpT9C2oB
Xke4mdGEti19P1FoUylCBU6HGflwmWeRqcRBU2MKYLQh/0Z4UpIMZTfLxGoP5ugO
5HvU3TTL6QN7ZlWwEeWinz9DQwbAZ9SKTIsMnmKclm7hoxl2E6ciIwyIwb1zwi6F
32xbK90Tv7L9g3xp6lVHKUBVLmNNlt67/eC1KG/5QjTBzgjpWKdQ+FxhrmDu2KCY
N6D2VO0CE2CYgiNWfy2T+3VVU9cMi3TVjCNgltNaPFKLvQRLypSzKUcN8E+glSXB
xTBQptQxAoGBAMNI64gy8OObN6JZ2AKI0xgJftlqTJG9r7Q73S18xlu5xUEuSrt4
brWJIO91KV+hAstPR9/sUhtF6Vv45iUDJazQVj9PfNQ6xXRBwOABx7Azbp3Uj+vE
J0a/HflC5RlDvulblbnpwB4+Hh03ngwjzBNZo180zZ4zyPi62KyQYsZXAoGBAOnG
Omjrk7ObMBM0p68NHl4dpsLHPuFuLeUUoFiIzGAOWYGWSjBRVvY8uldm9QTgFcCY
Q5Vm/MHEzijDdyVm4AZ8OdAo6s0Y2eFRg13vxhAzfTfEeU5A+w+m2NxgX31jKVD8
CS3cn895Gc895RW59KkfJYKIJPK7XqssYHPnij9zAoGBAKZIZSXoGm30MP8w+VnP
H9AL5dNDTEec+QvDHaYp5M0d9fR2cnQHLF2vkjfTz5L6CuMLcuwc0h0e/oGuAmnp
sB7il035ZRVhpdhEVPNpEJvcb4g9Av+CWt9GxUMFzwYgRWN29JsFiu01f3bpcM5t
fMvNO+tJjvvypGIG15C4sSKpAoGBAISPHQvlmWlc9FPPpc4DWFbx5V/Jb4SQ34Fo
wsuxNoK5YFO1Dvh7Bx61rMxm20Uac/sz3A7ZyTVvDZ/bxrH0tc+3nWjF6u7FyraT
bDRWBC+nord+t3KCMlrbt1ivZTZKhmEfqWzSD+SrvNDhPgel3HAR80kjPiBN0w6p
KUBP6Pq1AoGAfLe6HZUWriqXfgVAYcsINmmQOqh92P1KWutpXoA6NhR3Q5YqUdD2
LAf395xEPxgZpHUynCEa/60qVUB+ItL18m1cjJ5rt8txY+cAOpbcKHgbxkvyeVhd
lKwA6NpvCf9+llmU0jHFFzC+cMCtlW/+3vMXOH33yVVG38wODnCn728=
-----END RSA PRIVATE KEY-----
`

const providerCert = `
-----BEGIN CERTIFICATE-----
MIIEkTCCAnmgAwIBAgIQL8jpSo9OhkzAyrD6PK5LKDANBgkqhkiG9w0BAQsFADBj
MQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3IFlvcmsxDzANBgNVBAcMBkFybW9u
azEaMBgGA1UECgwRSUJNIENsb3VkIFByaXZhdGUxFDASBgNVBAMMC3d3dy5pYm0u
Y29tMB4XDTIwMDEwMzA5MTI1MFoXDTIwMDQwMjA5MTI1MFowPDEVMBMGA1UEChMM
Y2VydC1tYW5hZ2VyMSMwIQYDVQQDExpwbGF0Zm9ybS1pZGVudGl0eS1wcm92aWRl
cjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAN3EgTnqWIv9VHUsTRDz
f/RmdqhgMNKoyVLKdeIb9BAkSyDtAGrw4s3nU26cKRv/Rg9+BDDtadHW5LmANSLB
CDRB22bVKC8Sd7bEakeI47162vA1joPx8rK57Z49fkJJczYhipp/7vMKmdVrCF4Z
BAX8MtYW/1INKSYl/cHj4jRTQ4LybWwzh0djXfHWVZ16CkBaNwfXC1GpHGfdRCCB
U+wGTBtMTPfQbZWldBYGcyHYZycjBkkeQOzMI8EFi36bLhtk8dW5a2Vdg2SY02Ia
MS+1lO0inWUGafr0g7prrraGqf0ru//kavCIz6QlNR/TZ73Nt2yY5b9uWiEp+Zqf
sXMCAwEAAaNoMGYwDgYDVR0PAQH/BAQDAgWgMAwGA1UdEwEB/wQCMAAwHwYDVR0j
BBgwFoAUwthwBFAcoy91DauQMPcMKc0kJmcwJQYDVR0RBB4wHIIacGxhdGZvcm0t
aWRlbnRpdHktcHJvdmlkZXIwDQYJKoZIhvcNAQELBQADggIBALLjSwAPxVJvBvBR
tQDAG69zViTkpQly4Yxf67SjaKGcKDp1aKI2+yef50nvVJKbSwx/vclbtA72gTSF
apENpxJw4Wfy+2Makha6Kxqmv3D+4g34LGJEAzij0G74uB6J6x3xZ2EoNMLfBLAf
atJwn8sp4jI+n0hOODXX9wMWU8HKrxuQ9y89TrJrOTNSfeB2/ZemAL89vaiKSbip
b8UkM0pZVW/qFm3m/bODCWRXDvjMBpLW+W94DwVvYqbQJqVyg4CcWMWV0kQIUZr3
2OeJeiiaDndIFTQkcyrLbdGgp8711t+yqStuYyF0sjbbM+lzwtaAR19MOae06XhO
O1KEejgOjak5SHw/0HYV4eo6ekTWLdZSNrv4fmR59w8cNFk8MQUiUlbKckTDb3GA
kx5UNFfq8kRzukpp1KI1Yqydf3ILMhGdPAgj7+udoaK5oY5X2/h10Oq7KcjM0U8M
bafbQFqomJ9agWiR8neTz6etUdWPHzVTNoaC2J3u/D/MHZ8DzDiDa4RuEa4ToXis
vCAaLYLrgNoZADvMTjF/uXEYdicBysvO0kIENB/14bDtrH/0garW2k33RWyOwNsH
9OGK3gd+qO0IbGEb0prCWfIMfYtEMB4jF5oN/A5ujzjuVbRhOBF1rT+I16hZEFdq
Vmo3JP8DaE4Ccc8EBRDZWaUlf0H7
-----END CERTIFICATE-----
`

const providerKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEA3cSBOepYi/1UdSxNEPN/9GZ2qGAw0qjJUsp14hv0ECRLIO0A
avDizedTbpwpG/9GD34EMO1p0dbkuYA1IsEINEHbZtUoLxJ3tsRqR4jjvXra8DWO
g/Hysrntnj1+QklzNiGKmn/u8wqZ1WsIXhkEBfwy1hb/Ug0pJiX9wePiNFNDgvJt
bDOHR2Nd8dZVnXoKQFo3B9cLUakcZ91EIIFT7AZMG0xM99BtlaV0FgZzIdhnJyMG
SR5A7MwjwQWLfpsuG2Tx1blrZV2DZJjTYhoxL7WU7SKdZQZp+vSDumuutoap/Su7
/+Rq8IjPpCU1H9Nnvc23bJjlv25aISn5mp+xcwIDAQABAoIBAFpAMVMPgu+drjRt
TgO7BPTCikGlu9jPZfcye90EAURccNK6gPQEili/bp3SIYzKS4ncf/oldG44ZSb0
/SrioeMadh6YWt0lR+DyeEoadIxVJiqhsRkTavC/Z+uBHRP+7ks2RUDxFZ6FTuui
vmoFBjUlWM711vxebMOgrG+uvQ0Ln0ezsp+AePmgaXPpNA/k1Bczqd+bBK/FdZ4H
NY/uq/YAcBNn4pKfKZyia155jZQbiDxZfvHHv2ft9t8ZO1eEaH7a1Kb6kD9OOGBU
5l5RS2K6b5V3IhjGOKdSFUnYUXw4qJkX3LDaKu5eS4SF7/rIWqhLNnWoyNlj7SLu
noG12vECgYEA3o07Rsbip8mFKIMcnseOtzMQjQhEmcTpH0NbCT81oMnkx7qFBYkA
3L9Pd+PK3SD+h2SO3RoplBosXFsG5IsdXdfURqgnC886ORdQ9LgxQGiD4g3/hO3J
xggTbZC2q/G3XDdOgvXWOfI4YavxidFVhDr4QfnrnJ+0kM48CfhpP6kCgYEA/xka
54i7GCjv8iT0g7ahHniq+OoD1jN5NMopBqnGrQ0RcGBbQiYa6777ifxC8BaDgQZM
Xw6Ldvz6fuQlYyGHKYvMwnc1fLMmQlvR2akskHyAM4ycx3mayZQ7Rvs01vTR3koW
fTq2FuFpPh0LvpLIDTg8af0X+h4TsxXxe7g/SbsCgYAT6b8sDg+pw7ZOcazV6DU0
3zTT6bF5sMzLJ1O1+BnsUSRPWrkcTa3uEpAhwhgcwR8FLlI6JmUiA1WEzgBH4sq1
4uHzkIgt3lTvVs6/ltuRT8i9KVN7qWssdPyICTLPH8mwTqKPxT9t3+Z7ls797eUg
98XO4XXbhc66RjnJCpkgQQKBgE+im7aKBbYuGXHjPX3cxoGHtAAOtfHpkb38GkCd
QYuGV9pYCkUlNiRAM7BI0vC2Zokh2TeXh8w2f1CbklxW/CXpIZvXsRfjqV7v4Rvs
5x+X7kGpdMwE12B0aOjoxP8R4G5ffJPNf0i3R65/TMpKci9GTf3cZ97KAuZWRFOW
GhO/AoGAZ7Qw/ajO5Ny623ZuGX8/C78PV6xxeL3XXbeSs0rvaIySZQ6hVmStMhgs
Df94gj0SUz+DR7szeidI656Pl/JfQRLIN8U73F0t32DqLI5gJPSZWGxgwBC3X2jZ
RzwWwNYWcO620+ilij0tvmPOb+F/hAk+2nbEW5zqlOoQQEKSqbA=
-----END RSA PRIVATE KEY-----
`

const authCert = `
-----BEGIN CERTIFICATE-----
MIIEnzCCAoegAwIBAgIQH2k6Qg4keqrSyaXIfk01xzANBgkqhkiG9w0BAQsFADBj
MQswCQYDVQQGEwJVUzERMA8GA1UECAwITmV3IFlvcmsxDzANBgNVBAcMBkFybW9u
azEaMBgGA1UECgwRSUJNIENsb3VkIFByaXZhdGUxFDASBgNVBAMMC3d3dy5pYm0u
Y29tMB4XDTIwMDEwMzA5MTI0OVoXDTIwMDQwMjA5MTI0OVowNzEVMBMGA1UEChMM
Y2VydC1tYW5hZ2VyMR4wHAYDVQQDExVwbGF0Zm9ybS1hdXRoLXNlcnZpY2UwggEi
MA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDID30+mQ7TCakhVJiVb8wLBOXP
gYSSytm1v6ozxrwkuNNzF9JKeSzqRNUS2wIplFJvS4TEznX/Gbi9qVNWDAeYOyUp
d3pzGYNlMIDC1/nbtcK5wFWgTbOntLjbPWxqnYm3K5pWgdDVUEMVyvFKbGw5V3sw
lzNffi2gVqwg1yPlzgnlc5VJrNrNsrGhn6Kt1FALgCG5bDagXGuDGKRQnq3zfwmd
IDXbvnDDfwdRkq9eyxU6M8RQKNKb0CR/AtAKxVtuzcF6SumUp6R4TIk3Nis2Rwj7
wzYv8B/oUzFf+N3g4uDs4VOEs5dFMDGVhboQFEwkIACMBLkGE3Ai4MJ5Gbp5AgMB
AAGjezB5MA4GA1UdDwEB/wQEAwIFoDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaA
FMLYcARQHKMvdQ2rkDD3DCnNJCZnMDgGA1UdEQQxMC+CFXBsYXRmb3JtLWF1dGgt
c2VydmljZYcEfwAAAYcQAAAAAAAAAAAAAAAAAAAAATANBgkqhkiG9w0BAQsFAAOC
AgEAHP24Nh5Ogk60ydGJezksUisqy4egHywBogXAfhhhwooxcBNLvnEXN1fE/cHF
APYqWDFhfArG7YkvNtA35V1Gfz/iQnvYRFOaLeY50D3TiLkiZadYLf1OMpORQWlX
xy45lGMdqcHYsWOdvDF6K5dn6hcn7zy65cGa4lVdg74/YlzFk7W9rvh7ymQtHSQm
1aqvgGjhWzd/SNctbdJ30vQ0irfBdx7s+IeC8+o58Nn6EYIQn30lQuZ9+PfBCAxx
5cUdv2PSXU7VCmgtvVgeEG91mW/SRKtU8XHUzLCQpIGVlal/zbseYIw0qHdQ+ynE
VXbFVt3NYK6BtmTahwKVeaP/xt2GSNvk3W6c5oGOQLiR9PogTtgL/Q0Mzk6Voe0i
YJINnuv3IISp8IHr3gKD+FX6bpfyCRljTVAbkek++Mo4TyOn1mq8gzYPjCLI2B8R
zZ6XFHw8bRYc/E078X65jy+Cfm0Ft59YsF9X5IaEIXWhnqlO/nqwpAP9LxC1OY4e
9gDdvf9uiAIRP4jBO5QIYVbMIxWA1ZB6hMootO9imlCgjdJzNxkG6jvuJeGWsIaB
EoZ8At/x+YkKI7byVYnddlyuk8XD0h9YbiB8fvbyImNydVgLhQ6TvRGzBVrUW0Aj
JM1y82EUGayoE3ls/P4t13IotIVdLJG7m8RvHB0np2OW+m0=
-----END CERTIFICATE-----
`

const authKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAyA99PpkO0wmpIVSYlW/MCwTlz4GEksrZtb+qM8a8JLjTcxfS
Snks6kTVEtsCKZRSb0uExM51/xm4valTVgwHmDslKXd6cxmDZTCAwtf527XCucBV
oE2zp7S42z1sap2JtyuaVoHQ1VBDFcrxSmxsOVd7MJczX34toFasINcj5c4J5XOV
SazazbKxoZ+irdRQC4AhuWw2oFxrgxikUJ6t838JnSA1275ww38HUZKvXssVOjPE
UCjSm9AkfwLQCsVbbs3BekrplKekeEyJNzYrNkcI+8M2L/Af6FMxX/jd4OLg7OFT
hLOXRTAxlYW6EBRMJCAAjAS5BhNwIuDCeRm6eQIDAQABAoIBABDG+8HcJBoenUhm
R8WAcqQZ7QvR2MRWrptHT4a4C98ogNMOFFyafOHIk/XY1/2UqrUaqC+5ALw78/oz
0QFOUPS7QQqRMnukkBhUUPgGFLYQMBxTvNqMqZxO05HPPOBUBdCOn4rRqCm7aWTS
eBm9rfXBGFui0nQXqMvheN3oLOKAluv2MbWlViW/XL/LKUc7QFfinfRHzXzeGvNS
jY/FLvw71swtLeIEc+SGp/2AXOeps7tis0AD9PAuz6udqAWqNc5GYvQQ2/JylRBM
3F2iWs+mE8OVvl4AanIUxypuUDkW0Q7SJH8arykOEF3d52GNLeLdBVwV/pT9rIWR
oUk3jmECgYEA58zPa6Tql1QRaZCZUuA2CFSdNCG+cFnJZZX97VDnedU7j5OPwkGi
ibJCYT9r6T821hm57uZlxrXQaVQb7Qc4cF7uSY9xyPcjrG+1mCav+TWJbSgvL/Vr
iK2e4wuMBQ5+G0r4rX9BFUnf/yQRh5uL/unJdqJ9NQQCmTWwpeTdJ18CgYEA3PJk
7FpGvqxsn6zKRINAxMYGzTmBuKSu+q4+0bhqIE105Vr8a3ECqSoVrzFvsr7CS484
r6aJzXnCoDmPJe7zmoziBQgb1BI3+WPLo3C5ObwibwhsD74dWaDXpe+PAXeD+SqE
QVoS0kvugCQGxUxytw3b7njWY9KK/kZVba5IJScCgYAE/LERvdCWRObDAKtrzwsa
jsd7EabJQAVHPDzkWDNCucW4RxJ2uXbHw6AN+5FUyVlirdcEPsy9w7eiLH8VUGx0
bnZt+roV+ss3sIeVqG9syywTvyOUwpD3tdSPSfZPPYqVB73l7bD2xkodcSc3Za9T
YrBE7yYGd/meVjtgtlXTvQKBgD/roGcAdxcSqxpG0v2fO02yPKWCzZStkDHV1za1
M37E6dywxfYPa8Wk90EH/Fip92wEs1GdzKA1aakQ/ejriG5d88Wg4uwnmSn9RRyz
r+KUGyIkPZ8QUU1syvamp6i0+ulK85g0ht1wOj+4LH1H8KgkykvjiaK3dp65ml1w
3k5hAoGBAMm2/rG5ZUaVQK1yDML7pJDzonghdpXIhf2sdl2hy041WMgjKbuCEMFM
JAphihBzLg6TK9LUMxT17hjH4nooGyNfKuYTvT/12DwlpJ4WyNOt/a9YjOLxGRwg
tj2wAwR6zzLBAux0hrfYVtqtsnHRGCzrnsZx4oVLQD4mpUmZeZiq
-----END RSA PRIVATE KEY-----
`

const mgmtCert = `
-----BEGIN CERTIFICATE-----
MIIExjCCAq6gAwIBAgIRAPNdqhcVBngJvsUEm3jpjHwwDQYJKoZIhvcNAQELBQAw
YzELMAkGA1UEBhMCVVMxETAPBgNVBAgMCE5ldyBZb3JrMQ8wDQYDVQQHDAZBcm1v
bmsxGjAYBgNVBAoMEUlCTSBDbG91ZCBQcml2YXRlMRQwEgYDVQQDDAt3d3cuaWJt
LmNvbTAeFw0yMDAxMDMwOTEyNTFaFw0yMDA0MDIwOTEyNTFaMD4xFTATBgNVBAoT
DGNlcnQtbWFuYWdlcjElMCMGA1UEAxMccGxhdGZvcm0taWRlbnRpdHktbWFuYWdl
bWVudDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAK/oWKHF4aFQ6eIf
lO7At3TVI83lkCEWRiCMiDR9wG1AAPLevBXifTlLS2wIcLN0lwudmoSqNrB5hkw8
IkOklxX1RD7dL2gGxlOegE0kLX2JUNKuj8nS49JtzRK0CtSZe4XxkO+xhnEtkB0V
9zEYjgiVoCHM1nUfY1TZuP+RRxhJQj5FRt4jmOPJk9pg74lAC9YkbQoFBRD1jFNE
pX66PQAwd8B6HnsMEOydnJaNG1BuWYaMioP8qoo6MeSVkOrFqH5hh61Jp9h+uD7x
iH88g3AxZLL+Tl2/Ua6F1Og0q9gBw4NDps4LA3nV1pK0ZT+pqjYnI18bLTJXP+Qa
TJPEBOcCAwEAAaOBmTCBljAOBgNVHQ8BAf8EBAMCBaAwDAYDVR0TAQH/BAIwADAf
BgNVHSMEGDAWgBTC2HAEUByjL3UNq5Aw9wwpzSQmZzBVBgNVHREETjBMghxwbGF0
Zm9ybS1pZGVudGl0eS1tYW5hZ2VtZW50gixwbGF0Zm9ybS1pZGVudGl0eS1tYW5h
Z2VtZW50Lmt1YmUtc3lzdGVtLnN2YzANBgkqhkiG9w0BAQsFAAOCAgEATMwZcSsC
r66gcYArbdg6ftU0ZFDRfBB4d4dSHi8Qle0zq4iX5HY0vtEIyJ1lnCtAoOSBG1GT
qoghTr/azMO9YUzhCn4osE6LuYVK6pPiX1yBNdx3Gb98EP+EuucFBjXIh83rrAr6
qD8gXgfCDyBkvhqXfItZy8hmpQyBj4lrfKjXZwh4xIr/bRGj1vFu2NIXNTshdlTf
dsyOyLrqsNxexdkQvPKIK89jTEPT7QiFnvIjTrI6ysjhTVvedl7va9A6PEdjKB2N
IRzvyvj3AAg3RTLbd6FVSlPJG/qHbdy0NeaWwsGPn3CzRhX40iR0O5r7w/FiIELK
lmMwyQB/l5yhWadzBpTtYLCkQA96keq2e7FzvZquzdMPILWuTbpcURfO3df0QAxR
/VJlMr2UPXTRHtjT67bAYzQ/Jkp6BxB0+c3zriJAMGm/KAWQzr44FDnWqTCmJiql
ojhTuzBiBVYr7YW+gyBsCXggetnRLIZx2nX/OcbPp0Xs9JxHwL42ROrybkLLTOn9
YWKryd2uqeqcIbsam7A0/DOXcGWTtdconokziOng81Imm2yfOmMT7uTN80R32yeV
Qghoc4/+K3ZaDbykSbbVnFxUxHihrpzAZ3jBrc4YIKbuNmSkGT7rjvqz28kU9qEl
a5Q3l6jlyME4/uNTB8Igjrzal7LuqDHY1/g=
-----END CERTIFICATE-----
`

const mgmtKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAr+hYocXhoVDp4h+U7sC3dNUjzeWQIRZGIIyINH3AbUAA8t68
FeJ9OUtLbAhws3SXC52ahKo2sHmGTDwiQ6SXFfVEPt0vaAbGU56ATSQtfYlQ0q6P
ydLj0m3NErQK1Jl7hfGQ77GGcS2QHRX3MRiOCJWgIczWdR9jVNm4/5FHGElCPkVG
3iOY48mT2mDviUAL1iRtCgUFEPWMU0Slfro9ADB3wHoeewwQ7J2clo0bUG5ZhoyK
g/yqijox5JWQ6sWofmGHrUmn2H64PvGIfzyDcDFksv5OXb9RroXU6DSr2AHDg0Om
zgsDedXWkrRlP6mqNicjXxstMlc/5BpMk8QE5wIDAQABAoIBAC5OV6aeVLsGdcsU
zTvcd27F/Nrip4JOPr+qvuAyikk4JAfLInAQYy6F56DUx/Y1pt1j650owaTLGNNQ
sSRlW3BoVELILwaPYf35J4UnFOKLFz+evDNA1MA8L1PRGVjF3oJwCNtQPDfm6pz9
f0jpUdpoqPK83WePj7JybgS41aJKlqONOZP6H0COF6d+s+zVHstNABTIodbs7R4m
/8ZagrUwob4ctnozC9MdEphKGBfnAhWw7dirWi+sZXxeJoWhbA/aZJm7RLl7yCah
fXauuvPeXWVwKGvLvKygKpCyCoWAMnoCaawEBD+INPmlFwKbWUMwWIawcfSgFlz1
VSuMC0ECgYEA5EoTynCvR1g8xCRB0EAave4joawVqLwkJYJeJ6u8Ho5X8VZTJO1M
WWhuwiuUc9W6YyxSZH+EivL2iKsCKktcHhDt2IFzFXuJN7ptqqdFFRAmybWs7r8K
X5MjNFGd5TMgES+vaK98Dxe0S9TuE+UEmBKZtaN6J7UsInRAwQMAxj8CgYEAxUKL
ZIyAcEBi2A6V5qrjfBQbdEQa1EuTq9wtGIsysS0A195z1NQuYnzAEhKzFX6iPrJE
EVaKPmkqiVb2TbJY9JhyulQGFDi+oiLpEPIy4Yul63p5tSbdqBXaTMP8pDmG0uE/
T0vhg51caj8UXEes6F/Vblc+mb4xEMdYHmXUp1kCgYEAm8BymjahT5LC+qsl/9gM
prKXutD2ggXCv5ifdn9zvIEvPOfBmDsSURmwH2euRA14ufreR9S8cDRtwTazJnn9
4kF4lLNU+j4rk0rtUXYH6uP3Ir2Bu+25PRXc+fAJQvvLkA4xpcG8aWupxecDe61P
ohVmq4daa0bvUp5wI39SAu8CgYBKJ6r5QKxkBoBiEpGdpm5gTbrIaXXgiwzXlazK
RkHnQzDG9hR7VMyfL6CQ6sCx+uoJQcC+99Z+gdCA0tVC/iHcyZaPn1itr2tSzmbp
fPNNwM6+CEvKovK+5oArTZ4jKpEZo53GNJNxg+2i71W7HyLMNIKquVeKOmp9kUL6
InPf+QKBgBRldGHF8aFqWtM3LYY6XWHAw7jujg7wlPFzcBJ+ergePjzIYSLdx3zd
Q6YnL8QwKvc7gIPdyExIlOPV/eXuFwAmg26k5TntW5k+JZcM4pU26CUYVQmGn5H4
YzpJKXZQwcQ+SQ6m9YebyOVZ1gkQIGdSjz5gGdgS63MG8nweiS4D
-----END RSA PRIVATE KEY-----
`

const registerClientScript = `#!/bin/sh
HTTP_CODE=""
while true
do
  HTTP_CODE=$(curl -k -o /dev/null -I  -w "%{http_code}"  -X GET -u oauthadmin:$WLP_CLIENT_REGISTRATION_SECRET -H "Content-Type: application/json" https://platform-auth-service:9443/oidc/endpoint/OP/registration/$WLP_CLIENT_ID)
  if [ $HTTP_CODE = "404" -o $HTTP_CODE = "200" ] ; then
	 break;
  fi
done
if [ $HTTP_CODE = "404" ]
then
  echo "Running new client id registration"
  sed "s/WLP_CLIENT_ID/$WLP_CLIENT_ID/g" /jsons/platform-oidc-registration.json
  sed "s/WLP_CLIENT_SECRET/$WLP_CLIENT_SECRET/g" /jsons/platform-oidc-registration.json
  sed "s/OIDC_ISSUER_URL/$OIDC_ISSUER_URL/g" /jsons/platform-oidc-registration.json
  until curl -i -k -X POST -u oauthadmin:$WLP_CLIENT_REGISTRATION_SECRET \
   -H "Content-Type: application/json" \
   --data @/jsons/platform-oidc-registration.json \
   https://platform-auth-service:9443/oidc/endpoint/OP/registration | grep '201 Created'; do sleep 1; done;
else
  echo "Running update client id registration."
  sed "s/WLP_CLIENT_ID/$WLP_CLIENT_ID/g" /jsons/platform-oidc-registration.json
  sed "s/WLP_CLIENT_SECRET/$WLP_CLIENT_SECRET/g" /jsons/platform-oidc-registration.json
  sed "s/OIDC_ISSUER_URL/$OIDC_ISSUER_URL/g" /jsons/platform-oidc-registration.json
  until curl -i -k -X PUT -u oauthadmin:$WLP_CLIENT_REGISTRATION_SECRET \
   -H "Content-Type: application/json" \
   --data @/jsons/platform-oidc-registration.json \
   https://platform-auth-service:9443/oidc/endpoint/OP/registration/$WLP_CLIENT_ID | grep '200 OK'; do sleep 1; done;
fi
`

const registrationJson = `{
   "token_endpoint_auth_method":"client_secret_basic",
   "client_id": "WLP_CLIENT_ID",
   "client_secret": "WLP_CLIENT_SECRET",
   "scope":"openid profile email",
   "grant_types":[
      "authorization_code",
      "client_credentials",
      "password",
      "implicit",
      "refresh_token",
      "urn:ietf:params:oauth:grant-type:jwt-bearer"
   ],
   "response_types":[
      "code",
      "token",
      "id_token token"
   ],
   "application_type":"web",
   "subject_type":"public",
   "post_logout_redirect_uris":["OIDC_ISSUER_URL/console/logout"],
   "preauthorized_scope":"openid profile email general",
   "introspect_tokens":true,
   "trusted_uri_prefixes":["OIDC_ISSUER_URL"],
   "redirect_uris":["OIDC_ISSUER_URL/auth/liberty/callback","https://127.0.0.1:443/idauth/oidc/endpoint/OP"]
}
`
