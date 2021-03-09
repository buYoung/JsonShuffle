# JsonShuffle
서버와 클라이언트간의 통신을 할때 AES 암호화를 사용할려했으나, 조금만 생각해보니 취약점이 바로보이더라구요..

그래서 생각한게 JSONOBJECT로 통신을할꺼니까 거기다가 시간정보를 넣자! (세션개념으로 접근)

그중 필요한게 AES 암호화를한다해도 정보가 비슷하면 암호화돼서 출력되는 내용이 비슷합니다.

그래서 JSONOBJECT를 Shuffle해줄 필요가있더라구요..

처음에는 rand의 Shuffle을 사용해서 map[stirng]interface{}를 섞어볼려했으나  map[int]는 해볼만했지만 map[string]이라 시도도 해봤다가 실패의 경험만 쌓았습니다. :D

그러다 접근한게 JSONOBJECT의 시작은 "{" 이고 끝은 "}" 이니까 이걸기준으로 정보를 나누고 ","로 split으로 배열을 출력후 rand.Shuffle을 쓰면 어떨까? 하고 생각하게됩니다.

성능은 나름 나쁘지 않네요 100000번기준 최소 135ms (1회 시도시 1us 1마이크로초가 나옵니다.)

## Running Test
- 100000번 기준입니다.
* 2021/03/09 10:53:02 1번방식  136.2519ms (Trim)
* 2021/03/09 10:53:02 2번방식  139.0153ms (Replace)
* 2021/03/09 10:53:02 3번방식  137.0612ms (ReplaceAll)
* 2021/03/09 10:53:04 4번방식  1.7672294s (RegExp Replace)
* 2021/03/09 10:53:04 5번방식  610.8187ms (rune array cut)
* 2021/03/09 10:53:04 6번방식  113.6091ms (char array cut)
### Environment
* Windows 10 Home
* AMD 1700
* 24GB 12800 ddr4

### ETC
main.go에있는 UniqueRand는 Stack Overflow를 참조했습니다. 
코드가 너무 이상하면 저를위해 태클을 부탁드립니다! :D
