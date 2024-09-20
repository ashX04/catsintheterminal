package library

import (
	"time"

	"golang.org/x/exp/rand"
)

var splart = "safasfsaf"

func Ascii_archive() string {
	// Seed the random number generator
	rand.Seed(uint64(time.Now().UnixNano()))

	//array of strings
	art := []string{
		`

Cute Cat brings you fortune
	 /\_/\  
	( o.o ) 
	 > ^ < 
`,
		`
Tiny Bear wishes you good luck

ʕ•ᴥ•ʔ
`,
		`
Bunny says relax

(\(\ 
( -.-)
o_(\")(\")
`,
		`
Koala says you are cute

  (\/)
 (o.o)
  (> <)
`,
		`
Cute Bird brings you food

  \\
  (o>
\\_//)
 \_/_)
  _|_
`,
		`
Smiling Frog tells you to be immune to failure like his skin

   @..@
  (----)
 ( >__< )
 ^^  ^^
`,
		`
Sleepy Owl grants you crazy amounts of luck

  ,___,
  [O.o]
  /)  )
 -"-"-
`,
		`
Chubby Pig hopes you grow chubby like him from success

  ^  ^
 ( 'oo' )
  (    )
  ^^ ^^
`,
		`
Jumping Bunny wishes you luck

	(\(\ 
	(='.')
   c(_(\")(\")
`,
		`
Tiny Bird tells you to be proud of yourself
   
	 \\
	 (o>
   \\_//)
	\_/_)
	 _|_
   `,
		`
cute ducky wants you to feed her

   <(o )___
	(  ._> /
	 '---' 
`,
		`
Sleepy Bird suggests you sleep it off

   zZz
  ~(_o)
   / \
`,
		` 
Singing Bird sings your glory

 ♪ ♫ ♪
 (o>
 / ) )
/  | |
`,
		`
♡ ∩_∩ 
  („• ֊ •„)♡
|￣U U￣￣￣￣￣￣￣￣￣|
|        I love you    |   
￣￣￣￣￣￣￣￣￣￣￣￣
`,
		`
    /)  /)
ପ(˶•-•˶)ଓ ♡
   /づ  づ
`,
		`
bunny wished the other one were you                 




                                                        ★ ⁺.

                                        (\_(\    /)_/)
                                        (    )  (     )
                                       /    |    |    \
                                      ( O   |    |   O )
`,
		`
cute cat brings you money
  ^~^  ,
 ("Y") )
 /   \/ 
(\|||/)
`,
		`
ʕ•́ᴥ•̀ʔっ
`,
		`ᕙ('▿')ᕗ
`,

		`≧◠‿●‿◠≦
`,

		`ᓚᘏᗢ`,
		// 		`
		//         へ   ♡   ╱|、
		//      ૮ - ՛ )    ('  -7
		//      /  ⁻ ៸|     |、⁻〵
		//  乀 (ˍ, ل ل      じしˍ,)ノ
		// `,
		// 		`'⎚⩊⎚' -✧`,
		// 		`
		//   ∧,,,∧
		// (  ̳• · • ̳)
		// /    づ♡ I love you
		// `,
	}
	randomart := art[rand.Intn(len(art))]
	return randomart

}
